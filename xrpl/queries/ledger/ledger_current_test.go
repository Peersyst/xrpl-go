package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/testutil"
)

// Ledger Current request has no fields to test

func TestLedgerCurrentResponse(t *testing.T) {
	s := CurrentResponse{
		LedgerCurrentIndex: 123,
	}
	j := `{
	"ledger_current_index": 123
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
