package integration

import (
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
)

// EnvKey is the key for the integration environment.
type EnvKey string

const (
	LocalnetEnv EnvKey = "localnet"
	TestnetEnv  EnvKey = "testnet"
	DevnetEnv   EnvKey = "devnet"
)

// IntegrationWebsocketEnvs is the map of integration environments.
var IntegrationWebsocketEnvs = map[EnvKey]Env{
	LocalnetEnv: {
		Host:           "wss://0.0.0.0:6006",
		FaucetProvider: nil,
	},
	TestnetEnv: {
		Host:           "wss://s.altnet.rippletest.net:51233",
		FaucetProvider: faucet.NewTestnetFaucetProvider(),
	},
	DevnetEnv: {
		Host:           "wss://s.devnet.rippletest.net:51233",
		FaucetProvider: faucet.NewTestnetFaucetProvider(),
	},
}

var IntegrationRPCEnvs = map[EnvKey]Env{
	LocalnetEnv: {
		Host:           "http://0.0.0.0:5005",
		FaucetProvider: nil,
	},
	TestnetEnv: {
		Host:           "https://s.altnet.rippletest.net:51234",
		FaucetProvider: faucet.NewTestnetFaucetProvider(),
	},
	DevnetEnv: {
		Host:           "https://s.devnet.rippletest.net:51234",
		FaucetProvider: faucet.NewTestnetFaucetProvider(),
	},
}

// Env is the environment for the integration tests.
// It contains the host and the faucet provider.
type Env struct {
	Host           string
	FaucetProvider FaucetProvider
}
