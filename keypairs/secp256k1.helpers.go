package keypairs

import (
	"crypto/sha512"
	"fmt"
	"math"
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

	x, y := btcec.S256().ScalarBaseMult(privateGenerator.Bytes())

	// Convert x and y to Uint8Array
	xBytes := x.Bytes()
	yBytes := y.Bytes()

	// Create a Uint8Array with the concatenated x and y bytes
	uint8Array := make([]uint8, len(xBytes)+len(yBytes))
	copy(uint8Array[:len(xBytes)], xBytes)
	copy(uint8Array[len(xBytes):], yBytes)

	// Get the public generator point
	// publicGenerator := secp256k1.NewPublicKey(&secp256k1.FieldVal{X: *x}, &secp256k1.FieldVal{Y: *y})
	// var publicGenerator = btcec.NewPublicKey(x, y)

	result := new(big.Int).Add(deriveScalar(uint8Array, accountIndex), privateGenerator)
	// module := math.Mod(result, order)
	resultFloat := new(big.Float).SetInt(result)
	resultFloat64, _ := resultFloat.Float64()

	orderFloat := new(big.Float).SetInt(order)
	orderFloat64, _ := orderFloat.Float64()

	modResult := math.Mod(resultFloat64, orderFloat64)
	// convert modResult to big.Int
	modResultInt := new(big.Int).SetInt64(int64(modResult))
	fmt.Println(modResultInt)
	return modResultInt
}

func deriveScalar(bytes []uint8, discrim uint) *big.Int {
	var order = btcec.S256().N

	for i := 0; i <= 0xffff_ffff; i++ {
		// We hash the bytes to find a 256-bit number, looping until we are sure it
		// is less than the order of the curve.
		var hasher = sha512.New()
		hasher.Write(bytes)

		// If the optional discriminator index was passed in, update the hash.
		if discrim != 0 {
			hasher.Write([]byte{byte(discrim)})
		}

		hasher.Write([]byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)})
		key := new(big.Int).SetBytes(hasher.Sum(nil))

		fmt.Println(key)
		// Check if the key is within the valid range
		if key.Cmp(big.NewInt(0)) > 0 && key.Cmp(order) < 0 {
			return key
		}
	}

	panic("failed to derive scalar")
}
