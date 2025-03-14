package transaction

import (
	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// A CredentialDelete transaction removes a credential from the ledger, effectively revoking it.
// Users may also want to delete an unwanted credential to reduce their reserve requirement.
type CredentialDelete struct {
	BaseTx

	// Arbitrary data defining the type of credential this entry represents. The minimum length is 1 byte and the maximum length is 64 bytes.
	CredentialType types.CredentialType

	// The subject of the credential to delete. If omitted, use the Account (sender of the transaction) as the subject of the credential.
	Subject types.Address `json:",omitempty"`

	// The issuer of the credential to delete. If omitted, use the Account (sender of the transaction) as the issuer of the credential.
	Issuer types.Address `json:",omitempty"`
}

// TxType returns the type of the CredentialDelete transaction.
func (*CredentialDelete) TxType() TxType {
	return CredentialDeleteTx
}

// Flatten returns a flattened version of the CredentialDelete transaction.
func (c *CredentialDelete) Flatten() FlatTransaction {
	flattened := c.BaseTx.Flatten()

	flattened["TransactionType"] = c.TxType().String()

	if c.CredentialType != "" {
		flattened["CredentialType"] = c.CredentialType.String()
	}

	if c.Subject != "" {
		flattened["Subject"] = c.Subject.String()
	}

	if c.Issuer != "" {
		flattened["Issuer"] = c.Issuer.String()
	}

	return flattened
}

// Validate validates the CredentialDelete transaction.
func (c *CredentialDelete) Validate() (bool, error) {
	// validate the base transaction
	_, err := c.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if !c.CredentialType.IsValid() {
		return false, ErrInvalidCredentialType
	}

	if c.Subject != "" && !addresscodec.IsValidAddress(c.Subject.String()) {
		return false, ErrInvalidSubject
	}

	if c.Issuer != "" && !addresscodec.IsValidAddress(c.Issuer.String()) {
		return false, ErrInvalidIssuer
	}

	return true, nil
}
