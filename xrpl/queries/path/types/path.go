package types

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type Alternative struct {
	PathsComputed     [][]transaction.PathStep `json:"paths_computed"`
	SourceAmount      types.CurrencyAmount     `json:"source_amount"`
	DestinationAmount types.CurrencyAmount     `json:"destination_amount,omitempty"`
}

type RippleAlternative struct {
	PathsComputed [][]transaction.PathStep `json:"paths_computed"`
	SourceAmount  types.CurrencyAmount     `json:"source_amount"`
}
