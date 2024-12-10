package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

func main() {
	wallet, err := xrpl.NewWalletFromSeed("sEdSMVV4dJ1JbdBxmakRR4Puu3XVZz2", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := websocket.NewClient(
		websocket.NewClientConfig().
			WithHost("wss://s.altnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	defer client.Disconnect()

	fmt.Println("Connecting to server...")
	if err := client.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connection: ", client.IsConnected())

	accountInfo, err := client.GetAccountInfo(&account.InfoRequest{
		Account:    types.Address(wallet.GetAddress()),
		SignerList: true,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Account address: ", accountInfo.AccountData.Account)
	fmt.Println("Account sequence: ", accountInfo.AccountData.Sequence)
	fmt.Println("Account balance: ", accountInfo.AccountData.Balance)
}
