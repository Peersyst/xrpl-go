package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/v1/pkg/crypto"
	"github.com/Peersyst/xrpl-go/v1/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/v1/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/v1/xrpl/wallet"
)

func main() {

	fmt.Println("Funding wallet on testnet:")

	cfg, err := rpc.NewClientConfig(
		"https://s.altnet.rippletest.net:51234/",
		rpc.WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}

	client := rpc.NewClient(cfg)

	wallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	balance, err := client.GetXrpBalance(wallet.ClassicAddress)
	if err != nil {
		balance = "0"
	}

	fmt.Println("Balance", wallet.ClassicAddress, balance)

	fmt.Println("Funding wallet", wallet.ClassicAddress)
	err = client.FundWallet(&wallet)
	if err != nil {
		fmt.Println(err)
		return
	}

	balance, err = client.GetXrpBalance(wallet.ClassicAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Balance", wallet.ClassicAddress, balance)

	fmt.Println("Funding wallet on devnet:")

	balance, err = client.GetXrpBalance(wallet.ClassicAddress)
	if err != nil {
		balance = "0"
	}

	fmt.Println("Balance", wallet.ClassicAddress, balance)

	fmt.Println("Funding wallet", wallet.ClassicAddress)
	err = client.FundWallet(&wallet)
	if err != nil {
		fmt.Println(err)
		return
	}

	balance, err = client.GetXrpBalance(wallet.ClassicAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Balance", wallet.ClassicAddress, balance)
}
