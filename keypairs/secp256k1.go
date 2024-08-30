package keypairs

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/isomorphic/utils"
	"github.com/btcsuite/btcd/btcec/v2"
)

const SECP256K1_PREFIX = "00"

type secp256k1Alg struct{}

func (c *secp256k1Alg) deriveKeypair(decodedSeed []byte, validator bool) (string, string, error) {
	derived := derivePrivateKey(decodedSeed, validator, 0)
	fmt.Println("derived: ", derived)

	privateKeyBytes, err := utils.NumberToBytesBE(derived, 32)
	if err != nil {
		return "", "", err
	}

	p, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	publicKey := p.PubKey()

	privateKey := SECP256K1_PREFIX + formatKey(privateKeyBytes)

	return privateKey, formatKey(publicKey.SerializeCompressed()), nil
}

func (c *secp256k1Alg) sign(msg, privKey string) (string, error) {
	// TODO
	return "", nil
}

func (c *secp256k1Alg) validate(msg, pubkey, sig string) bool {
	// TODO
	return true
}

type secp256k1ValidatorError struct{}

func (e *secp256k1ValidatorError) Error() string {
	return "validator keypairs can not use secp256k1"
}
