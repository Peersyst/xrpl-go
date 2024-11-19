package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// Create a Check object in the ledger, which is a deferred payment that can be cashed by its intended destination. The sender of this transaction is the sender of the Check.
//
// Example:
//
// ```json
//
//	{
//	    "Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
//	    "TransactionType": "EscrowCreate",
//	    "Amount": "10000",
//	    "Destination": "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
//	    "CancelAfter": 533257958,
//	    "FinishAfter": 533171558,
//	    "Condition": "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
//	    "DestinationTag": 23480,
//	    "SourceTag": 11747
//	}
//
// / ```
type CheckCreate struct {
	BaseTx
	// The unique address of the account that can cash the Check.
	Destination types.Address
	// Maximum amount of source currency the Check is allowed to debit the sender, including transfer fees on non-XRP currencies.
	// The Check can only credit the destination with the same currency (from the same issuer, for non-XRP currencies). For non-XRP amounts, the nested field names MUST be lower-case.
	SendMax types.CurrencyAmount
	// (Optional) Arbitrary tag that identifies the reason for the Check, or a hosted recipient to pay.
	DestinationTag uint32 `json:",omitempty"`
	// (Optional) Time after which the Check is no longer valid, in seconds since the Ripple Epoch.
	Expiration uint32 `json:",omitempty"`
	// (Optional) Arbitrary 256-bit hash representing a specific reason or identifier for this Check.
	InvoiceID types.Hash256 `json:",omitempty"`
}

// TxType returns the type of the transaction (CheckCreate).
func (*CheckCreate) TxType() TxType {
	return CheckCreateTx
}

// Flatten returns the flattened map of the CheckCreate transaction.
func (c *CheckCreate) Flatten() FlatTransaction {
	flattened := c.BaseTx.Flatten()

	flattened["TransactionType"] = c.TxType().String()
	flattened["Destination"] = c.Destination.String()
	flattened["SendMax"] = c.SendMax.Flatten()
	if c.DestinationTag != 0 {
		flattened["DestinationTag"] = c.DestinationTag
	}
	if c.Expiration != 0 {
		flattened["Expiration"] = c.Expiration
	}
	if c.InvoiceID != "" {
		flattened["InvoiceID"] = c.InvoiceID.String()
	}
	return flattened
}
