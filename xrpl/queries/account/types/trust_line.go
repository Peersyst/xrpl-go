package types

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type TrustLine struct {
	Account        types.Address `json:"account"`
	Balance        string        `json:"balance"`
	Currency       string        `json:"currency"`
	Limit          string        `json:"limit"`
	LimitPeer      string        `json:"limit_peer"`
	QualityIn      uint          `json:"quality_in"`
	QualityOut     uint          `json:"quality_out"`
	NoRipple       bool          `json:"no_ripple,omitempty"`
	NoRipplePeer   bool          `json:"no_ripple_peer,omitempty"`
	Authorized     bool          `json:"authorized,omitempty"`
	PeerAuthorized bool          `json:"peer_authorized,omitempty"`
	Freeze         bool          `json:"freeze,omitempty"`
	FreezePeer     bool          `json:"freeze_peer,omitempty"`
}
