package transaction

import (
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

type Signer struct {
	SignerData SignerData `json:"Signer"`
}

func (s *Signer) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})
	flattened["Signer"] = s.SignerData.Flatten()
	return flattened
}

type SignerData struct {
	Account       types.Address
	TxnSignature  string
	SigningPubKey string
}

type FlatSignerData map[string]interface{}

func (sd *SignerData) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})
	if sd.Account != "" {
		flattened["Account"] = sd.Account.String()
	}
	if sd.TxnSignature != "" {
		flattened["TxnSignature"] = sd.TxnSignature
	}
	if sd.SigningPubKey != "" {
		flattened["SigningPubKey"] = sd.SigningPubKey
	}
	return flattened
}
