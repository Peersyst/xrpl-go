package types

type XChainBridge struct {
	// The door account on the issuing chain. For an XRP-XRP bridge, this must be the
	//genesis account (the account that is created when the network is first started, which contains all of the XRP).
	IssuingChainDoor Address
	// The asset that is minted and burned on the issuing chain. For an IOU-IOU bridge,
	//the issuer of the asset must be the door account on the issuing chain, to avoid supply issues.
	IssuingChainIssue Address
	// The door account on the locking chain.
	LockingChainDoor Address
	// The asset that is locked and unlocked on the locking chain.
	LockingChainIssue Address
}
