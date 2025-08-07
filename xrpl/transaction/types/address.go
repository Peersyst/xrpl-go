// Package types provides core transaction types and helpers for the XRPL Go library.
package types

// Address represents a classic XRPL account address as a string.
type Address string

func (a Address) String() string {
	return string(a)
}
