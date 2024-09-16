package transactions

import (
	"errors"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type NFTokenMint struct {
	BaseTx
	NFTokenTaxon uint
	Issuer       types.Address    `json:",omitempty"`
	TransferFee  uint16           `json:",omitempty"`
	URI          types.NFTokenURI `json:",omitempty"`
}

func (*NFTokenMint) TxType() TxType {
	return NFTokenMintTx
}

// TODO: Implement flatten
func (s *NFTokenMint) Flatten() FlatTransaction {
	return nil
}

func ValidateNFTokenMint(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	// check Account and Issuer are not the same
	if tx["Account"] == tx["Issuer"] {
		return errors.New("field 'Issuer' and 'Account' must be different")
	}

	err = ValidateRequiredField(tx, "NFTokenTaxon", typecheck.IsUint)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "Issuer", typecheck.IsString)
	if err != nil {
		return err
	}

	// Validate URI is a string
	err = ValidateOptionalField(tx, "URI", typecheck.IsString)
	if err != nil {
		return err
	}

	// Validate URI is a hex string
	if !typecheck.IsHex(tx["URI"].(string)) {
		return errors.New("URI must be a hex string")
	}

	return nil
}
