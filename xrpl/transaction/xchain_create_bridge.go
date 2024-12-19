package transaction

import "github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"

// (Requires the XChainBridge amendment )
//
// The XChainCreateBridge transaction creates a new Bridge ledger object and defines a new cross-chain bridge entrance on
// the chain that the transaction is submitted on. It includes information about door accounts and assets for the bridge.
//
// The transaction must be submitted first by the locking chain door account. To set up a valid bridge, door accounts on
// both chains must submit this transaction, in addition to setting up witness servers.
//
// The complete production-grade setup would also include a SignerListSet transaction on the two door accounts for the
// witnesses’ signing keys, as well as disabling the door accounts’ master key. This ensures that the witness servers are truly in control of the funds.
//
// ```json
//
//	{
//	  "Account": "rahDmoXrtPdh7sUdrPjini3gcnTVYjbjjw",
//	  "OtherChainSource": "rMTi57fNy2UkUb4RcdoUeJm7gjxVQvxzUo",
//	  "TransactionType": "XChainCreateClaimID",
//	  "SignatureReward": "100",
//	  "XChainBridge": {
//	    "LockingChainDoor": "rMAXACCrp3Y8PpswXcg3bKggHX76V3F8M4",
//	    "LockingChainIssue": {
//	      "currency": "XRP"
//	    },
//	    "IssuingChainDoor": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
//	    "IssuingChainIssue": {
//	      "currency": "XRP"
//	    }
//	  }
//	}
//
// ```
type XChainCreateBridge struct {
	BaseTx
	// The minimum amount, in XRP, required for a XChainAccountCreateCommit transaction.
	// If this isn't present, the XChainAccountCreateCommit transaction will fail.
	// This field can only be present on XRP-XRP bridges.
	MinAccountCreateAmount types.CurrencyAmount `json:",omitempty"`
	// The total amount to pay the witness servers for their signatures. This amount will be split among the signers.
	SignatureReward types.CurrencyAmount
	// The bridge (door accounts and assets) to create.
	XChainBridge types.XChainBridge
}

// Returns the type of the transaction.
func (x *XChainCreateBridge) TxType() TxType {
	return XChainCreateBridgeTx
}

// Returns a flattened version of the transaction.
func (x *XChainCreateBridge) Flatten() FlatTransaction {
	flatTx := x.BaseTx.Flatten()

	flatTx["TransactionType"] = x.TxType().String()

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
func (x *XChainCreateBridge) Validate() (bool, error) {
	_, err := x.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if ok, err := IsAmount(x.MinAccountCreateAmount, "MinAccountCreateAmount", false); !ok {
		return false, err
	}

	if ok, err := IsAmount(x.SignatureReward, "SignatureReward", true); !ok {
		return false, err
	}

	return x.XChainBridge.Validate()
}
