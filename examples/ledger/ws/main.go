package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	ledgerqueries "github.com/Peersyst/xrpl-go/xrpl/queries/ledger"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

func main() {
	client := websocket.NewClient(
		websocket.NewClientConfig().
			WithHost("wss://s.altnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	defer client.Disconnect()

	fmt.Println("⏳ Connecting to server...")
	if err := client.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("✅ Connected to server")
	fmt.Println()

	ledger, err := client.GetLedger(&ledgerqueries.Request{
		LedgerIndex: common.LedgerIndex(5115183),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ledger.Ledger.LedgerHash)
	fmt.Println(ledger.Ledger.LedgerIndex)
}
