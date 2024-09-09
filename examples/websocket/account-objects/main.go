package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl"
	public_servers "github.com/Peersyst/xrpl-go/xrpl/client/public_servers"
	"github.com/Peersyst/xrpl-go/xrpl/client/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/account"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

func main() {
	wallet, err := xrpl.NewWalletFromSeed("sEdSMVV4dJ1JbdBxmakRR4Puu3XVZz2", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	// init public urls for websocket client
	publicServers := public_servers.NewServerUrls()

	client := websocket.NewWebsocketClient(
		websocket.NewWebsocketClientConfig().
			WithHost(publicServers.TestnetWebSocket().Ripple()).
			WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)

	accountObjRes, err := client.GetAccountObjects(&account.AccountObjectsRequest{
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
}
