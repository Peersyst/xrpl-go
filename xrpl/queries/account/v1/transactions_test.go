package v1

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
)

func TestAccountTransactionsRequest(t *testing.T) {
	s := TransactionsRequest{
		Account:        "abc",
		LedgerIndexMin: 100,
		LedgerIndexMax: 120,
		LedgerHash:     "def",
		LedgerIndex:    common.LedgerIndex(10),
		Marker:         "123",
	}

	j := `{
	"account": "abc",
	"ledger_index_min": 100,
	"ledger_index_max": 120,
	"ledger_hash": "def",
	"ledger_index": 10,
	"marker": "123"
}`

	if err := testutil.Serialize(t, s, j); err != nil {
		t.Error(err)
	}
}
