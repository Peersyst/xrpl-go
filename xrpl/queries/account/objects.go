package account

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type ObjectType string

const (
	CheckObject          ObjectType = "check"
	DepositPreauthObject ObjectType = "deposit_preauth"
	EscrowObject         ObjectType = "escrow"
	NFTOfferObject       ObjectType = "nft_offer"
	OfferObject          ObjectType = "offer"
	PaymentChannelObject ObjectType = "payment_channel"
	SignerListObject     ObjectType = "signer_list"
	StateObject          ObjectType = "state"
	TicketObject         ObjectType = "ticket"
)

// ############################################################################
// Request
// ############################################################################

type ObjectsRequest struct {
	Account              types.Address          `json:"account"`
	Type                 ObjectType             `json:"type,omitempty"`
	DeletionBlockersOnly bool                   `json:"deletion_blockers_only,omitempty"`
	LedgerHash           common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex          common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Limit                int                    `json:"limit,omitempty"`
	Marker               any                    `json:"marker,omitempty"`
}

func (*ObjectsRequest) Method() string {
	return "account_objects"
}

// TODO: Implement (V2)
func (*ObjectsRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type ObjectsResponse struct {
	Account            types.Address             `json:"account"`
	AccountObjects     []ledger.FlatLedgerObject `json:"account_objects"`
	LedgerHash         common.LedgerHash         `json:"ledger_hash,omitempty"`
	LedgerIndex        common.LedgerIndex        `json:"ledger_index,omitempty"`
	LedgerCurrentIndex common.LedgerIndex        `json:"ledger_current_index,omitempty"`
	Limit              int                       `json:"limit,omitempty"`
	Marker             any                       `json:"marker,omitempty"`
	Validated          bool                      `json:"validated,omitempty"`
}
