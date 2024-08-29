package transactions

type AMMWithdraw struct {
	BaseTx
}

func (*AMMWithdraw) TxType() TxType {
	return AMMWithdrawTx
}

// TODO: Implement flatten
func (s *AMMWithdraw) Flatten() map[string]interface{} {
	return nil
}