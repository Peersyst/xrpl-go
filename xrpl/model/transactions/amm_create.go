package transactions

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

const AMM_MAX_TRADING_FEE = 1000

type AMMCreate struct {
	BaseTx

	/**
	 * The first of the two assets to fund this AMM with. This must be a positive amount.
	 */
	Amount types.CurrencyAmount

	/**
	* The second of the two assets to fund this AMM with. This must be a positive amount.
	 */
	Amount2 types.CurrencyAmount

	/**
	* The fee to charge for trades against this AMM instance, in units of 1/100,000; a value of 1 is equivalent to 0.001%.
	* The maximum value is 1000, indicating a 1% fee.
	* The minimum value is 0.
	 */
	TradingFee uint16
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
