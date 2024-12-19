package transaction

import (
	"github.com/Peersyst/xrpl-go/v1/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

// (Requires the XChainBridge amendment)
//
// The XChainCommit is the second step in a cross-chain transfer. It puts assets into trust on the locking chain
// so that they can be wrapped on the issuing chain, or burns wrapped assets on the issuing chain so that they can
// be returned on the locking chain.
//
// ```json
//
//	{
//	  "Account": "rMTi57fNy2UkUb4RcdoUeJm7gjxVQvxzUo",
//	  "TransactionType": "XChainCommit",
//	  "XChainBridge": {
//	    "LockingChainDoor": "rMAXACCrp3Y8PpswXcg3bKggHX76V3F8M4",
//	    "LockingChainIssue": {
//	      "currency": "XRP"
//	    },
//	    "IssuingChainDoor": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
//	    "IssuingChainIssue": {
//	      "currency": "XRP"
//	    }
//	  },
//	  "Amount": "10000",
//	  "XChainClaimID": "13f"
//	}
//
// ```
type XChainCommit struct {
	BaseTx

	// The asset to commit, and the quantity. This must match the door account's LockingChainIssue
	// (if on the locking chain) or the door account's IssuingChainIssue (if on the issuing chain).
	Amount types.CurrencyAmount
	// The destination account on the destination chain. If this is not specified, the account that
	// submitted the XChainCreateClaimID transaction on the destination chain will need to submit a
	// XChainClaim transaction to claim the funds.
	OtherChainDestination types.Address `json:",omitempty"`
	// The bridge to use to transfer funds.
	XChainBridge types.XChainBridge
	// The unique integer ID for a cross-chain transfer. This must be acquired on the destination
	// chain (via a XChainCreateClaimID transaction) and checked from a validated ledger before
	// submitting this transaction. If an incorrect sequence number is specified, the funds will
	// be lost.
	XChainClaimID string
}

// Returns the type of the transaction.
func (x *XChainCommit) TxType() TxType {
	return XChainCommitTx
}

// Returns a flattened version of the transaction.
func (x *XChainCommit) Flatten() FlatTransaction {
	flatTx := x.BaseTx.Flatten()

	flatTx["TransactionType"] = x.TxType().String()

	if x.Amount != nil {
		flatTx["Amount"] = x.Amount.Flatten()
	}

	if x.OtherChainDestination != "" {
		flatTx["OtherChainDestination"] = x.OtherChainDestination.String()
	}

	if x.XChainBridge != (types.XChainBridge{}) {
		flatTx["XChainBridge"] = x.XChainBridge.Flatten()
	}

	if x.XChainClaimID != "" {
		flatTx["XChainClaimID"] = x.XChainClaimID
	}

	return flatTx
}

// Validates the transaction.
func (x *XChainCommit) Validate() (bool, error) {
	_, err := x.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if ok, err := IsAmount(x.Amount, "Amount", true); !ok {
		return false, err
	}

	if ok, err := x.XChainBridge.Validate(); !ok {
		return false, err
	}

	if x.XChainClaimID == "" || !typecheck.IsHex(x.XChainClaimID) {
		return false, ErrInvalidXChainClaimID
	}

	return true, nil
}
