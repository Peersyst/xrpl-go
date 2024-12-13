package websocket

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"sync/atomic"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/currency"
	transaction "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/mitchellh/mapstructure"

	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/server"
	requests "github.com/Peersyst/xrpl-go/xrpl/queries/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/gorilla/websocket"
)

const (
	DefaultFeeCushion float32 = 1.2
	DefaultMaxFeeXRP  float32 = 2
)

var (
	ErrIncorrectID          = errors.New("incorrect id")
	ErrNotConnectedToServer = errors.New("not connected to server")
)

type Client struct {
	cfg  ClientConfig
	conn *websocket.Conn

	idCounter atomic.Uint32
	NetworkID uint32
}

// Creates a new websocket client with cfg.
// This client will open and close a websocket connection for each request.
func NewClient(cfg ClientConfig) *Client {
	return &Client{
		cfg: cfg,
	}
}

// Connect opens a websocket connection to the server.
func (c *Client) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(c.cfg.host, nil)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

// Disconnect closes the websocket connection.
func (c *Client) Disconnect() {
	if c.conn == nil {
		return
	}
	if err := c.conn.Close(); err != nil {
		fmt.Println("Error closing websocket connection: ", err)
	}
	c.conn = nil
}

// IsConnected returns true if the client is connected to the server.
func (c *Client) IsConnected() bool {
	return c.conn != nil
}

// Conn returns the websocket connection.
func (c *Client) Conn() *websocket.Conn {
	return c.conn
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
		err := c.calculateFeePerTransactionType(tx)
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

func (c *Client) FundWallet(wallet *xrpl.Wallet) error {
	if wallet.ClassicAddress == "" {
		return errors.New("fund wallet: cannot fund a wallet without a classic address")
	}

	err := c.cfg.faucetProvider.FundWallet(wallet.ClassicAddress)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Request(req XRPLRequest) (XRPLResponse, error) {
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

	err = c.Conn().WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return nil, err
	}

	_, v, err := c.Conn().ReadMessage()
	if err != nil {
		return nil, err
	}
	jDec := json.NewDecoder(bytes.NewReader(json.RawMessage(v)))
	jDec.UseNumber()
	var res ClientXrplResponse
	err = jDec.Decode(&res)
	if err != nil {
		return nil, err
	}

	if res.ID != int(id) {
		return nil, ErrIncorrectID
	}
	if err := res.CheckError(); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Submit(txBlob string, failHard bool) (*requests.SubmitResponse, error) {
	tx, err := binarycodec.Decode(txBlob)
	if err != nil {
		return nil, err
	}

	_, okTxSig := tx["TxSignature"].(string)
	_, okPubKey := tx["SigningPubKey"].(string)
	signers, okSigners := tx["Signers"].([]transaction.Signer)

	if okSigners && len(signers) > 0 {
		for _, signer := range signers {
			if signer.SignerData.SigningPubKey == "" && signer.SignerData.TxnSignature == "" {
				return nil, errors.New("signer data is empty")
			}
		}
	} else if !okTxSig && !okPubKey {
		return nil, errors.New("transaction must have a TxSignature or SigningPubKey set")
	}

	return c.submitRequest(&requests.SubmitRequest{
		TxBlob:   txBlob,
		FailHard: failHard,
	})
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

func (c *Client) formatRequest(req XRPLRequest, id int, marker any) ([]byte, error) {
	m := make(map[string]any)
	m["id"] = id
	m["command"] = req.Method()
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
func (c *Client) calculateFeePerTransactionType(tx *transaction.FlatTransaction) error {
	fee, err := c.getFeeXrp(c.cfg.feeCushion)
	if err != nil {
		return err
	}

	feeDrops, err := currency.XrpToDrops(fee)
	if err != nil {
		return err
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

	(*tx)["LastLedgerSequence"] = uint32(index.Int() + int(LedgerOffset))
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
