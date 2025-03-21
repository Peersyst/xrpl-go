package types

import (
	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
)

type AuthorizeCredentialsWrapper struct {
	Credential AuthorizeCredentials
}

type AuthorizeCredentials struct {
	// The issuer of the credential.
	Issuer Address
	// The credential type of the credential.
	CredentialType CredentialType
}

// IsValid returns true if the authorize credentials are valid.
func (a *AuthorizeCredentials) IsValid() bool {
	return addresscodec.IsValidAddress(a.Issuer.String()) && a.CredentialType.IsValid()
}

// Flatten returns a map of the authorize credentials.
func (a *AuthorizeCredentialsWrapper) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})

	flattened["Credential"] = a.Credential.Flatten()

	return flattened
}

// Flatten returns a map of the authorize credentials.
func (a *AuthorizeCredentials) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})

	if a.Issuer != "" {
		flattened["Issuer"] = a.Issuer.String()
	}
	if a.CredentialType != "" {
		flattened["CredentialType"] = a.CredentialType.String()
	}

	return flattened
}
