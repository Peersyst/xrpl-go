package types

import (
	"encoding/hex"
	"errors"
	"strings"

	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

// ErrNoLengthPrefix error is raised when no length prefix size is given.
var ErrNoLengthPrefix error = errors.New("no length prefix size given")

// Blob struct is used for manipulating hexadecimal data.
type Blob struct{}

// FromJson method for Blob converts a hexadecimal string from JSON to a byte array.
func (b *Blob) FromJson(json any) ([]byte, error) {
	// Convert hexadecimal string to byte array.
	// Return an error if the conversion fails.
	v, err := hex.DecodeString(json.(string))
	if err != nil {
		return nil, err
	}
	return v, nil
}

// ToJson method for Blob reads a certain number of bytes from a BinaryParser
// and converts it into a hexadecimal string.
// It returns an error if no length prefix is specified or if the read operation fails.
func (b *Blob) ToJson(p interfaces.BinaryParser, opts ...int) (any, error) {
	// If no length prefix is specified, return an error.
	if opts == nil {
		return nil, ErrNoLengthPrefix
	}
	// Read the specified number of bytes.
	// If the read operation fails, return an error.
	val, err := p.ReadBytes(opts[0])
	if err != nil {
		return nil, err
	}
	// Convert the bytes to a hexadecimal string and return it.
	return strings.ToUpper(hex.EncodeToString(val)), nil
}
