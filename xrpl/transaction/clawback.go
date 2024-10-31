package transaction

import (
	"errors"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
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

// Validates the Clawback struct.
func (c *Clawback) Validate() (bool, error) {
	// validate the base transaction
	_, err := c.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	// check if the field Amount is set
	if c.Amount == nil {
		return false, errors.New("clawback: missing field Amount")
	}

	// check if the Amount is a valid currency amount
	if ok, _ := IsIssuedCurrency(c.Amount); !ok {
		return false, errors.New("clawback: invalid Amount")
	}

	// check if Account is not the same as the issuer
	if c.Account.String() == c.Amount.Flatten().(map[string]interface{})["issuer"] {
		return false, errors.New("clawback: Account and Amount.issuer cannot be the same")
	}

	return true, nil
}
