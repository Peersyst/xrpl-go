package transaction

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

var (
	// ErrInvalidCredentialURI is returned when the URI field does not meet the maximum allowed hex-encoded length of 512 characters (256 bytes).
	ErrInvalidCredentialURI = errors.New("credential create: invalid URI, must have a maximum hex string length of 512 characters (256 bytes)")
)

const (
	// Minimum length of a credential type is 1 byte (1 byte = 2 hex characters).
	MinURILength = 2

	// Maximum of 256 bytes for the URI field (1 byte = 2 hex characters)
	MaxURILength = 512
)

// A CredentialCreate transaction creates a credential in the ledger.
// The issuer of the credential uses this transaction to provisionally issue a credential.
// The credential is not valid until the subject of the credential accepts it with a CredentialAccept transaction.
type CredentialCreate struct {
	// Base transaction fields
	BaseTx

	// The subject of the credential (the address that will receive the credential).
	Subject types.Address

	// Arbitrary data defining the type of credential this entry represents. The minimum length is 1 byte and the maximum length is 64 bytes.
	CredentialType types.CredentialType

	// Time after which this credential expires, in seconds since the Ripple Epoch.
	// The Ripple Epoch is January 1, 2000 (00:00 UTC), represented as Unix time 946684800.
	// https://xrpl.org/docs/references/protocol/data-types/basic-data-types#specifying-time
	Expiration uint32 `json:",omitempty"`

	// Arbitrary additional data about the credential, such as the URL where users can look up an associated Verifiable Credential document. If present, the minimum length is 1 byte and the maximum is 256 bytes.
	URI string `json:",omitempty"`
}

// TxType implements the TxType method for the CredentialCreate struct.
func (*CredentialCreate) TxType() TxType {
	return CredentialCreateTx
}

// Flatten implements the Flatten method for the CredentialCreate struct.
func (c *CredentialCreate) Flatten() FlatTransaction {
	flattened := c.BaseTx.Flatten()

	flattened["TransactionType"] = c.TxType().String()

	if c.Subject != "" {
		flattened["Subject"] = c.Subject.String()
	}

	if c.CredentialType != "" {
		flattened["CredentialType"] = c.CredentialType.String()
	}

	if c.Expiration != 0 {
		flattened["Expiration"] = c.Expiration
	}

	if c.URI != "" {
		flattened["URI"] = c.URI
	}

	return flattened
}

// Validate implements the Validate method for the CredentialCreate struct.
func (c *CredentialCreate) Validate() (bool, error) {
	// validate the base transaction
	_, err := c.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if !addresscodec.IsValidAddress(c.Subject.String()) {
		return false, ErrInvalidSubject
	}

	if !c.CredentialType.IsValid() {
		return false, ErrInvalidCredentialType
	}

	if c.URI != "" && (len(c.URI) < MinURILength || len(c.URI) > MaxURILength) {
		return false, ErrInvalidCredentialURI
	}

	return true, nil
}
