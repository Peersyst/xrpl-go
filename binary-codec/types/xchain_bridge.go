package types

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	"github.com/Peersyst/xrpl-go/v1/binary-codec/types/interfaces"
)

// Errors
var (
	errNotValidXChainBridge = errors.New("not a valid xchain bridge")
)

// XChainBridge is a struct that represents an xchain bridge.
type XChainBridge struct{}

// FromJSON converts a json XChainBridge object to its byte slice representation.
// It returns an error if the json is not valid or if the classic addresses are not valid.
func (x *XChainBridge) FromJSON(json any) ([]byte, error) {
	v, ok := json.(map[string]any)
	if !ok {
		return nil, errNotValidJSON
	}

	if v["LockingChainDoor"] == nil || v["LockingChainIssue"] == nil || v["IssuingChainDoor"] == nil || v["IssuingChainIssue"] == nil {
		return nil, errNotValidXChainBridge
	}

	_, lockingChainDoor, err := addresscodec.DecodeClassicAddressToAccountID(v["LockingChainDoor"].(string))
	if err != nil {
		return nil, errDecodeClassicAddress
	}

	_, lockingChainIssue, err := addresscodec.DecodeClassicAddressToAccountID(v["LockingChainIssue"].(string))
	if err != nil {
		return nil, errDecodeClassicAddress
	}

	_, issuingChainDoor, err := addresscodec.DecodeClassicAddressToAccountID(v["IssuingChainDoor"].(string))
	if err != nil {
		return nil, errDecodeClassicAddress
	}

	_, issuingChainIssue, err := addresscodec.DecodeClassicAddressToAccountID(v["IssuingChainIssue"].(string))
	if err != nil {
		return nil, errDecodeClassicAddress
	}

	bytes := make([]byte, 0, 80)

	bytes = append(bytes, lockingChainDoor...)
	bytes = append(bytes, lockingChainIssue...)
	bytes = append(bytes, issuingChainDoor...)
	bytes = append(bytes, issuingChainIssue...)

	return bytes, nil
}

// ToJSON converts a byte slice representation of an XChainBridge object to its json representation.
// It returns an error if the bytes are not valid or if the classic addresses are not valid.
func (x *XChainBridge) ToJSON(p interfaces.BinaryParser, opts ...int) (any, error) {
	if opts == nil {
		return nil, ErrNoLengthPrefix
	}

	bytes, err := p.ReadBytes(opts[0])
	if err != nil {
		return nil, errReadBytes
	}

	json := make(map[string]string)

	json["LockingChainDoor"], err = addresscodec.Encode(bytes[:20], []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength)
	if err != nil {
		return nil, err
	}
	json["LockingChainIssue"], err = addresscodec.Encode(bytes[20:40], []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength)
	if err != nil {
		return nil, err
	}
	json["IssuingChainDoor"], err = addresscodec.Encode(bytes[40:60], []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength)
	if err != nil {
		return nil, err
	}
	json["IssuingChainIssue"], err = addresscodec.Encode(bytes[60:80], []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength)
	if err != nil {
		return nil, err
	}

	return json, nil
}
