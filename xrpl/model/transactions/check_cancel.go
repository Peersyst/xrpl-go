package transactions

import (
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

func (s *CheckCancel) Flatten() map[string]interface{} {
	return nil
}
