package transactions

import (
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
