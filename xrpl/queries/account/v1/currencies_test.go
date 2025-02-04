package v1

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
)

func TestAccountCurrenciesRequest(t *testing.T) {
	s := CurrenciesRequest{
		Account:     "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
		Strict:      true,
		LedgerIndex: common.LedgerIndex(1234),
	}

	j := `{
	"account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
	"ledger_index": 1234,
	"strict": true
}`
	if err := testutil.Serialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestAccountCurrenciesResponse(t *testing.T) {
	s := CurrenciesResponse{
		LedgerHash:  "abc",
		LedgerIndex: 123,
		ReceiveCurrencies: []string{
			"USD",
			"JPY",
		},
		SendCurrencies: []string{
			"USD",
			"CAD",
		},
		Validated: true,
	}
	j := `{
	"ledger_hash": "abc",
	"ledger_index": 123,
	"receive_currencies": [
		"USD",
		"JPY"
	],
	"send_currencies": [
		"USD",
		"CAD"
	],
	"validated": true
}`
	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
