package types

import (
	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/binary-codec/types/errors"
	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

// ===============================
// Errors
// ===============================

// ErrNotValidXChainBridge is an error that occurs when the
// xchain bridge objectis not valid.
type ErrNotValidXChainBridge struct{}

// Error returns the error message.
func (e *ErrNotValidXChainBridge) Error() string {
	return "not a valid xchain bridge"
}

// ErrReadBytes is an error that occurs when trying to read bytes with a parser.
type ErrReadBytes struct{}

// Error returns the error message.
func (e *ErrReadBytes) Error() string {
	return "read bytes error"
}

// ErrDecodeClassicAddress is an error that occurs
// when trying to decode a classic address.
type ErrDecodeClassicAddress struct{}

// Error returns the error message.
func (e *ErrDecodeClassicAddress) Error() string {
	return "decode classic address error"
}

// ===============================
// XChainBridge
// ===============================

// XChainBridge is a struct that represents an xchain bridge.
type XChainBridge struct{}

// FromJson converts a json XChainBridge object to its byte slice representation.
// It returns an error if the json is not valid or if the classic addresses are not valid.
func (x *XChainBridge) FromJson(json any) ([]byte, error) {
	v, ok := json.(map[string]any)
	if !ok {
		return nil, &errors.ErrNotValidJson{}
	}

	if v["LockingChainDoor"] == nil || v["LockingChainIssue"] == nil || v["IssuingChainDoor"] == nil || v["IssuingChainIssue"] == nil {
		return nil, &ErrNotValidXChainBridge{}
	}

	_, lockingChainDoor, err := addresscodec.DecodeClassicAddressToAccountID(v["LockingChainDoor"].(string))
	if err != nil {
		return nil, &ErrDecodeClassicAddress{}
	}

	_, lockingChainIssue, err := addresscodec.DecodeClassicAddressToAccountID(v["LockingChainIssue"].(string))
	if err != nil {
		return nil, &ErrDecodeClassicAddress{}
	}

	_, issuingChainDoor, err := addresscodec.DecodeClassicAddressToAccountID(v["IssuingChainDoor"].(string))
	if err != nil {
		return nil, &ErrDecodeClassicAddress{}
	}

	_, issuingChainIssue, err := addresscodec.DecodeClassicAddressToAccountID(v["IssuingChainIssue"].(string))
	if err != nil {
		return nil, &ErrDecodeClassicAddress{}
	}

	bytes := make([]byte, 0, 80)

	bytes = append(bytes, lockingChainDoor...)
	bytes = append(bytes, lockingChainIssue...)
	bytes = append(bytes, issuingChainDoor...)
	bytes = append(bytes, issuingChainIssue...)

	return bytes, nil
}

// ToJson converts a byte slice representation of an XChainBridge object to its json representation.
// It returns an error if the bytes are not valid or if the classic addresses are not valid.
func (x *XChainBridge) ToJson(p interfaces.BinaryParser, opts ...int) (any, error) {
	if opts == nil {
		return nil, ErrNoLengthPrefix
	}

	bytes, err := p.ReadBytes(opts[0])
	if err != nil {
		return nil, &ErrReadBytes{}
	}

	json := make(map[string]string)

	json["LockingChainDoor"] = addresscodec.Encode(bytes[:20], []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength)
	json["LockingChainIssue"] = addresscodec.Encode(bytes[20:40], []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength)
	json["IssuingChainDoor"] = addresscodec.Encode(bytes[40:60], []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength)
	json["IssuingChainIssue"] = addresscodec.Encode(bytes[60:80], []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength)

	return json, nil
}
