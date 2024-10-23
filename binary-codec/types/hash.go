package types

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

// ErrInvalidHashLength struct is used when the hash length does not meet the expected value.
type ErrInvalidHashLength struct {
	Expected int
}

type ErrInvalidHashType struct{}

type ErrInvalidHexString struct {
	Err error
}

// Error method for ErrInvalidHashLength formats the error message.
func (e *ErrInvalidHashLength) Error() string {
	return fmt.Sprintf("invalid hash length expected length %v", e.Expected)
}

func (e *ErrInvalidHashType) Error() string {
	return "invalid hash type"
}

func (e *ErrInvalidHexString) Error() string {
	return "error decoding hex string: " + e.Err.Error()
}

// hashI interface combines the SerializedType interface and getLength method for hashes.
type hashI interface {
	SerializedType
	getLength() int
}

// hash struct represents a hash with a specific length.
type hash struct {
	Length int
}

// newHash is a constructor for creating a new hash with a specified length.
func newHash(l int) hash {
	return hash{
		Length: l,
	}
}

// getLength method for hash returns the hash's length.
func (h hash) getLength() int {
	return h.Length
}

// FromJson method for hash converts a hexadecimal string from JSON to a byte array.
// It returns an error if the conversion fails or the length of the decoded byte array is not as expected.
func (h hash) FromJson(json any) ([]byte, error) {
	v, ok := json.(string)
	if !ok {
		return nil, &ErrInvalidHashType{}
	}
	decoded, err := hex.DecodeString(v)
	if err != nil {
		return nil, &ErrInvalidHexString{Err: err}
	}
	if h.getLength() != len(decoded) {
		return nil, &ErrInvalidHashLength{Expected: h.getLength()}
	}
	return decoded, nil
}

// ToJson method for hash reads a certain number of bytes from a BinaryParser and converts it into a hexadecimal string.
// It returns an error if the read operation fails.
func (h hash) ToJson(p interfaces.BinaryParser, opts ...int) (any, error) {
	b, err := p.ReadBytes(h.Length)
	if err != nil {
		return nil, err
	}
	return strings.ToUpper(hex.EncodeToString(b)), nil
}
