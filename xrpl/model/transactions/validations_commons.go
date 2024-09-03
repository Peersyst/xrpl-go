package transactions

import (
	"fmt"

	typeoffns "github.com/Peersyst/xrpl-go/xrpl/utils/typeof-fns"
)

func ValidateBaseTransaction(tx FlatTransaction) error {
	ValidateRequiredField(tx, "TransactionType", typeoffns.IsString)
	ValidateRequiredField(tx, "Account", typeoffns.IsString)

	// optional fields
	ValidateOptionalField(tx, "Fee", typeoffns.IsString)
	ValidateOptionalField(tx, "Sequence", typeoffns.IsInt)
	ValidateOptionalField(tx, "AccountTxnID", typeoffns.IsString)
	ValidateOptionalField(tx, "LastLedgerSequence", typeoffns.IsInt)
	ValidateOptionalField(tx, "SourceTag", typeoffns.IsInt)
	ValidateOptionalField(tx, "SigningPubKey", typeoffns.IsString)
	ValidateOptionalField(tx, "TicketSequence", typeoffns.IsInt)
	ValidateOptionalField(tx, "TxnSignature", typeoffns.IsString)
	ValidateOptionalField(tx, "NetworkID", typeoffns.IsInt)

	// memos
	validateMemos(tx)

	// signers
	validateSigners(tx)

	return nil
}

func ValidateRequiredField(tx FlatTransaction, field string, checkValidity func(interface{}) bool) {
	// Check if the field is present in the transaction map.
	if _, ok := tx[field]; !ok {
		panic(field + " is missing")
	}

	// Check if the field is valid.
	if !checkValidity(tx[field]) {
		transactionType, _ := tx["TransactionType"].(string)
		panic(fmt.Errorf("%s: invalid field %s", transactionType, field))
	}
}

// ValidateOptionalField validates an optional field in the transaction map.
func ValidateOptionalField(tx FlatTransaction, paramName string, checkValidity func(interface{}) bool) {
	// Check if the field is present in the transaction map.
	if value, ok := tx[paramName]; ok {
		// Check if the field is valid.
		if !checkValidity(value) {
			transactionType, _ := tx["TransactionType"].(string)
			panic(fmt.Errorf("%s: invalid field %s", transactionType, paramName))
		}
	}
}

func validateMemos(tx FlatTransaction) {
	if tx["Memos"] != nil {
		memos, ok := tx["Memos"].([]FlatTransaction)
		if !ok {
			panic("BaseTransaction: invalid Memos")
		}
		for _, memo := range memos {
			if !IsMemo(memo) {
				panic("BaseTransaction: invalid Memos. A memo can only have hexadecimals values. See https://xrpl.org/docs/references/protocol/transactions/common-fields#memos-field")
			}
		}
	}
}

func validateSigners(tx FlatTransaction) {
	if tx["Signers"] != nil {
		signers, ok := tx["Signers"].([]FlatTransaction)
		if !ok {
			panic("BaseTransaction: invalid Signers")
		}
		for _, signer := range signers {
			if !IsSigner(signer) {
				panic("BaseTransaction: invalid Signers")
			}
		}
	}
}
