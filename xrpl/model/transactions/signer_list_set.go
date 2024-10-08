package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/ledger"
)

type SignerListSet struct {
	BaseTx
	SignerQuorum  uint
	SignerEntries []ledger.SignerEntryWrapper
}

func (*SignerListSet) TxType() TxType {
	return SignerListSetTx
}

// TODO: Implement flatten
func (s *SignerListSet) Flatten() FlatTransaction {
	return nil
}
