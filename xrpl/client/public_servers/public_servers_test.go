package publicservers

import (
	"testing"
)

func TestNewServerUrls(t *testing.T) {
	serverUrls := NewServerUrls()

	// Check mainnet WebSocket URLs
	mainnetWebSocket := serverUrls.MainnetWebSocket()
	if mainnetWebSocket.XRPLedgerFoundation() != "wss://xrplcluster.com/" {
		t.Errorf("Expected XRPLedgerFoundation URL to be 'wss://xrplcluster.com/', got %s", mainnetWebSocket.XRPLedgerFoundation())
	}
	if mainnetWebSocket.RippleS1() != "wss://s1.ripple.com" {
		t.Errorf("Expected RippleS1 URL to be 'wss://s1.ripple.com', got %s", mainnetWebSocket.RippleS1())
	}
	if mainnetWebSocket.RippleS2() != "wss://s2.ripple.com" {
		t.Errorf("Expected RippleS2 URL to be 'wss://s2.ripple.com', got %s", mainnetWebSocket.RippleS2())
	}

	// Check mainnet JSON-RPC URLs
	mainnetJsonRpc := serverUrls.MainnetJsonRpc()
	if mainnetJsonRpc.XRPLedgerFoundation() != "https://xrplcluster.com" {
		t.Errorf("Expected XRPLedgerFoundation URL to be 'https://xrplcluster.com', got %s", mainnetJsonRpc.XRPLedgerFoundation())
	}
	if mainnetJsonRpc.RippleS1() != "https://s1.ripple.com:51234" {
		t.Errorf("Expected RippleS1 URL to be 'https://s1.ripple.com:51234', got %s", mainnetJsonRpc.RippleS1())
	}
	if mainnetJsonRpc.RippleS2() != "https://s2.ripple.com:51234" {
		t.Errorf("Expected RippleS2 URL to be 'https://s2.ripple.com:51234', got %s", mainnetJsonRpc.RippleS2())
	}

	// Check testnet WebSocket URLs
	testnetWebSocket := serverUrls.TestnetWebSocket()
	if testnetWebSocket.Ripple() != "wss://s.altnet.rippletest.net:51233" {
		t.Errorf("Expected Ripple URL to be 'wss://s.altnet.rippletest.net:51233', got %s", testnetWebSocket.Ripple())
	}
	if testnetWebSocket.XRPLLabs() != "wss://testnet.xrpl-labs.com" {
		t.Errorf("Expected XRPLLabs URL to be 'wss://testnet.xrpl-labs.com', got %s", testnetWebSocket.XRPLLabs())
	}
	if testnetWebSocket.Clio() != "wss://clio.altnet.rippletest.net:51233" {
		t.Errorf("Expected Clio URL to be 'wss://clio.altnet.rippletest.net:51233', got %s", testnetWebSocket.Clio())
	}

	// Check testnet JSON-RPC URLs
	testnetJsonRpc := serverUrls.TestnetJsonRpc()
	if testnetJsonRpc.Ripple() != "https://s.altnet.rippletest.net:51234" {
		t.Errorf("Expected Ripple URL to be 'https://s.altnet.rippletest.net:51234', got %s", testnetJsonRpc.Ripple())
	}
	if testnetJsonRpc.XRPLLabs() != "https://testnet.xrpl-labs.com" {
		t.Errorf("Expected XRPLLabs URL to be 'https://testnet.xrpl-labs.com', got %s", testnetJsonRpc.XRPLLabs())
	}
	if testnetJsonRpc.Clio() != "https://clio.altnet.rippletest.net:51234" {
		t.Errorf("Expected Clio URL to be 'https://clio.altnet.rippletest.net:51234', got %s", testnetJsonRpc.Clio())
	}

	// Check devnet WebSocket URLs
	devnetWebSocket := serverUrls.DevnetWebSocket()
	if devnetWebSocket.Ripple() != "wss://s.devnet.rippletest.net:51233" {
		t.Errorf("Expected Ripple URL to be 'wss://s.devnet.rippletest.net:51233', got %s", devnetWebSocket.Ripple())
	}
	if devnetWebSocket.RippleDevnetClio() != "wss://clio.devnet.rippletest.net:51233" {
		t.Errorf("Expected RippleDevnetClio URL to be 'wss://clio.devnet.rippletest.net:51233', got %s", devnetWebSocket.RippleDevnetClio())
	}

	// Check devnet JSON-RPC URLs
	devnetJsonRpc := serverUrls.DevnetJsonRpc()
	if devnetJsonRpc.Ripple() != "https://s.devnet.rippletest.net:51234" {
		t.Errorf("Expected Ripple URL to be 'https://s.devnet.rippletest.net:51234', got %s", devnetJsonRpc.Ripple())
	}
	if devnetJsonRpc.RippleClio() != "https://clio.devnet.rippletest.net:51234" {
		t.Errorf("Expected RippleClio URL to be 'https://clio.devnet.rippletest.net:51234', got %s", devnetJsonRpc.RippleClio())
	}
}
