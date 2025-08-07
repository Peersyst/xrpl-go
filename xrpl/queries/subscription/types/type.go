// revive:disable:var-naming
// Package types defines types for subscription streams in XRPL.
package types

// Type represents a subscription stream type for server events.
type Type string

// Stream types used in subscription requests.
const (
	LedgerStreamType      Type = "ledgerClosed"
	ValidationStreamType  Type = "validationReceived"
	TransactionStreamType Type = "transaction"
	PeerStatusStreamType  Type = "peerStatusChange"
	OrderBookStreamType   Type = TransactionStreamType
	ConsensusStreamType   Type = "consensusPhase"
)
