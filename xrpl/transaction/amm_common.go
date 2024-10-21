package transaction

// Common flags for AMM transactions (Deposit and Withdraw).
const (
	tfLPToken         uint = 65536
	tfSingleAsset     uint = 524288
	tfTwoAsset        uint = 1048576
	tfOneAssetLPToken uint = 2097152
	tfLimitLPToken    uint = 4194304
)
