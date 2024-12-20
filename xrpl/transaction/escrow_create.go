package transaction

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

var (
	ErrEscrowCreateInvalidDestinationAddress   = errors.New("escrow create: invalid destination address")
	ErrEscrowCreateNoCancelOrFinishAfterSet    = errors.New("escrow create: either CancelAfter or FinishAfter must be set")
	ErrEscrowCreateNoConditionOrFinishAfterSet = errors.New("escrow create: either Condition or FinishAfter must be specified")
)

// Sequester XRP until the escrow process either finishes or is canceled.
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
// ```
type EscrowCreate struct {
	BaseTx
	// Amount of XRP, in drops, to deduct from the sender's balance and escrow. Once escrowed, the XRP can either go to the Destination address (after the FinishAfter time) or returned to the sender (after the CancelAfter time).
	Amount types.XRPCurrencyAmount
	// Address to receive escrowed XRP.
	Destination types.Address
	// (Optional) The time, in seconds since the Ripple Epoch, when this escrow expires. This value is immutable; the funds can only be returned to the sender after this time.
	CancelAfter uint32 `json:",omitempty"`
	// (Optional) The time, in seconds since the Ripple Epoch, when the escrowed XRP can be released to the recipient. This value is immutable, and the funds can't be accessed until this time.
	FinishAfter uint32 `json:",omitempty"`
	// (Optional) Hex value representing a PREIMAGE-SHA-256 crypto-condition. The funds can only be delivered to the recipient if this condition is fulfilled. If the condition is not fulfilled before the expiration time specified in the CancelAfter field, the XRP can only revert to the sender.
	Condition string `json:",omitempty"`
	// (Optional) Arbitrary tag to further specify the destination for this escrowed payment, such as a hosted recipient at the destination address.
	DestinationTag uint32 `json:",omitempty"`
}

// TxType returns the transaction type for this transaction (EscrowCreate).
func (*EscrowCreate) TxType() TxType {
	return EscrowCreateTx
}

// Flatten returns the flattened map of the EscrowCreate transaction.
func (e *EscrowCreate) Flatten() FlatTransaction {
	flattened := e.BaseTx.Flatten()

	flattened["TransactionType"] = "EscrowCreate"

	flattened["Amount"] = e.Amount.Flatten()

	if e.Destination != "" {
		flattened["Destination"] = e.Destination
	}
	if e.CancelAfter != 0 {
		flattened["CancelAfter"] = e.CancelAfter
	}
	if e.FinishAfter != 0 {
		flattened["FinishAfter"] = e.FinishAfter
	}
	if e.Condition != "" {
		flattened["Condition"] = e.Condition
	}
	if e.DestinationTag != 0 {
		flattened["DestinationTag"] = e.DestinationTag
	}

	return flattened
}

// Validates the EscrowCreate transaction and makes sure all the fields are correct.
func (e *EscrowCreate) Validate() (bool, error) {
	ok, err := e.BaseTx.Validate()
	if err != nil || !ok {
		return false, err
	}

	if !addresscodec.IsValidClassicAddress(e.Destination.String()) {
		return false, ErrInvalidDestinationAddress
	}

	if (e.FinishAfter == 0 && e.CancelAfter == 0) || (e.Condition == "" && e.FinishAfter == 0) {
		return false, ErrEscrowCreateNoConditionOrFinishAfterSet
	}

	return true, nil
}
