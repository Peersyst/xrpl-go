package transaction

import (
	"encoding/json"
	"errors"
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

// UnmarshalJSON unmarshals the Payment transaction from JSON.
func (p *AMMCreate) UnmarshalJSON(data []byte) error {
	type pHelper struct {
		BaseTx
		Amount     json.RawMessage
		Amount2    json.RawMessage
		TradingFee uint16
	}
	var h pHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*p = AMMCreate{
		BaseTx:     h.BaseTx,
		TradingFee: h.TradingFee,
	}
	var amount, amount2 types.CurrencyAmount
	var err error

	amount, err = types.UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return err
	}
	amount2, err = types.UnmarshalCurrencyAmount(h.Amount2)
	if err != nil {
		return err
	}

	p.Amount = amount
	p.Amount2 = amount2

	return nil
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

	if a.Amount == nil {
		return false, errors.New("missing field Amount")
	}

	if !IsAmount(a.Amount) {
		return false, errors.New("invalid field Amount")
	}

	if a.Amount2 == nil {
		return false, errors.New("missing field Amount2")
	}

	if !IsAmount(a.Amount2) {
		return false, errors.New("invalid field Amount2")
	}

	if a.TradingFee > AMM_MAX_TRADING_FEE {
		return false, fmt.Errorf("trading fee is too high, max value is %d", AMM_MAX_TRADING_FEE)
	}

	return true, nil
}
