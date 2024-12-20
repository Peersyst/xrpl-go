package stream

import "github.com/Peersyst/xrpl-go/v1/xrpl/queries/common"

type LedgerStream struct {
	Type             Type               `json:"type"`
	FeeBase          int                `json:"fee_base"`
	FeeRef           int                `json:"fee_ref"`
	LedgerHash       common.LedgerHash  `json:"ledger_hash"`
	LedgerIndex      common.LedgerIndex `json:"ledger_index"`
	LedgerTime       uint64             `json:"ledger_time"`
	ReserveBase      uint               `json:"reserve_base"`
	ReserveInc       uint               `json:"reserve_inc"`
	TxnCount         int                `json:"txn_count"`
	ValidatedLedgers string             `json:"validated_ledgers,omitempty"`
}
