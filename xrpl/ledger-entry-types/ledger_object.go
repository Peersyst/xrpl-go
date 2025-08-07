package ledger

import (
	"encoding/json"
	"fmt"
)

// EntryType represents the type of a ledger entry as a string identifier.
type EntryType string

const (
	AccountRootEntry                     EntryType = "AccountRoot"
	AmendmentsEntry                      EntryType = "Amendments"
	AMMEntry                             EntryType = "AMM"
	BridgeEntry                          EntryType = "Bridge"
	CheckEntry                           EntryType = "Check"
	CredentialEntry                      EntryType = "Credential"
	DelegateEntry                        EntryType = "Delegate"
	DepositPreauthObjEntry               EntryType = "DepositPreauth"
	DIDEntry                             EntryType = "DID"
	DirectoryNodeEntry                   EntryType = "DirectoryNode"
	EscrowEntry                          EntryType = "Escrow"
	FeeSettingsEntry                     EntryType = "FeeSettings"
	LedgerHashesEntry                    EntryType = "LedgerHashes"
	NegativeUNLEntry                     EntryType = "NegativeUNL"
	NFTokenOfferEntry                    EntryType = "NFTokenOffer"
	NFTokenPageEntry                     EntryType = "NFTokenPage"
	OfferEntry                           EntryType = "Offer"
	OracleEntry                          EntryType = "Oracle"
	PayChannelEntry                      EntryType = "PayChannel"
	RippleStateEntry                     EntryType = "RippleState"
	SignerListEntry                      EntryType = "SignerList"
	TicketEntry                          EntryType = "Ticket"
	XChainOwnedClaimIDEntry              EntryType = "XChainOwnedClaimID"
	XChainOwnedCreateAccountClaimIDEntry EntryType = "XChainOwnedCreateAccountClaimID"
)

// FlatLedgerObject represents a generic ledger entry as a flat map of field names to values.
type FlatLedgerObject map[string]interface{}

// EntryType returns the LedgerEntryType string stored in this flat object.
func (f FlatLedgerObject) EntryType() EntryType {
	return EntryType(f["LedgerEntryType"].(string))
}

// Object represents a generic ledger entry object with an EntryType method.
type Object interface {
	EntryType() EntryType
}

func EmptyLedgerObject(t string) (Object, error) {
	switch EntryType(t) {
	case AccountRootEntry:
		return &AccountRoot{}, nil
	case AmendmentsEntry:
		return &Amendments{}, nil
	case AMMEntry:
		return &AMM{}, nil
	case BridgeEntry:
		return &Bridge{}, nil
	case CheckEntry:
		return &Check{}, nil
	case CredentialEntry:
		return &Credential{}, nil
	case DelegateEntry:
		return &Delegate{}, nil
	case DepositPreauthObjEntry:
		return &DepositPreauthObj{}, nil
	case DIDEntry:
		return &DID{}, nil
	case DirectoryNodeEntry:
		return &DirectoryNode{}, nil
	case EscrowEntry:
		return &Escrow{}, nil
	case FeeSettingsEntry:
		return &FeeSettings{}, nil
	case LedgerHashesEntry:
		return &Hashes{}, nil
	case NegativeUNLEntry:
		return &NegativeUNL{}, nil
	case NFTokenOfferEntry:
		return &NFTokenOffer{}, nil
	case NFTokenPageEntry:
		return &NFTokenPage{}, nil
	case OfferEntry:
		return &Offer{}, nil
	case OracleEntry:
		return &Oracle{}, nil
	case PayChannelEntry:
		return &PayChannel{}, nil
	case RippleStateEntry:
		return &RippleState{}, nil
	case SignerListEntry:
		return &SignerList{}, nil
	case TicketEntry:
		return &Ticket{}, nil
	case XChainOwnedClaimIDEntry:
		return &XChainOwnedClaimID{}, nil
	case XChainOwnedCreateAccountClaimIDEntry:
		return &XChainOwnedCreateAccountClaimID{}, nil
	}
	return nil, fmt.Errorf("unrecognized LedgerObject type \"%s\"", t)
}

// UnmarshalLedgerObject parses the provided JSON data into the correct ledger entry object based on its LedgerEntryType.
func UnmarshalLedgerObject(data []byte) (Object, error) {
	if len(data) == 0 {
		return nil, nil
	}
	type helper struct {
		LedgerEntryType EntryType
	}
	var h helper
	if err := json.Unmarshal(data, &h); err != nil {
		return nil, err
	}
	var o Object
	switch h.LedgerEntryType {
	case AccountRootEntry:
		o = &AccountRoot{}
	case AmendmentsEntry:
		o = &Amendments{}
	case BridgeEntry:
		o = &Bridge{}
	case CheckEntry:
		o = &Check{}
	case CredentialEntry:
		o = &Credential{}
	case DelegateEntry:
		o = &Delegate{}
	case DepositPreauthObjEntry:
		o = &DepositPreauthObj{}
	case DIDEntry:
		o = &DID{}
	case DirectoryNodeEntry:
		o = &DirectoryNode{}
	case EscrowEntry:
		o = &Escrow{}
	case FeeSettingsEntry:
		o = &FeeSettings{}
	case LedgerHashesEntry:
		o = &Hashes{}
	case NegativeUNLEntry:
		o = &NegativeUNL{}
	case NFTokenOfferEntry:
		o = &NFTokenOffer{}
	case NFTokenPageEntry:
		o = &NFTokenPage{}
	case OfferEntry:
		o = &Offer{}
	case OracleEntry:
		o = &Oracle{}
	case PayChannelEntry:
		o = &PayChannel{}
	case RippleStateEntry:
		o = &RippleState{}
	case SignerListEntry:
		o = &SignerList{}
	case TicketEntry:
		o = &Ticket{}
	case XChainOwnedClaimIDEntry:
		o = &XChainOwnedClaimID{}
	case XChainOwnedCreateAccountClaimIDEntry:
		o = &XChainOwnedCreateAccountClaimID{}
	default:
		return nil, fmt.Errorf("unsupported ledger object of type %s", h.LedgerEntryType)
	}
	if err := json.Unmarshal(data, o); err != nil {
		return nil, err
	}
	return o, nil

}
