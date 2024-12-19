package main

import (
	"encoding/hex"
	"fmt"

	"github.com/Peersyst/xrpl-go/v1/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/v1/xrpl/rpc"
	transactions "github.com/Peersyst/xrpl-go/v1/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/v1/xrpl/wallet"
)

func main() {

	cfg, err := rpc.NewClientConfig(
		"https://s.altnet.rippletest.net:51234/",
		rpc.WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}

	client := rpc.NewClient(cfg)

	w, err := wallet.FromSeed("sEdSMVV4dJ1JbdBxmakRR4Puu3XVZz2", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	receiverWallet, err := wallet.FromSeed("sEd7d8Ci9nevdLCeUMctF3uGXp9WQqJ", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	payment := transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: types.Address(w.GetAddress()),
			Signers: []transactions.Signer{
				{
					SignerData: transactions.SignerData{
						Account:       types.Address(w.GetAddress()),
						SigningPubKey: w.PublicKey,
						TxnSignature:  "",
					},
				},
				{
					SignerData: transactions.SignerData{
						Account:       types.Address(w.GetAddress()),
						SigningPubKey: w.PublicKey,
						TxnSignature:  "",
					},
				},
			},
			Memos: []types.MemoWrapper{
				{
					Memo: types.Memo{
						MemoData:   hex.EncodeToString([]byte("Hello, World!")),
						MemoFormat: hex.EncodeToString([]byte("text/plain")),
						MemoType:   hex.EncodeToString([]byte("message")),
					},
				},
				{
					Memo: types.Memo{
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

	err = client.Autofill(&tx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction after autofill", tx)
}
