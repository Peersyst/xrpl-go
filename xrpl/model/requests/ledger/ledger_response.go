package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/ledger"
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type LedgerResponse struct {
	Ledger      LedgerHeader       `json:"ledger"`
	LedgerHash  string             `json:"ledger_hash"`
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
	Validated   bool               `json:"validated,omitempty"`
	QueueData   []LedgerQueueData  `json:"queue_data,omitempty"`
}

type LedgerHeader struct {
	AccountHash         string                         `json:"account_hash"`
	AccountState        []ledger.FlatLedgerObject      `json:"accountState,omitempty"`
	CloseFlags          int                            `json:"close_flags"`
	CloseTime           int                            `json:"close_time"`
	CloseTimeHuman      string                         `json:"close_time_human"`
	CloseTimeResolution int                            `json:"close_time_resolution"`
	Closed              bool                           `json:"closed"`
	LedgerHash          string                         `json:"ledger_hash"`
	LedgerIndex         string                         `json:"ledger_index"`
	ParentCloseTime     int                            `json:"parent_close_time"`
	ParentHash          string                         `json:"parent_hash"`
	TotalCoins          types.XRPCurrencyAmount        `json:"total_coins"`
	TransactionHash     string                         `json:"transaction_hash"`
	Transactions        []transactions.FlatTransaction `json:"transactions,omitempty"`
}

type LedgerQueueData struct {
	Account          types.Address                `json:"account"`
	Tx               transactions.FlatTransaction `json:"tx"`
	RetriesRemaining int                          `json:"retries_remaining"`
	PreflightResult  string                       `json:"preflight_result"`
	LastResult       string                       `json:"last_result,omitempty"`
	AuthChange       bool                         `json:"auth_change,omitempty"`
	Fee              types.XRPCurrencyAmount      `json:"fee,omitempty"`
	FeeLevel         types.XRPCurrencyAmount      `json:"fee_level,omitempty"`
	MaxSpendDrops    types.XRPCurrencyAmount      `json:"max_spend_drops,omitempty"`
}