package account

import (
	"errors"

	accounttypes "github.com/Peersyst/xrpl-go/xrpl/queries/account/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ErrNoAccountID is returned when no account ID is specified in a request.
var (
	ErrNoAccountID = errors.New("no account ID specified")
)

// ############################################################################
// Request
// ############################################################################

// ChannelsRequest returns information about an account's payment channels where the account is the source for a specific ledger version.
type ChannelsRequest struct {
	common.BaseRequest
	Account            types.Address          `json:"account"`
	DestinationAccount types.Address          `json:"destination_account,omitempty"`
	LedgerIndex        common.LedgerSpecifier `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash      `json:"ledger_hash,omitempty"`
	Limit              int                    `json:"limit,omitempty"`
	Marker             any                    `json:"marker,omitempty"`
}

// Method returns the method name for the ChannelsRequest.
func (*ChannelsRequest) Method() string {
	return "account_channels"
}

// APIVersion returns the API version supported by ChannelsRequest.
func (*ChannelsRequest) APIVersion() int {
	return version.RippledAPIV2
}

// Validate method to be added to each request struct
func (r *ChannelsRequest) Validate() error {
	if r.Account == "" {
		return ErrNoAccountID
	}

	return nil
}

// ############################################################################
// Response
// ############################################################################

// ChannelsResponse represents the response from the account_channels method, including payment channel results and pagination data.
type ChannelsResponse struct {
	Account     types.Address                `json:"account"`
	Channels    []accounttypes.ChannelResult `json:"channels"`
	LedgerIndex common.LedgerIndex           `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash            `json:"ledger_hash,omitempty"`
	Validated   bool                         `json:"validated,omitempty"`
	Limit       int                          `json:"limit,omitempty"`
	Marker      any                          `json:"marker,omitempty"`
}
