package common

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

// FaucetProvider defines an interface for interacting with XRPL faucets.
// Implementations of this interface can be used to fund wallets on different
// XRPL networks (e.g., Devnet, Testnet) by requesting XRP from their respective faucets.
type FaucetProvider interface {
	// FundWallet sends a request to the faucet to fund the specified wallet address.
	// It returns an error if the funding request fails.
	FundWallet(address types.Address) error
}
