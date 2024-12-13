package clio

import (
	"testing"

	cliotypes "github.com/Peersyst/xrpl-go/xrpl/queries/clio/types"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
)

func TestNFTsByIssuerRequest(t *testing.T) {
	s := NFTsByIssuerRequest{
		Issuer:   "abc",
		Marker:   "123",
		Limit:    10,
		NftTaxon: 1,
	}

	j := `{
	"issuer": "abc",
	"marker": "123",
	"limit": 10,
	"nft_taxon": 1
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestNFTsByIssuerResponse(t *testing.T) {
	s := NFTsByIssuerResponse{
		Issuer: "abc",
		NFTs: []cliotypes.NFToken{
			{
				NFTokenID:       "123",
				LedgerIndex:     1,
				Owner:           "abc",
				IsBurned:        false,
				Flags:           0,
				TransferFee:     0,
				Issuer:          "abc",
				NFTokenTaxon:    1,
				NFTokenSequence: 1,
				URI:             "abc",
			},
		},
		Marker:       "123",
		Limit:        10,
		NFTokenTaxon: 1,
	}

	j := `{
	"issuer": "abc",
	"nfts": [
		{
			"nft_id": "123",
			"ledger_index": 1,
			"owner": "abc",
			"is_burned": false,
			"flags": 0,
			"transfer_fee": 0,
			"issuer": "abc",
			"nft_taxon": 1,
			"nft_sequence": 1,
			"uri": "abc"
		}
	],
	"marker": "123",
	"limit": 10,
	"nft_taxon": 1
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
