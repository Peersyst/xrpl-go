package transactions

import (
	"encoding/json"
	"errors"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type CheckCash struct {
	BaseTx

	// The ID of the Check ledger object to cash as a 64-character hexadecimal string.
	CheckID types.Hash256

	// Redeem the Check for exactly this amount, if possible. The currency must
	// match that of the SendMax of the corresponding CheckCreate transaction. You.
	// must provide either this field or DeliverMin.
	Amount types.CurrencyAmount `json:",omitempty"`

	//Redeem the Check for at least this amount and for as much as possible. The
	//currency must match that of the SendMax of the corresponding CheckCreate.
	//transaction. You must provide either this field or Amount.
	DeliverMin types.CurrencyAmount `json:",omitempty"`
}

func (*CheckCash) TxType() TxType {
	return CheckCashTx
}

// TODO: Implement flatten
func (s *CheckCash) Flatten() FlatTransaction {
	return nil
}

func (tx *CheckCash) UnmarshalJSON(data []byte) error {
	type ccHelper struct {
		BaseTx
		CheckID    types.Hash256
		Amount     json.RawMessage `json:",omitempty"`
		DeliverMin json.RawMessage `json:",omitempty"`
	}
	var h ccHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*tx = CheckCash{
		BaseTx:  h.BaseTx,
		CheckID: h.CheckID,
	}

	var amount, min types.CurrencyAmount
	var err error
	amount, err = types.UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return err
	}
	min, err = types.UnmarshalCurrencyAmount(h.DeliverMin)
	if err != nil {
		return err
	}
	tx.Amount = amount
	tx.DeliverMin = min
	return nil

}

func ValidateCheckCash(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	_, hasCheckID := tx["CheckID"]
	_, hasAmount := tx["Amount"]
	_, hasDeliverMin := tx["DeliverMin"]

	// Check if either Amount or DeliverMin are set
	if !hasAmount && !hasDeliverMin {
		return errors.New("CheckCash: must provide either Amount or DeliverMin")
	}

	// Check that both Amount and DeliverMin are not set
	if hasAmount && hasDeliverMin {
		return errors.New("CheckCash: cannot have both Amount and DeliverMin")
	}

	// Check if the field Amount is an Amount
	if hasAmount && !IsAmount(tx["Amount"]) {
		return errors.New("CheckCash: Amount must be an Amount - https://xrpl.org/docs/references/protocol/data-types/basic-data-types#specifying-currency-amounts")
	}

	// Check if the field DeliverMin is an Amount
	if hasDeliverMin && !IsAmount(tx["DeliverMin"]) {
		return errors.New("CheckCash: DeliverMin must be an Amount - https://xrpl.org/docs/references/protocol/data-types/basic-data-types#specifying-currency-amounts")
	}

	// Check if the field CheckID is set
	if hasCheckID && !typecheck.IsString(tx["CheckID"]) {
		return errors.New("CheckCash: CheckID must be a string")
	}

	return nil
}
