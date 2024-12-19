package types

import "github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"

type FeeLevels struct {
	MedianLevel     types.XRPCurrencyAmount `json:"median_level"`
	MinimumLevel    types.XRPCurrencyAmount `json:"minimum_level"`
	OpenLedgerLevel types.XRPCurrencyAmount `json:"open_ledger_level"`
	ReferenceLevel  types.XRPCurrencyAmount `json:"reference_level"`
}
