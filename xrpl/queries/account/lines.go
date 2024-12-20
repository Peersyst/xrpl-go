package account

import (
	accounttypes "github.com/Peersyst/xrpl-go/v1/xrpl/queries/account/types"
	"github.com/Peersyst/xrpl-go/v1/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The account_lines method returns information about an account's trust lines,
// including balances in all non-XRP currencies and assets. All information
// retrieved is relative to a particular version of the ledger.
type LinesRequest struct {
	Account     types.Address          `json:"account"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Peer        types.Address          `json:"peer,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
}

func (*LinesRequest) Method() string {
	return "account_lines"
}

// TODO: Implement (V2)
func (*LinesRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the account_lines method.
type LinesResponse struct {
	Account            types.Address            `json:"account"`
	Lines              []accounttypes.TrustLine `json:"lines"`
	LedgerCurrentIndex common.LedgerIndex       `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex       `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash        `json:"ledger_hash,omitempty"`
	Marker             any                      `json:"marker,omitempty"`
}
