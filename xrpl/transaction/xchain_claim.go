package transaction

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	"github.com/Peersyst/xrpl-go/v1/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

var (
	ErrInvalidDestinationAddress = errors.New("xchainClaim: invalid destination address")
	ErrMissingXChainClaimID      = errors.New("xchainClaim: missing XChainClaimID")
)

// (Requires the XChainBridge amendment)
// The XChainClaim transaction completes a cross-chain transfer of value. It allows a user to claim the value on the
// destination chain - the equivalent of the value locked on the source chain. A user can only claim the value if they
// own the cross-chain claim ID associated with the value locked on the source chain (the Account field).
// The user can send the funds to anyone (the Destination field). This transaction is only needed if an
// OtherChainDestination isn't specified in the XChainCommit transaction, or if something goes wrong with the automatic transfer of funds.
//
// If the transaction succeeds in moving funds, the referenced XChainOwnedClaimID ledger object will be destroyed.
// This prevents transaction replay. If the transaction fails, the XChainOwnedClaimID won't be
// destroyed and the transaction can be re-run with different parameters.
//
// ```json
//
//	{
//	  "Account": "rahDmoXrtPdh7sUdrPjini3gcnTVYjbjjw",
//	  "Amount": "10000",
//	  "TransactionType": "XChainClaim",
//	  "XChainClaimID": "13f",
//	  "Destination": "rahDmoXrtPdh7sUdrPjini3gcnTVYjbjjw",
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
type XChainClaim struct {
	BaseTx

	// The amount to claim on the destination chain. This must match the amount attested to on the
	// attestations associated with this XChainClaimID.
	Amount types.CurrencyAmount
	// The destination account on the destination chain. It must exist or the transaction will fail.
	// However, if the transaction fails in this case, the sequence number and collected signatures
	// won't be destroyed, and the transaction can be rerun with a different destination.
	Destination types.Address
	// An integer destination tag.
	DestinationTag uint32 `json:",omitempty"`
	// The bridge to use for the transfer.
	XChainBridge types.XChainBridge
	// The unique integer ID for the cross-chain transfer that was referenced in the corresponding XChainCommit transaction.
	XChainClaimID string
}

// TxType returns the type of the transaction.
func (x *XChainClaim) TxType() TxType {
	return XChainClaimTx
}

// Flatten returns a flattened version of the transaction.
func (x *XChainClaim) Flatten() FlatTransaction {
	flatTx := x.BaseTx.Flatten()

	flatTx["TransactionType"] = x.TxType()

	if x.Amount != nil {
		flatTx["Amount"] = x.Amount.Flatten()
	}

	if x.Destination != "" {
		flatTx["Destination"] = x.Destination.String()
	}

	if x.DestinationTag != 0 {
		flatTx["DestinationTag"] = x.DestinationTag
	}

	if x.XChainBridge != (types.XChainBridge{}) {
		flatTx["XChainBridge"] = x.XChainBridge.Flatten()
	}

	if x.XChainClaimID != "" {
		flatTx["XChainClaimID"] = x.XChainClaimID
	}

	return flatTx
}

// Validate validates the transaction.
func (x *XChainClaim) Validate() (bool, error) {
	_, err := x.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if ok, err := IsAmount(x.Amount, "Amount", true); !ok {
		return false, err
	}

	if !addresscodec.IsValidClassicAddress(x.Destination.String()) {
		return false, ErrInvalidDestinationAddress
	}

	if x.XChainClaimID == "" || !typecheck.IsHex(x.XChainClaimID) {
		return false, ErrMissingXChainClaimID
	}

	if ok, err := x.XChainBridge.Validate(); !ok {
		return false, err
	}

	return true, nil
}
