package transaction

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// The maximum value is 1000, indicating a 1% fee. The minimum value is 0. https://xrpl.org/docs/references/protocol/transactions/types/ammcreate#ammcreate-fields
const AMM_MAX_TRADING_FEE = 1000

type AMMCreate struct {
	BaseTx
	// The first of the two assets to fund this AMM with. This must be a positive amount.
	Amount types.CurrencyAmount
	// The second of the two assets to fund this AMM with. This must be a positive amount.
	Amount2 types.CurrencyAmount
	// The fee to charge for trades against this AMM instance, in units of 1/100,000; a value of 1 is equivalent to 0.001%. The maximum value is 1000, indicating a 1% fee. The minimum value is 0.
	TradingFee uint16
}

func (*AMMCreate) TxType() TxType {
	return AMMCreateTx
}

func (s *AMMCreate) Flatten() FlatTransaction {
	// Add BaseTx fields
	flattened := s.BaseTx.Flatten()

	// Add AMMCreate-specific fields
	flattened["TransactionType"] = "AMMCreate"

	if s.Amount != nil {
		flattened["Amount"] = s.Amount.Flatten()
	}

	if s.Amount2 != nil {
		flattened["Amount2"] = s.Amount2.Flatten()
	}

	flattened["TradingFee"] = s.TradingFee

	return flattened
}

func (a *AMMCreate) Validate() (bool, error) {
	_, err := a.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if ok, err := IsAmount(IsAmountArgs{field: a.Amount, fieldName: "Amount", isFieldRequired: true}); !ok {
		return false, err
	}

	if ok, err := IsAmount(IsAmountArgs{field: a.Amount2, fieldName: "Amount2", isFieldRequired: true}); !ok {
		return false, err
	}

	if a.TradingFee > AMM_MAX_TRADING_FEE {
		return false, fmt.Errorf("trading fee is too high, max value is %d", AMM_MAX_TRADING_FEE)
	}

	return true, nil
}
