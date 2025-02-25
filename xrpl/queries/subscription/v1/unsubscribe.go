package v1

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type UnsubscribeOrderBook struct {
	TakerGets types.IssuedCurrencyAmount `json:"taker_gets,omitempty"`
	TakerPays types.IssuedCurrencyAmount `json:"taker_pays,omitempty"`
	Both      bool                       `json:"both,omitempty"`
}

// ############################################################################
// Request
// ############################################################################

// The unsubscribe command tells the server to stop sending messages for a
// particular subscription or set of subscriptions.
type UnsubscribeRequest struct {
	Streams          []string               `json:"streams,omitempty"`
	Accounts         []types.Address        `json:"accounts,omitempty"`
	AccountsProposed []types.Address        `json:"accounts_proposed,omitempty"`
	Books            []UnsubscribeOrderBook `json:"books,omitempty"`
}

func (*UnsubscribeRequest) Method() string {
	return "unsubscribe"
}

// TODO: Implement V2
func (*UnsubscribeRequest) Validate() error {
	return nil
}

func (*UnsubscribeRequest) APIVersion() int {
	return version.RippledAPIV1
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the unsubscribe method.
type UnsubscribeResponse struct {
}
