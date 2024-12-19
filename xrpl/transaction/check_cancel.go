package transaction

import (
	"github.com/Peersyst/xrpl-go/v1/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

// Cancels an unredeemed Check, removing it from the ledger without sending any money.
// The source or the destination of the check can cancel a Check at any time using this transaction type.
// If the Check has expired, any address can cancel it.
//
// Example:
//
// ```json
//
//	{
//		"Account": "rUn84CUYbNjRoTQ6mSW7BVJPSVJNLb1QLo",
//		"TransactionType": "CheckCancel",
//		"CheckID": "49647F0D748DC3FE26BDACBC57F251AADEFFF391403EC9BF87C97F67E9977FB0",
//		"Fee": "12"
//	}
//
// ```
type CheckCancel struct {
	BaseTx
	// The ID of the Check ledger object to cancel, as a 64-character hexadecimal string.
	CheckID types.Hash256
}

// TxType returns the type of the transaction (CheckCancel).
func (*CheckCancel) TxType() TxType {
	return CheckCancelTx
}

// Flatten returns the flattened map of the CheckCancel transaction.
func (c *CheckCancel) Flatten() FlatTransaction {
	flattened := c.BaseTx.Flatten()

	flattened["TransactionType"] = c.TxType().String()
	flattened["CheckID"] = c.CheckID.String()

	return flattened
}

// Validate checks the validity of the CheckCancel transaction.
func (c *CheckCancel) Validate() (bool, error) {
	ok, err := c.BaseTx.Validate()
	if err != nil || !ok {
		return false, err
	}

	if !typecheck.IsHex(c.CheckID.String()) || len(c.CheckID.String()) != 64 {
		return false, ErrInvalidCheckID
	}

	return true, nil
}
