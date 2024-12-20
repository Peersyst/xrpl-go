package transaction

import (
	"errors"

	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

var (
	ErrTrustSetMissingLimitAmount  = errors.New("missing field LimitAmount")
	ErrTrustSetQualityInNotNumber  = errors.New("QualityIn must be a number")
	ErrTrustSetQualityOutNotNumber = errors.New("QualityOut must be a number")
)

const (
	// Authorize the other party to hold currency issued by this account. (No
	// effect unless using the asfRequireAuth AccountSet flag.) Cannot be unset.
	tfSetAuth uint32 = 0x00010000
	// Enable the No Ripple flag, which blocks rippling between two trust lines.
	// of the same currency if this flag is enabled on both.
	tfSetNoRipple uint32 = 0x00020000
	// Disable the No Ripple flag, allowing rippling on this trust line.
	tfClearNoRipple uint32 = 0x00040000
	// Freeze the trust line.
	tfSetFreeze uint32 = 0x00100000
	// Unfreeze the trust line.
	tfClearFreeze uint32 = 0x00200000
)

// Create or modify a trust line linking two accounts.
type TrustSet struct {
	// Base transaction fields
	BaseTx
	// Object defining the trust line to create or modify, in the format of a Currency Amount.
	LimitAmount types.CurrencyAmount
	// (Optional) Value incoming balances on this trust line at the ratio of this number per 1,000,000,000 units.
	// A value of 0 is shorthand for treating balances at face value. For example, if you set the value to 10,000,000, 1% of incoming funds remain with the sender.
	// If an account sends 100 currency, the sender retains 1 currency unit and the destination receives 99 units. This option is included for parity: in practice, you are much more likely to set a QualityOut value.
	// Note that this fee is separate and independent from token transfer fees.
	QualityIn uint32 `json:",omitempty"`
	// (Optional) Value outgoing balances on this trust line at the ratio of this number per 1,000,000,000 units.
	// A value of 0 is shorthand for treating balances at face value. For example, if you set the value to 10,000,000, 1% of outgoing funds would remain with the issuer.
	// If the sender sends 100 currency units, the issuer retains 1 currency unit and the destination receives 99 units. Note that this fee is separate and independent from token transfer fees.
	QualityOut uint32 `json:",omitempty"`
}

// TxType returns the type of the transaction (TrustSet).
func (*TrustSet) TxType() TxType {
	return TrustSetTx
}

// Flatten returns a flattened map of the TrustSet transaction.
func (t *TrustSet) Flatten() FlatTransaction {
	flattened := t.BaseTx.Flatten()

	flattened["TransactionType"] = "TrustSet"

	if t.LimitAmount != nil {
		flattened["LimitAmount"] = t.LimitAmount.Flatten()
	}
	if t.QualityIn != 0 {
		flattened["QualityIn"] = t.QualityIn
	}
	if t.QualityOut != 0 {
		flattened["QualityOut"] = t.QualityOut
	}

	return flattened
}

// Set the SetAuth flag
//
// SetAuth: Authorize the other party to hold currency issued by this account. (No
// effect unless using the asfRequireAuth AccountSet flag.) Cannot be unset.
func (t *TrustSet) SetSetAuthFlag() {
	t.Flags |= tfSetAuth
}

// Set the SetNoRipple flag
//
// SetNoRipple: Enable the No Ripple flag, which blocks rippling between two trust lines.
// of the same currency if this flag is enabled on both.
func (t *TrustSet) SetSetNoRippleFlag() {
	t.Flags |= tfSetNoRipple
}

// Set the ClearNoRipple flag
//
// ClearNoRipple: Disable the No Ripple flag, allowing rippling on this trust line.
func (t *TrustSet) SetClearNoRippleFlag() {
	t.Flags |= tfClearNoRipple
}

// Set the SetFreeze flag
//
// SetFreeze: Freeze the trust line
func (t *TrustSet) SetSetFreezeFlag() {
	t.Flags |= tfSetFreeze
}

// Set the ClearFreeze flag
//
// ClearFreeze: Unfreeze the trust line
func (t *TrustSet) SetClearFreezeFlag() {
	t.Flags |= tfClearFreeze
}

// Validates the TrustSet transaction.
func (t *TrustSet) Validate() (bool, error) {
	// Validate the base transaction
	_, err := t.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	// Check if the field LimitAmount is set
	if t.LimitAmount == nil {
		return false, ErrTrustSetMissingLimitAmount
	}

	if ok, err := IsAmount(t.LimitAmount, "LimitAmount", true); !ok {
		return false, err
	}

	return true, nil
}
