package types

import (
	"encoding/hex"
	"errors"
	"strings"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

const (
	MPTIssuanceIDBytesLength = 24
)

var (
	ErrInvalidIssueObject = errors.New("invalid issue object")
	ErrInvalidCurrency    = errors.New("invalid currency")
	ErrInvalidIssuer      = errors.New("invalid issuer")

	XRPBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
)

// Issue represents an XRPL Issue, which is essentially an AccountID.
// It is used to identify the issuer of a currency in the XRPL.
// The FromJson method converts a classic address string to an AccountID byte slice.
// The ToJson method converts an AccountID byte slice back to a classic address string.
// This type is crucial for handling currency issuers in XRPL transactions and ledger entries.
type Issue struct {
	length int
}

// FromJSON parses a classic address string and returns the corresponding AccountID byte slice.
// It uses the addresscodec package to decode the classic address.
// If the input is not a valid classic address, it returns an error.
func (i *Issue) FromJSON(json any) ([]byte, error) {
	if !i.isIssueObject(json) {
		return nil, ErrInvalidIssueObject
	}

	mapObj, ok := json.(map[string]any)
	if !ok {
		return nil, ErrInvalidIssueObject
	}

	currency, ok := mapObj["currency"]
	if !ok {
		mptIssuanceId, ok := mapObj["mpt_issuance_id"].(string)
		if !ok {
			return nil, ErrInvalidCurrency
		}

		mptIssuanceIdBytes, err := hex.DecodeString(mptIssuanceId)
		if err != nil {
			return nil, err
		}

		i.length = MPTIssuanceIDBytesLength

		return mptIssuanceIdBytes, nil
	}

	currencyCodec := &Currency{}

	currencyBytes, err := currencyCodec.FromJSON(currency)
	if err != nil {
		return nil, err
	}

	issuer, ok := mapObj["issuer"]
	if issuerString, okstring := issuer.(string); ok && okstring {
		_, issuerBytes, err := addresscodec.DecodeClassicAddressToAccountID(issuerString)
		if err != nil {
			return nil, err
		}

		return append(currencyBytes, issuerBytes...), nil
	}

	return currencyBytes, nil
}

// ToJSON converts an AccountID byte slice back to a classic address string.
// It uses the addresscodec package to encode the byte slice.
// If the input is not a valid AccountID byte slice, it returns an error.
func (i *Issue) ToJSON(p interfaces.BinaryParser, opts ...int) (any, error) {
	currencyCodec := &Currency{}

	if i.length == MPTIssuanceIDBytesLength || opts[0] == MPTIssuanceIDBytesLength {
		b, err := p.ReadBytes(MPTIssuanceIDBytesLength)
		if err != nil {
			return nil, err
		}

		id := hex.EncodeToString(b)

		return map[string]any{
			"mpt_issuance_id": strings.ToUpper(id),
		}, nil
	}

	currencyStr, err := currencyCodec.ToJSON(p, opts...)
	if err != nil {
		return nil, err
	}

	if currencyStr == "XRP" {
		return map[string]any{
			"currency": "XRP",
		}, nil
	}

	issuer, err := p.ReadBytes(20)
	if err != nil {
		return nil, err
	}

	address, err := addresscodec.Encode(issuer, []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"currency": currencyStr,
		"issuer":   address,
	}, nil
}

func (i *Issue) isIssueObject(obj any) bool {
	mapObj, ok := obj.(map[string]any)
	if !ok {
		return false
	}

	nKeys := len(mapObj)

	_, okMptIssuanceId := mapObj["mpt_issuance_id"]
	if nKeys == 1 && okMptIssuanceId {
		return true
	}

	_, okCurrency := mapObj["currency"]
	if nKeys == 1 && okCurrency {
		return true
	}

	_, okIssuer := mapObj["issuer"]
	if nKeys == 2 && okCurrency && okIssuer {
		return true
	}

	return false
}
