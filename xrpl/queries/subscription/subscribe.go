package subscribe

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	streamtypes "github.com/Peersyst/xrpl-go/xrpl/queries/subscription/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The subscribe method requests periodic notifications from the server when
// certain events happen.
type Request struct {
	common.BaseRequest
	Streams          []string                `json:"streams,omitempty"`
	Accounts         []types.Address         `json:"accounts,omitempty"`
	AccountsProposed []types.Address         `json:"accounts_proposed,omitempty"`
	Books            []streamtypes.OrderBook `json:"books,omitempty"`
	URL              string                  `json:"url,omitempty"`
	URLUsername      string                  `json:"url_username,omitempty"`
	URLPassword      string                  `json:"url_password,omitempty"`
}

func (*Request) Method() string {
	return "subscribe"
}

// TODO: Implement V2
func (*Request) Validate() error {
	return nil
}

func (*Request) APIVersion() int {
	return version.RippledAPIV2
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the subscribe method.
type Response struct {
	LoadBase         uint               `json:"load_base,omitempty"`
	LoadFactor       uint               `json:"load_factor,omitempty"`
	Random           string             `json:"random,omitempty"`
	ServerStatus     string             `json:"server_status,omitempty"`
	FeeBase          uint               `json:"fee_base,omitempty"`
	FeeRef           uint               `json:"fee_ref,omitempty"`
	LedgerHash       common.LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerIndex      common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerTime       uint64             `json:"ledger_time,omitempty"`
	ReserveBase      uint               `json:"reserve_base,omitempty"`
	ReserveInc       uint               `json:"reserve_inc,omitempty"`
	ValidatedLedgers string             `json:"validated_ledgers,omitempty"`
	Offers           []ledger.Offer     `json:"offers,omitempty"`
}
