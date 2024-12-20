package websocket

type ClientConfig struct {
	// Connection config
	host string

	// Fee config
	feeCushion float32
	maxFeeXRP  float32

	// Faucet config
	faucetProvider FaucetProvider
}

func NewClientConfig() *ClientConfig {
	return &ClientConfig{
		host:       "localhost",
		feeCushion: DefaultFeeCushion,
		maxFeeXRP:  DefaultMaxFeeXRP,
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
func (wc ClientConfig) WithFaucetProvider(fp FaucetProvider) ClientConfig {
	wc.faucetProvider = fp
	return wc
}
