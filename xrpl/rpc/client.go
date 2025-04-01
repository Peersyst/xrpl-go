package rpc

import (
	"bytes"
	"context"
	"net/http"
	"time"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/xrpl/common"
	"github.com/Peersyst/xrpl-go/xrpl/hash"
	requests "github.com/Peersyst/xrpl-go/xrpl/queries/transactions"
	rpctypes "github.com/Peersyst/xrpl-go/xrpl/rpc/types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"

	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

type Client struct {
	cfg *Config

	NetworkID uint32
}

func NewClient(cfg *Config) *Client {
	return &Client{
		cfg: cfg,
	}
}

// Request sends a request to the XRPL server and returns the response and any error encountered.
func (c *Client) Request(reqParams XRPLRequest) (XRPLResponse, error) {

	err := reqParams.Validate()
	if err != nil {
		return nil, err
	}

	body, err := createRequest(reqParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.cfg.URL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// add timeout context to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req = req.WithContext(ctx)

	req.Header = c.cfg.Headers

	var response *http.Response

	response, err = c.cfg.HTTPClient.Do(req)
	if err != nil || response == nil {
		return nil, err
	}

	// allow client to reuse persistent connection
	defer response.Body.Close()

	// Check for service unavailable response and retry if so
	if response.StatusCode == 503 {

		maxRetries := 3
		backoffDuration := 1 * time.Second

		for i := 0; i < maxRetries; i++ {
			time.Sleep(backoffDuration)

			// Make request again after waiting
			response, err = c.cfg.HTTPClient.Do(req)
			if err != nil {
				return nil, err
			}

			if response.StatusCode != 503 {
				break
			}

			// Increase backoff duration for the next retry
			backoffDuration *= 2
		}

		if response.StatusCode == 503 {
			// Return service unavailable error here after retry 3 times
			return nil, &ClientError{ErrorString: "Server is overloaded, rate limit exceeded"}
		}

	}

	var jr Response
	jr, err = checkForError(response)
	if err != nil {
		return nil, err
	}

	return &jr, nil
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

// SubmitTx signs the transaction (if necessary) and submits it to the server
// via a submission request. It applies the provided submit options to decide whether
// to autofill missing fields and enforce failHard mode during submission.
func (c *Client) SubmitTx(tx transaction.FlatTransaction, opts *rpctypes.SubmitOptions) (*requests.SubmitResponse, error) {
	txBlob, err := getSignedTx(c, tx, opts.Autofill, opts.Wallet)
	if err != nil {
		return nil, err
	}

	return c.submitRequest(&requests.SubmitRequest{
		TxBlob:   txBlob,
		FailHard: opts.FailHard,
	})
}

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
				return nil, ErrSignerDataIsEmpty
			}
		}
	}

	return c.submitMultisignedRequest(&requests.SubmitMultisignedRequest{
		Tx:       tx,
		FailHard: failHard,
	})
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

// FaucetProvider returns the faucet provider for the client.
func (c *Client) FaucetProvider() common.FaucetProvider {
	return c.cfg.faucetProvider
}

// FundWallet funds a wallet with the client's faucet provider.
func (c *Client) FundWallet(wallet *wallet.Wallet) error {
	if wallet.ClassicAddress == "" {
		return ErrCannotFundWalletWithoutClassicAddress
	}

	err := c.cfg.faucetProvider.FundWallet(wallet.ClassicAddress)
	if err != nil {
		return err
	}

	return nil
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

	lastLedgerSequence, ok := tx["LastLedgerSequence"].(uint32)
	if !ok {

		return nil, ErrMissingLastLedgerSequenceInTransaction

	}

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

// SubmitTxAndWait prepares a transaction by ensuring it is fully signed,
// submits it to the server, and waits for ledger confirmation.
// It validates that the transaction's EngineResult is successful before returning
// the transaction response.
func (c *Client) SubmitTxAndWait(tx transaction.FlatTransaction, opts *rpctypes.SubmitOptions) (*requests.TxResponse, error) {
	// Get the signed transaction blob.
	txBlob, err := getSignedTx(c, tx, opts.Autofill, opts.Wallet)
	if err != nil {
		return nil, err
	}

	// Delegate to SubmitTxBlobAndWait to handle submission, engine result check,
	// ledger sequence validation, and waiting for confirmation.
	return c.SubmitTxBlobAndWait(txBlob, opts.FailHard)
}

// getSignedTx ensures the transaction is fully signed and returns the transaction blob.
// If the transaction is already signed, it encodes and returns it. Otherwise, it autofills (if enabled)
// and signs the transaction using the provided wallet.
func getSignedTx(client *Client, tx transaction.FlatTransaction, autofill bool, wallet *wallet.Wallet) (string, error) {
	// Check if the transaction is already signed: both fields must be non-empty.
	sig, sigOk := tx["TxSignature"].(string)
	pubKey, pubKeyOk := tx["SigningPubKey"].(string)
	if sigOk && sig != "" && pubKeyOk && pubKey != "" {
		blob, err := binarycodec.Encode(tx)
		if err != nil {
			return "", err
		}
		return blob, nil
	}

	// If not signed, ensure a wallet is provided.
	if wallet == nil {
		return "", ErrMissingWallet
	}

	// Optionally autofill the transaction.
	if autofill {
		if err := client.Autofill(&tx); err != nil {
			return "", err
		}
	}

	// Sign the transaction.
	txBlob, _, err := wallet.Sign(tx)
	if err != nil {
		return "", err
	}
	return txBlob, nil
}
