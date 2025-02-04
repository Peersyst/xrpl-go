package v1

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The `account_currencies` command retrieves a list of currencies that an
// account can send or receive, based on its trust lines.
type CurrenciesRequest struct {
	common.BaseRequest
	Account     types.Address          `json:"account"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Strict      bool                   `json:"strict,omitempty"`
}

func (*CurrenciesRequest) Method() string {
	return "account_currencies"
}

func (*CurrenciesRequest) APIVersion() int {
	return version.RippledAPIV1
}

// TODO: Implement (V2)
func (*CurrenciesRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the account_currencies method.
type CurrenciesResponse struct {
	LedgerHash        common.LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerIndex       common.LedgerIndex `json:"ledger_index"`
	ReceiveCurrencies []string           `json:"receive_currencies"`
	SendCurrencies    []string           `json:"send_currencies"`
	Validated         bool               `json:"validated"`
}
