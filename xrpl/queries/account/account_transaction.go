package account

import (
	transactions "github.com/Peersyst/xrpl-go/xrpl/transaction"
)

const (
	ErrAccountTxUnmarshal string = "Unmarshal JSON AccountTransaction"
)

type Transaction struct {
	LedgerIndex uint64                       `json:"ledger_index"`
	Meta        transactions.TxMeta          `json:"meta"`
	Tx          transactions.FlatTransaction `json:"tx"`
	TxBlob      string                       `json:"tx_blob"`
	Validated   bool                         `json:"validated"`
}
