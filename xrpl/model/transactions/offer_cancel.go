package transactions

import "github.com/Peersyst/xrpl-go/pkg/typecheck"

type OfferCancel struct {
	BaseTx
	OfferSequence uint
}

func (*OfferCancel) TxType() TxType {
	return OfferCancelTx
}

// TODO: Implement flatten
func (s *OfferCancel) Flatten() FlatTransaction {
	return nil
}

func ValidateOfferCancel(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "OfferSequence", typecheck.IsUint)
	if err != nil {
		return err
	}

	return nil
}
