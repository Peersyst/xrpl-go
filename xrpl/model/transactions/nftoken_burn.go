package transactions

import (
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type NFTokenBurn struct {
	BaseTx
	NFTokenID types.NFTokenID
	Owner     types.Address `json:",omitempty"`
}

func (*NFTokenBurn) TxType() TxType {
	return NFTokenBurnTx
}

// TODO: Implement flatten
func (s *NFTokenBurn) Flatten() FlatTransaction {
	return nil
}

func ValidateNFTokenBurn(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "NFTokenID", typecheck.IsString)
	if err != nil {
		return err
	}

	// TODO: replace by IsAccount to check it's a correct xrpl account when it's implemented
	err = ValidateOptionalField(tx, "Owner", typecheck.IsString)
	if err != nil {
		return err
	}

	return nil
}
