package account

import (
	"github.com/Peersyst/xrpl-go/v1/xrpl/ledger-entry-types"
	accounttypes "github.com/Peersyst/xrpl-go/v1/xrpl/queries/account/types"
	"github.com/Peersyst/xrpl-go/v1/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The `account_info` command retrieves information about an account, its
// activity, and its XRP balance. All information retrieved is relative to a
// particular version of the ledger.
type InfoRequest struct {
	Account     types.Address          `json:"account"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	Queue       bool                   `json:"queue,omitempty"`
	SignerList  bool                   `json:"signer_list,omitempty"`
	Strict      bool                   `json:"strict,omitempty"`
}

func (*InfoRequest) Method() string {
	return "account_info"
}

// TODO: Implement (V2)
func (*InfoRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the account_info method.
type InfoResponse struct {
	AccountData        ledger.AccountRoot     `json:"account_data"`
	SignerLists        []ledger.SignerList    `json:"signer_lists,omitempty"`
	LedgerCurrentIndex common.LedgerIndex     `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex     `json:"ledger_index,omitempty"`
	QueueData          accounttypes.QueueData `json:"queue_data,omitempty"`
	Validated          bool                   `json:"validated"`
}
