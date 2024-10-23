package types

import (
	"bytes"
	"encoding/binary"

	"github.com/Peersyst/xrpl-go/binary-codec/definitions"
	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

// UInt16 represents a 16-bit unsigned integer.
type UInt16 struct{}

// FromJson converts a JSON value into a serialized byte slice representing a 16-bit unsigned integer.
// If the input value is a string, it's assumed to be a transaction type or ledger entry type name, and the
// method will attempt to convert it into a corresponding type code. If the conversion fails, an error is returned.
func (u *UInt16) FromJson(value any) ([]byte, error) {

	if _, ok := value.(string); ok {
		tc, err := definitions.Get().GetTransactionTypeCodeByTransactionTypeName(value.(string))
		if err != nil {
			tc, err = definitions.Get().GetLedgerEntryTypeCodeByLedgerEntryTypeName(value.(string))
			if err != nil {
				return nil, err
			}
		}
		value = int(tc)
	}

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint16(value.(int)))

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToJson takes a BinaryParser and optional parameters, and converts the serialized byte data
// back into a JSON integer value. This method assumes the parser contains data representing
// a 16-bit unsigned integer. If the parsing fails, an error is returned.
func (u *UInt16) ToJson(p interfaces.BinaryParser, opts ...int) (any, error) {
	b, err := p.ReadBytes(2)
	if err != nil {
		return nil, err
	}
	return int(binary.BigEndian.Uint16(b)), nil
}
