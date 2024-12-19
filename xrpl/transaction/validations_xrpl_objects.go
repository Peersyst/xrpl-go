package transaction

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	maputils "github.com/Peersyst/xrpl-go/v1/pkg/map_utils"
	"github.com/Peersyst/xrpl-go/v1/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/v1/xrpl/currency"
	"github.com/Peersyst/xrpl-go/v1/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

const (
	// The Memos field includes arbitrary messaging data with the transaction.
	// It is presented as an array of objects. Each object has only one field, Memo,
	// which in turn contains another object with one or more of the following fields:
	// MemoData, MemoFormat, and MemoType. https://xrpl.org/docs/references/protocol/transactions/common-fields#memos-field
	MemoSize   = 3
	SignerSize = 3
	// For a token, must have the following fields: currency, issuer, value. https://xrpl.org/docs/references/protocol/data-types/basic-data-types#specifying-currency-amounts
	IssuedCurrencySize      = 3
	StandardCurrencyCodeLen = 3
)

// *************************
// Errors
// *************************

var (
	// ErrEmptyPath is returned when the path is empty.
	ErrEmptyPath = errors.New("path(s) should have at least one path")
	// ErrInvalidTokenCurrency is returned when the token currency is XRP.
	ErrInvalidTokenCurrency = errors.New("invalid or missing token currency, it also cannot have a similar standard code as XRP")
	// ErrInvalidTokenFields is returned when the issued currency object does not have the required fields (currency, issuer and value).
	ErrInvalidTokenFields = errors.New("issued currency object should have 3 fields: currency, issuer, value")
	// ErrInvalidPathStepCombination is returned when the path step is invalid. The fields combination is invalid.
	ErrInvalidPathStepCombination = errors.New("invalid path step, check the valid fields combination at https://xrpl.org/docs/concepts/tokens/fungible-tokens/paths#path-specifications")
	// ErrInvalidTokenValue is returned when the value field is not a valid positive number.
	ErrInvalidTokenValue = errors.New("value field should be a valid positive number")
	// ErrInvalidTokenType is returned when an issued currency is of type XRP.
	ErrInvalidTokenType = errors.New("an issued currency cannot be of type XRP")
	// ErrMissingTokenCurrency is returned when the currency field is missing for an issued currency.
	ErrMissingTokenCurrency = errors.New("currency field is missing for the issued currency")
	// ErrInvalidAssetFields is returned when the asset object does not have the required fields (currency, or currency and issuer).
	ErrInvalidAssetFields = errors.New("asset object should have at least one field 'currency', or two fields 'currency' and 'issuer'")
	// ErrMissingAssetCurrency is returned when the currency field is missing for an asset.
	ErrMissingAssetCurrency = errors.New("currency field is required for an asset")
	// ErrInvalidAssetIssuer is returned when the issuer field is invalid for an asset.
	ErrInvalidAssetIssuer = errors.New("issuer field must be a valid XRPL classic address")
)

// ErrMissingAmount is a function that returns an error when a field of type CurrencyAmount is missing.
func ErrMissingAmount(fieldName string) error {
	return fmt.Errorf("missing field %s", fieldName)
}

// *************************
// Validations
// *************************

