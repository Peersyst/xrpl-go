package v1

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
)

func TestAccountObjectsRequest(t *testing.T) {
	s := ObjectsRequest{
		Account:     "rsuHaTvJh1bDmDoxX9QcKP7HEBSBt4XsHx",
		Type:        SignerListObject,
		LedgerIndex: common.LedgerIndex(123),
	}

	j := `{
	"account": "rsuHaTvJh1bDmDoxX9QcKP7HEBSBt4XsHx",
	"type": "signer_list",
	"ledger_index": 123
}`
	if err := testutil.Serialize(t, s, j); err != nil {
		t.Error(err)
	}
}
