package transaction

import (
	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

// (Requires the XChainBridge amendment )
//
// This transaction can only be used for XRP-XRP bridges.
//
// The XChainAccountCreateCommit transaction creates a new account for a witness server to submit transactions on an issuing chain.
//
// ```json
//
//	{
//	  "Account": "rwEqJ2UaQHe7jihxGqmx6J4xdbGiiyMaGa",
//	  "Destination": "rD323VyRjgzzhY4bFpo44rmyh2neB5d8Mo",
//	  "TransactionType": "XChainAccountCreateCommit",
//	  "Amount": "20000000",
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
type XChainAccountCreateCommit struct {
	BaseTx

	// The amount, in XRP, to use for account creation. This must be greater than or equal to
	// the MinAccountCreateAmount specified in the Bridge ledger object.
	Amount types.CurrencyAmount
	// The destination account on the destination chain.
	Destination types.Address
	// The amount, in XRP, to be used to reward the witness servers for providing signatures.
	// This must match the amount on the Bridge ledger object.
	SignatureReward types.CurrencyAmount `json:",omitempty"`
	// The bridge to create accounts for.
	XChainBridge types.XChainBridge
}

// Returns the type of the transaction.
func (x *XChainAccountCreateCommit) TxType() TxType {
	return XChainAccountCreateCommitTx
}

// Returns a flattened version of the transaction.
func (x *XChainAccountCreateCommit) Flatten() FlatTransaction {
	flatTx := x.BaseTx.Flatten()

	flatTx["TransactionType"] = x.TxType().String()

	if x.Amount != nil {
		flatTx["Amount"] = x.Amount.Flatten()
	}

	if x.Destination != "" {
		flatTx["Destination"] = x.Destination.String()
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
func (x *XChainAccountCreateCommit) Validate() (bool, error) {
	_, err := x.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if ok, err := IsAmount(x.Amount, "Amount", true); !ok {
		return false, err
	}

	if !addresscodec.IsValidClassicAddress(x.Destination.String()) {
		return false, ErrInvalidAccount
	}

	if ok, err := IsAmount(x.SignatureReward, "SignatureReward", false); !ok {
		return false, err
	}

	return x.XChainBridge.Validate()
}
