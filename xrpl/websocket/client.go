package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync/atomic"
	"time"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/xrpl/currency"
	"github.com/Peersyst/xrpl-go/xrpl/hash"
	transaction "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/mitchellh/mapstructure"

	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/server"
	streamtypes "github.com/Peersyst/xrpl-go/xrpl/queries/subscription/types"
	requests "github.com/Peersyst/xrpl-go/xrpl/queries/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
	"github.com/Peersyst/xrpl-go/xrpl/websocket/interfaces"
	wstypes "github.com/Peersyst/xrpl-go/xrpl/websocket/types"
	ws "github.com/gorilla/websocket"

	xrplCommon "github.com/Peersyst/xrpl-go/xrpl/common"
)

const (
	DefaultFeeCushion float32 = 1.2
	DefaultMaxFeeXRP  float32 = 2
)

var (
	ErrIncorrectID          = errors.New("incorrect id")
	ErrNotConnectedToServer = errors.New("not connected to server")
	ErrRequestTimedOut      = errors.New("request timed out")
)

type Client struct {
	cfg  ClientConfig
	conn *Connection

	// Channels
	errChan          chan error
	requestChan      chan *ClientResponse
	ledgerClosedChan chan *streamtypes.LedgerStream
	validationChan   chan *streamtypes.ValidationStream
	transactionChan  chan *streamtypes.TransactionStream
	peerStatusChan   chan *streamtypes.PeerStatusStream
	orderBookChan    chan *streamtypes.OrderBookStream
	bookChangesChan  chan *streamtypes.BookChangesStream
	consensusChan    chan *streamtypes.ConsensusStream

	idCounter atomic.Uint32
	NetworkID uint32
}

// Creates a new websocket client with cfg.
// This client will open and close a websocket connection for each request.
func NewClient(cfg ClientConfig) *Client {
	return &Client{
		cfg:         cfg,
		requestChan: make(chan *ClientResponse),
		errChan:     make(chan error),
		conn:        NewConnection(cfg.host),
	}
}

// Connect opens a websocket connection to the server. It starts reading messages in a goroutine.
func (c *Client) Connect() error {
	err := c.conn.Connect()
	if err != nil {
		return err
	}
	go c.readMessages()
	return nil
}

// Disconnect closes the websocket connection.
func (c *Client) Disconnect() error {
	return c.conn.Disconnect()
}

// IsConnected returns true if the client is connected to the server.
func (c *Client) IsConnected() bool {
	return c.conn.IsConnected()
}

func (c *Client) FaucetProvider() xrplCommon.FaucetProvider {
	return c.cfg.faucetProvider
}

