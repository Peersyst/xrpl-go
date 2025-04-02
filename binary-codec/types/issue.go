package types

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

var (
	ErrUnexpectedIssueJSONType = errors.New("unexpected type for Issue JSON")
	ErrCurrencyFieldMissing    = errors.New("currency field missing")
	ErrCurrencyFieldNotString  = errors.New("currency field must be a string")
	ErrFailedToDecodeCurrency  = errors.New("failed to decode currency")
	ErrIssuerFieldNotString    = errors.New("issuer field must be a string")
	ErrFailedToDecodeIssuer    = errors.New("failed to decode issuer")
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
	// Handle classic address string.
	if s, ok := json.(string); ok {
		_, accountID, err := addresscodec.DecodeClassicAddressToAccountID(s)
		if err != nil {
			return nil, err
		}
		return accountID, nil
	}

	// Otherwise, expect a map.
	m, ok := json.(map[string]interface{})
	if !ok {
		return nil, ErrUnexpectedIssueJSONType
	}

	// Extract the currency field.
	currencyVal, ok := m["currency"]
	if !ok {
		return nil, ErrCurrencyFieldMissing
	}
	currencyStr, ok := currencyVal.(string)
	if !ok {
		return nil, ErrCurrencyFieldNotString
	}

	// If currency is "XRP", no issuer is needed.
	if currencyStr == "XRP" {
		return []byte("XRP"), nil
	}

	// Use the helper from the Amount type to decode the currency.
	currencyBytes, err := serializeIssuedCurrencyCode(currencyStr)
	if err != nil {
		return nil, ErrFailedToDecodeCurrency
	}

	// If an issuer is provided, decode it and concatenate.
	if issuerVal, exists := m["issuer"]; exists {
		issuerStr, ok := issuerVal.(string)
		if !ok {
			return nil, ErrIssuerFieldNotString
		}
		_, issuerBytes, err := addresscodec.DecodeClassicAddressToAccountID(issuerStr)
		if err != nil {
			return nil, ErrFailedToDecodeIssuer
		}
		combined := append(currencyBytes, issuerBytes...)
		return combined, nil
	}

	// Return just the currency bytes if no issuer.
	return currencyBytes, nil
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
