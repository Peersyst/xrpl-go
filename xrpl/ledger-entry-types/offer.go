package ledger

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type OfferFlags uint

const (
	PassiveOffer OfferFlags = 0x00010000
	SellOffer    OfferFlags = 0x00020000
)

type Offer struct {
	Account           types.Address
	BookDirectory     types.Hash256
	BookNode          string
	Expiration        uint `json:",omitempty"`
	Flags             OfferFlags
	LedgerEntryType   EntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint
	Sequence          uint
	TakerPays         types.CurrencyAmount
	TakerGets         types.CurrencyAmount
}

func (*Offer) EntryType() EntryType {
	return OfferEntry
}

func (o *Offer) UnmarshalJSON(data []byte) error {
	type offerHelper struct {
		Account           types.Address
		BookDirectory     types.Hash256
		BookNode          string
		Expiration        uint
		Flags             OfferFlags
		LedgerEntryType   EntryType
		OwnerNode         string
		PreviousTxnID     types.Hash256
		PreviousTxnLgrSeq uint
		Sequence          uint
		TakerPays         json.RawMessage
		TakerGets         json.RawMessage
	}
	var h offerHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*o = Offer{
		Account:           h.Account,
		BookDirectory:     h.BookDirectory,
		BookNode:          h.BookNode,
		Expiration:        h.Expiration,
		Flags:             h.Flags,
		LedgerEntryType:   h.LedgerEntryType,
		OwnerNode:         h.OwnerNode,
		PreviousTxnID:     h.PreviousTxnID,
		PreviousTxnLgrSeq: h.PreviousTxnLgrSeq,
		Sequence:          h.Sequence,
	}
	pays, err := types.UnmarshalCurrencyAmount(h.TakerPays)
	if err != nil {
		return err
	}
	gets, err := types.UnmarshalCurrencyAmount(h.TakerGets)
	if err != nil {
		return err
	}
	o.TakerPays = pays
	o.TakerGets = gets
	return nil
}
