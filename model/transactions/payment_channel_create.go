package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type PaymentChannelCreate struct {
	BaseTx
	Amount         XRPCurrencyAmount
	Destination    Address
	SettleDelay    uint
	PublicKey      []byte
	CancelAfter    uint `json:",omitempty"`
	DestinationTag uint `json:",omitempty"`
}

func (*PaymentChannelCreate) TxType() TxType {
	return PaymentChannelCreateTx
}
