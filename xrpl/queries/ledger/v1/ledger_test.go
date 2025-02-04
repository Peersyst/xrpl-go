package v1

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	ledgertypes "github.com/Peersyst/xrpl-go/xrpl/queries/ledger/types"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestLedgerRequest(t *testing.T) {
	s := Request{
		LedgerHash:  "abc",
		LedgerIndex: common.LedgerIndex(123),
	}
	j := `{
	"ledger_hash": "abc",
	"ledger_index": 123
}`
	if err := testutil.Serialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestLedgerResponse(t *testing.T) {
	s := Response{
		Ledger: ledgertypes.BaseLedger{
			AccountHash:         "53BD4650A024E27DEB52DBB6A52EDB26528B987EC61C895C48D1EB44CEDD9AD3",
			CloseTime:           638329241,
			CloseTimeHuman:      "2020-Mar-24 01:40:41.000000000 UTC",
			CloseTimeResolution: 10,
			Closed:              true,
			LedgerHash:          "1723099E269C77C4BDE86C83FA6415D71CF20AA5CB4A94E5C388ED97123FB55B",
			LedgerIndex:         "54300932",
			ParentCloseTime:     638329240,
			ParentHash:          "DF68B3BCABD31097634BABF0BDC87932D43D26E458BFEEFD36ADF2B3D94998C0",
			TotalCoins:          types.XRPCurrencyAmount(99991024049648900),
			TransactionHash:     "50B3A8FE2C5620E43AA57564209AEDFEA3E868CFA2F6E4AB4B9E55A7A62AAF7B",
		},
		LedgerHash:  "1723099E269C77C4BDE86C83FA6415D71CF20AA5CB4A94E5C388ED97123FB55B",
		LedgerIndex: 54300932,
		Validated:   true,
	}
	j := `{
	"ledger": {
		"account_hash": "53BD4650A024E27DEB52DBB6A52EDB26528B987EC61C895C48D1EB44CEDD9AD3",
		"close_flags": 0,
		"close_time": 638329241,
		"close_time_human": "2020-Mar-24 01:40:41.000000000 UTC",
		"close_time_resolution": 10,
		"closed": true,
		"ledger_hash": "1723099E269C77C4BDE86C83FA6415D71CF20AA5CB4A94E5C388ED97123FB55B",
		"ledger_index": "54300932",
		"parent_close_time": 638329240,
		"parent_hash": "DF68B3BCABD31097634BABF0BDC87932D43D26E458BFEEFD36ADF2B3D94998C0",
		"total_coins": "99991024049648900",
		"transaction_hash": "50B3A8FE2C5620E43AA57564209AEDFEA3E868CFA2F6E4AB4B9E55A7A62AAF7B"
	},
	"ledger_hash": "1723099E269C77C4BDE86C83FA6415D71CF20AA5CB4A94E5C388ED97123FB55B",
	"ledger_index": 54300932,
	"validated": true
}`
	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
