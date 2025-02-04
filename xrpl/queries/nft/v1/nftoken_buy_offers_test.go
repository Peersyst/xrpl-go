package v1

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	nfttypes "github.com/Peersyst/xrpl-go/xrpl/queries/nft/types"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestNFTokenBuyOffersRequest(t *testing.T) {
	s := NFTokenBuyOffersRequest{
		NFTokenID:   "00090000D0B007439B080E9B05BF62403911301A7B1F0CFAA048C0A200000007",
		LedgerIndex: common.Validated,
	}

	j := `{
	"nft_id": "00090000D0B007439B080E9B05BF62403911301A7B1F0CFAA048C0A200000007",
	"ledger_index": "validated"
}`

	if err := testutil.Serialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestNFTokenBuyOffersResponse(t *testing.T) {
	s := NFTokenBuyOffersResponse{
		NFTokenID: "00090000D0B007439B080E9B05BF62403911301A7B1F0CFAA048C0A200000007",
		Offers: []nfttypes.NFTokenOffer{
			{
				Amount:            types.XRPCurrencyAmount(1500),
				Flags:             0,
				NFTokenOfferIndex: "3212D26DB00031889D4EF7D9129BB0FA673B5B40B1759564486C0F0946BA203F",
				Owner:             "rsuHaTvJh1bDmDoxX9QcKP7HEBSBt4XsHx",
			},
		},
	}

	j := `{
	"nft_id": "00090000D0B007439B080E9B05BF62403911301A7B1F0CFAA048C0A200000007",
	"offers": [
		{
			"amount": "1500",
			"flags": 0,
			"nft_offer_index": "3212D26DB00031889D4EF7D9129BB0FA673B5B40B1759564486C0F0946BA203F",
			"owner": "rsuHaTvJh1bDmDoxX9QcKP7HEBSBt4XsHx"
		}
	]
}`

	if err := testutil.Serialize(t, s, j); err != nil {
		t.Error(err)
	}
}
