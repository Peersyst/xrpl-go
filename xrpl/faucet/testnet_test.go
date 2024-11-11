package faucet

import (
	"testing"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/pkg/random"
	"github.com/Peersyst/xrpl-go/xrpl"
)

// Note: This test interacts with the actual Testnet faucet.
// The result and behavior may vary based on the faucet's current state and rate limits.
// Manual verification of the printed result is recommended.
func TestTestnetFaucetProvider_FundWallet(t *testing.T) {

	// Create a new TestnetFaucetProvider
	provider := NewTestnetFaucetProvider()

	// Test wallet address
	testWallet, err := xrpl.NewWallet(crypto.ED25519(), random.NewRandomizer())
	if err != nil {
		t.Errorf("Wallet creation error: %v", err)
	}

	// Call FundWallet
	err = provider.FundWallet(testWallet.ClassicAddress)

	// Check for errors
	if err != nil {
		t.Errorf("FundWallet returned an error: %v", err)
	}
}
