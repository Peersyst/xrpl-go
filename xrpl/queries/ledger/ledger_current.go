package ledger

import "github.com/Peersyst/xrpl-go/xrpl/queries/common"

// ############################################################################
// Request
// ############################################################################

// The ledger_current method returns the unique identifiers of the current
// in-progress ledger.
type CurrentRequest struct {
}

func (*CurrentRequest) Method() string {
	return "ledger_current"
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
