package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/testutil"
)

func TestFeeSettings(t *testing.T) {
	var s Object = &FeeSettings{
		BaseFee:           "000000000000000A",
		Flags:             0,
		LedgerEntryType:   FeeSettingsEntry,
		ReferenceFeeUnits: 10,
		ReserveBase:       20000000,
		ReserveIncrement:  5000000,
	}

	j := `{
	"Flags": 0,
	"LedgerEntryType": "FeeSettings",
	"BaseFee": "000000000000000A",
	"ReferenceFeeUnits": 10,
	"ReserveBase": 20000000,
	"ReserveIncrement": 5000000
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
