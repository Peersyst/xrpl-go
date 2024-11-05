package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type Escrow struct {
	Account           types.Address
	Amount            types.XRPCurrencyAmount
	CancelAfter       uint   `json:",omitempty"`
	Condition         string `json:",omitempty"`
	Destination       types.Address
	DestinationNode   string `json:",omitempty"`
	DestinationTag    uint   `json:",omitempty"`
	FinishAfter       uint   `json:",omitempty"`
	Flags             uint
	LedgerEntryType   EntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint
	SourceTag         uint `json:",omitempty"`
}

func (*Escrow) EntryType() EntryType {
	return EscrowEntry
}
