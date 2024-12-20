package ledger

import "github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"

// The FeeSettings entry contains the current base transaction cost and reserve amounts as determined by fee voting.
// Each ledger version contains at most one FeeSettings entry.
//
// ```json
//
//	{
//	   "BaseFee": "000000000000000A",
//	   "Flags": 0,
//	   "LedgerEntryType": "FeeSettings",
//	   "ReferenceFeeUnits": 10,
//	   "ReserveBase": 20000000,
//	   "ReserveIncrement": 5000000,
//	   "index": "4BC50C9B0D8515D3EAAE1E74B29A95804346C491EE1A95BF25E4AAB854A6A651"
//	}
//
// ```
type FeeSettings struct {
	// The unique ID for this ledger entry. In JSON, this field is represented with different names depending on the
	// context and API method. (Note, even though this is specified as "optional" in the code, every ledger entry
	// should have one unless it's legacy data from very early in the XRP Ledger's history.)
	Index types.Hash256 `json:"index,omitempty"`
	// A bit-map of boolean flags enabled for this object. Currently, the protocol defines
	// no flags for FeeSettings objects. The value is always 0.
	Flags uint32
	// The value 0x0073, mapped to the string FeeSettings, indicates that this object
	// contains the ledger's fee settings.
	LedgerEntryType EntryType
	// The transaction cost of the "reference transaction" in drops of XRP as hexadecimal.
	BaseFee           string
	ReferenceFeeUnits uint32
	// The base reserve for an account in the XRP Ledger, as drops of XRP.
	ReserveBase uint32
	// The incremental owner reserve for owning objects, as drops of XRP.
	ReserveIncrement uint32
	// The identifying hash of the transaction that most recently modified this entry.
	// (Added by the fixPreviousTxnID amendment.)
	PreviousTxnID types.Hash256 `json:",omitempty"`
	// The index of the ledger that contains the transaction that most recently modified this entry. (Added by the fixPreviousTxnID amendment.)
	PreviousTxnLgrSeq uint32 `json:",omitempty"`
}

// EntryType returns the type of the ledger entry.
func (*FeeSettings) EntryType() EntryType {
	return FeeSettingsEntry
}
