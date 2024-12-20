package ledger

import "github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"

// (Added by the TicketBatch amendment.)
// A Ticket entry type represents a Ticket, which tracks an account sequence number that has
// been set aside for future use. You can create new tickets with a TicketCreate transaction.
//
// ```json
//
//	{
//	  "Account": "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
//	  "Flags": 0,
//	  "LedgerEntryType": "Ticket",
//	  "OwnerNode": "0000000000000000",
//	  "PreviousTxnID": "F19AD4577212D3BEACA0F75FE1BA1644F2E854D46E8D62E9C95D18E9708CBFB1",
//	  "PreviousTxnLgrSeq": 4,
//	  "TicketSequence": 3
//	}
//
// ```
type Ticket struct {
	// The unique ID for this ledger entry. In JSON, this field is represented with different names depending on the
	// context and API method. (Note, even though this is specified as "optional" in the code, every ledger entry
	// should have one unless it's legacy data from very early in the XRP Ledger's history.)
	Index types.Hash256 `json:"index,omitempty"`
	// The account that owns this Ticket.
	Account types.Address
	Flags   uint32
	// The value 0x0054, mapped to the string Ticket, indicates that this is a Ticket entry.
	LedgerEntryType EntryType
	// A hint indicating which page of the owner directory links to this entry, in case the
	// directory consists of multiple pages.
	OwnerNode string
	// The identifying hash of the transaction that most recently modified this entry.
	PreviousTxnID types.Hash256
	// The index of the ledger that contains the transaction that most recently modified this entry.
	PreviousTxnLgrSeq uint32
	// The Sequence Number this Ticket sets aside.
	TicketSequence uint32
}

// EntryType returns the type of the ledger entry.
func (*Ticket) EntryType() EntryType {
	return TicketEntry
}
