package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
)

// ############################################################################
// Request
// ############################################################################

// The ledger_current method returns the unique identifiers of the current
// in-progress ledger.
type CurrentRequest struct {
	common.BaseRequest
}

func (*CurrentRequest) Method() string {
	return "ledger_current"
}

func (*CurrentRequest) APIVersion() int {
	return version.RippledAPIV2
}

// TODO: Implement V2
func (*CurrentRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the ledger_current method.
type CurrentResponse struct {
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index"`
}
