package types

import "github.com/Peersyst/xrpl-go/pkg/typecheck"

const (
	// Minimum length of a credential type is 1 byte (1 byte = 2 hex characters).
	MinCredentialTypeLength = 2

	// Maximum length of a credential type is 64 bytes (1 byte = 2 hex characters).
	MaxCredentialTypeLength = 128
)

type CredentialType string

// String returns the string representation of a CredentialType.
func (c *CredentialType) String() string {
	return string(*c)
}

// IsValidCredentialType checks if a credential type meets all the requirements:
// - Not empty
// - Valid hex string
// - Length between MinCredentialTypeLength and MaxCredentialTypeLength
func (c *CredentialType) IsValid() bool {
	if c.String() == "" {
		return false
	}

	credTypeStr := c.String()
	if !typecheck.IsHex(credTypeStr) {
		return false
	}

	length := len(credTypeStr)
	return length >= MinCredentialTypeLength && length <= MaxCredentialTypeLength
}
