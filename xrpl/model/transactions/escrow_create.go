package transactions

import (
	"errors"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type EscrowCreate struct {
	BaseTx
	Amount         types.XRPCurrencyAmount
	Destination    types.Address
	CancelAfter    uint   `json:",omitempty"`
	FinishAfter    uint   `json:",omitempty"`
	Condition      string `json:",omitempty"`
	DestinationTag uint   `json:",omitempty"`
}

func (*EscrowCreate) TxType() TxType {
	return EscrowCreateTx
}

// TODO: Implement flatten
func (s *EscrowCreate) Flatten() FlatTransaction {
	return nil
}

func ValidateEscrowCreate(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "Amount", typecheck.IsString)
	if err != nil {
		return err
	}

	// TODO: update to IsAccount when that function exists
	err = ValidateRequiredField(tx, "Destination", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "DestinationTag", typecheck.IsUint)
	if err != nil {
		return err
	}

	_, hasCancelAfter := tx["CancelAfter"]
	_, hasFinishAfter := tx["FinishAfter"]

	if !hasCancelAfter && !hasFinishAfter {
		return errors.New("EscrowCreate: must provide either CancelAfter or FinishAfter field")
	}

	ValidateOptionalField(tx, "CancelAfter", typecheck.IsUint)
	ValidateOptionalField(tx, "FinishAfter", typecheck.IsUint)
	ValidateOptionalField(tx, "Condition", typecheck.IsString)
	return nil
}
