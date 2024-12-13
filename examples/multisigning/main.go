package main

import (
	"encoding/hex"
	"fmt"
	"maps"
	"strings"
	"time"

	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
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

	w1, err := xrpl.NewWalletFromSeed("sEdTtvLmJmrb7GaivhWoXRkvU4NDjVf", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Wallet 1:", w1.GetAddress())

	w2, err := xrpl.NewWalletFromSeed("sEdSFiKMQp7RvYLgH7t7FEpwNRWv2Gr", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Wallet 2:", w2.GetAddress())

	master, err := xrpl.NewWalletFromSeed("sEdTMm2yv8c8Rg8YHFHQA9TxVMFy1ze", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Master Wallet:", master.GetAddress())
	fmt.Println()
	fmt.Println("Funding wallets...")

	if err := client.FundWallet(&w1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Wallet 1 funded")

	if err := client.FundWallet(&w2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Wallet 2 funded")

	if err := client.FundWallet(&master); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Master wallet funded")
	fmt.Println()
	fmt.Println("Setting up signer list...")

	ss := &transaction.SignerListSet{
		BaseTx: transaction.BaseTx{
			Account: master.GetAddress(),
		},
		SignerQuorum: 2,
		SignerEntries: []ledger.SignerEntryWrapper{
			{
				SignerEntry: ledger.SignerEntry{
					Account:      w1.GetAddress(),
					SignerWeight: 1,
				},
			},
			{
				SignerEntry: ledger.SignerEntry{
					Account:      w2.GetAddress(),
					SignerWeight: 1,
				},
			},
		},
	}

	flatSs := ss.Flatten()

	if err := client.Autofill(&flatSs); err != nil {
		fmt.Println(err)
		return
	}

	blob, hash, err := master.Sign(flatSs)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Submit(blob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("SignerListSet transaction submitted")
	fmt.Println("Transaction hash:", hash)
	fmt.Println("Transaction result:", res.EngineResult)
	fmt.Println()

	time.Sleep(10 * time.Second)

	fmt.Println("Setting up AccountSet multisign transaction...")

	as := &transaction.AccountSet{
		BaseTx: transaction.BaseTx{
			Account: master.GetAddress(),
		},
		Domain: strings.ToUpper(hex.EncodeToString([]byte("example.com"))),
	}

	flatAs := as.Flatten()

	if err := client.AutofillMultisigned(&flatAs, 2); err != nil {
		fmt.Println(err)
		return
	}

	w1As := maps.Clone(flatAs)

	blob1, _, err := w1.Multisign(w1As)
	if err != nil {
		fmt.Println(err)
		return
	}

	w2As := maps.Clone(flatAs)

	blob2, _, err := w2.Multisign(w2As)
	if err != nil {
		fmt.Println(err)
		return
	}

	blob, err = xrpl.Multisign(blob1, blob2)
	if err != nil {
		fmt.Println(err)
		return
	}

	mRes, err := client.SubmitMultisigned(blob, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Multisigned transaction submitted")
	fmt.Println("Transaction result:", mRes.EngineResult)
}
