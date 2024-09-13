package transactions

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
)

const AMM_MAX_TRADING_FEE = 1000

type AMMCreate struct {
	BaseTx
}

func (*AMMCreate) TxType() TxType {
	return AMMCreateTx
}

// TODO: Implement flatten
func (s *AMMCreate) Flatten() FlatTransaction {
	return nil
}

func ValidateAMMCreate(tx FlatTransaction) error {
	// Validate base transaction fields
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "Amount", IsAmount)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "Amount2", IsAmount)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "TradingFee", typecheck.IsInt)
	if err != nil {
		return err
	}

	if tx["TradingFee"].(int) < 0 || tx["TradingFee"].(int) > AMM_MAX_TRADING_FEE {
		return fmt.Errorf("AMMCreate: TradingFee must be between 0 and %v", AMM_MAX_TRADING_FEE)
	}

	return nil
}
