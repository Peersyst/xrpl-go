package transaction

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

const (
	// Claw back the specified amount of Asset, and a corresponding amount of Asset2 based on
	// the AMM pool's asset proportion; both assets must be issued by the issuer in the Account
	// field. If this flag isn't enabled, the issuer claws back the specified amount of Asset,
	// while a corresponding proportion of Asset2 goes back to the Holder.
	tfClawTwoAssets uint32 = 0x00000001
)

var (
	// ErrInvalidHolder is returned when the holder is invalid.
	ErrInvalidHolder = errors.New("invalid holder")
	// ErrInvalidAmountIssuer is returned when the amount issuer is invalid.
	ErrInvalidAmountIssuer = errors.New("invalid amount issuer")
)

// Claw back tokens from a holder who has deposited your issued tokens into an AMM pool.
// Clawback is disabled by default. To use clawback, you must send an AccountSet transaction
// to enable the Allow Trust Line Clawback setting. An issuer with any existing tokens cannot
// enable clawback. You can only enable Allow Trust Line Clawback if you have a completely empty
// owner directory, meaning you must do so before you set up any trust lines, offers, escrows,
// payment channels, checks, or signer lists. After you enable clawback, it cannot reverted:
// the account permanently gains the ability to claw back issued assets on trust lines.
// (Added by the AMMClawback amendment)
//
// ```json
//
//	{
//	  "TransactionType": "AMMClawback",
//	  "Account": "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
//	  "Holder": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
//	  "Asset": {
//	      "currency" : "FOO",
//	      "issuer" : "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL"
//	  },
//	  "Asset2" : {
//	      "currency" : "BAR",
//	      "issuer" : "rHtptZx1yHf6Yv43s1RWffM3XnEYv3XhRg"
//	  },
//	  "Amount": {
//	      "currency" : "FOO",
//	      "issuer" : "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
//	      "value" : "1000"
//	  }
//	}
//
// ```
type AMMClawback struct {
	BaseTx
	// The account holding the asset to be clawed back.
	Holder string
	// Specifies the asset that the issuer wants to claw back from the AMM pool.
	// The asset can be XRP, a token, or an MPT (see: Specifying Without Amounts).
	// The issuer field must match with Account.
	Asset types.IssuedCurrency
	// Specifies the other asset in the AMM's pool. The asset can be XRP, a token,
	// or an MPT (see: Specifying Without Amounts).
	Asset2 types.CurrencyAmount `json:",omitempty"`
	// The maximum amount to claw back from the AMM account. The currency and issuer subfields
	// should match the Asset subfields. If this field isn't specified, or the value subfield
	// exceeds the holder's available tokens in the AMM, all of the holder's tokens are clawed back.
	Amount types.IssuedCurrencyAmount `json:",omitempty"`
}

// Returns the type of the transaction.
func (a *AMMClawback) TxType() TxType {
	return AMMClawbackTx
}

// Returns the flattened transaction.
func (a *AMMClawback) Flatten() FlatTransaction {
	flattened := a.BaseTx.Flatten()

	if a.Holder != "" {
		flattened["Holder"] = a.Holder
	}

	if a.Asset != (types.IssuedCurrency{}) {
		flattened["Asset"] = a.Asset.Flatten()
	}

	if a.Asset2 != nil {
		flattened["Asset2"] = a.Asset2.Flatten()
	}

	if a.Amount != (types.IssuedCurrencyAmount{}) {
		flattened["Amount"] = a.Amount.Flatten()
	}

	return flattened
}

// Validates the transaction.
func (a *AMMClawback) Validate() (bool, error) {
	_, err := a.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if !addresscodec.IsValidAddress(a.Holder) {
		return false, ErrInvalidHolder
	}

	if a.Asset != (types.IssuedCurrency{}) {
		if !addresscodec.IsValidAddress(a.Asset.Issuer) {
			return false, ErrInvalidAssetIssuer
		}
	}

	if a.Amount != (types.IssuedCurrencyAmount{}) {
		if !addresscodec.IsValidAddress(a.Amount.Issuer.String()) {
			return false, ErrInvalidAmountIssuer
		}
	}

	return true, nil
}

// Sets the clawback flag for two assets.
func (a *AMMClawback) SetClawTwoAssets() {
	a.Flags |= tfClawTwoAssets
}
