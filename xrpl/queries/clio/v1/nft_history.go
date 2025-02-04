package v1

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
)

// ############################################################################
// Request
// ############################################################################

// The nft_history method retrieves a list of transactions that involved the
// specified NFToken.
type NFTHistoryRequest struct {
	common.BaseRequest
	NFTokenID      string `json:"nft_id"`
	LedgerIndexMin uint   `json:"ledger_index_min,omitempty"`
	LedgerIndexMax uint   `json:"ledger_index_max,omitempty"`
	Binary         bool   `json:"binary,omitempty"`
	Forward        bool   `json:"forward,omitempty"`
	Limit          uint   `json:"limit,omitempty"`
	Marker         any    `json:"marker,omitempty"`
}

func (*NFTHistoryRequest) Method() string {
	return "nft_history"
}

func (*NFTHistoryRequest) APIVersion() int {
	return version.RippledAPIV1
}

// TODO: Implement V2
func (*NFTHistoryRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type NFTHistoryTransactions struct {
	LedgerIndex uint                        `json:"ledger_index"`
	Meta        transaction.TxObjMeta       `json:"meta"`
	Tx          transaction.FlatTransaction `json:"tx,omitempty"`
	TxBlob      string                      `json:"tx_blob,omitempty"`
	Validated   bool                        `json:"validated"`
}

// The expected response from the nft_history method.
type NFTHistoryResponse struct {
	NFTokenID      string                   `json:"nft_id"`
	LedgerIndexMin uint                     `json:"ledger_index_min,omitempty"`
	LedgerIndexMax uint                     `json:"ledger_index_max,omitempty"`
	Limit          uint                     `json:"limit,omitempty"`
	Marker         any                      `json:"marker,omitempty"`
	Transactions   []NFTHistoryTransactions `json:"transactions"`
	Validated      bool                     `json:"validated,omitempty"`
}
