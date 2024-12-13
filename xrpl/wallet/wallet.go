package wallet

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
	"github.com/Peersyst/xrpl-go/keypairs"
	"github.com/Peersyst/xrpl-go/pkg/random"
	"github.com/Peersyst/xrpl-go/xrpl/hash"
	"github.com/Peersyst/xrpl-go/xrpl/interfaces"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// A utility for deriving a wallet composed of a keypair (publicKey/privateKey).
// A wallet can be derived from either a seed, mnemonic, or entropy (array of random numbers).
// It provides functionality to sign/verify transactions offline.
type Wallet struct {
	PublicKey      string
	PrivateKey     string
	ClassicAddress types.Address
	Seed           string
}

// Creates a new random Wallet. In order to make this a valid account on ledger, you must
// Send XRP to it.
func New(alg interfaces.CryptoImplementation) (Wallet, error) {
	seed, err := keypairs.GenerateSeed("", alg, random.NewRandomizer())
	if err != nil {
		return Wallet{}, err
	}
	return FromSeed(seed, "")
}

// Derives a wallet from a seed.
// Returns a Wallet object. If an error occurs, it will be returned.
func FromSeed(seed string, masterAddress string) (Wallet, error) {
	privKey, pubKey, err := keypairs.DeriveKeypair(seed, false)
	if err != nil {
		return Wallet{}, err
	}

	var classicAddr types.Address
	if ok := addresscodec.IsValidClassicAddress(masterAddress); ok {
		classicAddr = types.Address(masterAddress)
	} else {
		addr, err := keypairs.DeriveClassicAddress(pubKey)
		if err != nil {
			return Wallet{}, err
		}
		classicAddr = types.Address(addr)
	}

	return Wallet{
		PublicKey:      pubKey,
		PrivateKey:     privKey,
		Seed:           seed,
		ClassicAddress: classicAddr,
	}, nil

}

// Derives a wallet from a secret (AKA a seed).
// Returns a Wallet object. If an error occurs, it will be returned.
func FromSecret(seed string) (Wallet, error) {
	return FromSeed(seed, "")
}

// // Derives a wallet from a bip39 or RFC1751 mnemonic (Defaults to bip39).
// // Returns a Wallet object. If an error occurs, it will be returned.
func FromMnemonic(mnemonic string) (*Wallet, error) {
	// Validate the mnemonic
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("invalid mnemonic")
	}

	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "")

	// Derive the master key
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	// Derive the key using the path m/44'/144'/0'/0/0
	path := []uint32{
		44 + bip32.FirstHardenedChild,
		144 + bip32.FirstHardenedChild,
		bip32.FirstHardenedChild,
		0,
		0,
	}

	key := masterKey
	for _, childNum := range path {
		key, err = key.NewChildKey(childNum)
		if err != nil {
			return nil, err
		}
	}

	// Convert the private key to the format expected by the XRPL library
	privKey := strings.ToUpper(hex.EncodeToString(key.Key))
	pubKey := strings.ToUpper(hex.EncodeToString(key.PublicKey().Key))

	// Derive classic address
	classicAddr, err := keypairs.DeriveClassicAddress(pubKey)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		PublicKey:      pubKey,
		PrivateKey:     fmt.Sprintf("00%s", privKey),
		ClassicAddress: types.Address(classicAddr),
		Seed:           "", // We don't have the seed in this case
	}, nil
}

// Signs a transaction offline.
// In order for a transaction to be validated, it must be signed by the account sending the transaction to prove
// that the owner is actually the one deciding to take that action.
//
// TODO: Refactor to accept a `Transaction` object instead of a map.
func (w *Wallet) Sign(tx map[string]interface{}) (string, string, error) {
	tx["SigningPubKey"] = w.PublicKey

	encodedTx, err := binarycodec.EncodeForSigning(tx)
	if err != nil {
		return "", "", err
	}

	txHash, err := w.computeSignature(encodedTx)
	if err != nil {
		return "", "", err
	}

	tx["TxnSignature"] = txHash

	txBlob, err := binarycodec.Encode(tx)
	if err != nil {
		return "", "", err
	}

	txHash, err = hash.TxBlob(txBlob)
	if err != nil {
		return "", "", err
	}

	return txBlob, txHash, nil
}

// Returns the classic address of the wallet.
func (w *Wallet) GetAddress() types.Address {
	return types.Address(w.ClassicAddress)
}

// Signs a multisigned transaction offline.
// Returns the transaction blob and the transaction hash.
func (w *Wallet) Multisign(tx map[string]interface{}) (string, string, error) {
	encodedTx, err := binarycodec.EncodeForMultisigning(tx, w.ClassicAddress.String())
	if err != nil {
		return "", "", err
	}

	txHash, err := w.computeSignature(encodedTx)
	if err != nil {
		return "", "", err
	}

	signer := transaction.Signer{
		SignerData: transaction.SignerData{
			Account:       w.ClassicAddress,
			TxnSignature:  txHash,
			SigningPubKey: w.PublicKey,
		},
	}

	tx["Signers"] = []any{signer.Flatten()}
	blob, err := binarycodec.Encode(tx)
	if err != nil {
		return "", "", err
	}
	blobHash, err := hash.TxBlob(blob)
	if err != nil {
		return "", "", err
	}

	return blob, blobHash, nil
}

// Computes the signature of a transaction.
// Returns the signature of the transaction. If an error occurs, it will return an error.
func (w *Wallet) computeSignature(encodedTx string) (string, error) {
	hexTx, err := hex.DecodeString(encodedTx)
	if err != nil {
		return "", err
	}

	txHash, err := keypairs.Sign(string(hexTx), w.PrivateKey)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

// Verifies a signed transaction offline.
// Returns a boolean indicating if the transaction is valid and an error if it is not.
// If the transaction is signed with a public key, the public key must match the one in the transaction.
// func (w *Wallet) VerifyTransaction(tx map[string]any) (bool, error) {
// 	return false, errors.New("not implemented")
// }

// // Gets an X-address in Testnet/Mainnet format.
// func (w *Wallet) GetXAddress() (string, error) {
// 	return "", errors.New("not implemented")
// }
