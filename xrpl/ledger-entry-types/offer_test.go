package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
	"github.com/stretchr/testify/require"
)

func TestOffer(t *testing.T) {
	var s Object = &Offer{
		Account:           "rBqb89MRQJnMPq8wTwEbtz4kvxrEDfcYvt",
		BookDirectory:     "ACC27DE91DBA86FC509069EAF4BC511D73128B780F2E54BF5E07A369E2446000",
		BookNode:          "0000000000000000",
		Flags:             131072,
		LedgerEntryType:   OfferEntry,
		OwnerNode:         "0000000000000000",
		PreviousTxnID:     "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF",
		PreviousTxnLgrSeq: 14524914,
		Sequence:          866,
		TakerGets: types.IssuedCurrencyAmount{
			Issuer:   "r9Dr5xwkeLegBeXq6ujinjSBLQzQ1zQGjH",
			Currency: "XAG",
			Value:    "37",
		},
		TakerPays: types.XRPCurrencyAmount(79550000000),
	}

	j := `{
	"Flags": 131072,
	"LedgerEntryType": "Offer",
	"Account": "rBqb89MRQJnMPq8wTwEbtz4kvxrEDfcYvt",
	"BookDirectory": "ACC27DE91DBA86FC509069EAF4BC511D73128B780F2E54BF5E07A369E2446000",
	"BookNode": "0000000000000000",
	"OwnerNode": "0000000000000000",
	"PreviousTxnID": "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF",
	"PreviousTxnLgrSeq": 14524914,
	"Sequence": 866,
	"TakerPays": "79550000000",
	"TakerGets": {
		"issuer": "r9Dr5xwkeLegBeXq6ujinjSBLQzQ1zQGjH",
		"currency": "XAG",
		"value": "37"
	}
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestOffer_SetLsfPassive(t *testing.T) {
	o := &Offer{}
	o.SetLsfPassive()
	require.Equal(t, o.Flags, uint32(0x00010000))
}

func TestOffer_SetLsfSell(t *testing.T) {
	o := &Offer{}
	o.SetLsfSell()
	require.Equal(t, o.Flags, uint32(0x00020000))
}

func TestOffer_EntryType(t *testing.T) {
	o := &Offer{}
	require.Equal(t, o.EntryType(), OfferEntry)
}
