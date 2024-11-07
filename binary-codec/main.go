package binarycodec

import (
	"bytes"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/Peersyst/xrpl-go/binary-codec/definitions"

	"github.com/Peersyst/xrpl-go/binary-codec/serdes"
	"github.com/Peersyst/xrpl-go/binary-codec/types"
)

var ErrSigningClaimFieldNotFound = errors.New("'Channel' & 'Amount' fields are both required, but were not found")

const (
	txMultiSigPrefix          = "534D5400"
	paymentChannelClaimPrefix = "434C4D00"
	txSigPrefix               = "53545800"
)

// Encode converts a JSON transaction object to a hex string in the canonical binary format.
// The binary format is defined in XRPL's core codebase.
func Encode(json map[string]any) (string, error) {

	st := &types.STObject{}

	// Iterate over the keys in the provided JSON
	for k := range json {

		// Get the FieldIdNameMap from the definitions package
		fh := definitions.Get().Fields[k]

		// If the field is not found in the FieldIdNameMap, delete it from the JSON

		if fh == nil {
			delete(json, k)
			continue
		}
	}

	b, err := st.FromJSON(json)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString(b)), nil
}

// EncodeForMultiSign: encodes a transaction into binary format in preparation for providing one
// signature towards a multi-signed transaction.
// (Only encodes fields that are intended to be signed.)
func EncodeForMultisigning(json map[string]any, xrpAccountID string) (string, error) {

	st := &types.AccountID{}

	// SigningPubKey is required for multi-signing but should be set to empty string.

	json["SigningPubKey"] = ""

	suffix, err := st.FromJSON(xrpAccountID)
	if err != nil {
		return "", err
	}

	encoded, err := Encode(removeNonSigningFields(json))

	if err != nil {
		return "", err
	}

	return strings.ToUpper(txMultiSigPrefix + encoded + hex.EncodeToString(suffix)), nil
}

// Encodes a transaction into binary format in preparation for signing.
func EncodeForSigning(json map[string]any) (string, error) {

	encoded, err := Encode(removeNonSigningFields(json))

	if err != nil {
		return "", err
	}

	return strings.ToUpper(txSigPrefix + encoded), nil
}

// EncodeForPaymentChannelClaim: encodes a payment channel claim into binary format in preparation for signing.
func EncodeForSigningClaim(json map[string]any) (string, error) {

	if json["Channel"] == nil || json["Amount"] == nil {
		return "", ErrSigningClaimFieldNotFound
	}

	channel, err := types.NewHash256().FromJSON(json["Channel"])

	if err != nil {
		return "", err
	}

	t := &types.Amount{}
	amount, err := t.FromJSON(json["Amount"])

	if err != nil {
		return "", err

	}

	if bytes.HasPrefix(amount, []byte{0x40}) {
		amount = bytes.Replace(amount, []byte{0x40}, []byte{0x00}, 1)
	}

	return strings.ToUpper(paymentChannelClaimPrefix + hex.EncodeToString(channel) + hex.EncodeToString(amount)), nil
}

// removeNonSigningFields removes the fields from a JSON transaction object that should not be signed.
func removeNonSigningFields(json map[string]any) map[string]any {
	for k := range json {
		fi, _ := definitions.Get().GetFieldInstanceByFieldName(k)

		if fi != nil && !fi.IsSigningField {
			delete(json, k)
		}
	}

	return json
}

// Decode decodes a hex string in the canonical binary format into a JSON transaction object.
func Decode(hexEncoded string) (map[string]any, error) {
	b, err := hex.DecodeString(hexEncoded)
	if err != nil {
		return nil, err
	}
	p := serdes.NewBinaryParser(b)
	st := &types.STObject{}
	m, err := st.ToJSON(p)
	if err != nil {
		return nil, err
	}

	return m.(map[string]any), nil
}
