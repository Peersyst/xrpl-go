package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/v1/xrpl/testutil"
)

func TestLedgerDataRequest(t *testing.T) {
	s := DataRequest{
		LedgerIndex: common.Closed,
		Binary:      true,
		Limit:       5,
	}
	j := `{
	"ledger_index": "closed",
	"binary": true,
	"limit": 5
}`
	if err := testutil.Serialize(t, s, j); err != nil {
		t.Error(err)
	}
}