// IsMemo checks if the given object is a valid Memo object.
func IsMemo(memo types.Memo) (bool, error) {
	// Get the size of the Memo object.
	size := len(maputils.GetKeys(memo.Flatten()))

	if size == 0 {
		return false, errors.New("memo object should have at least one field, MemoData, MemoFormat or MemoType")
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
	if size != SignerSize {
		return false, errors.New("signers: Signer should have 3 fields: Account, TxnSignature, SigningPubKey")
	}

	validAccount := strings.TrimSpace(signerData.Account.String()) != "" && addresscodec.IsValidClassicAddress(signerData.Account.String())
	if !validAccount {
		return false, errors.New("signers: Account should be a string")
	}

	if strings.TrimSpace(signerData.TxnSignature) == "" {
		return false, errors.New("signers: TxnSignature should be a non-empty string")
	}

	if strings.TrimSpace(signerData.SigningPubKey) == "" {
		return false, errors.New("signers: SigningPubKey should be a non-empty string")
	}

	return true, nil

}

// IsAmount checks if the given object is a valid Amount object.
// It is a string for an XRP amount or a map for an IssuedCurrency amount.
func IsAmount(field types.CurrencyAmount, fieldName string, isFieldRequired bool) (bool, error) {
	if isFieldRequired && field == nil {
		return false, ErrMissingAmount(fieldName)
	}

	if !isFieldRequired && field == nil {
		// no need to check further properties on a nil field, will create a panic with tests otherwise
		return true, nil
	}

	if field.Kind() == types.XRP {
		return true, nil
	}

	if ok, err := IsIssuedCurrency(field); !ok {
		return false, err
	}

	return true, nil
}

// IsIssuedCurrency checks if the given object is a valid IssuedCurrency object.
func IsIssuedCurrency(input types.CurrencyAmount) (bool, error) {
	if input.Kind() == types.XRP {
		return false, ErrInvalidTokenType
	}

	// Get the size of the IssuedCurrency object.
	issuedAmount, _ := input.(types.IssuedCurrencyAmount)

	numOfKeys := len(maputils.GetKeys(issuedAmount.Flatten().(map[string]interface{})))
	if numOfKeys != IssuedCurrencySize {
		return false, ErrInvalidTokenFields
	}

	if strings.TrimSpace(issuedAmount.Currency) == "" {
		return false, ErrMissingTokenCurrency
	}
	if strings.ToUpper(issuedAmount.Currency) == currency.NativeCurrencySymbol {
		return false, ErrInvalidTokenCurrency
	}

	if !addresscodec.IsValidClassicAddress(issuedAmount.Issuer.String()) {
		return false, ErrInvalidIssuer
	}

	// Check if the value is a valid positive number
	value, err := strconv.ParseFloat(issuedAmount.Value, 64)
	if err != nil || value < 0 {
		return false, ErrInvalidTokenValue
	}

	return true, nil
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
		switch {
		case hasAccount && !hasCurrency && !hasIssuer:
			return true, nil
		case hasCurrency && !hasAccount && !hasIssuer:
			return true, nil
		case hasIssuer && !hasAccount && !hasCurrency:
			return true, nil
		case hasIssuer && hasCurrency && pathStep.Currency != currency.NativeCurrencySymbol:
			return true, nil
		default:
			return false, ErrInvalidPathStepCombination
		}

	}
	return true, nil
}

// IsPaths checks if the given slice of slices of maps is a valid Paths.
func IsPaths(pathsteps [][]PathStep) (bool, error) {
	if len(pathsteps) == 0 {
		return false, ErrEmptyPath
	}

	for _, path := range pathsteps {
		if len(path) == 0 {
			return false, ErrEmptyPath
		}

		if ok, err := IsPath(path); !ok {
			return false, err
		}
	}

	return true, nil
}

// IsAsset checks if the given object is a valid Asset object.
func IsAsset(asset ledger.Asset) (bool, error) {
	// Get the size of the Asset object.
	lenKeys := len(maputils.GetKeys(asset.Flatten()))

	if lenKeys == 0 {
		return false, ErrInvalidAssetFields
	}

	if strings.TrimSpace(asset.Currency) == "" {
		return false, ErrMissingAssetCurrency
	}

	if strings.ToUpper(asset.Currency) == currency.NativeCurrencySymbol && strings.TrimSpace(asset.Issuer.String()) == "" {
		return true, nil
	}

	if strings.ToUpper(asset.Currency) == currency.NativeCurrencySymbol && asset.Issuer != "" {
		return false, ErrInvalidAssetIssuer
	}

	if asset.Currency != "" && !addresscodec.IsValidClassicAddress(asset.Issuer.String()) {
		return false, ErrInvalidAssetIssuer
	}

	return true, nil
}
