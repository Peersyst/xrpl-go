package v1

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
)

func TestNFTInfoRequest(t *testing.T) {
	s := NFTInfoRequest{
		NFTokenID:   "00080000B4F4AFC5FBCBD76873F18006173D2193467D3EE70000099B00000000",
		LedgerIndex: common.Validated,
	}

	j := `{
	"nft_id": "00080000B4F4AFC5FBCBD76873F18006173D2193467D3EE70000099B00000000",
	"ledger_index": "validated"
}`

	if err := testutil.Serialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestNFTInfoResponse(t *testing.T) {
	s := NFTInfoResponse{
		NFTokenID:   "00080000B4F4AFC5FBCBD76873F18006173D2193467D3EE70000099B00000000",
		LedgerIndex: 270,
		Owner:       "rG9gdNygQ6npA9JvDFWBoeXbiUcTYJnEnk",
		IsBurned:    true,
		Flags:       8,
		TransferFee: 0,
		Issuer:      "rHVokeuSnjPjz718qdb47bGXBBHNMP3KDQ",
	}
	j := `{
	"nft_id": "00080000B4F4AFC5FBCBD76873F18006173D2193467D3EE70000099B00000000",
	"ledger_index": 270,
	"owner": "rG9gdNygQ6npA9JvDFWBoeXbiUcTYJnEnk",
	"is_burned": true,
	"flags": 8,
	"transfer_fee": 0,
	"issuer": "rHVokeuSnjPjz718qdb47bGXBBHNMP3KDQ",
	"nft_taxon": 0,
	"nft_sequence": 0,
	"uri": ""
}`
	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
