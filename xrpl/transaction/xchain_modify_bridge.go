package transaction

import (
	"errors"

	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

var (
	ErrInvalidFlags = errors.New("invalid flags")
)

const (
	tfClearAccountCreateAmount uint32 = 0x00010000
)

// (Requires the XChainBridge amendment )
//
// The XChainCreateClaimID transaction creates a new cross-chain claim ID that is used
// for a cross-chain transfer. A cross-chain claim ID represents one cross-chain transfer
// of value.
//
// This transaction is the first step of a cross-chain transfer of value and is submitted
// on the destination chain, not the source chain.
//
// It also includes the account on the source chain that locks or burns the funds on the
// source chain.
//
// ```json
//
//	{
//	  "TransactionType": "XChainModifyBridge",
//	  "Account": "rhWQzvdmhf5vFS35vtKUSUwNZHGT53qQsg",
//	  "XChainBridge": {
//	    "LockingChainDoor": "rhWQzvdmhf5vFS35vtKUSUwNZHGT53qQsg",
//	    "LockingChainIssue": {
//	      "currency": "XRP"
//	    },
//	    "IssuingChainDoor": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
//	    "IssuingChainIssue": {
//	      "currency": "XRP"
//	    }
//	  },
//	  "SignatureReward": 200,
//	  "MinAccountCreateAmount": 1000000
//	}
//
// ```
type XChainModifyBridge struct {
	BaseTx

	// Specifies the flags for this transaction.
	Flags uint32
	// The minimum amount, in XRP, required for a XChainAccountCreateCommit transaction.
	// If this is not present, the XChainAccountCreateCommit transaction will fail.
	// This field can only be present on XRP-XRP bridges.
	MinAccountCreateAmount types.CurrencyAmount `json:",omitempty"`
	// The signature reward split between the witnesses for submitting attestations.
	SignatureReward types.CurrencyAmount `json:",omitempty"`
	// The bridge to modify.
	XChainBridge types.XChainBridge
}

// Returns the type of the transaction.
func (x *XChainModifyBridge) TxType() TxType {
	return XChainModifyBridgeTx
}

// Sets the clear account create amount flag.
func (x *XChainModifyBridge) SetClearAccountCreateAmount() {
	x.Flags |= tfClearAccountCreateAmount
}

// Returns a flattened version of the transaction.
func (x *XChainModifyBridge) Flatten() FlatTransaction {
	flatTx := x.BaseTx.Flatten()

	flatTx["TransactionType"] = x.TxType().String()

	if x.Flags != 0 {
		flatTx["Flags"] = x.Flags
	}

	if x.MinAccountCreateAmount != nil {
		flatTx["MinAccountCreateAmount"] = x.MinAccountCreateAmount.Flatten()
	}

	if x.SignatureReward != nil {
		flatTx["SignatureReward"] = x.SignatureReward.Flatten()
	}

	if x.XChainBridge != (types.XChainBridge{}) {
		flatTx["XChainBridge"] = x.XChainBridge.Flatten()
	}

	return flatTx
}

// Validates the transaction.
func (x *XChainModifyBridge) Validate() (bool, error) {
	_, err := x.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if !IsFlagEnabled(x.Flags, tfClearAccountCreateAmount) {
		return false, ErrInvalidFlags
	}

	if ok, err := IsAmount(x.MinAccountCreateAmount, "MinAccountCreateAmount", false); !ok {
		return false, err
	}

	if ok, err := IsAmount(x.SignatureReward, "SignatureReward", false); !ok {
		return false, err
	}

	return x.XChainBridge.Validate()
}