// Autofill fills in the missing fields in a transaction.
func (c *Client) Autofill(tx *transaction.FlatTransaction) error {
	if err := c.setValidTransactionAddresses(tx); err != nil {
		return err
	}

	err := c.setTransactionFlags(tx)
	if err != nil {
		return err
	}

	if _, ok := (*tx)["NetworkID"]; !ok {
		if c.NetworkID != 0 {
			(*tx)["NetworkID"] = c.NetworkID
		}
	}
	if _, ok := (*tx)["Sequence"]; !ok {
		err := c.setTransactionNextValidSequenceNumber(tx)
		if err != nil {
			return err
		}
	}
	if _, ok := (*tx)["Fee"]; !ok {
		err := c.calculateFeePerTransactionType(tx, 0)
		if err != nil {
			return err
		}
	}
	if _, ok := (*tx)["LastLedgerSequence"]; !ok {
		err := c.setLastLedgerSequence(tx)
		if err != nil {
			return err
		}
	}
	if txType, ok := (*tx)["TransactionType"].(transaction.TxType); ok {
		if acc, ok := (*tx)["Account"].(types.Address); txType == transaction.AccountDeleteTx && ok {
			err := c.checkAccountDeleteBlockers(acc)
			if err != nil {
				return err
			}
		}
		if txType == transaction.PaymentTx {
			err := c.checkPaymentAmounts(tx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// AutofillMultisigned fills in the missing fields in a multisigned transaction.
// This function is used to fill in the missing fields in a multisigned transaction.
// It fills in the missing fields in the transaction and calculates the fee per number of signers.
func (c *Client) AutofillMultisigned(tx *transaction.FlatTransaction, nSigners uint64) error {
	err := c.Autofill(tx)
	if err != nil {
		return err
	}

	err = c.calculateFeePerTransactionType(tx, nSigners)
	if err != nil {
		return err
	}

	return nil
}

// FundWallet funds a wallet with XRP from the faucet.
// If the wallet does not have a classic address, it will return an error.
func (c *Client) FundWallet(wallet *wallet.Wallet) error {
	if wallet.ClassicAddress == "" {
		return errors.New("fund wallet: cannot fund a wallet without a classic address")
	}

	err := c.cfg.faucetProvider.FundWallet(wallet.ClassicAddress)
	if err != nil {
		return err
	}

	return nil
}

// Request sends a request to the server and returns the response.
// This function is used to send requests to the server.
// It returns the response from the server.
func (c *Client) Request(req interfaces.Request) (*ClientResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id := c.idCounter.Add(1)

	msg, err := c.formatRequest(req, int(id), nil)
	if err != nil {
		return nil, err
	}

	if !c.IsConnected() {
		return nil, ErrNotConnectedToServer
	}

	err = c.conn.WriteMessage(msg)
	if err != nil {
		return nil, err
	}

	res, err := c.awaitResponse(int(id))
	if err != nil {
		return nil, err
	}

	if res.ID != int(id) {
		return nil, ErrIncorrectID
	}
	if err := res.CheckError(); err != nil {
		return nil, err
	}

	return res, nil
}

// SubmitTxBlob sends a pre-signed transaction blob to the server.
// It decodes the blob to confirm that it contains either a signature
// or a signing public key, and then submits it using a submission request.
// The failHard flag determines how strictly errors are handled.
func (c *Client) SubmitTxBlob(txBlob string, failHard bool) (*requests.SubmitResponse, error) {
	tx, err := binarycodec.Decode(txBlob)
	if err != nil {
		return nil, err
	}

	_, okTxSig := tx["TxSignature"].(string)
	_, okPubKey := tx["SigningPubKey"].(string)

	if !okTxSig && !okPubKey {
		return nil, ErrMissingTxSignatureOrSigningPubKey
	}

	return c.submitRequest(&requests.SubmitRequest{
		TxBlob:   txBlob,
		FailHard: failHard,
	})
}

// SubmitTx signs the transaction (if needed) by using getSignedTx,
// and then submits it to the server via a submission request.
// It uses the provided submit options to decide whether to autofill missing
// fields and whether to enforce a failHard mode during submission.
func (c *Client) SubmitTx(tx transaction.FlatTransaction, opts *xrplCommon.SubmitOptions) (*requests.SubmitResponse, error) {
	txBlob, err := getSignedTx(c, tx, opts.Autofill, opts.Wallet)
	if err != nil {
		return nil, err
	}

	return c.submitRequest(&requests.SubmitRequest{
		TxBlob:   txBlob,
		FailHard: opts.FailHard,
	})
}

// SubmitMultisigned sends a multisigned transaction to the server and returns the response.
// This function is used to send multisigned transactions to the server.
// It returns the response from the server.
func (c *Client) SubmitMultisigned(txBlob string, failHard bool) (*requests.SubmitMultisignedResponse, error) {
	tx, err := binarycodec.Decode(txBlob)
	if err != nil {
		return nil, err
	}
	signers, okSigners := tx["Signers"].([]interface{})

	if okSigners && len(signers) > 0 {
		for _, sig := range signers {
			signer := sig.(map[string]any)
			signerData := signer["Signer"].(map[string]any)
			if signerData["SigningPubKey"] == "" && signerData["TxnSignature"] == "" {
				return nil, errors.New("signer data is empty")
			}
		}
	}

	return c.submitMultisignedRequest(&requests.SubmitMultisignedRequest{
		Tx:       tx,
		FailHard: failHard,
	})
}

// SubmitTxBlobAndWait sends a pre-signed transaction blob to the server,
// decodes it to retrieve the required LastLedgerSequence, submits the blob,
// and then waits until the transaction is confirmed in a ledger. It returns
// the transaction response if the submission is successful.
func (c *Client) SubmitTxBlobAndWait(txBlob string, failHard bool) (*requests.TxResponse, error) {
	tx, err := binarycodec.Decode(txBlob)
	if err != nil {
		return nil, err
	}

	lastLedgerSequence := tx["LastLedgerSequence"].(uint32)

	txResponse, err := c.SubmitTxBlob(txBlob, failHard)
	if err != nil {
		return nil, err
	}

	if txResponse.EngineResult != "tesSUCCESS" {
		return nil, &ClientError{ErrorString: "transaction failed to submit with engine result: " + txResponse.EngineResult}
	}

	txHash, err := hash.SignTxBlob(txBlob)
	if err != nil {
		return nil, err
	}

	return c.waitForTransaction(txHash, lastLedgerSequence)
}

// SubmitTxAndWait prepares a transaction by ensuring it is fully signed (using
// getSignedTx), submits it to the server, and waits for ledger confirmation.
// It validates that the transaction's EngineResult is successful before returning
// the transaction response.
func (c *Client) SubmitTxAndWait(tx transaction.FlatTransaction, opts *xrplCommon.SubmitOptions) (*requests.TxResponse, error) {
	txBlob, err := getSignedTx(c, tx, opts.Autofill, opts.Wallet)
	if err != nil {
		return nil, err
	}

	decodedTx, err := binarycodec.Decode(txBlob)
	if err != nil {
		return nil, err
	}

	lastLedgerSequence, ok := decodedTx["LastLedgerSequence"].(uint32)
	if !ok {
		return nil, errors.New("missing LastLedgerSequence in transaction")
	}

	txResponse, err := c.SubmitTx(tx, opts)
	if err != nil {
		return nil, err
	}

	if txResponse.EngineResult != "tesSUCCESS" {
		return nil, &ClientError{
			ErrorString: "transaction failed to submit with engine result: " + txResponse.EngineResult,
		}
	}

	txHash, err := hash.SignTxBlob(txBlob)
	if err != nil {
		return nil, err
	}

	return c.waitForTransaction(txHash, lastLedgerSequence)
}

func (c *Client) waitForTransaction(txHash string, lastLedgerSequence uint32) (*requests.TxResponse, error) {
	var txResponse *requests.TxResponse
	i := 0

	for i < c.cfg.maxRetries {
		// Get the current ledger index
		currentLedger, err := c.GetLedgerIndex()
		if err != nil {
			return nil, err
		}

		// Check if the transaction has been included in the current ledger
		if currentLedger.Int() >= int(lastLedgerSequence) {
			break
		}

		// Request the transaction from the server
		res, err := c.Request(&requests.TxRequest{
			Transaction: txHash,
		})
		if err != nil {
			return nil, err
		}

		err = res.GetResult(&txResponse)
		if err != nil {
			return nil, err
		}

		// Check if the transaction has been included in the current ledger
		if txResponse.LedgerIndex.Int() >= int(lastLedgerSequence) {
			break
		}

		// Wait for the retry delay before retrying
		time.Sleep(c.cfg.retryDelay)
		i++
	}

	if txResponse == nil {
		return nil, errors.New("transaction not found")
	}

	return txResponse, nil
}

func (c *Client) submitMultisignedRequest(req *requests.SubmitMultisignedRequest) (*requests.SubmitMultisignedResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var subRes requests.SubmitMultisignedResponse
	err = res.GetResult(&subRes)
	if err != nil {
		return nil, err
	}
	return &subRes, nil
}

func (c *Client) submitRequest(req *requests.SubmitRequest) (*requests.SubmitResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var subRes requests.SubmitResponse
	err = res.GetResult(&subRes)
	if err != nil {
		return nil, err
	}
	return &subRes, nil
}

func (c *Client) formatRequest(req interfaces.Request, id int, marker any) ([]byte, error) {
	m := make(map[string]any)
	m["id"] = id
	m["command"] = req.Method()
	m["api_version"] = req.APIVersion()
	if marker != nil {
		m["marker"] = marker
	}
	dec, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &m})
	err := dec.Decode(req)
	if err != nil {
		return nil, err
	}

	return json.Marshal(m)
}

// TODO: Implement this when IsValidXAddress is implemented
func (c *Client) getClassicAccountAndTag(address string) (string, uint32) {
	return address, 0
}

func (c *Client) convertTransactionAddressToClassicAddress(tx *transaction.FlatTransaction, fieldName string) {
	if address, ok := (*tx)[fieldName].(string); ok {
		classicAddress, _ := c.getClassicAccountAndTag(address)
		(*tx)[fieldName] = classicAddress
	}
}

func (c *Client) validateTransactionAddress(tx *transaction.FlatTransaction, addressField, tagField string) error {
	classicAddress, tag := c.getClassicAccountAndTag((*tx)[addressField].(string))
	(*tx)[addressField] = classicAddress

	if tag != uint32(0) {
		if txTag, ok := (*tx)[tagField].(uint32); ok && txTag != tag {
			return fmt.Errorf("the %s, if present, must be equal to the tag of the %s", addressField, tagField)
		}
		(*tx)[tagField] = tag
	}

	return nil
}

// Sets valid addresses for the transaction.
func (c *Client) setValidTransactionAddresses(tx *transaction.FlatTransaction) error {
	// Validate if "Account" address is an xAddress
	if err := c.validateTransactionAddress(tx, "Account", "SourceTag"); err != nil {
		return err
	}

	if _, ok := (*tx)["Destination"]; ok {
		if err := c.validateTransactionAddress(tx, "Destination", "DestinationTag"); err != nil {
			return err
		}
	}

	// DepositPreuaht
	c.convertTransactionAddressToClassicAddress(tx, "Authorize")
	c.convertTransactionAddressToClassicAddress(tx, "Unauthorize")
	// EscrowCancel, EscrowFinish
	c.convertTransactionAddressToClassicAddress(tx, "Owner")
	// SetRegularKey
	c.convertTransactionAddressToClassicAddress(tx, "RegularKey")

	return nil
}

// Sets the next valid sequence number for a given transaction.
func (c *Client) setTransactionNextValidSequenceNumber(tx *transaction.FlatTransaction) error {
	if _, ok := (*tx)["Account"].(string); !ok {
		return errors.New("missing Account in transaction")
	}
	res, err := c.GetAccountInfo(&account.InfoRequest{
		Account:     types.Address((*tx)["Account"].(string)),
		LedgerIndex: common.LedgerTitle("current"),
	})

	if err != nil {
		return err
	}

	(*tx)["Sequence"] = uint32(res.AccountData.Sequence)
	return nil
}

// Calculates the current transaction fee for the ledger.
// Note: This is a public API that can be called directly.
func (c *Client) getFeeXrp(cushion float32) (string, error) {
	res, err := c.GetServerInfo(&server.InfoRequest{})
	if err != nil {
		return "", err
	}

	if res.Info.ValidatedLedger.BaseFeeXRP == 0 {
		return "", errors.New("getFeeXrp: could not get BaseFeeXrp from ServerInfo")
	}

	loadFactor := res.Info.LoadFactor
	if res.Info.LoadFactor == 0 {
		loadFactor = 1
	}

	fee := res.Info.ValidatedLedger.BaseFeeXRP * float32(loadFactor) * cushion

	if fee > c.cfg.maxFeeXRP {
		fee = c.cfg.maxFeeXRP
	}

	// Round fee to NUM_DECIMAL_PLACES
	roundedFee := float32(math.Round(float64(fee)*math.Pow10(int(currency.MaxFractionLength)))) / float32(math.Pow10(int(currency.MaxFractionLength)))

	// Convert the rounded fee back to a string with NUM_DECIMAL_PLACES
	return fmt.Sprintf("%.*f", currency.MaxFractionLength, roundedFee), nil
}

// Calculates the fee per transaction type.
//
// TODO: Add fee support for `EscrowFinish` `AccountDelete`, `AMMCreate`, and multisigned transactions.
func (c *Client) calculateFeePerTransactionType(tx *transaction.FlatTransaction, nSigners uint64) error {
	fee, err := c.getFeeXrp(c.cfg.feeCushion)
	if err != nil {
		return err
	}

	feeDrops, err := currency.XrpToDrops(fee)
	if err != nil {
		return err
	}

	if nSigners > 0 {
		// Convert feeDrops to uint64 for safe arithmetic
		baseFee, err := strconv.ParseUint(feeDrops, 10, 64)
		if err != nil {
			return err
		}

		// Calculate total signers fee: fee * nSigners
		signersFee := baseFee * nSigners

		// Add base fee and signers fee
		totalFee := baseFee + signersFee

		// Convert back to string
		feeDrops = strconv.FormatUint(totalFee, 10)
	}

	(*tx)["Fee"] = feeDrops

	return nil
}

// Sets the latest validated ledger sequence for the transaction.
// Modifies the `LastLedgerSequence` field in the tx.
func (c *Client) setLastLedgerSequence(tx *transaction.FlatTransaction) error {
	index, err := c.GetLedgerIndex()
	if err != nil {
		return err
	}

	(*tx)["LastLedgerSequence"] = index.Uint32() + xrplCommon.LedgerOffset
	return err
}

// Checks for any blockers that prevent the deletion of an account.
// Returns nil if there are no blockers, otherwise returns an error.
func (c *Client) checkAccountDeleteBlockers(address types.Address) error {
	accObjects, err := c.GetAccountObjects(&account.ObjectsRequest{
		Account:              address,
		LedgerIndex:          common.LedgerTitle("validated"),
		DeletionBlockersOnly: true,
	})
	if err != nil {
		return err
	}

	if len(accObjects.AccountObjects) > 0 {
		return errors.New("account %s cannot be deleted; there are Escrows, PayChannels, RippleStates, or Checks associated with the account")
	}
	return nil
}

func (c *Client) checkPaymentAmounts(tx *transaction.FlatTransaction) error {
	if _, ok := (*tx)["DeliverMax"]; ok {
		if _, ok := (*tx)["Amount"]; !ok {
			(*tx)["Amount"] = (*tx)["DeliverMax"]
		} else if (*tx)["Amount"] != (*tx)["DeliverMax"] {
			return errors.New("payment transaction: Amount and DeliverMax fields must be identical when both are provided")
		}
	}
	return nil
}

// Sets a transaction's flags to its numeric representation.
// TODO: Add flag support for AMMDeposit, AMMWithdraw,
// NFTTOkenCreateOffer, NFTokenMint, OfferCreate, XChainModifyBridge (not supported).
func (c *Client) setTransactionFlags(tx *transaction.FlatTransaction) error {
	flags, ok := (*tx)["Flags"].(uint32)
	if !ok && flags > 0 {
		(*tx)["Flags"] = int(0)
		return nil
	}

	_, ok = (*tx)["TransactionType"].(string)
	if !ok {
		return errors.New("transaction type is missing in transaction")
	}

	return nil
}

func (c *Client) awaitResponse(id int) (*ClientResponse, error) {
	for {
		select {
		case res := <-c.requestChan:
			if res.ID == id {
				return res, nil
			}
		case <-time.After(c.cfg.timeout):
			return nil, ErrRequestTimedOut
		}
	}
}

func (c *Client) handleMessage(message []byte) {
	var stream wstypes.Message
	c.unmarshalMessage(message, &stream)
	if stream.IsRequest() {
		c.handleRequest(message)
	} else if stream.IsStream() {
		c.handleStream(stream.Type, message)
	}
}

func (c *Client) handleRequest(message []byte) {
	var res ClientResponse
	c.unmarshalMessage(message, &res)
	c.requestChan <- &res
}

func (c *Client) unmarshalMessage(message []byte, v any) {
	if err := json.Unmarshal(message, v); err != nil {
		if c.errChan == nil {
			c.errChan = make(chan error)
		}
		c.errChan <- err
	}
}

func (c *Client) handleStream(t streamtypes.Type, message []byte) {
	switch t {
	case streamtypes.LedgerStreamType:
		var ledger streamtypes.LedgerStream
		c.unmarshalMessage(message, &ledger)

		if c.ledgerClosedChan != nil {
			c.ledgerClosedChan <- &ledger
		}
	case streamtypes.TransactionStreamType:
		var transaction streamtypes.TransactionStream
		c.unmarshalMessage(message, &transaction)
		if c.transactionChan != nil {
			c.transactionChan <- &transaction
		}
	case streamtypes.ValidationStreamType:
		var validation streamtypes.ValidationStream
		c.unmarshalMessage(message, &validation)
		if c.validationChan != nil {
			c.validationChan <- &validation
		}
	case streamtypes.PeerStatusStreamType:
		var peerStatus streamtypes.PeerStatusStream
		c.unmarshalMessage(message, &peerStatus)
		if c.peerStatusChan != nil {
			c.peerStatusChan <- &peerStatus
		}
	case streamtypes.ConsensusStreamType:
		var consensus streamtypes.ConsensusStream
		c.unmarshalMessage(message, &consensus)
		if c.consensusChan != nil {
			c.consensusChan <- &consensus
		}
	default:
		if c.errChan == nil {
			c.errChan = make(chan error)
		}
		c.errChan <- fmt.Errorf("unknown stream type: %v", t)
	}
}

func (c *Client) readMessages() {
	retryCount := 0
	maxRetries := c.cfg.maxReconnects

	for {
		if c.conn == nil {
			return
		}
		message, err := c.conn.ReadMessage()
		switch {
		case ws.IsCloseError(err) || ws.IsUnexpectedCloseError(err):
			if retryCount >= maxRetries {
				if c.errChan == nil {
					c.errChan = make(chan error)
				}
				c.errChan <- fmt.Errorf("max reconnection attempts (%d) reached", maxRetries)
				return
			}
			retryCount++
			connErr := c.conn.Connect()
			if connErr != nil {
				if c.errChan == nil {
					c.errChan = make(chan error)
				}
				c.errChan <- connErr
				return
			}
		case err != nil:
			c.errChan <- err
			return
		default:
			// Send the message to the channel
			c.handleMessage(message)
			// Reset retry count on successful message
			retryCount = 0
		}
	}
}
