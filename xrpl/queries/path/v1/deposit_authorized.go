package v1

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The deposit_authorized command indicates whether one account is authorized to
// send payments directly to another.
type DepositAuthorizedRequest struct {
	common.BaseRequest
	SourceAccount      types.Address          `json:"source_account"`
	DestinationAccount types.Address          `json:"destination_account"`
	LedgerHash         common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex        common.LedgerSpecifier `json:"ledger_index,omitempty"`
}

func (*DepositAuthorizedRequest) Method() string {
	return "deposit_authorized"
}

func (*DepositAuthorizedRequest) APIVersion() int {
	return version.RippledAPIV1
}

// TODO: Implement V2
func (*DepositAuthorizedRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the deposit_authorized method.
type DepositAuthorizedResponse struct {
	DepositAuthorized  bool               `json:"deposit_authorized"`
	DestinationAccount types.Address      `json:"destination_account"`
	LedgerHash         common.LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerIndex        common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index,omitempty"`
	SourceAccount      types.Address      `json:"source_account"`
	Validated          bool               `json:"validated,omitempty"`
}
