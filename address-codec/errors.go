package addresscodec

import (
	"errors"
	"fmt"
)

var (
	// Static errors

	// ErrInvalidClassicAddress indicates an invalid classic address.
	ErrInvalidClassicAddress = errors.New("invalid classic address")
	// ErrInvalidSeed indicates an invalid seed; could not determine encoding algorithm.
	ErrInvalidSeed = errors.New("invalid seed; could not determine encoding algorithm")

	// ErrInvalidXAddress indicates an invalid x-address.
	ErrInvalidXAddress = errors.New("invalid x-address")
	// ErrInvalidTag indicates an invalid tag.
	ErrInvalidTag = errors.New("invalid tag")
	// ErrInvalidAccountID indicates an invalid account ID.
	ErrInvalidAccountID = errors.New("invalid accountId")

	// ErrInvalidAddressFormat indicates a general invalid XRPL address format.
	ErrInvalidAddressFormat = errors.New("invalid address format")

	// ErrChecksum indicates that the checksum of a check-encoded string does not verify against
	// the checksum.
	ErrChecksum = errors.New("checksum error")
	// ErrInvalidFormat indicates that the check-encoded string has an invalid format.
	ErrInvalidFormat = errors.New("invalid format: version and/or checksum bytes missing")
)

// Dynamic errors

// EncodeLengthError is an error that occurs when the length of the input does not match the expected length.
type EncodeLengthError struct {
	Instance string
	Input    int
	Expected int
}

// Error implements the error interface.
func (e *EncodeLengthError) Error() string {
	return fmt.Sprintf("`%v` length should be %v not %v", e.Instance, e.Expected, e.Input)
}
