package transactions

import (
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type SetRegularKey struct {
	BaseTx
	RegularKey types.Address `json:",omitempty"`
}

func (*SetRegularKey) TxType() TxType {
	return SetRegularKeyTx
}

// TODO: Implement flatten
func (s *SetRegularKey) Flatten() FlatTransaction {
	return nil
}

func ValidateSetRegularKey(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	ValidateOptionalField(tx, "RegularKey", typecheck.IsString)

	return nil
}
