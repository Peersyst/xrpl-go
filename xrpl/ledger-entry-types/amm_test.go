package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestAmm(t *testing.T) {
	amm := AMM{
		LedgerEntryCommonFields: LedgerEntryCommonFields{
			Flags: 0,
		},
		Account: "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
		Asset: Asset{
			Currency: "XRP",
		},
		Asset2: Asset{
			Currency: "TST",
			Issuer:   "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd",
		},
		AuctionSlot: AuctionSlot{
			Account: "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
			AuthAccounts: []AuthAccounts{
				{
					AuthAccount: AuthAccount{
						Account: "rMKXGCbJ5d8LbrqthdG46q3f969MVK2Qeg",
					},
				},
				{
					AuthAccount: AuthAccount{
						Account: "rBepJuTLFJt3WmtLXYAxSjtBWAeQxVbncv",
					},
				},
			},
			DiscountedFee: 60,
			Expiration:    721870180,
			Price: types.IssuedCurrencyAmount{
				Currency: "039C99CD9AB0B70B32ECDA51EAAE471625608EA2",
				Issuer:   "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
				Value:    "0.8696263565463045",
			},
		},
		LPTokenBalance: types.IssuedCurrencyAmount{
			Currency: "039C99CD9AB0B70B32ECDA51EAAE471625608EA2",
			Issuer:   "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
			Value:    "71150.53584131501",
		},
		TradingFee: 600,
		VoteSlots: []VoteSlots{
			{
				VoteEntry: VoteEntry{
					Account:    "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TradingFee: 600,
					VoteWeight: 100000,
				},
			},
		},
	}

	json := `{
	"Flags": 0,
	"Account": "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
	"Asset": {
		"currency": "XRP"
	},
	"Asset2": {
		"currency": "TST",
		"issuer": "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd"
	},
	"AuctionSlot": {
		"Account": "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
		"AuthAccounts": [
			{
				"AuthAccount": {
					"Account": "rMKXGCbJ5d8LbrqthdG46q3f969MVK2Qeg"
				}
			},
			{
				"AuthAccount": {
					"Account": "rBepJuTLFJt3WmtLXYAxSjtBWAeQxVbncv"
				}
			}
		],
		"DiscountedFee": 60,
		"Price": {
			"issuer": "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
			"currency": "039C99CD9AB0B70B32ECDA51EAAE471625608EA2",
			"value": "0.8696263565463045"
		},
		"Expiration": 721870180
	},
	"LPTokenBalance": {
		"issuer": "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
		"currency": "039C99CD9AB0B70B32ECDA51EAAE471625608EA2",
		"value": "71150.53584131501"
	},
	"TradingFee": 600,
	"VoteSlots": [
		{
			"VoteEntry": {
				"Account": "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
				"TradingFee": 600,
				"VoteWeight": 100000
			}
		}
	]
}`

	if err := testutil.SerializeAndDeserialize(t, amm, json); err != nil {
		t.Error(err)
	}
}
