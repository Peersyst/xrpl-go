package v1

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
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

// The account_objects command returns the raw ledger format for all objects
// owned by an account. For a higher-level view of an account's trust lines and
// balances, see the account_lines method instead.
type ObjectsRequest struct {
	common.BaseRequest
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

func (*ObjectsRequest) APIVersion() int {
	return version.RippledAPIV1
}

// TODO: Implement (V2)
func (*ObjectsRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the account_objects method.
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
