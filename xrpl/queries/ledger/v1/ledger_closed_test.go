package v1

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
)

// Ledger closed request does not have any fields to test

func TestLedgerClosedResponse(t *testing.T) {
	s := ClosedResponse{
		LedgerHash:  "abc",
		LedgerIndex: 123,
	}
	j := `{
	"ledger_hash": "abc",
	"ledger_index": 123
}`
	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
