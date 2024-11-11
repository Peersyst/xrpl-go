package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type PayChannel struct {
	Account           types.Address
	Amount            types.XRPCurrencyAmount
	Balance           types.XRPCurrencyAmount
	CancelAfter       uint `json:",omitempty"`
	Destination       types.Address
	DestinationTag    uint   `json:",omitempty"`
	DestinationNode   string `json:",omitempty"`
	Expiration        uint   `json:",omitempty"`
	Flags             uint
	LedgerEntryType   EntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint
	PublicKey         string
	SettleDelay       uint
	SourceTag         uint `json:",omitempty"`
}

func (*PayChannel) EntryType() EntryType {
	return PayChannelEntry
}
