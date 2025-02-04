package server

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	servertypes "github.com/Peersyst/xrpl-go/xrpl/queries/server/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
)

// ############################################################################
// Request
// ############################################################################

// The fee command reports the current state of the open-ledger requirements
// for the transaction cost. This requires the FeeEscalation amendment to be
// enabled.
type FeeRequest struct {
	common.BaseRequest
}

func (*FeeRequest) Method() string {
	return "fee"
}

func (*FeeRequest) APIVersion() int {
	return version.RippledAPIV2
}

// TODO: Implement V2
func (*FeeRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the fee method.
type FeeResponse struct {
	CurrentLedgerSize  string                `json:"current_ledger_size"`
	CurrentQueueSize   string                `json:"current_queue_size"`
	Drops              servertypes.FeeDrops  `json:"drops"`
	ExpectedLedgerSize string                `json:"expected_ledger_size"`
	LedgerCurrentIndex common.LedgerIndex    `json:"ledger_current_index"`
	Levels             servertypes.FeeLevels `json:"levels"`
	MaxQueueSize       string                `json:"max_queue_size"`
}
