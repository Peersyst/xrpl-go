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

	accountObjRes, err := client.GetAccountObjects(&account.ObjectsRequest{
		Account: types.Address(wallet.GetAddress()),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Account address: ", wallet.GetAddress())
	for _, obj := range accountObjRes.AccountObjects {
		fmt.Println("Object: ", obj)
	}

	accountLinesRes, err := client.GetAccountLines(&account.LinesRequest{
		Account: types.Address(wallet.GetAddress()),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Account lines: ", accountLinesRes)
}
