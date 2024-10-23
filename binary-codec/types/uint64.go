package types

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

// UInt64 represents a 64-bit unsigned integer.
type UInt64 struct{}

var ErrInvalidUInt64String error = errors.New("invalid UInt64 string, value should be a string representation of a UInt64")

// FromJson converts a JSON value into a serialized byte slice representing a 64-bit unsigned integer.
// The input value is assumed to be a string representation of an integer. If the serialization fails, an error is returned.
func (u *UInt64) FromJson(value any) ([]byte, error) {

	var buf = new(bytes.Buffer)

	if _, ok := value.(string); !ok {
		return nil, ErrInvalidUInt64String
	}

	if !isNumeric(value.(string)) {
		if hex, err := hex.DecodeString(value.(string)); err == nil {
			buf.Write(hex)
			return buf.Bytes(), nil
		}
		stringToUint64, err := strconv.ParseUint(value.(string), 10, 64)
		if err != nil {
			return nil, err
		}
		value = stringToUint64
		err = binary.Write(buf, binary.BigEndian, value)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	} else {
		value = strings.Repeat("0", 16-len(value.(string))) + value.(string) // right justify the string
		decoded, err := hex.DecodeString(value.(string))
		if err != nil {
			return nil, err
		}
		buf.Write(decoded)
	}

	return buf.Bytes(), nil
}

// ToJson takes a BinaryParser and optional parameters, and converts the serialized byte data
// back into a JSON string value. This method assumes the parser contains data representing
// a 64-bit unsigned integer. If the parsing fails, an error is returned.
func (u *UInt64) ToJson(p interfaces.BinaryParser, opts ...int) (any, error) {
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
