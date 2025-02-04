package types

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
)

var (
	ErrInvalidIssuingChainDoorAddress  = errors.New("xchainBridge: invalid issuing chain door address")
	ErrInvalidIssuingChainIssueAddress = errors.New("xchainBridge: invalid issuing chain issue address")
	ErrInvalidLockingChainDoorAddress  = errors.New("xchainBridge: invalid locking chain door address")
	ErrInvalidLockingChainIssueAddress = errors.New("xchainBridge: invalid locking chain issue address")
)

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

type FlatXChainBridge map[string]string

func (x *XChainBridge) Flatten() FlatXChainBridge {
	flat := make(FlatXChainBridge)

	flat["IssuingChainDoor"] = x.IssuingChainDoor.String()
	flat["IssuingChainIssue"] = x.IssuingChainIssue.String()
	flat["LockingChainDoor"] = x.LockingChainDoor.String()
	flat["LockingChainIssue"] = x.LockingChainIssue.String()

	return flat
}

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
