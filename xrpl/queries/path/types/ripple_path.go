package types

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type RipplePathFindCurrency struct {
	Currency string        `json:"currency"`
	Issuer   types.Address `json:"issuer,omitempty"`
}
