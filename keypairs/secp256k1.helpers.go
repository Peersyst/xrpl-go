package keypairs

import (
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
)

func derivePrivateKey(seed []byte, validator bool, accountIndex uint) *big.Int {
	root := validator
	order := btcec.S256().N

	// This private generator represents the `root` private key, and is what's
	// used by validators for signing when a keypair is generated from a seed.

	privateGenerator := deriveScalar(seed, 0)
	if root {
		// As returned by validation_create for a given seed
		return privateGenerator
	}

	publicKey := new(btcec.PublicKey)
	x, y := btcec.S256().ScalarBaseMult(privateGenerator.Bytes())

	publicKey.X().SetBytes(x.Bytes())
	publicKey.Y().SetBytes(y.Bytes())

	// A seed can generate many keypairs as a function of the seed and a uint32.
	// Almost everyone just uses the first account, `0`.
	var acctIndex uint
	if accountIndex != 0 {
		acctIndex = accountIndex
	}

	res := new(big.Int).Add(deriveScalar(publicKey.SerializeCompressed(), acctIndex), privateGenerator)

	result := new(big.Int)
	result.Mod(res, order)
	return result
}

func deriveScalar(bytes []uint8, discrim uint) *big.Int {
	var order = btcec.S256().N

	for i := 0; i <= 0xffff_ffff; i++ {
		// We hash the bytes to find a 256-bit number, looping until we are sure it
		// is less than the order of the curve.
		var hasher = NewSha512().Add(bytes)

		// If the optional discriminator index was passed in, update the hash.
		if discrim != 0 {
			hasher.AddU32(uint32(discrim))
		}

		hasher.AddU32(uint32(i))

		key := hasher.First256BigInt()

		// Check if the key is within the valid range
		if key.Cmp(big.NewInt(0)) > 0 && key.Cmp(order) < 0 {
			return key
		}
	}

	panic("failed to derive scalar")
}
