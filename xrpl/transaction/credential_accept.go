package transaction

import (
	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// A CredentialCreate transaction creates a credential in the ledger.
// The issuer of the credential uses this transaction to provisionally issue a credential.
// The credential is not valid until the subject of the credential accepts it with a CredentialAccept transaction.
type CredentialAccept struct {
	// Base transaction fields
	BaseTx

	// The address of the issuer that created the credential.
	Issuer types.Address

	// Arbitrary data defining the type of credential. The minimum size is 1 byte and the maximum is 64 bytes.
	CredentialType types.CredentialType
}

// TxType implements the TxType method for the CredentialCreate struct.
func (*CredentialAccept) TxType() TxType {
	return CredentialAcceptTx
}

// Flatten implements the Flatten method for the CredentialCreate struct.
func (c *CredentialAccept) Flatten() FlatTransaction {
	flattened := c.BaseTx.Flatten()

	flattened["TransactionType"] = c.TxType().String()

	if c.Issuer != "" {
		flattened["Issuer"] = c.Issuer.String()
	}

	if c.CredentialType != "" {
		flattened["CredentialType"] = c.CredentialType.String()
	}

	return flattened
}

// Validate implements the Validate method for the CredentialCreate struct.
func (c *CredentialAccept) Validate() (bool, error) {
	// validate the base transaction
	_, err := c.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if !addresscodec.IsValidAddress(c.Issuer.String()) {
		return false, ErrInvalidIssuer
	}

	if !c.CredentialType.IsValid() {
		return false, ErrInvalidCredentialType
	}

	return true, nil
}
