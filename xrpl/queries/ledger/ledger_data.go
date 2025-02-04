package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	ledgertypes "github.com/Peersyst/xrpl-go/xrpl/queries/ledger/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
)

// ############################################################################
// Request
// ############################################################################

// The ledger_data method retrieves contents of the specified ledger. You can
// iterate through several calls to retrieve the entire contents of a single
// ledger version.
type DataRequest struct {
	common.BaseRequest
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Binary      bool                   `json:"binary,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
	Type        ledger.EntryType       `json:"type,omitempty"`
}

func (*DataRequest) Method() string {
	return "ledger_data"
}

func (*DataRequest) APIVersion() int {
	return version.RippledAPIV2
}

// TODO: Implement V2
func (*DataRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the ledger_data method.
type DataResponse struct {
	LedgerIndex string              `json:"ledger_index"`
	LedgerHash  common.LedgerHash   `json:"ledger_hash"`
	State       []ledgertypes.State `json:"state"`
	Marker      any                 `json:"marker"`
}
