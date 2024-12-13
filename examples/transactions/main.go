package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	transactions "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

func main() {

	wallet, err := xrpl.NewWalletFromSeed("sEdSMVV4dJ1JbdBxmakRR4Puu3XVZz2", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = xrpl.NewWalletFromSeed("sEd7d8Ci9nevdLCeUMctF3uGXp9WQqJ", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = websocket.NewClient(
		websocket.NewClientConfig().
			WithHost("wss://s.altnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)

	payment := transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: "r9cZA1mLK5R5AmHZiRd6CCe83ACaut34Mf",
		},
		Amount: types.IssuedCurrencyAmount{
			Currency: "USD",
			Issuer:   "r9cZA1mLK5R5AmHZiRd6CCe83ACaut34Mf",
			Value:    "100",
		},
		Destination:    "r9cZA1mLK5R5AmHZiRd6CCe83ACaut34Mf",
		DestinationTag: 100,
	}

	fmt.Println(payment)
	fmt.Println(payment.Flatten())

	txBlob, hash, err := wallet.Sign(payment.Flatten())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(txBlob)
	fmt.Println(hash)
}
