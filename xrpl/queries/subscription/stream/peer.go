package stream

import "github.com/Peersyst/xrpl-go/v1/xrpl/queries/common"

type PeerStatusStream struct {
	Type           Type               `json:"type"`
	Action         string             `json:"action"`
	Date           uint64             `json:"date"`
	LedgerHash     common.LedgerHash  `json:"ledger_hash"`
	LedgerIndex    common.LedgerIndex `json:"ledger_index"`
	LedgerIndexMax common.LedgerIndex `json:"ledger_index_max,omitempty"`
	LedgerIndexMin common.LedgerIndex `json:"ledger_index_min,omitempty"`
}
