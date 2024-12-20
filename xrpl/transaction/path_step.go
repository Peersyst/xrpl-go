package transaction

import (
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

type PathStep struct {
	Account  types.Address `json:"account,omitempty"`
	Currency string        `json:"currency,omitempty"`
	Issuer   types.Address `json:"issuer,omitempty"`
}

func (p *PathStep) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})

	if p.Account != "" {
		flattened["account"] = p.Account.String()
	}

	if p.Currency != "" {
		flattened["currency"] = p.Currency
	}

	if p.Issuer != "" {
		flattened["issuer"] = p.Issuer.String()
	}

	return flattened

}
