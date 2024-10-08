package transaction

import (
	"encoding/json"
	"errors"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

const (
	// Authorize the other party to hold currency issued by this account. (No
	// effect unless using the asfRequireAuth AccountSet flag.) Cannot be unset.
	tfSetAuth uint = 65536 // 0x00010000
	// Enable the No Ripple flag, which blocks rippling between two trust lines.
	// of the same currency if this flag is enabled on both.
	tfSetNoRipple uint = 131072 // 0x00020000
	// Disable the No Ripple flag, allowing rippling on this trust line.
	tfClearNoRipple uint = 262144 // 0x00040000
	// Freeze the trust line
	tfSetFreeze uint = 1048576 // 0x00100000
	// Unfreeze the trust line
	tfClearFreeze uint = 2097152 // 0x00200000
)

// Create or modify a trust line linking two accounts.
type TrustSet struct {
	// Base transaction fields
	BaseTx
	//Object defining the trust line to create or modify, in the format of a Currency Amount.
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

func (t *TrustSet) UnmarshalJSON(data []byte) error {
	type tsHelper struct {
		BaseTx
		LimitAmount json.RawMessage
		QualityIn   uint32 `json:",omitempty"`
		QualityOut  uint32 `json:",omitempty"`
	}
	var h tsHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*t = TrustSet{
		BaseTx:     h.BaseTx,
		QualityIn:  h.QualityIn,
		QualityOut: h.QualityOut,
	}
	limit, err := types.UnmarshalCurrencyAmount(h.LimitAmount)
	if err != nil {
		return err
	}
	t.LimitAmount = limit

	return nil
}

// Validates the TrustSet transaction.
func (tx *TrustSet) Validate() (bool, error) {
	// Validate the base transaction
	_, err := tx.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	// Check if the field LimitAmount is set
	if tx.LimitAmount == nil {
		return false, errors.New("trustSet: missing field LimitAmount")
	}

	if !IsAmount(tx.LimitAmount) {
		return false, errors.New("trustSet: invalid LimitAmount")
	}

	// Check if QualityIn is a number
	if tx.QualityIn != 0 && !typecheck.IsUint32(tx.QualityIn) {
		return false, errors.New("trustSet: QualityIn must be a number")
	}

	// Check if QualityOut is a number
	if tx.QualityOut != 0 && !typecheck.IsUint32(tx.QualityOut) {
		return false, errors.New("trustSet: QualityOut must be a number")
	}

	return true, nil
}
