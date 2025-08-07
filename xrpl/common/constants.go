//revive:disable var-naming
package common

import "time"

const (
	// LedgerOffset is the number of ledgers to offset when querying ledger data.
	LedgerOffset uint32 = 20

	// DefaultHost is the default host for the XRPL server.
	DefaultHost = "localhost"
	// DefaultMaxRetries is the default maximum number of retries for RPC calls.
	DefaultMaxRetries = 10
	// DefaultMaxReconnects is the default maximum number of reconnect attempts for websocket.
	DefaultMaxReconnects = 3
	// DefaultRetryDelay is the default delay between retry attempts.
	DefaultRetryDelay = 1 * time.Second
	// DefaultFeeCushion is the default fee cushion multiplier.
	DefaultFeeCushion float32 = 1.2
	// DefaultMaxFeeXRP is the default maximum fee in XRP.
	DefaultMaxFeeXRP float32 = 2

	// DefaultTimeout is the default timeout for RPC calls (5 seconds).
	DefaultTimeout = 5 * time.Second
)
