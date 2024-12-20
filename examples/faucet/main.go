package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

func main() {

	fmt.Println("Funding wallet on testnet:")

	testnetFaucet := faucet.NewTestnetFaucetProvider()
	testnetClientCfg := websocket.NewClientConfig().
		WithHost("wss://s.altnet.rippletest.net:51233").
		WithFaucetProvider(testnetFaucet)
	testnetClient := websocket.NewClient(testnetClientCfg)

	wallet, err := xrpl.NewWallet(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	balance, err := testnetClient.GetXrpBalance(wallet.ClassicAddress)
	if err != nil {
		balance = "0"
	}

	fmt.Println("Balance", wallet.ClassicAddress, balance)

	fmt.Println("Funding wallet", wallet.ClassicAddress)
	err = testnetClient.FundWallet(&wallet)
	if err != nil {
		fmt.Println(err)
		return
	}

	balance, err = testnetClient.GetXrpBalance(wallet.ClassicAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Balance", wallet.ClassicAddress, balance)

	fmt.Println("Funding wallet on devnet:")

	devnetFaucet := faucet.NewDevnetFaucetProvider()

	clientCfg := websocket.NewClientConfig().
		WithHost("wss://s.devnet.rippletest.net:51233").
		WithFaucetProvider(devnetFaucet)

	devnetClient := websocket.NewClient(clientCfg)

	balance, err = devnetClient.GetXrpBalance(wallet.ClassicAddress)
	if err != nil {
		balance = "0"
	}

	fmt.Println("Balance", wallet.ClassicAddress, balance)

	fmt.Println("Funding wallet", wallet.ClassicAddress)
	err = devnetClient.FundWallet(&wallet)
	if err != nil {
		fmt.Println(err)
		return
	}

	balance, err = devnetClient.GetXrpBalance(wallet.ClassicAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Balance", wallet.ClassicAddress, balance)
}
