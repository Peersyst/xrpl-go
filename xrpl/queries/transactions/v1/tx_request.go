package v1

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The tx method retrieves information on a single transaction, by its
// identifying hash.
type TxRequest struct {
	common.BaseRequest
	Transaction string             `json:"transaction"`
	Binary      bool               `json:"binary,omitempty"`
	MinLedger   common.LedgerIndex `json:"min_ledger,omitempty"`
	MaxLedger   common.LedgerIndex `json:"max_ledger,omitempty"`
}

func (*TxRequest) Method() string {
	return "tx"
}

func (*TxRequest) APIVersion() int {
	return version.RippledAPIV1
}

// TODO: Implement V2
func (*TxRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the tx method.
type TxResponse struct {
	Date        uint               `json:"date"`
	CtID        string             `json:"ct_id,omitempty"`
	Hash        types.Hash256      `json:"hash"`
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
	// TODO: Improve Meta parsing
	Meta      any                         `json:"meta"`
	Validated bool                        `json:"validated"`
	Tx        transaction.FlatTransaction `json:",omitempty"`
}
