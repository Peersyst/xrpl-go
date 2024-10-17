package main

import (
	"fmt"
	"log"

	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/rpc"
)

func main() {
	fmt.Println("Starting JSON RPC client")
	// init new config object with desired node address
	cfg, err := rpc.NewJsonRpcConfig("https://s.altnet.rippletest.net:51234/")
	if err != nil {
		log.Panicln(err)
	}

	// Initialise new json client with json config
	client := rpc.NewJsonRpcClient(cfg)

	// call the desired method
	req := &account.AccountInfoRequest{
		Account: "rPUK1iYbtS6LP9sA2jbUDHtTnbnQqLBnac",
	}

	fmt.Println("Sending GetAccountInfo request")
	ac, err := client.GetAccountInfo(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("GetAccountInfo response: %v\n", ac)
}
