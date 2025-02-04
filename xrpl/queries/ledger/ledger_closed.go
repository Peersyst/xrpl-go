package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
)

// ############################################################################
// Request
// ############################################################################

// The ledger_closed method returns the unique identifiers of the most recently
// closed ledger.
type ClosedRequest struct {
	common.BaseRequest
}

func (*ClosedRequest) Method() string {
	return "ledger_closed"
}

func (*ClosedRequest) APIVersion() int {
	return version.RippledAPIV2
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
