package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

// Every entry in a ledger's state data has the same set of common fields, plus additional fields based on the ledger entry type.
// Field names are case-sensitive.
type LedgerEntryCommonFields struct {
	// The unique ID for this ledger entry.
	// In JSON, this field is represented with different names depending on the context and API method.
	// (Note, even though this is specified as "optional" in the code, every ledger entry should have one unless it's legacy data from very early in the XRP Ledger's history.)
	Index       types.Hash256 `json:"index,omitempty"`
	LedgerIndex types.Hash256 `json:",omitempty"`
	// The type of ledger entry. Valid ledger entry types include AccountRoot, Offer, RippleState, and others.
	LedgerEntryType string `json:",omitempty"`
	// Set of bit-flags for this ledger entry.
	Flags uint32
}
