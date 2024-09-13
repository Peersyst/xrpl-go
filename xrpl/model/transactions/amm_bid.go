package transactions

import (
	"errors"
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
)

const MAX_AUTH_ACCOUNTS = 4

// TODO: Implement AMMBid
type AMMBid struct {
	BaseTx
}

func (*AMMBid) TxType() TxType {
	return AMMBidTx
}

// TODO: Implement flatten
func (s *AMMBid) Flatten() FlatTransaction {
	return nil
}

// ValidateAMMBid validates an AMMBid transaction.
func ValidateAMMBid(tx FlatTransaction) error {
	// Validate base transaction fields
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "Asset", IsAsset)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "Asset2", IsAsset)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "Amount", IsAmount)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "BidMin", IsAmount)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "BidMax", IsAmount)
	if err != nil {
		return err
	}

	if tx["AuthAccounts"] != nil {
		if !typecheck.IsArrayOrSlice(tx["AuthAccounts"]) {
			return errors.New("AMMBid: AuthAccounts must be an array of strings")
		}

		authAccounts := tx["AuthAccounts"].([]map[string]interface{})

		if len(authAccounts) > MAX_AUTH_ACCOUNTS {
			return fmt.Errorf("AMMBid: AuthAccounts must have at most %v accounts", MAX_AUTH_ACCOUNTS)
		}

		err := validateAuthAccounts(tx["Account"].(string), authAccounts)
		if err != nil {
			return err
		}
	}

	return nil
}

// validateAuthAccounts validates the AuthAccounts field of an AMMBid transaction.
func validateAuthAccounts(senderAddress string, authAccounts []map[string]interface{}) error {
	for _, account := range authAccounts {

		// check if account is nil or not an object
		if account == nil || !typecheck.IsMap(account) {
			return errors.New("AMMBid: AuthAccounts must be an array of objects")
		}

		// check if account has the required fields, TODO: should check if Account is a valid xrpl account
		if account["Account"] == nil || !typecheck.IsString(account["Account"]) {
			return errors.New("AMMBid: AuthAccounts must have an Account field as valid xrpl account")
		}

		// check that the authAccount field is not the same as tx.Account
		if account["Account"] == senderAddress {
			return errors.New("AMMBid: AuthAccounts must not include the transaction Account")
		}
	}

	return nil
}
