package xrpl

import (
	"errors"
	"sort"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
)

var (
	ErrNoTxToMultisign = errors.New("no transaction to multisign")
)

// Multisign is a utility for signing a transaction offline.
// It takes a list of transaction blobs and returns the multisigned transaction blob.
// These transaction blobs must be signed with the wallet.Multisign method.
// They cannot contain SigningPubKey, otherwise the transaction will fail to submit.
// If an error occurs, it will return an error.
func Multisign(blobs ...string) (string, error) {
	if len(blobs) == 0 {
		return "", ErrNoTxToMultisign
	}

	signers := make([]interface{}, 0)
	for _, blob := range blobs {
		tx, err := binarycodec.Decode(blob)
		if err != nil {
			return "", err
		}

		signers = append(signers, tx["Signers"].([]interface{})...)
	}

	tx, err := binarycodec.Decode(blobs[0])
	if err != nil {
		return "", err
	}

	tx["Signers"] = sortSigners(signers)

	blob, err := binarycodec.Encode(tx)
	if err != nil {
		return "", err
	}

	return blob, nil
}

// sortSigners sorts the signers of a transaction.
// It sorts the signers by account.
func sortSigners(signers []interface{}) []interface{} {
	sort.Slice(signers, func(i, j int) bool {
		iSigner := signers[i].(map[string]interface{})["Signer"].(map[string]interface{})
		jSigner := signers[j].(map[string]interface{})["Signer"].(map[string]interface{})

		return iSigner["Account"].(string) > jSigner["Account"].(string)
	})
	return signers
}
