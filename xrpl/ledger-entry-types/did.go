package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

// A DID ledger entry holds references to, or data associated with, a single DID.
// Requires the "did" amendment to be enabled.
// Example:
// ```json
//
//	{
//	    "Account": "rpfqJrXg5uidNo2ZsRhRY6TiF1cvYmV9Fg",
//	    "DIDDocument": "646F63",
//	    "Data": "617474657374",
//	    "Flags": 0,
//	    "LedgerEntryType": "DID",
//	    "OwnerNode": "0",
//	    "PreviousTxnID": "A4C15DA185E6092DF5954FF62A1446220C61A5F60F0D93B4B09F708778E41120",
//	    "PreviousTxnLgrSeq": 4,
//	    "URI": "6469645F6578616D706C65",
//	    "index": "46813BE38B798B3752CA590D44E7FEADB17485649074403AD1761A2835CE91FF"
//	}
//
// ```
type DID struct {
	// The account that controls the DID.
	Account types.Address
	// The W3C standard DID document associated with the DID.
	// The DIDDocument field isn't checked for validity and is limited to a maximum length of 256 bytes.
	DIDDocument string `json:",omitempty"`
	// The public attestations of identity credentials associated with the DID.
	// The Data field isn't checked for validity and is limited to a maximum length of 256 bytes.
	Data string `json:",omitempty"`
	// The value 0x0049, mapped to the string DID, indicates that this object is a DID object.
	LedgerEntryType string
	// A hint indicating which page of the sender's owner directory links to this entry, in case the directory consists of multiple pages.
	OwnerNode string
	// The identifying hash of the transaction that most recently modified this object.
	PreviousTxnID string
	// The index of the ledger that contains the transaction that most recently modified this object.
	PreviousTxnLgrSeq uint32
	// The Universal Resource Identifier that points to the corresponding DID document or the data associated with the DID.
	// This field can be an HTTP(S) URL or IPFS URI.
	// This field isn't checked for validity and is limited to a maximum length of 256 bytes.
	URI string `json:",omitempty"`
}

// EntryType returns the type of the ledger entry.
func (*DID) EntryType() LedgerEntryType {
	return DIDEntry
}
