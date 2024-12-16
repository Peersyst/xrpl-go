package websocket

import (
	"testing"
	"time"

	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/stretchr/testify/require"
)

func TestNewClientConfig(t *testing.T) {
	config := NewClientConfig()
	require.Equal(t, config.maxRetries, 10)
	require.Equal(t, config.retryDelay, 1*time.Second)
	require.Equal(t, config.host, "localhost")
	require.Equal(t, config.feeCushion, float32(1.2))
	require.Equal(t, config.maxFeeXRP, float32(2.0))
}

func TestWithMaxRetries(t *testing.T) {
	config := NewClientConfig().WithMaxRetries(20)
	require.Equal(t, config.maxRetries, 20)
}

func TestWithRetryDelay(t *testing.T) {
	config := NewClientConfig().WithRetryDelay(2 * time.Second)
	require.Equal(t, config.retryDelay, 2*time.Second)
}

func TestWithFeeCushion(t *testing.T) {
	config := NewClientConfig().WithFeeCushion(1.5)
	require.Equal(t, config.feeCushion, float32(1.5))
}

func TestWithMaxFeeXRP(t *testing.T) {
	config := NewClientConfig().WithMaxFeeXRP(3.0)
	require.Equal(t, config.maxFeeXRP, float32(3.0))
}

func TestWithFaucetProvider(t *testing.T) {
	config := NewClientConfig().WithFaucetProvider(faucet.NewTestnetFaucetProvider())
	require.NotNil(t, config.faucetProvider)
}