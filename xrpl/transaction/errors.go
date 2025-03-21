package transaction

import (
	"errors"
)

var (
	// ErrDestinationAccountConflict is returned when the Destination matches the Account.
	ErrDestinationAccountConflict = errors.New("destination cannot be the same as the Account")
	// ErrInvalidAccount is returned when the Account field does not meet XRPL address standards.
	ErrInvalidAccount = errors.New("invalid xrpl address for Account")
	// ErrInvalidCheckID is returned when the CheckID is not a valid 64-character hexadecimal string.
	ErrInvalidCheckID = errors.New("invalid CheckID, must be a valid 64-character hexadecimal string")
	// ErrInvalidCredentialType is returned when the CredentialType is not a valid hexadecimal string between 1 and 64 bytes.
	ErrInvalidCredentialType = errors.New("invalid credential type, must be a hexadecimal string between 1 and 64 bytes")
	// ErrInvalidCredentialIDs is returned when the CredentialIDs field is empty or not a valid hexadecimal string array.
	ErrInvalidCredentialIDs = errors.New("invalid credential IDs, must be a valid hexadecimal string array")
	// ErrInvalidDestination is returned when the Destination field does not meet XRPL address standards.
	ErrInvalidDestination = errors.New("invalid xrpl address for Destination")
	// ErrInvalidIssuer is returned when the issuer address is an invalid xrpl address.
	ErrInvalidIssuer = errors.New("invalid xrpl address for Issuer")
	// ErrInvalidOwner is returned when the Owner field does not meet XRPL address standards.
	ErrInvalidOwner = errors.New("invalid xrpl address for Owner")
	// ErrInvalidHexPublicKey is returned when the PublicKey is not a valid hexadecimal string.
	ErrInvalidHexPublicKey = errors.New("invalid PublicKey, must be a valid hexadecimal string")
	// ErrInvalidTransactionType is returned when the TransactionType field is invalid or missing.
	ErrInvalidTransactionType = errors.New("invalid or missing TransactionType")
	// ErrInvalidSubject is returned when the Subject field is an invalid xrpl address.
	ErrInvalidSubject = errors.New("invalid xrpl address for Subject")
)
