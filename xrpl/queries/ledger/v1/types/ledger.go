package types

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// BaseLedger represents the core fields of a closed or validated ledger,
// including metadata such as hashes, timestamps, transaction lists, and totals.
type BaseLedger struct {
	AccountHash         string                    `json:"account_hash"`
	AccountState        []ledger.FlatLedgerObject `json:"accountState,omitempty"`
	CloseFlags          int                       `json:"close_flags"`
	CloseTime           int                       `json:"close_time"`
	CloseTimeHuman      string                    `json:"close_time_human"`
	CloseTimeResolution int                       `json:"close_time_resolution"`
	Closed              bool                      `json:"closed"`
	LedgerHash          string                    `json:"ledger_hash"`
	LedgerIndex         string                    `json:"ledger_index"`
	ParentCloseTime     int                       `json:"parent_close_time"`
	ParentHash          string                    `json:"parent_hash"`
	TotalCoins          types.XRPCurrencyAmount   `json:"total_coins"`
	TransactionHash     string                    `json:"transaction_hash"`
	Transactions        []interface{}             `json:"transactions,omitempty"`
}

// QueueData contains information about transactions queued in the ledger,
// including retry counts, preflight and last results, and optional fee details.
type QueueData struct {
	Account          types.Address               `json:"account"`
	Tx               transaction.FlatTransaction `json:"tx"`
	RetriesRemaining int                         `json:"retries_remaining"`
	PreflightResult  string                      `json:"preflight_result"`
	LastResult       string                      `json:"last_result,omitempty"`
	AuthChange       bool                        `json:"auth_change,omitempty"`
	Fee              types.XRPCurrencyAmount     `json:"fee,omitempty"`
	FeeLevel         types.XRPCurrencyAmount     `json:"fee_level,omitempty"`
	MaxSpendDrops    types.XRPCurrencyAmount     `json:"max_spend_drops,omitempty"`
}
