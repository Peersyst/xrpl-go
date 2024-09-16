package transactions

import (
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type PaymentChannelFund struct {
	BaseTx
	Channel    types.Hash256
	Amount     types.XRPCurrencyAmount
	Expiration uint `json:",omitempty"`
}

func (*PaymentChannelFund) TxType() TxType {
	return PaymentChannelFundTx
}

// TODO: Implement flatten
func (s *PaymentChannelFund) Flatten() FlatTransaction {
	return nil
}

func ValidatePaymentChannelFund(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "Channel", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "Amount", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "Expiration", typecheck.IsUint)
	if err != nil {
		return err
	}

	return nil
}
