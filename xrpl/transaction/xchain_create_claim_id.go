package transaction

import (
	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

// (Requires the XChainBridge amendment )
//
// The XChainCreateClaimID transaction creates a new cross-chain claim ID that is used for a cross-chain transfer.
// A cross-chain claim ID represents one cross-chain transfer of value.
//
// This transaction is the first step of a cross-chain transfer of value and is submitted on the destination chain,
// not the source chain.
//
// It also includes the account on the source chain that locks or burns the funds on the source chain.
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
type XChainCreateClaimID struct {
	BaseTx

	// The account that must send the XChainCommit transaction on the source chain.
	OtherChainSource types.Address
	// The amount, in XRP, to reward the witness servers for providing signatures.
	// This must match the amount on the Bridge ledger object.
	SignatureReward types.CurrencyAmount
	// The bridge to create the claim ID for.
	XChainBridge types.XChainBridge
}

// Returns the type of the transaction.
func (x *XChainCreateClaimID) TxType() TxType {
	return XChainCreateClaimIDTx
}

// Returns a flattened version of the transaction.
func (x *XChainCreateClaimID) Flatten() FlatTransaction {
	flatTx := x.BaseTx.Flatten()

	flatTx["TransactionType"] = x.TxType().String()

	if x.OtherChainSource != "" {
		flatTx["OtherChainSource"] = x.OtherChainSource.String()
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
func (x *XChainCreateClaimID) Validate() (bool, error) {
	_, err := x.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if !addresscodec.IsValidClassicAddress(x.OtherChainSource.String()) {
		return false, ErrInvalidAccount
	}

	if ok, err := IsAmount(x.SignatureReward, "SignatureReward", true); !ok {
		return false, err
	}

	return x.XChainBridge.Validate()
}
