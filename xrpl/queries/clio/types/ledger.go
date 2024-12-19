package types

import (
	"github.com/Peersyst/xrpl-go/v1/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/v1/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

type Ledger struct {
	AccountHash         string                        `json:"account_hash"`
	AccountState        []ledger.FlatLedgerObject     `json:"accountState,omitempty"`
	CloseFlags          int                           `json:"close_flags"`
	CloseTime           uint                          `json:"close_time"`
	CloseTimeHuman      string                        `json:"close_time_human"`
	CloseTimeResolution int                           `json:"close_time_resolution"`
	Closed              bool                          `json:"closed"`
	LedgerHash          common.LedgerHash             `json:"ledger_hash"`
	LedgerIndex         string                        `json:"ledger_index"`
	ParentCloseTime     uint                          `json:"parent_close_time"`
	ParentHash          string                        `json:"parent_hash"`
	TotalCoins          types.XRPCurrencyAmount       `json:"total_coins"`
	TransactionHash     string                        `json:"transaction_hash"`
	Transactions        []transaction.FlatTransaction `json:"transactions,omitempty"`
}
