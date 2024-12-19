package types

import (
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction"
)

type Alternative struct {
	PathsComputed [][]transaction.PathStep `json:"paths_computed"`
	// SourceAmount      types.CurrencyAmount     `json:"source_amount"`
	SourceAmount any `json:"source_amount"`
	// DestinationAmount types.CurrencyAmount     `json:"destination_amount,omitempty"`
	DestinationAmount any `json:"destination_amount,omitempty"`
}

type RippleAlternative struct {
	PathsComputed [][]transaction.PathStep `json:"paths_computed"`
	// SourceAmount  types.CurrencyAmount     `json:"source_amount"`
	SourceAmount any `json:"source_amount"`
}
