package signing

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
)

type SignResponse struct {
	TxBlob string                       `json:"tx_blob"`
	TxJson transactions.FlatTransaction `json:"tx_json"`
}
