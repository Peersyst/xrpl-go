package transaction

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	ledger "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

var (
	ErrHolderMatchesIssuer    = errors.New("AMMClawback: Holder and Asset.issuer must be distinct")
	ErrAccountMismatch        = errors.New("AMMClawback: Account must be the same as Asset.issuer")
	ErrAmountCurrencyMismatch = errors.New("AMMClawback: Amount.currency must match Asset.currency")
	ErrAmountIssuerMismatch   = errors.New("AMMClawback: Amount.issuer must match Asset.issuer")
)

// Claw back tokens from a holder who has deposited your issued tokens into an AMM pool.

// Clawback is disabled by default. To use clawback, you must send an AccountSet transaction to enable the Allow Trust Line Clawback setting. 
// An issuer with any existing tokens cannot enable clawback. You can only enable Allow Trust Line Clawback if you have a completely empty owner directory,
// meaning you must do so before you set up any trust lines, offers, escrows, payment channels, checks, or signer lists. 
// After you enable clawback, it cannot reverted: the account permanently gains the ability to claw back issued assets on trust lines.
// Example:
// {
//   "TransactionType": "AMMClawback",
//   "Account": "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
//   "Holder": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
//   "Asset": {
//       "currency" : "FOO",
//       "issuer" : "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL"
//   },
//   "Asset2" : {
//       "currency" : "BAR",
//       "issuer" : "rHtptZx1yHf6Yv43s1RWffM3XnEYv3XhRg"
//   },
//   "Amount": {
//       "currency" : "FOO",
//       "issuer" : "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
//       "value" : "1000"
//   }
// }

type AMMClawback struct {
	BaseTx

	// The issuer of the asset being clawed back. Only the issuer can submit this transaction.
	Account types.Address
	// Holder is the account holding the asset to be clawed back.
	Holder types.Address
	// The definition for one of the assets in the AMM's pool. In JSON, this is an object with currency and issuer fields (omit issuer for XRP).
	Asset ledger.Asset
	// The definition for the other asset in the AMM's pool. In JSON, this is an object with currency and issuer fields (omit issuer for XRP).
	Asset2 ledger.Asset
	// Amount is the maximum amount to claw back. It is optional.
	// If omitted, or if the value exceeds the holder's balance, all tokens will be clawed back.
	Amount *types.IssuedCurrencyAmount `json:",omitempty"`
}

const (
	// Perform a special double-asset deposit to an AMM with an empty pool.
	tfClawTwoAssets uint32 = 1
)


func (a *AMMClawback) SetTwoAssetFlag() {
	a.Flags |= tfClawTwoAssets
}


// TxType returns the transaction type.
func (*AMMClawback) TxType() TxType {
	return AMMClawbackTx
}

// Flatten returns a map representation of the AMMClawback transaction.
// This is useful for serialization.
func (a *AMMClawback) Flatten() FlatTransaction {
	flattened := a.BaseTx.Flatten()
	flattened["TransactionType"] = "AMMClawback"


	flattened["Account"] = a.Account.String()

	flattened["Holder"] = a.Holder.String()

	flattened["Asset"] = a.Asset.Flatten()

	flattened["Asset2"] = a.Asset2.Flatten()

	if !a.Amount.IsZero() {
		flattened["Amount"] = a.Amount.Flatten()
	}
	return flattened
}


func (a *AMMClawback) Validate() (bool, error) {
	_, err := a.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if !addresscodec.IsValidAddress(a.Holder.String()) {
		return false, ErrInvalidAccount
	}

	if a.Holder == a.Asset.Issuer {
		return false, ErrHolderMatchesIssuer
	}

	if a.Account != a.Asset.Issuer {
		return false, ErrAccountMismatch
	}

	if ok, err := IsAsset(a.Asset); !ok {
		return false, err
	}

	if ok, err := IsAsset(a.Asset2); !ok {
		return false, err
	}
	if ok, err := IsIssuedCurrency(a.Amount); !ok {
		return false, err
	}

	if a.Amount != nil {
		if a.Amount.Currency != a.Asset.Currency {
			return false, ErrAmountCurrencyMismatch
		}
		if a.Amount.Issuer != a.Asset.Issuer {
			return false, ErrAmountIssuerMismatch
		}
		
	}

	return true, nil
}
