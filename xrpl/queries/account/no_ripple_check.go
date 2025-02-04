package account

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The `noripple_check` command provides a quick way to check the status of the
// default ripple field for an account and the No Ripple flag of its trust
// lines, compared with the recommended settings. Expects a response in the form
// of an NoRippleCheckResponse.
type NoRippleCheckRequest struct {
	common.BaseRequest
	Account      types.Address          `json:"account"`
	Role         string                 `json:"role"`
	Transactions bool                   `json:"transactions,omitempty"`
	Limit        int                    `json:"limit,omitempty"`
	LedgerHash   common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex  common.LedgerSpecifier `json:"ledger_index,omitempty"`
}

func (*NoRippleCheckRequest) Method() string {
	return "noripple_check"
}

func (*NoRippleCheckRequest) APIVersion() int {
	return version.RippledAPIV2
}

func (*NoRippleCheckRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// Response expected by a NoRippleCheckRequest
type NoRippleCheckResponse struct {
	LedgerCurrentIndex common.LedgerIndex            `json:"ledger_current_index"`
	Problems           []string                      `json:"problems"`
	Transactions       []transaction.FlatTransaction `json:"transactions"`
}
