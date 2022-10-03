package types

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	addresscodec "github.com/xyield/xrpl-go/address-codec"
	bigdecimal "github.com/xyield/xrpl-go/pkg/big-decimal"
)

const (
	MinIOUExponent  = -96
	MaxIOUExponent  = 80
	MaxIOUPrecision = 16
	MinIOUMantissa  = 1000000000000000
	MaxIOUMantissa  = 9999999999999999

	NotXRPBitMask            = 0x80
	PosSignBitMask           = 0x4000000000000000
	ZeroCurrencyAmountHex    = 0x8000000000000000
	NativeAmountByteLength   = 8
	CurrencyAmountByteLength = 48

	MinXRP   = 1e-6
	MaxDrops = 1e17 // 100 billion XRP in drops aka 10^17

	IOUCodeRegex = `[0-9A-Za-z?!@#$%^&*<>(){}\[\]|]{3}`
)

var (
	ErrInvalidXRPValue     = errors.New("invalid XRP value")
	ErrInvalidCurrencyCode = errors.New("invalid currency code")
)

type InvalidAmountError struct {
	Amount string
}

func (e *InvalidAmountError) Error() string {
	return fmt.Sprintf("value '%s' is an invalid amount", e.Amount)
}

type OutOfRangeError struct {
	Type string
}

func (e *OutOfRangeError) Error() string {
	return fmt.Sprintf("%s is out of range", e.Type)
}

type InvalidCodeError struct {
	Disallowed string
}

func (e *InvalidCodeError) Error() string {
	return fmt.Sprintf("'%s' is/are disallowed or invalid", e.Disallowed)
}

type Amount struct{}

// Serializes an issued currency amount to its bytes representation from json
func (a *Amount) SerializeJson(value any) ([]byte, error) {

	switch value := value.(type) {
	case string:
		return SerializeXrpAmount(value)
	case map[string]any:
		return SerializeIssuedCurrencyAmount(value["value"].(string), value["currency"].(string), value["issuer"].(string))
	default:
		return nil, errors.New("invalid amount type")
	}
}

// validates the format of an XRP amount value
// XRP values shouldn't contain a decimal point BECAUSE they are represented as integers as drops
func verifyXrpValue(value string) error {

	r := regexp.MustCompile(`\d+`) // regex to match only digits
	m := r.FindAllString(value, -1)

	if len(m) != 1 {
		return ErrInvalidXRPValue
	}

	decimal := new(big.Float)
	decimal, ok := decimal.SetString(value) // bigFloat for precision

	if !ok {
		return errors.New("failed to convert string to big.Float")
	}

	if decimal.Sign() == 0 {
		return nil
	}

	if decimal.Cmp(big.NewFloat(MinXRP)) == -1 || decimal.Cmp(big.NewFloat(MaxDrops)) == 1 {
		return &InvalidAmountError{value}
	}

	return nil
}

// validates the format of an issued currency amount value
func verifyIOUValue(value string) error {

	bigDecimal, err := bigdecimal.NewBigDecimal(value)

	if err != nil {
		return err
	}

	if bigDecimal.UnscaledValue == "" {
		return nil
	}

	exp := bigDecimal.Scale

	if bigDecimal.Precision > MaxIOUPrecision {
		return &OutOfRangeError{Type: "Precision"} // if the precision is greater than 16, return an error
	}
	if exp < MinIOUExponent {
		return &OutOfRangeError{Type: "Exponent"} // if the scale is less than -96 or greater than 80, return an error
	}
	if exp > MaxIOUExponent {
		return &OutOfRangeError{Type: "Exponent"} // if the scale is less than -96 or greater than 80, return an error
	}

	return err
}

// Serializes an XRP amount value
func SerializeXrpAmount(value string) ([]byte, error) {

	if verifyXrpValue(value) != nil {
		return nil, verifyXrpValue(value)
	}

	val, err := strconv.ParseUint(value, 10, 64)

	if err != nil {
		return nil, err
	}

	valWithPosBit := val | PosSignBitMask
	valBytes := make([]byte, NativeAmountByteLength)

	binary.BigEndian.PutUint64(valBytes, uint64(valWithPosBit))

	return valBytes, nil
}

// XRPL definition of precision is number of significant digits:
// Tokens can represent a wide variety of assets, including those typically measured in very small or very large denominations.
// This format uses significant digits and a power-of-ten exponent in a similar way to scientific notation.
// The format supports positive and negative significant digits and exponents within the specified range.
// Unlike typical floating-point representations of non-whole numbers, this format uses integer math for all calculations,
// so it always maintains 15 decimal digits of precision. Multiplication and division have adjustments to compensate for
// over-rounding in the least significant digits.

