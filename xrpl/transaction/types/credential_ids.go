package types

import "github.com/Peersyst/xrpl-go/pkg/typecheck"

type CredentialIDs []string

func (c CredentialIDs) IsValid() bool {
	if len(c) == 0 {
		return false
	}

	for _, id := range c {
		if !typecheck.IsHex(id) {
			return false
		}
	}

	return true
}

func (c CredentialIDs) Flatten() []string {
	return c
}
