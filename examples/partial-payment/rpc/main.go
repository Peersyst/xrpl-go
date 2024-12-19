package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/v1/pkg/crypto"
	"github.com/Peersyst/xrpl-go/v1/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/v1/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/v1/xrpl/wallet"
)

func main() {
	cfg, err := rpc.NewClientConfig(
		"https://s.altnet.rippletest.net:51234/",
		rpc.WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}

	client := rpc.NewClient(cfg)

	fmt.Println("⏳ Funding wallets...")
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
	if err := client.FundWallet(&w1); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("💸 Wallet 1 funded")
	if err := client.FundWallet(&w2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("💸 Wallet 2 funded")
	fmt.Println()

	fmt.Println("⏳ Sending TrustSet transaction...")
	ts := &transaction.TrustSet{
		BaseTx: transaction.BaseTx{
			Account: w2.ClassicAddress,
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   w1.ClassicAddress,
			Value:    "10000000000",
		},
	}

	flatTs := ts.Flatten()

	err = client.Autofill(&flatTs)
	if err != nil {
		fmt.Println(err)
		return
	}

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

	fmt.Println("✅ TrustSet transaction submitted!")
	fmt.Printf("🌐 Hash: %s\n", res.Hash.String())
	fmt.Printf("🌐 Validated: %t\n", res.Validated)
	fmt.Println()

	fmt.Println("⏳ Issuing tokens for wallet 2...")
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

	fmt.Println("✅ Payment transaction submitted!")
	fmt.Printf("🌐 Hash: %s\n", res.Hash.String())
	fmt.Printf("🌐 Validated: %t\n", res.Validated)
	fmt.Println()

	fmt.Println("⏳ Submitting Partial Payment transaction...")
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

	fmt.Println("✅ Partial Payment transaction submitted!")
	fmt.Printf("🌐 Hash: %s\n", res.Hash.String())
	fmt.Printf("🌐 Validated: %t\n", res.Validated)
	fmt.Println()
}
