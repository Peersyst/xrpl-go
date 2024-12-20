package types

import (
	"bytes"
	"encoding/hex"
	"errors"
	"regexp"
	"strings"

	"github.com/Peersyst/xrpl-go/v1/binary-codec/types/interfaces"
)

// UInt64 represents a 64-bit unsigned integer.
type UInt64 struct{}

var ErrInvalidUInt64String = errors.New("invalid UInt64 string, value should be a string representation of a UInt64")

// FromJSON converts a JSON value into a serialized byte slice representing a 64-bit unsigned integer.
// The input value is assumed to be a string representation of an integer. If the serialization fails, an error is returned.
func (u *UInt64) FromJSON(value any) ([]byte, error) {

	var buf = new(bytes.Buffer)

	if _, ok := value.(string); !ok {
		return nil, ErrInvalidUInt64String
	}

	if !isNumeric(value.(string)) {
		hex, err := hex.DecodeString(value.(string))
		if err != nil {
			return nil, err
		}
		buf.Write(hex)
		return buf.Bytes(), nil
	}
	value = strings.Repeat("0", 16-len(value.(string))) + value.(string) // right justify the string
	decoded, err := hex.DecodeString(value.(string))
	if err != nil {
		return nil, err
	}
	buf.Write(decoded)

	return buf.Bytes(), nil
}

// ToJSON takes a BinaryParser and optional parameters, and converts the serialized byte data
// back into a JSON string value. This method assumes the parser contains data representing
// a 64-bit unsigned integer. If the parsing fails, an error is returned.
func (u *UInt64) ToJSON(p interfaces.BinaryParser, _ ...int) (any, error) {
	b, err := p.ReadBytes(8)
	if err != nil {
		return nil, err
	}
	return strings.ToUpper(hex.EncodeToString(b)), nil
}

// isNumeric checks if a string only contains numerical values.
func isNumeric(s string) bool {
	match, _ := regexp.MatchString("^[0-9]+$", s)
	return match
}
