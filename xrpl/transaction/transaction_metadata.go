package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
)

type TxMeta interface {
	TxMeta()
}

// TODO: Improve CurrencyAmount parsing
type TxObjMeta struct {
	AffectedNodes []AffectedNode `json:"AffectedNodes,omitempty"`
	// PartialDeliveredAmount types.CurrencyAmount `json:"DeliveredAmount,omitempty"`
	PartialDeliveredAmount any    `json:"DeliveredAmount,omitempty"`
	TransactionIndex       uint64 `json:"TransactionIndex,omitempty"`
	TransactionResult      string `json:"TransactionResult,omitempty"`
	// DeliveredAmount        types.CurrencyAmount `json:"delivered_amount,omitempty"`
	DeliveredAmount any `json:"delivered_amount,omitempty"`
}

func (TxObjMeta) TxMeta() {}

type AffectedNode struct {
	CreatedNode  *CreatedNode  `json:"CreatedNode,omitempty"`
	ModifiedNode *ModifiedNode `json:"ModifiedNode,omitempty"`
	DeletedNode  *DeletedNode  `json:"DeletedNode,omitempty"`
}

type CreatedNode struct {
	LedgerEntryType ledger.EntryType        `json:"LedgerEntryType,omitempty"`
	LedgerIndex     string                  `json:"LedgerIndex,omitempty"`
	NewFields       ledger.FlatLedgerObject `json:"NewFields,omitempty"`
}

type ModifiedNode struct {
	LedgerEntryType   ledger.EntryType        `json:"LedgerEntryType,omitempty"`
	LedgerIndex       string                  `json:"LedgerIndex,omitempty"`
	FinalFields       ledger.FlatLedgerObject `json:"FinalFields,omitempty"`
	PreviousFields    ledger.FlatLedgerObject `json:"PreviousFields,omitempty"`
	PreviousTxnID     string                  `json:"PreviousTxnID,omitempty"`
	PreviousTxnLgrSeq uint64                  `json:"PreviousTxnLgrSeq,omitempty"`
}

type DeletedNode struct {
	LedgerEntryType ledger.EntryType        `json:"LedgerEntryType,omitempty"`
	LedgerIndex     string                  `json:"LedgerIndex,omitempty"`
	FinalFields     ledger.FlatLedgerObject `json:"FinalFields,omitempty"`
}
