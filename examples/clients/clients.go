package clients

import (
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

// GetDevnetRPCClient returns a new rpc client for the devnet
func GetDevnetRPCClient() *rpc.Client {
	cfg, err := rpc.NewClientConfig(
		"https://s.devnet.rippletest.net:51234",
		rpc.WithFaucetProvider(faucet.NewDevnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}

	return rpc.NewClient(cfg)
}

// GetTestnetRPCClient returns a new rpc client for the testnet
func GetTestnetRPCClient() *rpc.Client {
	cfg, err := rpc.NewClientConfig(
		"https://s.altnet.rippletest.net:51234",
		rpc.WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}

	return rpc.NewClient(cfg)
}

// GetDevnetWebsocketClient returns a new websocket client for the devnet
func GetDevnetWebsocketClient() *websocket.Client {
	client := websocket.NewClient(
		websocket.NewClientConfig().
			WithHost("wss://s.devnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewDevnetFaucetProvider()),
	)

	return client
}

// GetTestnetWebsocketClient returns a new websocket client for the testnet
func GetTestnetWebsocketClient() *websocket.Client {
	client := websocket.NewClient(
		websocket.NewClientConfig().
			WithHost("wss://s.altnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)

	return client
}
