package ledger

import "github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"

// (Added by the NegativeUNL amendment.)
// The NegativeUNL ledger entry type contains the current status of the Negative UNL, a list of trusted validators currently believed to be offline.
//
// Each ledger version contains at most one NegativeUNL entry. If no validators are
// currently disabled or scheduled to be disabled, there is no NegativeUNL entry.
//
// ```json
//
//	{
//	    "DisabledValidators": [
//	      {
//	        "DisabledValidator": {
//	          "FirstLedgerSequence": 91371264,
//	          "PublicKey": "ED58F6770DB5DD77E59D28CB650EC3816E2FC95021BB56E720C9A12DA79C58A3AB"
//	        }
//	      }
//	    ],
//	    "Flags": 0,
//	    "LedgerEntryType": "NegativeUNL",
//	    "PreviousTxnID": "8D47FFE664BE6C335108DF689537625855A6A95160CC6D351341B92624D9C5E3",
//	    "PreviousTxnLgrSeq": 91442944,
//	    "index": "2E8A59AA9D3B5B186B0B9E0F62E6C02587CA74A4D778938E957B6357D364B244"
//	}
//
// ```
type NegativeUNL struct {
	// The unique ID for this ledger entry. In JSON, this field is represented with different names depending on the
	// context and API method. (Note, even though this is specified as "optional" in the code, every ledger entry
	// should have one unless it's legacy data from very early in the XRP Ledger's history.)
	Index types.Hash256 `json:"index,omitempty"`
	// Set of bit-flags for this ledger entry.
	Flags uint32
	// The value 0x004E, mapped to the string NegativeUNL, indicates that this entry is the Negative UNL.
	LedgerEntryType EntryType
	// A list of DisabledValidator objects (see below), each representing a trusted validator
	// that is currently disabled.
	DisabledValidators []DisabledValidatorEntry `json:",omitempty"`
	// The identifying hash of the transaction that most recently modified this entry.
	// (Added by the fixPreviousTxnID amendment.)
	PreviousTxnID types.Hash256 `json:",omitempty"`
	// The index of the ledger that contains the transaction that most recently modified this entry.
	// (Added by the fixPreviousTxnID amendment.)
	PreviousTxnLgrSeq uint32 `json:",omitempty"`
	// The public key of a trusted validator that is scheduled to be disabled in the next flag ledger.
	ValidatorToDisable string `json:",omitempty"`
	// The public key of a trusted validator in the Negative UNL that is scheduled to be re-enabled in the next flag ledger.
	ValidatorToReEnable string `json:",omitempty"`
}

// EntryType returns the type of the ledger entry.
func (*NegativeUNL) EntryType() EntryType {
	return NegativeUNLEntry
}

// Each DisabledValidator object represents one disabled validator. In JSON, a
// DisabledValidator object has one field, DisabledValidator, which in turn contains
// another object with the following fields:
type DisabledValidatorEntry struct {
	DisabledValidator DisabledValidator
}

type DisabledValidator struct {
	// The ledger index when the validator was added to the Negative UNL.
	FirstLedgerSequence uint32
	// The master public key of the validator, in hexadecimal.
	PublicKey string
}
