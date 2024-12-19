package keypairs

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	"github.com/Peersyst/xrpl-go/v1/keypairs/interfaces"
)

var (
	// Static errors

	// ErrInvalidSignature is returned when the derived keypair did not generate a verifiable signature.
	ErrInvalidSignature = errors.New("derived keypair did not generate verifiable signature")
)

const (
	// verificationMessage is the message that is used to verify the signature of the derived keypair.
	// Only used for testing purposes.
	verificationMessage = "This test message should verify."
)

// GenerateSeed generates a seed from a given entropy, a crypto algorithm implementation and a randomizer.
// If the entropy is empty, it generates a random seed. Otherwise, it uses the entropy to generate the seed.
// The seed is encoded using the addresscodec package.
func GenerateSeed(entropy string, alg interfaces.KeypairCryptoAlg, r interfaces.Randomizer) (string, error) {
	var pe []byte
	var err error
	if entropy == "" {
		pe, err = r.GenerateBytes(addresscodec.FamilySeedLength)
		if err != nil {
			return "", err
		}
	} else {
		pe = []byte(entropy)[:addresscodec.FamilySeedLength]
	}
	return addresscodec.EncodeSeed(pe, alg)
}

// Derives a keypair from a given seed. Returns a tuple of private key and public key.
// The seed has to be encoded using the addresscodec package. Otherwise, it returns an error.
func DeriveKeypair(seed string, validator bool) (private, public string, err error) {
	ds, alg, err := addresscodec.DecodeSeed(seed)
	if err != nil {
		return "", "", err
	}
	private, public, err = alg.DeriveKeypair(ds, validator)
	if err != nil {
		return "", "", err
	}
	signature, err := alg.Sign(verificationMessage, private)
	if err != nil {
		return "", "", err
	}
	if !alg.Validate(verificationMessage, public, signature) {
		return "", "", ErrInvalidSignature
	}
	return private, public, nil
}

// DeriveClassicAddress derives a classic address from a given public key.
// The public key has to be encoded using the addresscodec package. Otherwise, it returns an error.
func DeriveClassicAddress(pubKey string) (string, error) {
	return addresscodec.EncodeClassicAddressFromPublicKeyHex(pubKey)
}

// DeriveNodeAddress derives a node address from a given public key.
// The public key has to be encoded using the addresscodec package. Otherwise, it returns an error.
func DeriveNodeAddress(pubKey string, alg interfaces.NodeDerivationCryptoAlg) (string, error) {
	decoded, err := addresscodec.DecodeNodePublicKey(pubKey)
	if err != nil {
		return "", err
	}
	accountPubKey, err := alg.DerivePublicKeyFromPublicGenerator(decoded)
	if err != nil {
		return "", err
	}

	accountID := addresscodec.Sha256RipeMD160(accountPubKey)

	return addresscodec.EncodeAccountIDToClassicAddress(accountID)
}

// Sign signs a message with a given private key.
// The private key needs to satisfy a crypto algorithm implementation. Otherwise, it returns an error.
// Currently, only ED25519 and SECP256K1 are supported.
// If the message is empty, it returns an error.
func Sign(msg, privKey string) (string, error) {
	alg := getCryptoImplementationFromKey(privKey)
	if alg == nil {
		return "", ErrInvalidCryptoImplementation
	}
	return alg.Sign(msg, privKey)
}

// Validate validates a signature of a message with a given public key.
// The public key needs to satisfy a crypto algorithm implementation. Otherwise, it returns an error.
// Currently, only ED25519 and SECP256K1 are supported.
// If the message is empty, it returns an error.
func Validate(msg, pubKey, sig string) (bool, error) {
	alg := getCryptoImplementationFromKey(pubKey)
	if alg == nil {
		return false, ErrInvalidCryptoImplementation
	}
	return alg.Validate(msg, pubKey, sig), nil
}
