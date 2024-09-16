package transactions

import (
	"errors"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/ledger"
)

const MAX_SIGNERS = 32

type SignerListSet struct {
	BaseTx
	SignerQuorum  uint
	SignerEntries []ledger.SignerEntryWrapper
}

func (*SignerListSet) TxType() TxType {
	return SignerListSetTx
}

// TODO: Implement flatten
func (s *SignerListSet) Flatten() FlatTransaction {
	return nil
}

func ValidateSignerListSet(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "SignerQuorum", typecheck.IsUint)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "SignerEntries", typecheck.IsArrayOrSlice)
	if err != nil {
		return err
	}

	signerEntries, _ := tx["SignerEntries"].([]interface{})
	if len(signerEntries) == 0 {
		return errors.New("field 'SignerEntries' must not be empty")
	}

	if len(signerEntries) > MAX_SIGNERS {
		return errors.New("field 'SignerEntries' must not exceed 32 members")
	}

	// check the WalletLocator of a SignerEntry
	for _, entry := range signerEntries {
		signerEntry := entry.(map[string]interface{})
		walletLocator, ok := signerEntry["WalletLocator"].(string)
		if ok && !typecheck.IsHex(walletLocator) {
			return errors.New("WalletLocator in SignerEntry must be a 256-bit (32-byte) hexadecimal value")
		}
	}

	return nil
}
