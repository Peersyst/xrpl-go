package transactions

import (
	"errors"
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

	if tx["Amount"] == nil {
		return errors.New("AMMCreate: missing field Amount")
	}

	if !IsAmount(tx["Amount"]) {
		return errors.New("AMMCreate: Amount must be an Amount")
	}

	if tx["Amount2"] == nil {
		return errors.New("AMMCreate: missing field Amount2")
	}

	if !IsAmount(tx["Amount2"]) {
		return errors.New("AMMCreate: Amount2 must be an Amount")
	}

	if tx["TradingFee"] == nil {
		return errors.New("AMMCreate: missing field TradingFee")
	}

	// TODO: Check later if the type check is correct
	if !typecheck.IsInt(tx["TradingFee"]) {
		return errors.New("AMMCreate: TradingFee must be a number")
	}

	if tx["TradingFee"].(int) < 0 || tx["TradingFee"].(int) > AMM_MAX_TRADING_FEE {
		return fmt.Errorf("AMMCreate: TradingFee must be between 0 and %v", AMM_MAX_TRADING_FEE)
	}

	return nil
}
