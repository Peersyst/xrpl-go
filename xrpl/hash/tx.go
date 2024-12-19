package hash

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"strings"

	binarycodec "github.com/Peersyst/xrpl-go/v1/binary-codec"
	"github.com/Peersyst/xrpl-go/v1/pkg/crypto"
)

// SignTxBlob hashes a signed transaction blob
// It takes a transaction blob and returns the hash of the signed transaction.
// It returns an error if the transaction blob is invalid.
func SignTxBlob(txBlob string) (string, error) {
	tx, err := binarycodec.Decode(txBlob)
	if err != nil {
		return "", err
	}

	// Check if the transaction has at least one of the required fields
	hasTxnSignature := tx["TxnSignature"] != nil
	hasSigners := tx["Signers"] != nil
	hasSigningPubKey := tx["SigningPubKey"] != nil

	if !hasTxnSignature && !hasSigners && !hasSigningPubKey {
		return "", errors.New("transaction must have at least one of TxnSignature, Signers, or SigningPubKey")
	}

	// Create a byte slice with the correct capacity
	payload := make([]byte, 4+len(txBlob)/2)

	// Convert TRANSACTION_PREFIX to big-endian bytes
	binary.BigEndian.PutUint32(payload[:4], TransactionPrefix)

	// Decode the txBlob into the rest of the payload
	_, err = hex.Decode(payload[4:], []byte(txBlob))
	if err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString(crypto.Sha512Half(payload))), nil
}
