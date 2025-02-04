package account

import (
	"errors"

	accounttypes "github.com/Peersyst/xrpl-go/xrpl/queries/account/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

var (
	ErrNoAccountID = errors.New("no account ID specified")
)

// ############################################################################
// Request
// ############################################################################

//   - The account_channels method returns information about an account's Payment
//
// Channels. This includes only channels where the specified account is the
// channel's source, not the destination. (A channel's "source" and "owner" are
// the same.) All information retrieved is relative to a particular version of
// the ledger.
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

// The expected response from the account_channels method.
type ChannelsResponse struct {
	Account     types.Address                `json:"account"`
	Channels    []accounttypes.ChannelResult `json:"channels"`
	LedgerIndex common.LedgerIndex           `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash            `json:"ledger_hash,omitempty"`
	Validated   bool                         `json:"validated,omitempty"`
	Limit       int                          `json:"limit,omitempty"`
	Marker      any                          `json:"marker,omitempty"`
}
