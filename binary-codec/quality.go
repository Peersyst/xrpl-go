package binarycodec

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"

	bigdecimal "github.com/Peersyst/xrpl-go/pkg/big-decimal"
)

const (
	// PosSignBitMask is the mask for the positive sign bit.
	PosSignBitMask = 0x4000000000000000
	// ZeroQualityHex is the hex representation of the zero quality.
	ZeroQualityHex = 0x5500000000000000
	// MaxIOUPrecision is the maximum precision for an IOU.
	MaxIOUPrecision = 16
	// MinIOUExponent is the minimum exponent for an IOU.
	MinIOUExponent = -96
	// MaxIOUExponent is the maximum exponent for an IOU.
	MaxIOUExponent = 80
)

var (
	ErrInvalidQuality = errors.New("invalid quality")
)

// EncodeQuality encodes a quality amount to a hex string.
func EncodeQuality(quality string) (string, error) {
	if len(quality) == 0 {
		return "", ErrInvalidQuality
	}
	if len(strings.Trim(strings.Trim(quality, "0"), ".")) == 0 {
		zeroAmount := make([]byte, 8)
		binary.BigEndian.PutUint64(zeroAmount, uint64(ZeroQualityHex))
		return hex.EncodeToString(zeroAmount), nil
	}

	bigDecimal, err := bigdecimal.NewBigDecimal(quality)
	if err != nil {
		return "", err
	}

	if !isValidQuality(*bigDecimal) {
		return "", ErrInvalidQuality
	}

	if bigDecimal.UnscaledValue == "" {
		zeroAmount := make([]byte, 8)
		binary.BigEndian.PutUint64(zeroAmount, uint64(ZeroQualityHex))
		// if the value is zero, then return the zero currency amount hex
		return hex.EncodeToString(zeroAmount), nil
	}

	// convert the unscaled value to an unsigned integer
	mantissa, err := strconv.ParseUint(bigDecimal.UnscaledValue, 10, 64)

	if err != nil {
		return "", err
	}

	// get the scale
	exp := bigDecimal.Scale

	// set first bit to 1 because it is not XRP
	serial := uint64(ZeroQualityHex)
	if bigDecimal.Sign == 0 {
		// if the sign is positive, set the sign (second) bit to 1
		serial |= PosSignBitMask
	}

	serialized := make([]byte, 8)
	binary.BigEndian.PutUint64(serialized, mantissa)
	serialized[0] += byte(exp) + 100
	return strings.ToUpper(hex.EncodeToString(serialized)), nil
}

// Decode a quality amount from a hex string to a string.
func DecodeQuality(quality string) (string, error) {
	if quality == "" {
		return "", ErrInvalidQuality
	}

	decoded, err := hex.DecodeString(quality)
	if err != nil {
		return "", err
	}

	bytes := decoded[len(decoded)-8:]
	exp := int(bytes[0]) - 100
	mantissaBytes := append([]byte{0}, bytes[1:]...)
	mantissa := binary.BigEndian.Uint64(mantissaBytes)

	// Convert mantissa to string
	mantissaStr := strconv.FormatUint(mantissa, 10)

	// Add decimal point based on exponent
	if exp < 0 {
		// Need to add leading zeros
		if len(mantissaStr) <= -exp {
			zeros := strings.Repeat("0", -exp-len(mantissaStr)+1)
			mantissaStr = "0." + zeros + mantissaStr
		} else {
			// Insert decimal point from right to left
			insertPos := len(mantissaStr) + exp
			mantissaStr = mantissaStr[:insertPos] + "." + mantissaStr[insertPos:]
		}
	} else if exp > 0 {
		// Add trailing zeros
		mantissaStr += strings.Repeat("0", exp)
	}

	// Trim trailing zeros after decimal point
	if strings.Contains(mantissaStr, ".") {
		mantissaStr = strings.TrimRight(mantissaStr, "0")
		mantissaStr = strings.TrimRight(mantissaStr, ".")
	}

	return mantissaStr, nil
}

func isValidQuality(quality bigdecimal.BigDecimal) bool {
	return quality.Precision <= MaxIOUPrecision && quality.Scale >= MinIOUExponent && quality.Scale <= MaxIOUExponent
}
