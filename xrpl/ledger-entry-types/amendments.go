package ledger

import (
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

// The Amendments ledger entry type contains a list of Amendments that are currently active.
// Each ledger version contains at most one Amendments entry.
// ```json
//
//	{
//	    "Amendments": [
//	        "42426C4D4F1009EE67080A9B7965B44656D7714D104A72F9B4369F97ABF044EE",
//	        "4C97EBA926031A7CF7D7B36FDE3ED66DDA5421192D63DE53FFB46E43B9DC8373",
//	        // (... Long list of enabled amendment IDs ...)
//	        "03BDC0099C4E14163ADA272C1B6F6FABB448CC3E51F522F978041E4B57D9158C",
//	        "35291ADD2D79EB6991343BDA0912269C817D0F094B02226C1C14AD2858962ED4"
//	    ],
//	    "Flags": 0,
//	    "LedgerEntryType": "Amendments",
//	    "Majorities": [
//	        {
//	            "Majority": {
//	            "Amendment": "7BB62DC13EC72B775091E9C71BF8CF97E122647693B50C5E87A80DFD6FCFAC50",
//	                "CloseTime": 779561310
//	            }
//	        },
//	        {
//	            "Majority": {
//	                "Amendment": "755C971C29971C9F20C6F080F2ED96F87884E40AD19554A5EBECDCEC8A1F77FE",
//	                "CloseTime": 779561310
//	            }
//	        }
//	    ],
//	    "index": "7DB0788C020F02780A673DC74757F23823FA3014C1866E72CC4CD8B226CD6EF4"
//	}
//
// ```
type Amendments struct {
	// The unique ID for this ledger entry. In JSON, this field is represented with different names depending on the
	// context and API method. (Note, even though this is specified as "optional" in the code, every ledger entry
	// should have one unless it's legacy data from very early in the XRP Ledger's history.)
	Index types.Hash256 `json:"index,omitempty"`
	// A bit-map of boolean flags enabled for this object. Currently, the protocol defines no flags for Amendments objects.
	// The value is always 0.
	Flags uint32
	// The value 0x0066, mapped to the string Amendments, indicates that this object describes the status of amendments to the XRP Ledger.
	LedgerEntryType EntryType
	// Array of 256-bit amendment IDs for all currently enabled amendments. If omitted, there are no enabled amendments.
	Amendments []types.Hash256 `json:",omitempty"`
	// 	Array of objects describing the status of amendments that have majority support but are not yet enabled.
	// If omitted, there are no pending amendments with majority support.
	Majorities []MajorityEntry `json:",omitempty"`
	// The identifying hash of the transaction that most recently modified this entry. (Added by the fixPreviousTxnID amendment.)
	PreviousTxnID types.Hash256 `json:",omitempty"`
	// The index of the ledger that contains the transaction that most recently modified this entry. (Added by the fixPreviousTxnID amendment.)
	PreviousTxnLgrSeq uint32 `json:",omitempty"`
}

// Returns the type of the ledger entry.
func (*Amendments) EntryType() EntryType {
	return AmendmentsEntry
}

type MajorityEntry struct {
	Majority Majority
}

// A Majority object describes the status of a pending amendment that has majority support but is not yet enabled.
type Majority struct {
	// The Amendment ID of the pending amendment.
	Amendment types.Hash256
	// The close time of the ledger version that reached majority support for this amendment.
	CloseTime uint32
}
