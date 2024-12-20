package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/v1/pkg/crypto"
	"github.com/Peersyst/xrpl-go/v1/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/v1/xrpl/ledger-entry-types"
	transactionquery "github.com/Peersyst/xrpl-go/v1/xrpl/queries/transactions"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/v1/xrpl/wallet"
	"github.com/Peersyst/xrpl-go/v1/xrpl/websocket"
)

func main() {
	fmt.Println("Connecting to testnet...")
	client := websocket.NewClient(
		websocket.NewClientConfig().
			WithHost("wss://s.altnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	defer client.Disconnect()

	if err := client.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	if !client.IsConnected() {
		fmt.Println("Failed to connect to testnet")
		return
	}

	fmt.Println("Connected to testnet")
	fmt.Println()

	w, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	receiverWallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Wallet: ", w.GetAddress())
	fmt.Println("Requesting XRP from faucet...")
	if err := client.FundWallet(&w); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Wallet %s funded", w.GetAddress())
	fmt.Println()

	fmt.Println("Wallet: ", receiverWallet.GetAddress())
	fmt.Println("Requesting XRP from faucet...")
	if err := client.FundWallet(&receiverWallet); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Wallet %s funded", receiverWallet.GetAddress())
	fmt.Println()

	cc := &transaction.CheckCreate{
		BaseTx: transaction.BaseTx{
			Account: w.GetAddress(),
		},
		Destination: receiverWallet.GetAddress(),
		SendMax:     types.XRPCurrencyAmount(1000000),
		InvoiceID:   "46060241FABCF692D4D934BA2A6C4427CD4279083E38C77CBE642243E43BE291",
	}

	flatCc := cc.Flatten()

	if err := client.Autofill(&flatCc); err != nil {
		fmt.Println(err)
		return
	}

	blob, hash, err := w.Sign(flatCc)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.SubmitAndWait(blob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("CheckCreate transaction submitted")
	fmt.Println("Transaction hash: ", res.Hash.String())
	fmt.Println("Validated: ", res.Validated)
	fmt.Println()

	r, err := client.Request(&transactionquery.TxRequest{
		Transaction: hash,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	var txRes transactionquery.TxResponse
	if err := r.GetResult(&txRes); err != nil {
		fmt.Println(err)
		return
	}

	meta, ok := txRes.Meta.(map[string]interface{})
	if !ok {
		fmt.Println(txRes.Meta)
		fmt.Println("Meta is not of type TxObjMeta")
		return
	}

	var checkID string

	affectedNodes := meta["AffectedNodes"].([]interface{})

	for _, node := range affectedNodes {
		affectedNode, ok := node.(map[string]interface{})
		if !ok {
			fmt.Println("Node is not of type map[string]interface{}")
			return
		}

		createdNode, ok := affectedNode["CreatedNode"].(map[string]interface{})
		if !ok {
			continue
		}

		if createdNode["LedgerEntryType"] == string(ledger.CheckEntry) {

			checkID = createdNode["LedgerIndex"].(string)
		}
	}

	if checkID == "" {
		fmt.Println("Check not found")
		return
	}

	fmt.Println("Check created with ID: ", checkID)

	checkCash := &transaction.CheckCash{
		BaseTx: transaction.BaseTx{
			Account: receiverWallet.GetAddress(),
		},
		CheckID: types.Hash256(checkID),
		Amount:  types.XRPCurrencyAmount(1000000),
	}

	flatCheckCash := checkCash.Flatten()

	if err := client.Autofill(&flatCheckCash); err != nil {
		fmt.Println(err)
		return
	}

	blob, _, err = receiverWallet.Sign(flatCheckCash)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err = client.SubmitAndWait(blob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("CheckCash transaction submitted")
	fmt.Println("Transaction hash: ", res.Hash.String())
	fmt.Println("Validated: ", res.Validated)
}
