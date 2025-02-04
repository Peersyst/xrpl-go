package v1

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	pathtypes "github.com/Peersyst/xrpl-go/xrpl/queries/path/types"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestBookOffersRequest(t *testing.T) {
	s := BookOffersRequest{
		Taker: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
		TakerGets: pathtypes.BookOfferCurrency{
			Currency: "XRP",
		},
		TakerPays: pathtypes.BookOfferCurrency{
			Currency: "USD",
			Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
		},
		Limit: 10,
	}
	j := `{
	"taker_gets": {
		"currency": "XRP"
	},
	"taker_pays": {
		"currency": "USD",
		"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
	},
	"taker": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
	"limit": 10
}`

	if err := testutil.Serialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestBookOffersResponse(t *testing.T) {
	s := BookOffersResponse{
		LedgerCurrentIndex: 7035305,
		Offers: []pathtypes.BookOffer{
			{
				Account:           "rM3X3QSr8icjTGpaF52dozhbT2BZSXJQYM",
				BookDirectory:     "7E5F614417C2D0A7CEFEB73C4AA773ED5B078DE2B5771F6D55055E4C405218EB",
				BookNode:          "0000000000000000",
				Flags:             0,
				LedgerEntryType:   ledger.OfferEntry,
				OwnerNode:         "0000000000000AE0",
				PreviousTxnID:     "6956221794397C25A53647182E5C78A439766D600724074C99D78982E37599F1",
				PreviousTxnLgrSeq: 7022646,
				Sequence:          264542,
				TakerGets: types.IssuedCurrencyAmount{
					Currency: "EUR",
					Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
					Value:    "17.90363633316433",
				},
				TakerPays: types.IssuedCurrencyAmount{
					Currency: "USD",
					Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
					Value:    "27.05340557506234",
				},
				Quality: "1.511056473200875",
			},
			{
				Account:           "rhsxKNyN99q6vyYCTHNTC1TqWCeHr7PNgp",
				BookDirectory:     "7E5F614417C2D0A7CEFEB73C4AA773ED5B078DE2B5771F6D5505DCAA8FE12000",
				BookNode:          "0000000000000000",
				Flags:             131072,
				LedgerEntryType:   ledger.OfferEntry,
				OwnerNode:         "0000000000000001",
				PreviousTxnID:     "8AD748CD489F7FF34FCD4FB73F77F1901E27A6EFA52CCBB0CCDAAB934E5E754D",
				PreviousTxnLgrSeq: 7007546,
				Sequence:          265,
				TakerGets: types.IssuedCurrencyAmount{
					Currency: "EUR",
					Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
					Value:    "2.542743233917848",
				},
				TakerPays: types.IssuedCurrencyAmount{
					Currency: "USD",
					Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
					Value:    "4.19552633596446",
				},
				Quality: "1.65",
			},
		},
	}

	j := `{
	"ledger_current_index": 7035305,
	"offers": [
		{
			"Flags": 0,
			"LedgerEntryType": "Offer",
			"Account": "rM3X3QSr8icjTGpaF52dozhbT2BZSXJQYM",
			"BookDirectory": "7E5F614417C2D0A7CEFEB73C4AA773ED5B078DE2B5771F6D55055E4C405218EB",
			"BookNode": "0000000000000000",
			"OwnerNode": "0000000000000AE0",
			"PreviousTxnID": "6956221794397C25A53647182E5C78A439766D600724074C99D78982E37599F1",
			"PreviousTxnLgrSeq": 7022646,
			"Sequence": 264542,
			"TakerPays": {
				"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				"currency": "USD",
				"value": "27.05340557506234"
			},
			"TakerGets": {
				"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				"currency": "EUR",
				"value": "17.90363633316433"
			},
			"quality": "1.511056473200875"
		},
		{
			"Flags": 131072,
			"LedgerEntryType": "Offer",
			"Account": "rhsxKNyN99q6vyYCTHNTC1TqWCeHr7PNgp",
			"BookDirectory": "7E5F614417C2D0A7CEFEB73C4AA773ED5B078DE2B5771F6D5505DCAA8FE12000",
			"BookNode": "0000000000000000",
			"OwnerNode": "0000000000000001",
			"PreviousTxnID": "8AD748CD489F7FF34FCD4FB73F77F1901E27A6EFA52CCBB0CCDAAB934E5E754D",
			"PreviousTxnLgrSeq": 7007546,
			"Sequence": 265,
			"TakerPays": {
				"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				"currency": "USD",
				"value": "4.19552633596446"
			},
			"TakerGets": {
				"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				"currency": "EUR",
				"value": "2.542743233917848"
			},
			"quality": "1.65"
		}
	]
}`
	if err := testutil.Serialize(t, s, j); err != nil {
		t.Error(err)
	}
}
