package transactions

import (
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type PaymentChannelCreate struct {
	BaseTx
	Amount         types.XRPCurrencyAmount
	Destination    types.Address
	SettleDelay    uint
	PublicKey      string
	CancelAfter    uint `json:",omitempty"`
	DestinationTag uint `json:",omitempty"`
}

func (*PaymentChannelCreate) TxType() TxType {
	return PaymentChannelCreateTx
}

// TODO: Implement flatten
func (s *PaymentChannelCreate) Flatten() FlatTransaction {
	return nil
}

func ValidatePaymentChannelCreate(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "Amount", typecheck.IsString)
	if err != nil {
		return err
	}

	// TODO: update to IsAccount when it's implemented
	err = ValidateRequiredField(tx, "Destination", typecheck.IsString)
	if err != nil {
		return err
	}

	// TODO: check if IsUint is correct or if it should be IsInt
	err = ValidateRequiredField(tx, "SettleDelay", typecheck.IsUint)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "PublicKey", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "CancelAfter", typecheck.IsUint)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "DestinationTag", typecheck.IsUint)
	if err != nil {
		return err
	}

	return nil
}
