package main

import (
	"encoding/hex"
	"fmt"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl"
	transactions "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

const (
	AccountSeed        = "sEd7MgLAff94dLx91rVRByUbLrdSrdj"
	DestinationAddress = "rDwvihpE48E48F8rvNrqTb2UGWv62xqYTg"
	Currency           = "USD"
	Value              = "100"
	Issuer             = "rDwvihpE48E48F8rvNrqTb2UGWv62xqYTg"
)

func main() {
	mnemonicWallet, err := xrpl.NewWalletFromMnemonic("monster march exile fee forget response seven push dragon oil clinic attack black miss craft surface patient stomach tank float cabbage visual image resource")
	if err != nil {
		panic(err)
	}

	fmt.Println("Wallet generated from mnemonic")

	fmt.Printf("Private key: %s\n", mnemonicWallet.PrivateKey)
	fmt.Printf("Public 	key: %s\n", mnemonicWallet.PublicKey)
	fmt.Printf("Classic address: %s\n", mnemonicWallet.ClassicAddress)
	fmt.Printf("Seed: %s\n", mnemonicWallet.Seed)

	wallet, err := xrpl.NewWallet(addresscodec.ED25519)
	if err != nil {
		panic(err)
	}
	fmt.Println("Wallet generated from random seed")

	fmt.Printf("Private key: %s\n", wallet.PrivateKey)
	fmt.Printf("Public 	key: %s\n", wallet.PublicKey)
	fmt.Printf("Classic address: %s\n", wallet.ClassicAddress)
	fmt.Printf("Seed: %s\n", wallet.Seed)

	walletFromSeed, _ := xrpl.NewWalletFromSeed(wallet.Seed, "")

	fmt.Println("\nWallet generated from previous seed")

	fmt.Printf("Private key: %s\n", walletFromSeed.PrivateKey)
	fmt.Printf("Public 	key: %s\n", walletFromSeed.PublicKey)
	fmt.Printf("Classic address: %s\n", walletFromSeed.ClassicAddress)
	fmt.Printf("Seed: %s\n", walletFromSeed.Seed)

	walletFromSecret, _ := xrpl.NewWalletFromSecret(wallet.Seed)

	fmt.Println("\nWallet generated from previous seed")

	fmt.Printf("Private key: %s\n", walletFromSecret.PrivateKey)
	fmt.Printf("Public 	key: %s\n", walletFromSecret.PublicKey)
	fmt.Printf("Classic address: %s\n", walletFromSecret.ClassicAddress)
	fmt.Printf("Seed: %s\n", walletFromSecret.Seed)

	tx := transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: types.Address(wallet.ClassicAddress),
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
						MemoType:   hex.EncodeToString([]byte("message2")),
					},
				},
			},
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   Issuer,
			Currency: Currency,
			Value:    Value,
		},
		Destination: types.Address(DestinationAddress),
	}

	fmt.Println(tx.Flatten())

	fmt.Println("\nSigning a transaction with wallet generated from seed")

	txBlob, hash, err := wallet.Sign(tx.Flatten())
	if err != nil {
		panic(err)
	}

	fmt.Printf("txBlob: %s\n", txBlob)
	fmt.Printf("hash: %s\n", hash)

	fmt.Println("\nSigning a transaction with wallet generated from mnemonic")

	mnemonicTx := transactions.Payment{
		BaseTx: transactions.BaseTx{
			TransactionType: "Payment",
			Account:         types.Address(mnemonicWallet.ClassicAddress),
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   Issuer,
			Currency: Currency,
			Value:    Value,
		},
		Destination: types.Address(DestinationAddress),
	}

	mnemonicTxBlob, mnemonicHash, err := mnemonicWallet.Sign(mnemonicTx.Flatten())
	if err != nil {
		panic(err)
	}

	fmt.Printf("txBlob: %s\n", mnemonicTxBlob)
	fmt.Printf("hash: %s\n", mnemonicHash)
}
