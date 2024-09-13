package transactions

import (
	"errors"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type CheckCancel struct {
	BaseTx
	CheckID types.Hash256
}

// TODO: Implement flatten
func (*CheckCancel) TxType() TxType {
	return CheckCancelTx
}

func (s *CheckCancel) Flatten() FlatTransaction {
	return nil
}

func ValidateCheckCancel(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	_, hasCheckID := tx["CheckID"]
	if hasCheckID && !typecheck.IsString(tx["CheckID"]) {
		return errors.New("CheckCancel: invalid CheckID, must be a string")
	}

	return nil
}
