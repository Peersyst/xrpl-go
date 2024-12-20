package account

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestAccountOffersRequest(t *testing.T) {
	s := OffersRequest{
		Account:     "abc",
		LedgerIndex: common.LedgerIndex(10),
		Marker:      "123",
	}
	j := `{
	"account": "abc",
	"ledger_index": 10,
	"marker": "123"
}`
	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

}

func TestAccountOffersResponse(t *testing.T) {
	s := OffersResponse{
		Account: "abc",
		Offers: []OfferResult{
			{Flags: 0,
				Sequence: 1,
				TakerGets: types.IssuedCurrencyAmount{
					Issuer:   "def",
					Currency: "USD",
					Value:    "100",
				},
				TakerPays:  types.XRPCurrencyAmount(1),
				Quality:    "1",
				Expiration: 50000000,
			},
		},
		LedgerCurrentIndex: 54321,
		LedgerIndex:        54320,
		LedgerHash:         "def",
	}
	j := `{
	"account": "abc",
	"offers": [
		{
			"flags": 0,
			"seq": 1,
			"taker_gets": {
				"issuer": "def",
				"currency": "USD",
				"value": "100"
			},
			"taker_pays": "1",
			"quality": "1",
			"expiration": 50000000
		}
	],
	"ledger_current_index": 54321,
	"ledger_index": 54320,
	"ledger_hash": "def"
}`
	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
