package keypairs

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
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

	// convert this typescript code to Golang:
	// const publicGen = secp256k1.ProjectivePoint.BASE.multiply(privateGen).toRawBytes(true)
	publicKey := new(btcec.PublicKey)
	x, y := secp256k1.S256().ScalarBaseMult(privateGenerator.Bytes())

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

const HEX_REGEX = "^[A-F0-9]*$"

// bytesToHex converts a byte slice to a hexadecimal string in uppercase.
func BytesToHex(bytes []byte) string {
	hexStr := hex.EncodeToString(bytes)
	return strings.ToUpper(hexStr)
}

// bytesToNumberBE converts a big-endian byte slice to a big integer.
func BytesToNumberBE(bytes []byte) *big.Int {
	n := new(big.Int)
	n.SetBytes(bytes)
	return n
}

// bytesToNumberLE converts a little-endian byte slice to a big integer.
func BytesToNumberLE(bytes []byte) *big.Int {
	reversedBytes := ReverseBytes(bytes)
	return BytesToNumberBE(reversedBytes)
}

// numberToBytesBE converts a big integer to a big-endian byte slice of the specified length.
func NumberToBytesBE(n *big.Int, length int) []byte {
	bytes := n.Bytes()
	if len(bytes) < length {
		paddedBytes := make([]byte, length)
		copy(paddedBytes[length-len(bytes):], bytes)
		return paddedBytes
	}
	return bytes
}

// numberToBytesLE converts a big integer to a little-endian byte slice of the specified length.
func NumberToBytesLE(n *big.Int, length int) []byte {
	bytes := NumberToBytesBE(n, length)
	return ReverseBytes(bytes)
}

// numberToVarBytesBE converts a big integer to an unpadded big-endian byte slice.
func NumberToVarBytesBE(n *big.Int) []byte {
	return n.Bytes()
}

// reverseBytes reverses the order of bytes in a slice.
func ReverseBytes(bytes []byte) []byte {
	reversed := make([]byte, len(bytes))
	for i, b := range bytes {
		reversed[len(bytes)-1-i] = b
	}
	return reversed
}
