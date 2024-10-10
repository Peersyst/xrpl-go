package ledger

import (
	"encoding/json"
	"fmt"
)

type LedgerEntryType string

const (
	AccountRootEntry       LedgerEntryType = "AccountRoot"
	AmendmentsEntry        LedgerEntryType = "Amendments"
	AMMEntry               LedgerEntryType = "AMM"
	CheckEntry             LedgerEntryType = "Check"
	DepositPreauthObjEntry LedgerEntryType = "DepositPreauth"
	DirectoryNodeEntry     LedgerEntryType = "DirectoryNode"
	EscrowEntry            LedgerEntryType = "Escrow"
	FeeSettingsEntry       LedgerEntryType = "FeeSettings"
	LedgerHashesEntry      LedgerEntryType = "LedgerHashes"
	NegativeUNLEntry       LedgerEntryType = "NegativeUNL"
	NFTokenOfferEntry      LedgerEntryType = "NFTokenOffer"
	NFTokenPageEntry       LedgerEntryType = "NFTokenPage"
	OfferEntry             LedgerEntryType = "Offer"
	PayChannelEntry        LedgerEntryType = "PayChannel"
	RippleStateEntry       LedgerEntryType = "RippleState"
	SignerListEntry        LedgerEntryType = "SignerList"
	TicketEntry            LedgerEntryType = "Ticket"
)

type FlatLedgerObject map[string]interface{}

func (f FlatLedgerObject) EntryType() LedgerEntryType {
	return LedgerEntryType(f["LedgerEntryType"].(string))
}

type LedgerObject interface {
	EntryType() LedgerEntryType
}

func EmptyLedgerObject(t string) (LedgerObject, error) {
	switch LedgerEntryType(t) {
	case AccountRootEntry:
		return &AccountRoot{}, nil
	case AmendmentsEntry:
		return &Amendments{}, nil
	case AMMEntry:
		return &AMM{}, nil
	case CheckEntry:
		return &Check{}, nil
	case DepositPreauthObjEntry:
		return &DepositPreauthObj{}, nil
	case DirectoryNodeEntry:
		return &DirectoryNode{}, nil
	case EscrowEntry:
		return &Escrow{}, nil
	case FeeSettingsEntry:
		return &FeeSettings{}, nil
	case LedgerHashesEntry:
		return &LedgerHashes{}, nil
	case NegativeUNLEntry:
		return &NegativeUNL{}, nil
	case NFTokenOfferEntry:
		return &NFTokenOffer{}, nil
	case NFTokenPageEntry:
		return &NFTokenPage{}, nil
	case OfferEntry:
		return &Offer{}, nil
	case PayChannelEntry:
		return &PayChannel{}, nil
	case RippleStateEntry:
		return &RippleState{}, nil
	case SignerListEntry:
		return &SignerList{}, nil
	case TicketEntry:
		return &Ticket{}, nil
	}
	return nil, fmt.Errorf("unrecognized LedgerObject type \"%s\"", t)
}

func UnmarshalLedgerObject(data []byte) (LedgerObject, error) {
	if len(data) == 0 {
		return nil, nil
	}
	type helper struct {
		LedgerEntryType LedgerEntryType
	}
	var h helper
	if err := json.Unmarshal(data, &h); err != nil {
		return nil, err
	}
	var o LedgerObject
	switch h.LedgerEntryType {
	case AccountRootEntry:
		o = &AccountRoot{}
	case AmendmentsEntry:
		o = &Amendments{}
	case CheckEntry:
		o = &Check{}
	case DepositPreauthObjEntry:
		o = &DepositPreauthObj{}
	case DirectoryNodeEntry:
		o = &DirectoryNode{}
	case EscrowEntry:
		o = &Escrow{}
	case FeeSettingsEntry:
		o = &FeeSettings{}
	case LedgerHashesEntry:
		o = &LedgerHashes{}
	case NegativeUNLEntry:
		o = &NegativeUNL{}
	case NFTokenOfferEntry:
		o = &NFTokenOffer{}
	case NFTokenPageEntry:
		o = &NFTokenPage{}
	case OfferEntry:
		o = &Offer{}
	case PayChannelEntry:
		o = &PayChannel{}
	case RippleStateEntry:
		o = &RippleState{}
	case SignerListEntry:
		o = &SignerList{}
	case TicketEntry:
		o = &Ticket{}
	default:
		return nil, fmt.Errorf("unsupported ledger object of type %s", h.LedgerEntryType)
	}
	if err := json.Unmarshal(data, o); err != nil {
		return nil, err
	}
	return o, nil

}
