package types

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type QueueData struct {
	TxnCount           uint64                  `json:"txn_count"`
	AuthChangeQueued   bool                    `json:"auth_change_queued,omitempty"`
	LowestSequence     uint64                  `json:"lowest_sequence,omitempty"`
	HighestSequence    uint64                  `json:"highest_sequence,omitempty"`
	MaxSpendDropsTotal types.XRPCurrencyAmount `json:"max_spend_drops_total,omitempty"`
	Transactions       []QueueTransaction      `json:"transactions,omitempty"`
}
