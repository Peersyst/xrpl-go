package types

import (
	"bytes"
	"encoding/binary"

	"github.com/Peersyst/xrpl-go/binary-codec/definitions"
	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

// UInt8 represents an 8-bit unsigned integer.
type UInt8 struct{}

// FromJson converts a JSON value into a serialized byte slice representing an 8-bit unsigned integer.
// If the input value is a string, it's assumed to be a transaction result name, and the method will
// attempt to convert it into a transaction result type code. If the conversion fails, an error is returned.
func (u *UInt8) FromJson(value any) ([]byte, error) {
	if s, ok := value.(string); ok {
		tc, err := definitions.Get().GetTransactionResultTypeCodeByTransactionResultName(s)
		if err != nil {
			return nil, err
		}
		value = tc
	}

	var intValue int

	switch v := value.(type) {
	case int:
		intValue = v
	case int32:
		intValue = int(v)
	}

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint8(intValue))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToJson takes a BinaryParser and optional parameters, and converts the serialized byte data
// back into a JSON integer value. This method assumes the parser contains data representing
// an 8-bit unsigned integer. If the parsing fails, an error is returned.
func (u *UInt8) ToJson(p interfaces.BinaryParser, opts ...int) (any, error) {
	b, err := p.ReadBytes(1)
	if err != nil {
		return nil, err
	}
	return int(b[0]), nil
}
