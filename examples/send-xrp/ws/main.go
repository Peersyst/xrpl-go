package main

import (
	"fmt"
	"strconv"

	"github.com/Peersyst/xrpl-go/xrpl/common"
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

	fmt.Println("⏳ Connecting to testnet...")
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
		fmt.Println("❌ Failed to connect to testnet")
		return
	}

	fmt.Println("✅ Connected to testnet")
	fmt.Println()

	w, err := wallet.FromSeed(WalletSeed, "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("⏳ Funding wallet...")
	if err := client.FundWallet(&w); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("💸 Wallet funded")
	fmt.Println()

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

	fmt.Println("⏳ Sending 1 XRP to rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe...")
	p := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account: types.Address(w.GetAddress()),
		},
		Destination: "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		Amount:      types.XRPCurrencyAmount(xrpAmountInt),
		DeliverMax:  types.XRPCurrencyAmount(xrpAmountInt),
	}

	flattenedTx := p.Flatten()


	fmt.Println("⏳ Using SubmitTxBlobAndWait, expecting to succeed ...")

	if err := client.Autofill(&flattenedTx); err != nil {
		fmt.Println(err)
		return
	}

	txBlob, _, err := w.Sign(flattenedTx)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.SubmitTxBlobAndWait(txBlob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("✅ Payment submitted")
	fmt.Printf("🌐 Hash: %s\n", res.Hash)
	fmt.Printf("🌐 Validated: %t\n", res.Validated)


	fmt.Println("⏳ Using SubmitTxAndWait with wallet, expecting success ...")
	flattenedTx2 := p.Flatten()
	resp, err := client.SubmitTxAndWait(flattenedTx2, &common.SubmitOptions{
				Autofill: true,
				Wallet:   &w,
			})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("✅ Payment submitted via SubmitTxAndWait")
	fmt.Printf("🌐 Hash: %s\n", resp.Hash)
	fmt.Printf("🌐 Validated: %t\n", resp.Validated)


	fmt.Println("⏳ Using SubmitTxAndWait without wallet, expecting failure ...")
	flattenedTx3 := p.Flatten()
	resp1, err := client.SubmitTxAndWait(flattenedTx3, &common.SubmitOptions{
				Autofill: true,
				Wallet:  nil,
			})
	if err != nil {
		fmt.Printf("❌ Expected error triggered: %v\n", err)
	} else {
		fmt.Println("⚠️ Unexpected success:")
		fmt.Printf("🌐 Hash: %s\n", resp1.Hash)
		fmt.Printf("🌐 Validated: %t\n", resp1.Validated)
	}

}
