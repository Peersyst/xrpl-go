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

	regularKeyWallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Wallet 1:", w1.GetAddress())
	fmt.Println("Wallet 2:", w2.GetAddress())
	fmt.Println("Regular key wallet:", regularKeyWallet.GetAddress())

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

	fmt.Println("Requesting XRP from faucet for regular key wallet...")
	if err := client.FundWallet(&regularKeyWallet); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Regular key wallet funded")
	fmt.Println()

	rk := &transaction.SetRegularKey{
		BaseTx: transaction.BaseTx{
			Account: w1.GetAddress(),
		},
		RegularKey: regularKeyWallet.GetAddress(),
	}

	flatRk := rk.Flatten()

	err = client.Autofill(&flatRk)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Set regular key transaction autofill complete")
	fmt.Println()

	fmt.Println("Submitting SetRegularKey transaction...")
	blob, _, err := w1.Sign(flatRk)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.SubmitAndWait(blob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("SetRegularKey transaction submitted")
	fmt.Println("Transaction hash:", res.Hash.String())
	fmt.Println("Validated:", res.Validated)
	fmt.Println()

	fmt.Println("Checking if regular key is set...")
	p := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account: w1.GetAddress(),
		},
		Destination: w2.GetAddress(),
		Amount:      types.XRPCurrencyAmount(10000),
	}

	flatP := p.Flatten()

	err = client.Autofill(&flatP)
	if err != nil {
		fmt.Println(err)
		return
	}

	blob, _, err = regularKeyWallet.Sign(flatP)
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
}
