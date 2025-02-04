package v1

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
)

func TestNFTHistoryRequest(t *testing.T) {
	s := NFTHistoryRequest{
		NFTokenID:      "0000000000000000000000000000000000000000000000000000000000000000",
		LedgerIndexMin: 100,
		LedgerIndexMax: 200,
		Binary:         true,
		Forward:        true,
		Limit:          100,
		Marker:         "marker",
	}

	j := `{
	"nft_id": "0000000000000000000000000000000000000000000000000000000000000000",
	"ledger_index_min": 100,
	"ledger_index_max": 200,
	"binary": true,
	"forward": true,
	"limit": 100,
	"marker": "marker"
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Fatal(err)
	}
}
