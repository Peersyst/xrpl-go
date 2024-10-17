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

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
	bigdecimal "github.com/Peersyst/xrpl-go/pkg/big-decimal"
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
	zeroByteArray          = make([]byte, 20)
)

// InvalidAmountError is a custom error type for invalid amounts.
type InvalidAmountError struct {
	Amount string
}

// Error method for InvalidAmountError returns a formatted error string.
func (e *InvalidAmountError) Error() string {
	return fmt.Sprintf("value '%s' is an invalid amount", e.Amount)
}

// OutOfRangeError is a custom error type for out-of-range values.
type OutOfRangeError struct {
	Type string
}

// Error method for OutOfRangeError returns a formatted error string.
func (e *OutOfRangeError) Error() string {
	return fmt.Sprintf("%s is out of range", e.Type)
}

// InvalidCodeError is a custom error type for invalid currency codes.
type InvalidCodeError struct {
	Disallowed string
}

// Error method for InvalidCodeError returns a formatted error string.
func (e *InvalidCodeError) Error() string {
	return fmt.Sprintf("'%s' is/are disallowed or invalid", e.Disallowed)
}

// Amount is a struct that represents an XRPL Amount.
type Amount struct{}

// FromJson serializes an issued currency amount to its bytes representation from JSON.
func (a *Amount) FromJson(value any) ([]byte, error) {

	switch value := value.(type) {
	case string:
		return serializeXrpAmount(value)
	case map[string]any:
		return serializeIssuedCurrencyAmount(value["value"].(string), value["currency"].(string), value["issuer"].(string))
	default:
		return nil, errors.New("invalid amount type")
	}
}

// ToJson deserializes a binary-encoded Amount object from a BinaryParser into a JSON representation.
func (a *Amount) ToJson(p interfaces.BinaryParser, opts ...int) (any, error) {
	b, err := p.Peek()
	if err != nil {
		return nil, err
	}
	var sign string
	if !isPositive(b) {
		sign = "-"
	}
	if isNative(b) {
		xrp, err := p.ReadBytes(8)
		if err != nil {
			return nil, err
		}
		xrpVal := binary.BigEndian.Uint64(xrp)
		xrpVal = xrpVal & 0x3FFFFFFFFFFFFFFF
		return sign + strconv.FormatUint(xrpVal, 10), nil
	} else {
		token, err := p.ReadBytes(48)
		if err != nil {
			return nil, err
		}
		return deserializeToken(token)
	}
}

func deserializeToken(data []byte) (map[string]any, error) {

	var value string
	var err error
	if bytes.Equal(data[0:8], []byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) {
		value = "0"
	} else {
		value, err = deserializeValue(data[:8])
		if err != nil {
			return nil, err
		}
	}
	issuer, err := deserializeIssuer(data[28:])
	if err != nil {
		return nil, err
	}
	curr, err := deserializeCurrencyCode(data[8:28])
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"value":    value,
		"currency": curr,
		"issuer":   issuer,
	}, nil
}

func deserializeValue(data []byte) (string, error) {
	sign := ""
	if !isPositive(data[0]) {
		sign = "-"
	}
	value_bytes := data[:8]
	b1 := value_bytes[0]
	b2 := value_bytes[1]
	e1 := int((b1 & 0x3F) << 2)
	e2 := int(b2 >> 6)
	exponent := e1 + e2 - 97
	sig_figs := append([]byte{0, (b2 & 0x3F)}, value_bytes[2:]...)
	sig_figs_int := binary.BigEndian.Uint64(sig_figs)
	d, err := bigdecimal.NewBigDecimal(sign + strconv.Itoa(int(sig_figs_int)) + "e" + strconv.Itoa(exponent))
	if err != nil {
		return "", err
	}
	val := d.GetScaledValue()
	err = verifyIOUValue(val)
	if err != nil {
		return "", err
	}
	return val, nil
}

func deserializeCurrencyCode(data []byte) (string, error) {
	// Check for special xrp case
	if bytes.Equal(data, zeroByteArray) {
		return "XRP", nil
	}

	if bytes.Equal(data[0:12], make([]byte, 12)) && bytes.Equal(data[12:15], []byte{0x58, 0x52, 0x50}) && bytes.Equal(data[15:20], make([]byte, 5)) { // XRP in bytes
		return "", ErrInvalidCurrencyCode
	}
	iso := strings.ToUpper(string(data[12:15]))
	ok, _ := regexp.MatchString(IOUCodeRegex, iso)

	if !ok {
		return strings.ToUpper(hex.EncodeToString(data)), nil
	}
	return iso, nil
}

func deserializeIssuer(data []byte) (string, error) {
	return addresscodec.Encode(data, []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength), nil
}

// verifyXrpValue validates the format of an XRP amount value.
// XRP values should not contain a decimal point because they are represented as integers as drops.
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

// verifyIOUValue validates the format of an issued currency amount value.
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

// serializeXrpAmount serializes an XRP amount value.
func serializeXrpAmount(value string) ([]byte, error) {

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

// SerializeIssuedCurrencyValue serializes the value field of an issued currency amount to its bytes representation.
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

// serializeIssuedCurrencyCode serializes an issued currency code to its bytes representation.
// The currency code can be 3 allowed string characters, or 20 bytes of hex.
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

// SerializeIssuedCurrencyAmount serializes the currency field of an issued currency amount to its bytes representation
// from value, currency code, and issuer address in string form (e.g. "USD", "r123456789").
// The currency code can be 3 allowed string characters, or 20 bytes of hex in standard currency format (e.g. with "00" prefix)
// or non-standard currency format (e.g. without "00" prefix)
func serializeIssuedCurrencyAmount(value, currency, issuer string) ([]byte, error) {

	var valBytes []byte
	var err error
	if value == "0" {
		valBytes = make([]byte, 8)
		binary.BigEndian.PutUint64(valBytes, uint64(ZeroCurrencyAmountHex))
	} else {
		valBytes, err = SerializeIssuedCurrencyValue(value) // serialize the value
	}

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
func isNative(value byte) bool {
	x := value&NotXRPBitMask == 0 // & bitwise operator returns 1 if both first bits are 1, otherwise 0
	return x
}

// Determines if this AmountType is positive - 2nd bit in 1st byte is set to 1 for positive amounts
func isPositive(value byte) bool {
	x := value&0x40 > 0
	return x
}

func containsInvalidIOUCodeCharactersHex(currency []byte) bool {

	r := regexp.MustCompile(IOUCodeRegex) // regex to check if the currency code is valid
	m := r.FindAll(currency, -1)

	return len(m) != 1
}
