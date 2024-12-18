package main

import (
	"fmt"
	"strconv"

	"github.com/Peersyst/xrpl-go/xrpl/currency"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

const (
	WalletSeed = "sn3nxiW7v8KXzPzAqzyHXbSSKNuN9"
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

	w, err := wallet.FromSeed(WalletSeed, "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Requesting XRP from faucet...")
	if err := client.FundWallet(&w); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Wallet %s funded", w.GetAddress())
	fmt.Println()

	fmt.Println("Sending 1 XRP to rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe")
	xrpAmount, err := currency.XrpToDrops("1")
	if err != nil {
		fmt.Println(err)
		return
	}

	xrpAmountInt, err := strconv.ParseInt(xrpAmount, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account: types.Address(w.GetAddress()),
		},
		Destination: "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		Amount:      types.XRPCurrencyAmount(xrpAmountInt),
		DeliverMax:  types.XRPCurrencyAmount(xrpAmountInt),
	}

	flattenedTx := p.Flatten()

	fmt.Println("Autofilling transaction...")
	if err := client.Autofill(&flattenedTx); err != nil {
		fmt.Println(err)
		return
	}

	txBlob, hash, err := w.Sign(flattenedTx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Submitting transaction...")
	res, err := client.SubmitAndWait(txBlob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("Transaction hash:", hash)
	fmt.Printf("Transaction submitted: %d", res.LedgerIndex.Int())
	fmt.Println()
}
