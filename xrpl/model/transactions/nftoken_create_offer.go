package transactions

import (
	"encoding/json"
	"errors"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/model/utils"
)

type NFTokenCreateOffer struct {
	BaseTx
	Owner       types.Address `json:",omitempty"`
	NFTokenID   types.NFTokenID
	Amount      types.CurrencyAmount
	Expiration  uint          `json:",omitempty"`
	Destination types.Address `json:",omitempty"`
}

const (
	// Transaction Flags for an NFTokenCreateOffer Transaction.
	tfSellNFToken = 0x00000001
)

func (*NFTokenCreateOffer) TxType() TxType {
	return NFTokenCreateOfferTx
}

// TODO: Implement flatten
func (s *NFTokenCreateOffer) Flatten() FlatTransaction {
	return nil
}

func (n *NFTokenCreateOffer) UnmarshalJSON(data []byte) error {
	type ncoHelper struct {
		BaseTx
		Owner       types.Address `json:",omitempty"`
		NFTokenID   types.NFTokenID
		Amount      json.RawMessage
		Expiration  uint          `json:",omitempty"`
		Destination types.Address `json:",omitempty"`
	}
	var h ncoHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*n = NFTokenCreateOffer{
		BaseTx:      h.BaseTx,
		Owner:       h.Owner,
		NFTokenID:   h.NFTokenID,
		Expiration:  h.Expiration,
		Destination: h.Destination,
	}

	amount, err := types.UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return err
	}
	n.Amount = amount
	return nil
}

func ValidateNFTokenCreateOffer(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	// check if Owner is not the same as Account
	if tx["Owner"] == tx["Account"] {
		return errors.New("field 'Owner' must be different from 'Account'")
	}

	err = ValidateOptionalField(tx, "Destination", typecheck.IsString)
	if err != nil {
		return err
	}

	// TODO: replace by IsAccount when it's implemented
	err = ValidateOptionalField(tx, "Owner", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "NFTokenID", IsAmount)
	if err != nil {
		return err
	}

	isFlagsUint := typecheck.IsUint(tx["Flags"])

	if isFlagsUint && utils.IsFlagEnabled(tx["Flags"].(uint), tfSellNFToken) {
		// validate sell offer cases
		err = validateNFTokenSellOfferCases(tx)
		if err != nil {
			return err
		}
	} else {
		// validate buy offer cases
		err = validateNFTokenBuyOfferCases(tx)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateNFTokenSellOfferCases(tx FlatTransaction) error {
	// check if Owner is set with a valid xrpl address. TODO: replace by IsAccount when it's implemented
	err := ValidateRequiredField(tx, "Owner", typecheck.IsString)
	if err != nil {
		return err
	}

	return nil
}

func validateNFTokenBuyOfferCases(tx FlatTransaction) error {
	// check that Owner is set
	if tx["Owner"] == nil {
		return errors.New("field 'Owner' must be set for buy offers")
	}

	amount, err := ParseAmountValue((tx["Amount"]))
	if err != nil {
		return err
	}

	if amount <= 0 {
		return errors.New("field 'Amount' must be set for buy offers")
	}

	return nil
}
