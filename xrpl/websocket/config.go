package websocket

import (
	"time"

	"github.com/Peersyst/xrpl-go/v1/xrpl/common"
)

type ClientConfig struct {
	// Connection config
	host       string
	maxRetries int
	retryDelay time.Duration

	// Fee config
	feeCushion float32
	maxFeeXRP  float32

	// Faucet config
	faucetProvider common.FaucetProvider
}

func NewClientConfig() *ClientConfig {
	return &ClientConfig{
		host:       common.DefaultHost,
		feeCushion: common.DefaultFeeCushion,
		maxFeeXRP:  common.DefaultMaxFeeXRP,
		maxRetries: common.DefaultMaxRetries,
		retryDelay: common.DefaultRetryDelay,
	}
}

// WithHost sets the host of the websocket client.
// Default: "localhost"
func (wc ClientConfig) WithHost(host string) ClientConfig {
	wc.host = host
	return wc
}

// WithFeeCushion sets the fee cushion of the websocket client.
// Default: 1.2
func (wc ClientConfig) WithFeeCushion(feeCushion float32) ClientConfig {
	wc.feeCushion = feeCushion
	return wc
}

// WithMaxFeeXRP sets the maximum fee in XRP that the websocket client will use.
// Default: 2
func (wc ClientConfig) WithMaxFeeXRP(maxFeeXrp float32) ClientConfig {
	wc.maxFeeXRP = maxFeeXrp
	return wc
}

// WithFaucetProvider sets the faucet provider of the websocket client.
// Default: faucet.NewLocalFaucetProvider()
func (wc ClientConfig) WithFaucetProvider(fp common.FaucetProvider) ClientConfig {
	wc.faucetProvider = fp
	return wc
}

// WithMaxRetries sets the maximum number of retries for a transaction.
// Default: 10
func (wc ClientConfig) WithMaxRetries(maxRetries int) ClientConfig {
	wc.maxRetries = maxRetries
	return wc
}

// WithRetryDelay sets the delay between retries for a transaction.
// Default: 1 second
func (wc ClientConfig) WithRetryDelay(retryDelay time.Duration) ClientConfig {
	wc.retryDelay = retryDelay
	return wc
}
