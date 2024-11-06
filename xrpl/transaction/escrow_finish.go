package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type EscrowFinish struct {
	BaseTx
	Owner         types.Address
	OfferSequence uint32
	Condition     string `json:",omitempty"`
	Fulfillment   string `json:",omitempty"`
}

func (*EscrowFinish) TxType() TxType {
	return EscrowFinishTx
}

// TODO: Implement flatten
func (s *EscrowFinish) Flatten() FlatTransaction {
	return nil
}