// Serializes the value field of an issued currency amount to its bytes representation
func SerializeIssuedCurrencyValue(value string) ([]byte, error) {

	if verifyIOUValue(value) != nil {
		return nil, verifyIOUValue(value)
	}

	bigDecimal, err := bigdecimal.NewBigDecimal(value)

	if err != nil {
		return nil, err
	}

	if bigDecimal.UnscaledValue == "" {
		zeroAmount := make([]byte, 8)
		binary.BigEndian.PutUint64(zeroAmount, uint64(ZeroCurrencyAmountHex))
		return zeroAmount, nil // if the value is zero, then return the zero currency amount hex
	}

	mantissa, err := strconv.ParseUint(bigDecimal.UnscaledValue, 10, 64) // convert the unscaled value to an unsigned integer

	if err != nil {
		return nil, err
	}

	exp := bigDecimal.Scale // get the scale

	for mantissa < MinIOUMantissa && exp > MinIOUExponent {
		mantissa *= 10
		exp--
	}

	for mantissa > MaxIOUMantissa {
		if exp >= MaxIOUExponent {
			return nil, &OutOfRangeError{Type: "Exponent"} // if the scale is less than -96 or greater than 80, return an error
		}
		mantissa /= 10
		exp++

		if exp < MinIOUExponent || mantissa < MinIOUMantissa {
			// round down to zero
			zeroAmount := make([]byte, 8)
			binary.BigEndian.PutUint64(zeroAmount, uint64(ZeroCurrencyAmountHex))
			return zeroAmount, nil
		}

		if exp > MaxIOUExponent || mantissa > MaxIOUMantissa {
			return nil, &OutOfRangeError{Type: "Exponent"} // if the scale is less than -96 or greater than 80, return an error
		}
	}

	// convert components to bytes

	serial := uint64(ZeroCurrencyAmountHex) // set first bit to 1 because it is not XRP
	if bigDecimal.Sign == 0 {
		serial |= PosSignBitMask // if the sign is positive, set the sign (second) bit to 1
	}
	serial |= (uint64(exp+97) << 54) // if the exponent is positive, set the exponent bits to the exponent + 97
	serial |= uint64(mantissa)       // last 54 bits are mantissa

	serialReturn := make([]byte, 8)
	binary.BigEndian.PutUint64(serialReturn, serial)

	return serialReturn, nil
}

// Serializes an issued currency code to its bytes representation. The currency code can be 3 allowed string characters, or 20 bytes of hex
func serializeIssuedCurrencyCode(currency string) ([]byte, error) {

	currency = strings.TrimPrefix(currency, "0x")                                    // remove the 0x prefix if it exists
	if currency == "XRP" || currency == "0000000000000000000000005852500000000000" { // if the currency code is uppercase XRP, return an error
		return nil, &InvalidCodeError{Disallowed: "XRP uppercase"}
	}

	switch len(currency) {
	case 3: // if the currency code is 3 characters, it is standard
		return serializeIssuedCurrencyCodeChars(currency)
	case 40: // if the currency code is 40 characters, it is hex encoded
		return serializeIssuedCurrencyCodeHex(currency)
	}

	return nil, &InvalidCodeError{Disallowed: currency}

}

func serializeIssuedCurrencyCodeHex(currency string) ([]byte, error) {
	decodedHex, err := hex.DecodeString(currency)

	if err != nil {
		return nil, err
	}

	if bytes.HasPrefix(decodedHex, []byte{0x00}) {

		if bytes.Equal(decodedHex[12:15], []byte{0x00, 0x00, 0x00}) {
			return make([]byte, 20), nil
		}

		if containsInvalidIOUCodeCharactersHex(decodedHex[12:15]) {
			return nil, ErrInvalidCurrencyCode
		}
		return decodedHex, nil

	}
	return decodedHex, nil
}

func serializeIssuedCurrencyCodeChars(currency string) ([]byte, error) {

	r := regexp.MustCompile(IOUCodeRegex) // regex to check if the currency code is valid
	m := r.FindAllString(currency, -1)

	if len(m) != 1 {
		return nil, ErrInvalidCurrencyCode
	}

	currencyBytes := make([]byte, 20)
	copy(currencyBytes[12:], []byte(currency))
	return currencyBytes[:], nil
}

// Serializes the currency field of an issued currency amount to its bytes representation from value, currency code, and issuer address in string form (e.g. "USD", "r123456789")
// The currency code can be 3 allowed string characters, or 20 bytes of hex in standard currency format (e.g. with "00" prefix) or non-standard currency format (e.g. without "00" prefix)
func SerializeIssuedCurrencyAmount(value, currency, issuer string) ([]byte, error) {

	valBytes, err := SerializeIssuedCurrencyValue(value) // serialize the value

	if err != nil {
		return nil, err
	}
	currencyBytes, err := serializeIssuedCurrencyCode(currency) // serialize the currency code

	if err != nil {
		return nil, err
	}
	_, issuerBytes, err := addresscodec.DecodeClassicAddressToAccountID(issuer) // decode the issuer address
	if err != nil {
		return nil, err
	}

	// AccountIDs that appear as children of special fields (Amount issuer and PathSet account) are not length-prefixed.
	// So in Amount and PathSet fields, don't use the length indicator 0x14. This is in contrast to the AccountID fields where the length indicator prefix 0x14 is added.

	return append(append(valBytes, currencyBytes...), issuerBytes...), nil
}

// Returns true if this amount is a "native" XRP amount - first bit in first byte set to 0 for native XRP
func isNative(value []byte) bool {
	fmt.Printf("%08b", value)
	x := []byte(value)[0]&NotXRPBitMask == 0 // & bitwise operator returns 1 if both first bits are 1, otherwise 0
	return x
}

// Determines if this AmountType is positive - 2nd bit in 1st byte is set to 1 for positive amounts
func isPositive(value []byte) bool {
	fmt.Printf("%08b", value)
	x := []byte(value)[0]&0x40 > 0
	return x
}

func containsInvalidIOUCodeCharactersHex(currency []byte) bool {

	r := regexp.MustCompile(IOUCodeRegex) // regex to check if the currency code is valid
	m := r.FindAll(currency, -1)

	return len(m) != 1
}
