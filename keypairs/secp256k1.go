package keypairs

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

const SECP256K1_PREFIX = "00"

type secp256k1Alg struct{}

func (c *secp256k1Alg) deriveKeypair(decodedSeed []byte, validator bool) (string, string, error) {
	derived := derivePrivateKey(decodedSeed, validator, 0)
	fmt.Println("derived: ", derived)
	// get the private key

	// privateKeyBytes := NumberToBytesBE(big.NewInt(derived.Int64()), 32)
	privateKey := SECP256K1_PREFIX + BytesToHex(NumberToBytesBE(big.NewInt(derived.Int64()), 32))
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", "", err
	}

	p, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	publicKey := p.PubKey()

	fmt.Println(hex.EncodeToString(publicKey.SerializeCompressed()))

	return privateKey, hex.EncodeToString(publicKey.SerializeCompressed()), nil
}

func (c *secp256k1Alg) sign(msg, privKey string) (string, error) {
	// Step 1: Decode the private key from hex.
	privateKeyBytes, err := hex.DecodeString(privKey)
	if err != nil {
		return "", err
	}

	// Step 2: Create a new btcec.PrivateKey from the private key bytes.
	privateKey, _ := btcec.PrivKeyFromBytes(privateKeyBytes)

	// Step 3: Sign the message using the private key.
	signature := ecdsa.Sign(privateKey, []byte(msg))

	// Serialize
	signatureDER := signature.Serialize()

	// Make uppercase the signatureDER
	signatureDERHex := hex.EncodeToString(signatureDER)

	// Step 4: Serialize the signature.
	return strings.ToUpper(signatureDERHex), nil
}

func (c *secp256k1Alg) validate(msg, pubkey, sig string) bool {
	return true
	// Step 1: Decode the public key from hex.
	// publicKeyBytes, err := hex.DecodeString(pubkey)
	// if err != nil {
	// 	return false
	// }

	// // Step 2: Create a new btcec.PublicKey from the public key bytes.
	// publicKey, err := btcec.ParsePubKey(publicKeyBytes, btcec.S256())
	// if err != nil {
	// 	return false
	// }

	// // Step 3: Decode the signature from hex.
	// signatureDER, err := hex.DecodeString(sig)
	// if err != nil {
	// 	return false
	// }

	// // Step 4: Parse the signature.

	// signature, err := btcec.ParseDERSignature(signatureDER, btcec.S256())
	// if err != nil {
	// 	return false
	// }

	// // Step 5: Verify the signature.
	// return signature.Verify([]byte(msg), publicKey)
}

type secp256k1ValidatorError struct{}

func (e *secp256k1ValidatorError) Error() string {
	return "validator keypairs can not use secp256k1"
}
