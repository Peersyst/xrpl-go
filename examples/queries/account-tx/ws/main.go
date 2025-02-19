package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
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

	txs, err := client.GetAccountTransactions(&account.TransactionsRequest{
		Account: "rMCcNuTcajgw7YTgBy1sys3b89QqjUrMpH",
		LedgerIndex: common.LedgerIndex(4976692) ,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Number of transactions:", len(txs.Transactions))
	fmt.Println(txs.Transactions[0].Tx)
}