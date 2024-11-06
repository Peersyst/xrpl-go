package addresscodec

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	//nolint
	"golang.org/x/crypto/ripemd160"
)

const (
	// Lengths in bytes
	AccountAddressLength   = 20
	AccountPublicKeyLength = 33
	FamilySeedLength       = 16
	NodePublicKeyLength    = 33

	// Account/classic address prefix - value is 0
	AccountAddressPrefix = 0x00
	// Account public key prefix - value is 35
	AccountPublicKeyPrefix = 0x23
	// Family seed prefix - value is 33
	FamilySeedPrefix = 0x21
	// Node/validation public key prefix - value is 28
	NodePublicKeyPrefix = 0x1C
)

type EncodeLengthError struct {
	Instance string
	Input    int
	Expected int
}

func (e *EncodeLengthError) Error() string {
	return fmt.Sprintf("`%v` length should be %v not %v", e.Instance, e.Expected, e.Input)
}

type InvalidClassicAddressError struct {
	Input string
}

func (e *InvalidClassicAddressError) Error() string {
	return fmt.Sprintf("`%v` is an invalid classic address", e.Input)
}

// Returns the base58 encoding of byte slice, with the given type prefix, whilst ensuring that the byte slice is the expected length.
func Encode(b []byte, typePrefix []byte, expectedLength int) string {

	if len(b) != expectedLength {
		return ""
	}

	return Base58CheckEncode(b, typePrefix...)
}

// Returns the byte slice decoding of the base58-encoded string and prefix.
func Decode(b58string string, typePrefix []byte) ([]byte, error) {

	prefixLength := len(typePrefix)

	if !bytes.Equal(DecodeBase58(b58string)[:prefixLength], typePrefix) {
		return nil, errors.New("b58string prefix and typeprefix not equal")
	}

	result, err := Base58CheckDecode(b58string)
	result = result[prefixLength:]

	return result, err
}

// Returns the classic address from public key hex string.
func EncodeClassicAddressFromPublicKeyHex(pubkeyhex string) (string, error) {

	pubkey, err := hex.DecodeString(pubkeyhex)

	if err != nil {
		return "", err
	}

	if len(pubkey) == AccountPublicKeyLength-1 {
		pubkey = append([]byte{crypto.ED25519().Prefix()}, pubkey...)
	} else if len(pubkey) != AccountPublicKeyLength {
		return "", &EncodeLengthError{Instance: "PublicKey", Expected: AccountPublicKeyLength, Input: len(pubkey)}
	}

	accountID := Sha256RipeMD160(pubkey)

	if len(accountID) != AccountAddressLength {
		return "", &EncodeLengthError{Instance: "AccountID", Expected: AccountAddressLength, Input: len(accountID)}
	}

	address := Encode(accountID, []byte{AccountAddressPrefix}, AccountAddressLength)

	if !IsValidClassicAddress(address) {
		return "", &InvalidClassicAddressError{Input: address}
	}

	return address, nil
}

// Returns the decoded 'accountID' byte slice of the classic address.
func DecodeClassicAddressToAccountID(cAddress string) (typePrefix, accountID []byte, err error) {

	if len(DecodeBase58(cAddress)) != 25 {
		return nil, nil, &InvalidClassicAddressError{Input: cAddress}
	}

	return DecodeBase58(cAddress)[:1], DecodeBase58(cAddress)[1:21], nil

}

func IsValidClassicAddress(cAddress string) bool {
	_, _, c := DecodeClassicAddressToAccountID(cAddress)

	return c == nil
}

// Returns a base58 encoding of a seed.
func EncodeSeed(entropy []byte, encodingType CryptoImplementation) (string, error) {

	if len(entropy) != FamilySeedLength {
		return "", &EncodeLengthError{Instance: "Entropy", Input: len(entropy), Expected: FamilySeedLength}
	}

	if encodingType == crypto.ED25519() {
		prefix := []byte{0x01, 0xe1, 0x4b}
		return Encode(entropy, prefix, FamilySeedLength), nil
	} else if secp256k1 := crypto.SECP256K1(); encodingType == secp256k1 {
		prefix := []byte{secp256k1.FamilySeedPrefix()}
		return Encode(entropy, prefix, FamilySeedLength), nil
	}
	return "", errors.New("encoding type must be `ed25519` or `secp256k1`")

}

// Returns decoded seed and its algorithm.
func DecodeSeed(seed string) ([]byte, CryptoImplementation, error) {

	// decoded := DecodeBase58(seed)
	decoded, err := Base58CheckDecode(seed)

	if err != nil {
		return nil, crypto.Algorithm{}, errors.New("invalid seed; could not determine encoding algorithm")
	}

	if bytes.Equal(decoded[:3], []byte{0x01, 0xe1, 0x4b}) {
		return decoded[3:], crypto.ED25519(), nil
	}

	return decoded[1:], crypto.SECP256K1(), nil

}

// Returns byte slice of a double hashed given byte slice.
// The given byte slice is SHA256 hashed, then the result is RIPEMD160 hashed.
func Sha256RipeMD160(b []byte) []byte {
	sha256 := sha256.New()
	sha256.Write(b)

	ripemd160 := ripemd160.New()
	ripemd160.Write(sha256.Sum(nil))

	return ripemd160.Sum(nil)
}

// Returns the node public key encoding of the byte slice as a base58 string.
func EncodeNodePublicKey(b []byte) (string, error) {

	if len(b) != NodePublicKeyLength {
		return "", &EncodeLengthError{Instance: "NodePublicKey", Expected: NodePublicKeyLength, Input: len(b)}
	}

	npk := Base58CheckEncode(b, NodePublicKeyPrefix)

	return npk, nil
}

// Returns the decoded node public key encoding as a byte slice from a base58 string.
func DecodeNodePublicKey(key string) ([]byte, error) {

	decodedNodeKey, err := Decode(key, []byte{NodePublicKeyPrefix})
	if err != nil {
		return nil, err
	}

	return decodedNodeKey, nil
}

// Returns the account public key encoding of the byte slice as a base58 string.
func EncodeAccountPublicKey(b []byte) (string, error) {

	if len(b) != AccountPublicKeyLength {
		return "", &EncodeLengthError{Instance: "AccountPublicKey", Expected: AccountPublicKeyLength, Input: len(b)}
	}

	apk := Base58CheckEncode(b, AccountPublicKeyPrefix)

	return apk, nil
}

// Returns the decoded account public key encoding as a byte slice from a base58 string.
func DecodeAccountPublicKey(key string) ([]byte, error) {

	decodedAccountKey, err := Decode(key, []byte{AccountPublicKeyPrefix})
	if err != nil {
		return nil, err
	}

	return decodedAccountKey, nil
}
