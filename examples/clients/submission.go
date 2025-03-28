package clients

import (
	"fmt"

	requests "github.com/Peersyst/xrpl-go/xrpl/queries/transactions"
	transactions "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

type SubmittableTransaction interface {
	TxType() transactions.TxType
	Flatten() transactions.FlatTransaction
}

// Client interface that both RPC and WebSocket clients must implement
type TransactionClient interface {
	Autofill(tx *transactions.FlatTransaction) error
	SubmitAndWait(txBlob string, failHard bool) (*requests.TxResponse, error)
}

// SubmitAndWait submits a transaction and waits for it to be included in a validated ledger
func SubmitAndWait(client TransactionClient, txn SubmittableTransaction, wallet wallet.Wallet) *requests.TxResponse {
	fmt.Println()
	fmt.Printf("⏳ Submitting %s transaction...\n", txn.TxType())

	flattenedTx := txn.Flatten()

	err := client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error autofilling %s transaction: %s\n", txn.TxType(), err)
		fmt.Println()
		return nil
	}

	txBlob, _, err := wallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error signing %s transaction: %s\n", txn.TxType(), err)
		fmt.Println()
		return nil
	}

	response, err := client.SubmitAndWait(txBlob, false)
	if err != nil {
		fmt.Printf("❌ Error submitting %s transaction: %s\n", txn.TxType(), err)
		fmt.Println()
		return nil
	}

	fmt.Printf("✅ %s transaction submitted\n", txn.TxType())
	fmt.Printf("🌐 Hash: %s\n", response.Hash.String())
	fmt.Println()

	return response
}
