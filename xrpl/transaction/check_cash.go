package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// Attempts to redeem a Check object in the ledger to receive up to the amount authorized by the corresponding CheckCreate transaction.
// Only the Destination address of a Check can cash it with a CheckCash transaction.
// Cashing a check this way is similar to executing a Payment initiated by the destination.
//
// Since the funds for a check are not guaranteed, redeeming a Check can fail because the sender does not have a high enough balance or because there is not enough liquidity to deliver the funds.
// If this happens, the Check remains in the ledger and the destination can try to cash it again later, or for a different amount.
//
// ```json
//
//	{
//		"Account": "rfkE1aSy9G8Upk4JssnwBxhEv5p4mn2KTy",
//		"TransactionType": "CheckCash",
//		"Amount": "100000000",
//		"CheckID": "838766BA2B995C00744175F69A1B11E32C3DBC40E64801A4056FCBD657F57334",
//		"Fee": "12"
//	}
//
// ```
type CheckCash struct {
	BaseTx
	// The ID of the Check ledger object to cash, as a 64-character hexadecimal string.
	CheckID types.Hash256
	// (Optional) Redeem the Check for exactly this amount, if possible.
	// The currency must match that of the SendMax of the corresponding CheckCreate transaction.
	// You must provide either this field or DeliverMin.
	Amount types.CurrencyAmount `json:",omitempty"`
	// (Optional) Redeem the Check for at least this amount and for as much as possible.
	// The currency must match that of the SendMax of the corresponding CheckCreate transaction.
	// You must provide either this field or Amount.
	DeliverMin types.CurrencyAmount `json:",omitempty"`
}

// TxType returns the type of the transaction (CheckCash).
func (*CheckCash) TxType() TxType {
	return CheckCashTx
}

// Flatten returns the flattened map of the CheckCash transaction.
func (c *CheckCash) Flatten() FlatTransaction {
	flattened := c.BaseTx.Flatten()

	flattened["TransactionType"] = c.TxType().String()
	flattened["CheckID"] = c.CheckID.String()
	if c.Amount != nil {
		flattened["Amount"] = c.Amount.Flatten()
	}
	if c.DeliverMin != nil {
		flattened["DeliverMin"] = c.DeliverMin.Flatten()
	}

	return flattened
}
