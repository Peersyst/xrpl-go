package keypairs

import (
	"crypto/rand"
	"fmt"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
)

var r randomizer

const (
	VERIFICATIONMESSAGE = "This test message should verify."
)

func init() {
	r.Reader = rand.Reader
}

type CryptoImplementation interface {
	deriveKeypair(decodedSeed []byte, validator bool) (string, string, error)
	sign(msg, privKey string) (string, error)
	validate(msg, pubkey, sig string) bool
}

func GenerateSeed(entropy string, alg addresscodec.CryptoAlgorithm) (string, error) {
	var pe []byte
	if entropy == "" {
		b, err := r.generateBytes(addresscodec.FamilySeedLength)
		pe = b
		if err != nil {
			return "", err
		}
	} else {
		pe = []byte(entropy)[:addresscodec.FamilySeedLength]
	}
	return addresscodec.EncodeSeed(pe, alg)
}

// Derives a keypair from a given seed. Returns a tuple of private key and public key
func DeriveKeypair(seed string, validator bool) (private, public string, err error) {
	ds, alg, err := addresscodec.DecodeSeed(seed)
	if err != nil {
		return
	}
	ci := getCryptoImplementation(alg)
	if ci == nil {
		return "", "", &CryptoImplementationError{}
	}
	private, public, err = ci.deriveKeypair(ds, validator)
	if err != nil {
		return
	}
	signature, err := ci.sign(VERIFICATIONMESSAGE, private)

	if !ci.validate(VERIFICATIONMESSAGE, public, signature) {
		return "", "", &InvalidSignatureError{}
	}
	return
}

func DeriveClassicAddress(pubkey string) (string, error) {
	return addresscodec.EncodeClassicAddressFromPublicKeyHex(pubkey)
}

func Sign(msg, privKey string) (string, error) {
	alg := getCryptoImplementationFromKey(privKey)
	if alg == nil {
		return "", &CryptoImplementationError{}
	}
	return alg.sign(msg, privKey)
}

func Validate(msg, pubKey, sig string) (bool, error) {
	alg := getCryptoImplementationFromKey(pubKey)
	if alg == nil {
		return false, &CryptoImplementationError{}
	}
	return alg.validate(msg, pubKey, sig), nil
}

func getCryptoImplementation(alg addresscodec.CryptoAlgorithm) CryptoImplementation {
	switch alg {
	case addresscodec.ED25519:
		return &ed25519Alg{}
	case addresscodec.SECP256K1:
		return &secp256k1Alg{}
	default:
		return nil
	}
}

func getCryptoImplementationFromKey(k string) CryptoImplementation {
	switch deformatKey(k)[0] {
	case addresscodec.ED25519:
		return &ed25519Alg{}
	default:
		return nil
	}
}

type CryptoImplementationError struct{}

func (e *CryptoImplementationError) Error() string {
	return fmt.Sprintln("not a valid crypto implementation")
}

type InvalidSignatureError struct{}

func (e *InvalidSignatureError) Error() string {
	return "derived keypair did not generate verifiable signature"
}
