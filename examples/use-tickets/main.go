package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

const (
	WalletSeed = "sn3nxiW7v8KXzPzAqzyHXbSSKNuN9"
)

func main() {
	fmt.Println("Connecting to testnet...")
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
		fmt.Println("Failed to connect to testnet")
		return
	}

	fmt.Println("Connected to testnet")
	fmt.Println()

	wallet, err := xrpl.NewWalletFromSeed(WalletSeed, "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Wallet:", wallet.GetAddress())
	fmt.Println()

	fmt.Println("Requesting XRP from faucet...")
	if err := client.FundWallet(&wallet); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("XRP funded")
	fmt.Println()

	info, err := client.GetAccountInfo(&account.InfoRequest{
		Account: wallet.GetAddress(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Current wallet sequence:", info.AccountData.Sequence)
	fmt.Println()

	fmt.Println("Submitting TicketCreate transaction...")
	tc := &transaction.TicketCreate{
		BaseTx: transaction.BaseTx{
			Account:  wallet.GetAddress(),
			Sequence: info.AccountData.Sequence,
		},
		TicketCount: 10,
	}

	flatTc := tc.Flatten()

	if err := client.Autofill(&flatTc); err != nil {
		fmt.Println(err)
		return
	}

	blob, hash, err := wallet.Sign(flatTc)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Submit(blob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("TicketCreate transaction submitted")
	fmt.Println("Transaction hash:", hash)
	fmt.Println("Result:", res.EngineResult)
	fmt.Println()

	time.Sleep(3 * time.Second)

	objects, err := client.GetAccountObjects(&account.ObjectsRequest{
		Account: wallet.GetAddress(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Account objects:", objects.AccountObjects[0]["TicketSequence"])

	seq, err := objects.AccountObjects[0]["TicketSequence"].(json.Number).Int64()
	if err != nil {
		fmt.Println(err)
		return
	}

	as := &transaction.AccountSet{
		BaseTx: transaction.BaseTx{
			Account:        wallet.GetAddress(),
			Sequence:       0,
			TicketSequence: uint32(seq),
		},
	}

	flatAs := as.Flatten()

	if err := client.Autofill(&flatAs); err != nil {
		fmt.Println(err)
		return
	}

	flatAs["Sequence"] = uint32(0)

	blob, hash, err = wallet.Sign(flatAs)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err = client.Submit(blob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("AccountSet transaction submitted")
	fmt.Println("Transaction hash:", hash)
	fmt.Println("Result:", res.EngineResult)
}
