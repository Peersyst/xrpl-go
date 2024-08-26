package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/client/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/model/client/account"
)

func main() {
	client := websocket.NewClient(&websocket.WebsocketConfig{URL: "wss://s.altnet.rippletest.net"})

	acr, _, err := client.Account.GetAccountInfo(&account.AccountInfoRequest{Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59"})
	if err != nil {
		panic(err)
	}
	fmt.Println(acr)
	// Do stuff

}