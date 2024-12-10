package ledger

import "github.com/Peersyst/xrpl-go/xrpl/queries/common"

// ############################################################################
// Request
// ############################################################################

// The ledger_closed method returns the unique identifiers of the most recently
// closed ledger.
type ClosedRequest struct {
}

func (*ClosedRequest) Method() string {
	return "ledger_closed"
}

// TODO: Implement V2
func (*ClosedRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the ledger_closed method.
type ClosedResponse struct {
	LedgerHash  string             `json:"ledger_hash"`
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
}
