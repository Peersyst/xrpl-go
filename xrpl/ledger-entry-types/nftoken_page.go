package ledger

import "github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"

// The NFTokenPage object represents a collection of NFTs owned by the same account.
// An account can have multiple NFTokenPage entries, which form a doubly linked list.
// (Added by the NonFungibleTokensV1_1 amendment.)
//
// ```json
//
//	{
//	  "LedgerEntryType": "NFTokenPage",
//	  "PreviousPageMin":
//	    "8A244DD75DAF4AC1EEF7D99253A7B83D2297818B2297818B70E264D2000002F2",
//	  "NextPageMin":
//	    "8A244DD75DAF4AC1EEF7D99253A7B83D2297818B2297818BE223B0AE0000010B",
//	  "PreviousTxnID":
//	    "95C8761B22894E328646F7A70035E9DFBECC90EDD83E43B7B973F626D21A0822",
//	  "PreviousTxnLgrSeq":
//	    42891441,
//	  "NFTokens": [
//	    {
//	      "NFToken": {
//	        "NFTokenID":
//	          "000B013A95F14B0044F78A264E41713C64B5F89242540EE208C3098E00000D65",
//	        "URI": "697066733A2F2F62616679626569676479727A74357366703775646D37687537367568377932366E6634646675796C71616266336F636C67747179353566627A6469"
//	      }
//	    },
//	    /* potentially more objects */
//	  ]
//	}
//
// ```
type NFTokenPage struct {
	// The unique ID for this ledger entry. In JSON, this field is represented with different names depending on the
	// context and API method. (Note, even though this is specified as "optional" in the code, every ledger entry
	// should have one unless it's legacy data from very early in the XRP Ledger's history.)
	Index types.Hash256 `json:"index,omitempty"`
	// 	The value 0x0050, mapped to the string NFTokenPage, indicates that this is a
	// page containing NFToken objects.
	LedgerEntryType EntryType
	Flags           uint32
	// The locator of the next page, if any. Details about this field and how it should
	// be used are outlined below.
	NextPageMin types.Hash256 `json:",omitempty"`
	// The locator of the previous page, if any. Details about this field and how it should be used are outlined below.
	PreviousPageMin types.Hash256 `json:",omitempty"`
	// Identifies the transaction ID of the transaction that most recently modified this NFTokenPage object.
	PreviousTxnID types.Hash256 `json:",omitempty"`
	// The sequence of the ledger that contains the transaction that most recently modified this NFTokenPage object.
	PreviousTxnLgrSeq uint32 `json:",omitempty"`
	// The collection of NFToken objects contained in this NFTokenPage object. This
	// specification places an upper bound of 32 NFToken objects per page. Objects are
	// sorted from low to high with the NFTokenID used as the sorting parameter.
	NFTokens []types.NFToken
}

// EntryType returns the type of the ledger entry.
func (*NFTokenPage) EntryType() EntryType {
	return NFTokenPageEntry
}
