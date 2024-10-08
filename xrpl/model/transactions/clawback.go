package transactions

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

// Requires the Clawback amendment.
// Claw back tokens issued by your account.
// Clawback is disabled by default. To use clawback, you must send an AccountSet transaction to enable the Allow Trust Line Clawback setting.
// An issuer with any existing tokens cannot enable Clawback. You can only enable Allow Trust Line Clawback if you have a completely empty owner directory,
// meaning you must do so before you set up any trust lines, offers, escrows, payment channels, checks, or signer lists. After you enable Clawback,
// it cannot reverted: the account permanently gains the ability to claw back issued assets on trust lines.
type Clawback struct {
	// Base transaction fields
	BaseTx

	// Indicates the amount being clawed back, as well as the counterparty from which the amount is being clawed back.
	// The quantity to claw back, in the value sub-field, must not be zero. If this is more than the current balance,
	// the transaction claws back the entire balance. The sub-field issuer within Amount represents the token holder's
	// account ID, rather than the issuer's.
	Amount types.CurrencyAmount
}

func (*Clawback) TxType() TxType {
	return ClawbackTx
}

func (s *Clawback) Flatten() FlatTransaction {
	flattened := s.BaseTx.Flatten()

	flattened["TransactionType"] = "Clawback"

	if s.Amount != nil {
		flattened["Amount"] = s.Amount.Flatten()
	}

	return flattened
}

// UnmarshalJSON unmarshals the JSON data into a Clawback struct.
func (c *Clawback) UnmarshalJSON(data []byte) error {
	// Define a helper struct to hold the unmarshaled data
	type cHelper struct {
		BaseTx
		Amount json.RawMessage
	}

	var h cHelper

	// Unmarshal the JSON data into the helper struct
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}

	// Assign the values from the helper struct to the Clawback struct
	*c = Clawback{
		BaseTx: h.BaseTx,
	}

	// Unmarshal the Amount field into a CurrencyAmount struct
	var amount types.CurrencyAmount
	amount, err = types.UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return err
	}

	// Assign the unmarshaled CurrencyAmount to the Clawback struct
	c.Amount = amount

	return nil
}
