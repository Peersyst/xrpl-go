package transactions

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
)

func ValidateBaseTransaction(tx FlatTransaction) error {
	var err error

	err = ValidateRequiredField(tx, "TransactionType", typecheck.IsString)
	if err != nil {
		return err
	}

	ValidateRequiredField(tx, "Account", typecheck.IsString)

	// optional fields
	err = ValidateOptionalField(tx, "Fee", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "Sequence", typecheck.IsInt)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "AccountTxnID", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "LastLedgerSequence", typecheck.IsInt)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "SourceTag", typecheck.IsInt)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "SigningPubKey", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "TicketSequence", typecheck.IsInt)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "TxnSignature", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "NetworkID", typecheck.IsInt)
	if err != nil {
		return err
	}

	// memos
	err = validateMemos(tx)
	if err != nil {
		return err
	}

	// signers
	err = validateSigners(tx)
	if err != nil {
		return err
	}

	return nil
}

func ValidateRequiredField(tx FlatTransaction, field string, checkValidity func(interface{}) bool) error {
	// Check if the field is present in the transaction map.
	if _, ok := tx[field]; !ok {
		return fmt.Errorf("%s is missing", field)
	}

	// Check if the field is valid.
	if !checkValidity(tx[field]) {
		transactionType, _ := tx["TransactionType"].(string)
		return fmt.Errorf("%s: invalid field %s", transactionType, field)
	}

	return nil
}

// ValidateOptionalField validates an optional field in the transaction map.
func ValidateOptionalField(tx FlatTransaction, paramName string, checkValidity func(interface{}) bool) error {
	// Check if the field is present in the transaction map.
	if value, ok := tx[paramName]; ok {
		// Check if the field is valid.
		if !checkValidity(value) {
			transactionType, _ := tx["TransactionType"].(string)
			return fmt.Errorf("%s: invalid field %s", transactionType, paramName)
		}
	}

	return nil
}

// validateMemos validates the Memos field in the transaction map.
func validateMemos(tx FlatTransaction) error {
	// Check if the field Memos is set
	if tx["Memos"] != nil {
		flatMemoWrappers, ok := tx["Memos"].([]map[string]interface{})
		if !ok {
			return fmt.Errorf("BaseTransaction: invalid Memos conversion to []FlatMemoWrapper")
		}
		// loop through each memo and validate it
		for _, memo := range flatMemoWrappers {
			if !IsMemo(memo) {
				return fmt.Errorf("BaseTransaction: invalid Memos. A memo can only have hexadecimals values. See https://xrpl.org/docs/references/protocol/transactions/common-fields#memos-field")
			}
		}
	}

	return nil
}

// validateSigners validates the Signers field in the transaction map.
func validateSigners(tx FlatTransaction) error {
	if tx["Signers"] != nil {
		signers, ok := tx["Signers"].([]FlatSigner)
		if !ok {
			return fmt.Errorf("BaseTransaction: invalid Signers")
		}
		for _, signer := range signers {
			if !IsSigner(signer) {
				return fmt.Errorf("BaseTransaction: invalid Signers")
			}
		}
	}

	return nil
}
