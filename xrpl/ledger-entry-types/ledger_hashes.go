package ledger

import "github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"

// (Not to be confused with the "ledger hash" string data type, which uniquely identifies a ledger version.
// This section describes the LedgerHashes ledger object type.)
//
// The LedgerHashes object type contains a history of prior ledgers that led up to
// this ledger version, in the form of their hashes. Objects of this ledger type are
// modified automatically when closing a ledger. (This is one of the only times a
// ledger's state data is modified without a transaction or pseudo-transaction.)
// The LedgerHashes objects exist to make it possible to look up a previous ledger's
// hash with only the current ledger version and at most one lookup of a previous
// ledger version.
//
// There are two kinds of LedgerHashes object. Both types have the same fields. Each
// ledger version contains:
//   - Exactly one "recent history" LedgerHashes object
//   - A number of "previous history" LedgerHashes objects based on the current ledger index
//     (that is, the length of the ledger history). Specifically, the XRP Ledger adds a new
//     "previous history" object every 65536 ledger versions.
//
// ```json
//
//	{
//	  "LedgerEntryType": "LedgerHashes",
//	  "Flags": 0,
//	  "FirstLedgerSequence": 2,
//	  "LastLedgerSequence": 33872029,
//	  "Hashes": [
//	    "D638208ADBD04CBB10DE7B645D3AB4BA31489379411A3A347151702B6401AA78",
//	    "254D690864E418DDD9BCAC93F41B1F53B1AE693FC5FE667CE40205C322D1BE3B",
//	    "A2B31D28905E2DEF926362822BC412B12ABF6942B73B72A32D46ED2ABB7ACCFA",
//	    "AB4014846DF818A4B43D6B1686D0DE0644FE711577C5AB6F0B2A21CCEE280140",
//	    "3383784E82A8BA45F4DD5EF4EE90A1B2D3B4571317DBAC37B859836ADDE644C1",
//	    ... (up to 256 ledger hashes) ...
//	  ],
//	  "index": "B4979A36CDC7F3D3D5C31A4EAE2AC7D7209DDA877588B9AFC66799692AB0D66B"
//	}
//
// ```
type Hashes struct {
	// The unique ID for this ledger entry. In JSON, this field is represented with different names depending on the
	// context and API method. (Note, even though this is specified as "optional" in the code, every ledger entry
	// should have one unless it's legacy data from very early in the XRP Ledger's history.)
	Index types.Hash256 `json:"index,omitempty"`
	// A bit-map of boolean flags enabled for this object. Currently, the protocol defines
	// no flags for LedgerHashes objects. The value is always 0.
	Flags uint32
	// The value 0x0068, mapped to the string LedgerHashes, indicates that this object is a list of ledger hashes.
	LedgerEntryType EntryType
	// DEPRECATED Do not use. (The "recent hashes" object on Mainnet has the value 2 in this
	// field as a result of an old software bug. That value gets carried forward as the
	// "recent hashes" object is updated. New "previous history" objects do not have this field,
	// nor do "recent hashes" objects in parallel networks started with more recent versions of rippled.)
	FirstLedgerSequence uint32 `json:",omitempty"`
	// An array of up to 256 ledger hashes. The contents depend on which sub-type of LedgerHashes object this is.
	Hashes []types.Hash256
	// The Ledger Index of the last entry in this object's Hashes array.
	LastLedgerSequence uint32 `json:",omitempty"`
}

// EntryType returns the type of the ledger entry.
func (*Hashes) EntryType() EntryType {
	return LedgerHashesEntry
}
