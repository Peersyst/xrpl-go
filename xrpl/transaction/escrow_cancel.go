package transaction

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// Return escrowed XRP to the sender.
//
// Example:
//
// ```json
//
//	{
//	    "Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
//	    "TransactionType": "EscrowCancel",
//	    "Owner": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
//	    "OfferSequence": 7,
//	}
//
// ```
type EscrowCancel struct {
	BaseTx
	// Address of the source account that funded the escrow payment.
	Owner types.Address
	// Transaction sequence (or Ticket number) of EscrowCreate transaction that created the escrow to cancel.
	OfferSequence uint
}

// TxType returns the transaction type for this transaction (EscrowCancel).
func (*EscrowCancel) TxType() TxType {
	return EscrowCancelTx
}

// Flatten returns the flattened map of the EscrowCancel transaction.
func (e *EscrowCancel) Flatten() FlatTransaction {
	flattened := e.BaseTx.Flatten()

	flattened["TransactionType"] = "EscrowCancel"

	if e.Owner != "" {
		flattened["Owner"] = e.Owner
	}

	if e.OfferSequence != 0 {
		flattened["OfferSequence"] = e.OfferSequence
	}

	return flattened
}

// Validate checks if the EscrowCancel struct is valid.
func (e *EscrowCancel) Validate() (bool, error) {
	_, err := e.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if !addresscodec.IsValidClassicAddress(e.Owner.String()) {
		return false, errors.New("invalid xrpl address for the Owner field")
	}

	if e.OfferSequence == 0 {
		return false, errors.New("missing OfferSequence")
	}

	return true, nil
}
