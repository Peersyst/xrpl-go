package transaction

import (
	"errors"
	"strconv"

	maputils "github.com/Peersyst/xrpl-go/pkg/map_utils"
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

const (
	// The Memos field includes arbitrary messaging data with the transaction.
	// It is presented as an array of objects. Each object has only one field, Memo,
	// which in turn contains another object with one or more of the following fields:
	// MemoData, MemoFormat, and MemoType. https://xrpl.org/docs/references/protocol/transactions/common-fields#memos-field
	MEMO_SIZE   = 3
	SIGNER_SIZE = 3
	// For a token, must have the following fields: currency, issuer, value. https://xrpl.org/docs/references/protocol/data-types/basic-data-types#specifying-currency-amounts
	ISSUED_CURRENCY_SIZE       = 3
	STANDARD_CURRENCY_CODE_LEN = 3
)

// IsMemo checks if the given object is a valid Memo object.
func IsMemo(memo Memo) (bool, error) {
	// Get the size of the Memo object.
	size := len(maputils.GetKeys(memo.Flatten()))

	if size == 0 {
		return false, errors.New("memo object should have at least one field, MemoData, MemoFormat or MemoType")
	}

	if size > MEMO_SIZE {
		return false, errors.New("memo object should have at most three fields, MemoData, MemoFormat and MemoType")
	}

	validData := memo.MemoData == "" || typecheck.IsHex(memo.MemoData)
	if !validData {
		return false, errors.New("memoData should be a hexadecimal string")
	}

	validFormat := memo.MemoFormat == "" || typecheck.IsHex(memo.MemoFormat)
	if !validFormat {
		return false, errors.New("memoFormat should be a hexadecimal string")
	}

	validType := memo.MemoType == "" || typecheck.IsHex(memo.MemoType)
	if !validType {
		return false, errors.New("memoType should be a hexadecimal string")
	}

	return true, nil
}

// IsSigner checks if the given object is a valid Signer object.
func IsSigner(signerData SignerData) (bool, error) {
	size := len(maputils.GetKeys(signerData.Flatten()))
	if size != SIGNER_SIZE {
		return false, errors.New("signers: Signer should have 3 fields: Account, TxnSignature, SigningPubKey")
	}

	// TODO: Update to check if the account is valid when the "isAccount" function exists
	validAccount := signerData.Account != "" && typecheck.IsString(signerData.Account.String())
	if !validAccount {
		return false, errors.New("signers: Account should be a string")
	}

	validTxnSignature := signerData.TxnSignature != "" && typecheck.IsString(signerData.TxnSignature)
	if !validTxnSignature {
		return false, errors.New("signers: TxnSignature should be a string")
	}

	validSigningPubKey := signerData.SigningPubKey != "" && typecheck.IsString(signerData.SigningPubKey)
	if !validSigningPubKey {
		return false, errors.New("signers: SigningPubKey should be a string")
	}

	return true, nil

}

// IsAmount checks if the given object is a valid Amount object.
// It is a string for an XRP amount or a map for an IssuedCurrency amount.
func IsAmount(amount types.CurrencyAmount) bool {
	if amount == nil {
		return false
	}
	if amount.Kind() == types.XRP {
		return true
	}

	if ok, _ := IsIssuedCurrency(amount); ok {
		return true
	}

	return false
}

// IsIssuedCurrency checks if the given object is a valid IssuedCurrency object.
func IsIssuedCurrency(input types.CurrencyAmount) (bool, error) {
	if input.Kind() == types.XRP {
		return false, errors.New("an issued currency cannot be of type XRP")
	}

	issuedAmount, _ := input.(types.IssuedCurrencyAmount)
	if issuedAmount.Currency == "" {
		return false, errors.New("currency field is required for an issued currency")
	}
	if issuedAmount.Currency == "XRP" {
		return false, errors.New("cannot have an issued currency with a similar standard code as XRP")
	}

	// Check if the value is a valid positive number
	value, err := strconv.ParseFloat(issuedAmount.Value, 64)
	if err != nil || value < 0 {
		return false, errors.New("value field should be a valid positive number")
	}

	return true, nil
}

// IsPathStep checks if the given map is a valid PathStep.
func IsPathStep(pathStep map[string]interface{}) bool {
	if account, ok := pathStep["account"]; ok && !typecheck.IsString(account) {
		return false
	}
	if currency, ok := pathStep["currency"]; ok && !typecheck.IsString(currency) {
		return false
	}
	if issuer, ok := pathStep["issuer"]; ok && !typecheck.IsString(issuer) {
		return false
	}
	if _, ok := pathStep["account"]; ok {
		if _, ok := pathStep["currency"]; !ok {
			if _, ok := pathStep["issuer"]; !ok {
				return true
			}
		}
	}

	// check if the path step has either a currency or an issuer
	_, hasCurr := pathStep["currency"]
	_, hasIssuer := pathStep["issuer"]

	if !hasCurr && !hasIssuer {
		return true
	}
	return false
}

// IsPath checks if the given pathstep is valid.
func IsPath(path []PathStep) (bool, error) {
	for _, pathStep := range path {

		hasAccount := pathStep.Account != ""
		hasCurrency := pathStep.Currency != ""
		hasIssuer := pathStep.Issuer != ""

		/**
		In summary, the following combination of fields are valid, optionally with type, type_hex, or both (but these two are deprecated):

		- account by itself
		- currency by itself
		- currency and issuer as long as the currency is not XRP
		- issuer by itself

		Any other use of account, currency, and issuer fields in a path step is invalid.

		https://xrpl.org/docs/concepts/tokens/fungible-tokens/paths#path-specifications
		*/
		if (hasAccount && !hasCurrency && !hasIssuer) || (hasCurrency && !hasAccount && !hasIssuer) || (hasIssuer && !hasAccount && !hasCurrency) {
			return true, nil
		} else if hasIssuer && hasCurrency && pathStep.Currency != "XRP" {
			return true, nil
		} else {
			return false, errors.New("invalid path step, check the valid fields combination at https://xrpl.org/docs/concepts/tokens/fungible-tokens/paths#path-specifications")
		}

	}
	return true, nil
}

// IsPaths checks if the given slice of slices of maps is a valid Paths.
func IsPaths(pathsteps [][]PathStep) (bool, error) {
	if len(pathsteps) == 0 {
		return false, errors.New("paths should have at least one path")
	}

	for _, path := range pathsteps {
		if len(path) == 0 {
			return false, errors.New("path should have at least one path step")
		}

		if ok, err := IsPath(path); !ok {
			return false, err
		}
	}

	return true, nil
}
