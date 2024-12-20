package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/v1/pkg/crypto"
	"github.com/Peersyst/xrpl-go/v1/xrpl/faucet"
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

	w1, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	w2, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Wallet 1:", w1.GetAddress())
	fmt.Println("Wallet 2:", w2.GetAddress())

	fmt.Println()
	fmt.Println("Requesting XRP from faucet for wallet 1...")
	if err := client.FundWallet(&w1); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Wallet 1 funded")
	fmt.Println()

	fmt.Println("Requesting XRP from faucet for wallet 2...")
	if err := client.FundWallet(&w2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Wallet 2 funded")
	fmt.Println()

	ts := &transaction.TrustSet{
		BaseTx: transaction.BaseTx{
			Account: w2.GetAddress(),
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   w1.GetAddress(),
			Value:    "10000000000",
		},
	}

	flatTs := ts.Flatten()

	err = client.Autofill(&flatTs)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("TrustSet transaction autofill complete")
	fmt.Println()

	fmt.Println("Submitting TrustSet transaction...")
	blob, _, err := w2.Sign(flatTs)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.SubmitAndWait(blob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("TrustSet transaction submitted")
	fmt.Println("Transaction hash:", res.Hash.String())
	fmt.Println("Validated:", res.Validated)
	fmt.Println()

	fmt.Println("Issuing tokens for wallet 2...")
	p := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account: w1.GetAddress(),
		},
		Amount: types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   w1.GetAddress(),
			Value:    "50",
		},
		Destination: w2.GetAddress(),
	}

	flatP := p.Flatten()

	err = client.Autofill(&flatP)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Payment transaction autofill complete")
	fmt.Println()

	fmt.Println("Submitting Payment transaction...")
	blob, _, err = w1.Sign(flatP)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err = client.SubmitAndWait(blob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Payment transaction submitted")
	fmt.Println("Transaction hash:", res.Hash.String())
	fmt.Println("Validated:", res.Validated)
	fmt.Println()

	pp := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account: w2.GetAddress(),
		},
		Amount: types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   w1.GetAddress(),
			Value:    "10",
		},
		Destination: w1.GetAddress(),
	}

	pp.SetPartialPaymentFlag()

	flatPP := pp.Flatten()

	err = client.Autofill(&flatPP)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Partial Payment transaction autofill complete")
	fmt.Println()

	fmt.Println("Submitting Partial Payment transaction...")
	blob, _, err = w2.Sign(flatPP)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err = client.SubmitAndWait(blob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Partial Payment transaction submitted")
	fmt.Println("Transaction hash:", res.Hash.String())
	fmt.Println("Validated:", res.Validated)
	fmt.Println()
}
