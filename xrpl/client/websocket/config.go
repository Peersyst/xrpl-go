package websocket

type WebsocketClientConfig struct {
	// Connection config
	host string

	// Fee config
	feeCushion float32
	maxFeeXRP  float32

	// Faucet config
	faucetProvider FaucetProvider
}

func NewWebsocketClientConfig() *WebsocketClientConfig {
	return &WebsocketClientConfig{
		host:       "localhost",
		feeCushion: DEFAULT_FEE_CUSHION,
		maxFeeXRP:  DEFAULT_MAX_FEE_XRP,
	}
}

// WithHost sets the host of the websocket client.
// Default: "localhost"
func (wc WebsocketClientConfig) WithHost(host string) WebsocketClientConfig {
	wc.host = host
	return wc
}

// WithFeeCushion sets the fee cushion of the websocket client.
// Default: 1.2
func (wc WebsocketClientConfig) WithFeeCushion(feeCushion float32) WebsocketClientConfig {
	wc.feeCushion = feeCushion
	return wc
}

// WithMaxFeeXRP sets the maximum fee in XRP that the websocket client will use.
// Default: 2
func (wc WebsocketClientConfig) WithMaxFeeXRP(maxFeeXrp float32) WebsocketClientConfig {
	wc.maxFeeXRP = maxFeeXrp
	return wc
}

// WithFaucetProvider sets the faucet provider of the websocket client.
// Default: faucet.NewLocalFaucetProvider()
func (wc WebsocketClientConfig) WithFaucetProvider(fp FaucetProvider) WebsocketClientConfig {
	wc.faucetProvider = fp
	return wc
}
