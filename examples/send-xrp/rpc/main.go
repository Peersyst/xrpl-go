package main

import (
	"fmt"
	"strconv"

	"github.com/Peersyst/xrpl-go/v1/xrpl/currency"
	"github.com/Peersyst/xrpl-go/v1/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/v1/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/v1/xrpl/wallet"
)

const (
	WalletSeed = "sn3nxiW7v8KXzPzAqzyHXbSSKNuN9"
)

func main() {

	cfg, err := rpc.NewClientConfig(
		"https://s.altnet.rippletest.net:51234/",
		rpc.WithMaxFeeXRP(5.0),
		rpc.WithFeeCushion(1.5),
		rpc.WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}

	client := rpc.NewClient(cfg)

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

	txBlob, _, err := w.Sign(flattenedTx)
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
	fmt.Println("Transaction hash:", res.Hash)
	fmt.Println("Transaction submitted:", res.Validated)
	fmt.Println()
}
