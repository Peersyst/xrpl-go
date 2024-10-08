package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/test"
)

func TestNFTokenAcceptOfferTransaction(t *testing.T) {
	s := NFTokenAcceptOffer{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: NFTokenAcceptOfferTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		NFTokenSellOffer: "ABC",
		NFTokenBuyOffer:  "DEF",
		NFTokenBrokerFee: types.IssuedCurrencyAmount{
			Issuer:   "ty",
			Currency: "USD",
			Value:    "1",
		},
	}

	j := `{
	"Account": "abcdef",
	"TransactionType": "NFTokenAcceptOffer",
	"Fee": "1",
	"Sequence": 1234,
	"SigningPubKey": "ghijk",
	"TxnSignature": "A1B2C3D4E5F6",
	"NFTokenSellOffer": "ABC",
	"NFTokenBuyOffer": "DEF",
	"NFTokenBrokerFee": {
		"issuer": "ty",
		"currency": "USD",
		"value": "1"
	}
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

	tx, err := UnmarshalTx(json.RawMessage(j))
	if err != nil {
		t.Errorf("UnmarshalTx error: %s", err.Error())
	}
	if !reflect.DeepEqual(tx, &s) {
		t.Error("UnmarshalTx result differs from expected")
	}

}
