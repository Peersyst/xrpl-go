package main

import (
	"encoding/hex"
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

func main() {

	wsClient := websocket.NewWebsocketClient(
		websocket.NewWebsocketClientConfig().
			WithHost("wss://s.altnet.rippletest.net:51233"),
	)

	wallet, err := xrpl.NewWalletFromSeed("sEdSMVV4dJ1JbdBxmakRR4Puu3XVZz2", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	receiverWallet, err := xrpl.NewWalletFromSeed("sEd7d8Ci9nevdLCeUMctF3uGXp9WQqJ", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	payment := transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: types.Address(wallet.GetAddress()),
			Signers: []transactions.Signer{
				{
					SignerData: transactions.SignerData{
						Account:       types.Address(wallet.GetAddress()),
						SigningPubKey: wallet.PublicKey,
						TxnSignature:  "",
					},
				},
				{
					SignerData: transactions.SignerData{
						Account:       types.Address(wallet.GetAddress()),
						SigningPubKey: wallet.PublicKey,
						TxnSignature:  "",
					},
				},
			},
			Memos: []transactions.MemoWrapper{
				{
					Memo: transactions.Memo{
						MemoData:   hex.EncodeToString([]byte("Hello, World!")),
						MemoFormat: hex.EncodeToString([]byte("text/plain")),
						MemoType:   hex.EncodeToString([]byte("message")),
					},
				},
				{
					Memo: transactions.Memo{
						MemoData:   hex.EncodeToString([]byte("Hello, World 2!")),
						MemoFormat: hex.EncodeToString([]byte("text/plain")),
						MemoType:   hex.EncodeToString([]byte("message 2")),
					},
				},
			},
		},
		Destination: types.Address(receiverWallet.GetAddress()),
		Amount:      types.XRPCurrencyAmount(100000000),
	}

	tx := payment.Flatten()

	fmt.Println("Transaction before autofill", tx)

	err = wsClient.Autofill(&tx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction after autofill", tx)
}
