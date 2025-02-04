package v1

import (
	accounttypes "github.com/Peersyst/xrpl-go/xrpl/queries/account/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The account_offers method retrieves a list of offers made by a given account
// that are outstanding as of a particular ledger version.
type OffersRequest struct {
	common.BaseRequest
	Account     types.Address          `json:"account"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
	Strict      bool                   `json:"strict,omitempty"`
}

func (*OffersRequest) Method() string {
	return "account_offers"
}

func (*OffersRequest) APIVersion() int {
	return version.RippledAPIV1
}

// TODO: Implement (V2)
func (*OffersRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

type OffersResponse struct {
	Account            types.Address              `json:"account"`
	Offers             []accounttypes.OfferResult `json:"offers"`
	LedgerCurrentIndex common.LedgerIndex         `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex         `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash          `json:"ledger_hash,omitempty"`
	Marker             any                        `json:"marker,omitempty"`
}
