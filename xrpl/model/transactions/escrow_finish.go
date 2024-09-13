package transactions

import (
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type EscrowFinish struct {
	BaseTx
	Owner         types.Address
	OfferSequence uint
	Condition     string `json:",omitempty"`
	Fulfillment   string `json:",omitempty"`
}

func (*EscrowFinish) TxType() TxType {
	return EscrowFinishTx
}

// TODO: Implement flatten
func (s *EscrowFinish) Flatten() FlatTransaction {
	return nil
}

func ValidateEscrowFinish(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "Owner", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "OfferSequence", typecheck.IsUint)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "Condition", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "Fulfillment", typecheck.IsString)
	if err != nil {
		return err
	}

	return nil
}
