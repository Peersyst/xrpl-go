package account

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type Transaction struct {
	LedgerIndex uint64                      `json:"ledger_index"`
	Meta        transaction.TxObjMeta       `json:"meta"`
	Tx          transaction.FlatTransaction `json:"tx"`
	TxBlob      string                      `json:"tx_blob"`
	Validated   bool                        `json:"validated"`
}

// ############################################################################
// Request
// ############################################################################

// The account_tx method retrieves a list of transactions that involved the
// specified account.
type TransactionsRequest struct {
	common.BaseRequest
	Account        types.Address          `json:"account"`
	LedgerIndexMin int                    `json:"ledger_index_min,omitempty"`
	LedgerIndexMax int                    `json:"ledger_index_max,omitempty"`
	LedgerHash     common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex    common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Binary         bool                   `json:"binary,omitempty"`
	Forward        bool                   `json:"forward,omitempty"`
	Limit          int                    `json:"limit,omitempty"`
	Marker         any                    `json:"marker,omitempty"`
}

func (*TransactionsRequest) Method() string {
	return "account_tx"
}

func (*TransactionsRequest) APIVersion() int {
	return version.RippledAPIV2
}

// TODO: Implement (V2)
func (*TransactionsRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the account_tx method.
type TransactionsResponse struct {
	Account        types.Address      `json:"account"`
	LedgerIndexMin common.LedgerIndex `json:"ledger_index_min"`
	LedgerIndexMax common.LedgerIndex `json:"ledger_index_max"`
	Limit          int                `json:"limit"`
	Marker         any                `json:"marker,omitempty"`
	Transactions   []Transaction      `json:"transactions"`
	Validated      bool               `json:"validated"`
}
