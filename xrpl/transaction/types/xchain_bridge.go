package types

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
)

var (
	// ErrInvalidIssuingChainDoorAddress is returned when the issuing chain door address is invalid.
	ErrInvalidIssuingChainDoorAddress = errors.New("xchainBridge: invalid issuing chain door address")
	// ErrInvalidIssuingChainIssueAddress is returned when the issuing chain issue address is invalid.
	ErrInvalidIssuingChainIssueAddress = errors.New("xchainBridge: invalid issuing chain issue address")
	// ErrInvalidLockingChainDoorAddress is returned when the locking chain door address is invalid.
	ErrInvalidLockingChainDoorAddress = errors.New("xchainBridge: invalid locking chain door address")
	// ErrInvalidLockingChainIssueAddress is returned when the locking chain issue address is invalid.
	ErrInvalidLockingChainIssueAddress = errors.New("xchainBridge: invalid locking chain issue address")
)

// XChainBridge represents the bridge configuration for cross-chain transfers, including door accounts and assets.
type XChainBridge struct {
	// The door account on the issuing chain. For an XRP-XRP bridge, this must be the
	// genesis account (the account that is created when the network is first started, which contains all of the XRP).
	IssuingChainDoor Address
	// The asset that is minted and burned on the issuing chain. For an IOU-IOU bridge,
	// the issuer of the asset must be the door account on the issuing chain, to avoid supply issues.
	IssuingChainIssue Address
	// The door account on the locking chain.
	LockingChainDoor Address
	// The asset that is locked and unlocked on the locking chain.
	LockingChainIssue Address
}

// FlatXChainBridge is a flattened representation of XChainBridge for JSON serialization.
type FlatXChainBridge map[string]string

// Flatten returns a FlatXChainBridge mapping fields to their string values.
func (x *XChainBridge) Flatten() FlatXChainBridge {
	flat := make(FlatXChainBridge)

	flat["IssuingChainDoor"] = x.IssuingChainDoor.String()
	flat["IssuingChainIssue"] = x.IssuingChainIssue.String()
	flat["LockingChainDoor"] = x.LockingChainDoor.String()
	flat["LockingChainIssue"] = x.LockingChainIssue.String()

	return flat
}

// Validate checks each address in the XChainBridge and returns false with an error if any are invalid.
func (x *XChainBridge) Validate() (bool, error) {
	if !addresscodec.IsValidAddress(x.IssuingChainDoor.String()) {
		return false, ErrInvalidIssuingChainDoorAddress
	}
	if !addresscodec.IsValidAddress(x.IssuingChainIssue.String()) {
		return false, ErrInvalidIssuingChainIssueAddress
	}
	if !addresscodec.IsValidAddress(x.LockingChainDoor.String()) {
		return false, ErrInvalidLockingChainDoorAddress
	}
	if !addresscodec.IsValidAddress(x.LockingChainIssue.String()) {
		return false, ErrInvalidLockingChainIssueAddress
	}

	return true, nil
}
