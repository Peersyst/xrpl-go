package types

import (
	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	"github.com/Peersyst/xrpl-go/v1/binary-codec/types/interfaces"
)

// Issue represents an XRPL Issue, which is essentially an AccountID.
// It is used to identify the issuer of a currency in the XRPL.
// The FromJson method converts a classic address string to an AccountID byte slice.
// The ToJson method converts an AccountID byte slice back to a classic address string.
// This type is crucial for handling currency issuers in XRPL transactions and ledger entries.
type Issue struct{}

// FromJSON parses a classic address string and returns the corresponding AccountID byte slice.
// It uses the addresscodec package to decode the classic address.
// If the input is not a valid classic address, it returns an error.
func (i *Issue) FromJSON(json any) ([]byte, error) {
	_, accountID, err := addresscodec.DecodeClassicAddressToAccountID(json.(string))

	if err != nil {
		return nil, err
	}

	return accountID, nil
}

// ToJSON converts an AccountID byte slice back to a classic address string.
// It uses the addresscodec package to encode the byte slice.
// If the input is not a valid AccountID byte slice, it returns an error.
func (i *Issue) ToJSON(p interfaces.BinaryParser, opts ...int) (any, error) {
	if opts == nil {
		return nil, ErrNoLengthPrefix
	}
	b, err := p.ReadBytes(opts[0])
	if err != nil {
		return nil, err
	}
	return addresscodec.Encode(b, []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength)
}
