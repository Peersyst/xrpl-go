package transactions

import (
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type EscrowCancel struct {
	BaseTx
	// Address of the source account that funded the escrow payment.
	Owner types.Address
	// Transaction sequence (or Ticket  number) of EscrowCreate transaction that created the escrow to cancel.
	OfferSequence uint
}

func (*EscrowCancel) TxType() TxType {
	return EscrowCancelTx
}

// TODO: Implement flatten
func (s *EscrowCancel) Flatten() FlatTransaction {
	return nil
}

func ValidateEscrowCancel(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	// TODO: update to IsAccount when that function exists
	ValidateRequiredField(tx, "Owner", typecheck.IsString)

	ValidateRequiredField(tx, "OfferSequence", typecheck.IsUint)

	return nil
}
