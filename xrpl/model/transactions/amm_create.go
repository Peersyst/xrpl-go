package transactions

type AMMCreate struct {
	BaseTx
}

func (*AMMCreate) TxType() TxType {
	return AMMCreateTx
}

// TODO: Implement flatten
func (s *AMMCreate) Flatten() map[string]interface{} {
	return nil
}