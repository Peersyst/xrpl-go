package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestNFTokenOffer(t *testing.T) {
	var s LedgerObject = &NFTokenOffer{
		Amount:            types.XRPCurrencyAmount(1000000),
		Flags:             1,
		LedgerEntryType:   NFTokenOfferEntry,
		NFTokenID:         "00081B5825A08C22787716FA031B432EBBC1B101BB54875F0002D2A400000000",
		NFTokenOfferNode:  "0",
		Owner:             "rhRxL3MNvuKEjWjL7TBbZSDacb8PmzAd7m",
		OwnerNode:         "17",
		PreviousTxnID:     "BFA9BE27383FA315651E26FDE1FA30815C5A5D0544EE10EC33D3E92532993769",
		PreviousTxnLgrSeq: 75443565,
	}

	j := `{
	"Amount": "1000000",
	"Flags": 1,
	"LedgerEntryType": "NFTokenOffer",
	"NFTokenID": "00081B5825A08C22787716FA031B432EBBC1B101BB54875F0002D2A400000000",
	"NFTokenOfferNode": "0",
	"Owner": "rhRxL3MNvuKEjWjL7TBbZSDacb8PmzAd7m",
	"OwnerNode": "17",
	"PreviousTxnID": "BFA9BE27383FA315651E26FDE1FA30815C5A5D0544EE10EC33D3E92532993769",
	"PreviousTxnLgrSeq": 75443565
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
