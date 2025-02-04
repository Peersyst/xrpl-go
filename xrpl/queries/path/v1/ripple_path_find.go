package v1

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	pathtypes "github.com/Peersyst/xrpl-go/xrpl/queries/path/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The ripple_path_find method is a simplified version of the path_find method
// that provides a single response with a payment path you can use right away.
type RipplePathFindRequest struct {
	common.BaseRequest
	SourceAccount      types.Address                      `json:"source_account"`
	DestinationAccount types.Address                      `json:"destination_account"`
	DestinationAmount  types.CurrencyAmount               `json:"destination_amount"`
	SendMax            types.CurrencyAmount               `json:"send_max,omitempty"`
	SourceCurrencies   []pathtypes.RipplePathFindCurrency `json:"source_currencies,omitempty"`
	LedgerHash         common.LedgerHash                  `json:"ledger_hash,omitempty"`
	LedgerIndex        common.LedgerSpecifier             `json:"ledger_index,omitempty"`
}

func (*RipplePathFindRequest) Method() string {
	return "ripple_path_find"
}

func (*RipplePathFindRequest) APIVersion() int {
	return version.RippledAPIV1
}

// TODO: Implement V2
func (*RipplePathFindRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the ripple_path_find method.
type RipplePathFindResponse struct {
	Alternatives          []pathtypes.RippleAlternative `json:"alternatives"`
	DestinationAccount    types.Address                 `json:"destination_account"`
	DestinationCurrencies []string                      `json:"destination_currencies"`
	DestinationAmount     any                           `json:"destination_amount"`
	FullReply             bool                          `json:"full_reply,omitempty"`
	ID                    any                           `json:"id,omitempty"`
	LedgerCurrentIndex    int                           `json:"ledger_current_index,omitempty"`
	SourceAccount         types.Address                 `json:"source_account"`
	Validated             bool                          `json:"validated"`
}
