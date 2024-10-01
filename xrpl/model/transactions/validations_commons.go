package transactions

import (
	"fmt"
)

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
		memos, ok := tx["Memos"].([]map[string]interface{})
		if !ok {
			return fmt.Errorf("BaseTransaction: invalid Memos conversion to []map[string]interface{}")
		}
		// loop through each memo and validate it
		for _, memo := range memos {
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