package types

import "github.com/Peersyst/xrpl-go/xrpl/queries/common"

type PeerStatusEvents string

const (
	// The peer closed a ledger version with this Ledger Index, which usually means it is about
	// to start consensus.
	PeerStatusClosingLedger PeerStatusEvents = "CLOSING_LEDGER"
	// The peer built this ledger version as the result of a consensus round. Note: This ledger
	// is still not certain to become immutably validated.
	PeerStatusAcceptedLedger PeerStatusEvents = "ACCEPTED_LEDGER"
	// The peer concluded it was not following the rest of the network and switched to a different
	// ledger version.
	PeerStatusSwitchedLedger PeerStatusEvents = "SWITCHED_LEDGER"
	// The peer fell behind the rest of the network in tracking which ledger versions are validated
	// and which are undergoing consensus.
	PeerStatusLostSync PeerStatusEvents = "LOST_SYNC"
)

type PeerStatusStream struct {
	// `peerStatusChange` indicates this comes from the Peer Status stream.
	Type Type `json:"type"`
	// The type of event that prompted this message. See Peer Status Events for possible values.
	Action PeerStatusEvents `json:"action"`
	// The time this event occurred, in seconds since the Ripple Epoch.
	Date uint64 `json:"date"`
	// (May be omitted) The identifying Hash of a ledger version to which this message pertains.
	LedgerHash common.LedgerHash `json:"ledger_hash,omitempty"`
	// (May be omitted) The Ledger Index of a ledger version to which this message pertains.
	LedgerIndex common.LedgerIndex `json:"ledger_index,omitempty"`
	// (May be omitted) The largest Ledger Index the peer has currently available.
	LedgerIndexMax common.LedgerIndex `json:"ledger_index_max,omitempty"`
	// (May be omitted) The smallest Ledger Index the peer has currently available.
	LedgerIndexMin common.LedgerIndex `json:"ledger_index_min,omitempty"`
}
