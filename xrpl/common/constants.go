package common

import "time"

const (
	// Ledger constants
	LedgerOffset uint32 = 20

	// Config constants
	DefaultHost               = "localhost"
	DefaultMaxRetries         = 10
	DefaultRetryDelay         = 1 * time.Second
	DefaultFeeCushion float32 = 1.2
	DefaultMaxFeeXRP  float32 = 2

	// 5 seconds default timeout
	DefaultTimeout = 5 * time.Second
)
