package nft

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	nfttypes "github.com/Peersyst/xrpl-go/xrpl/queries/nft/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The nft_sell_offers method retrieves all of sell offers for the specified
// NFToken.
type NFTokenSellOffersRequest struct {
	common.BaseRequest
	NFTokenID   types.NFTokenID        `json:"nft_id"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
}

func (*NFTokenSellOffersRequest) Method() string {
	return "nft_sell_offers"
}

func (*NFTokenSellOffersRequest) APIVersion() int {
	return version.RippledAPIV2
}

// TODO: Implement V2
func (*NFTokenSellOffersRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the nft_sell_offers method.
type NFTokenSellOffersResponse struct {
	NFTokenID types.NFTokenID         `json:"nft_id"`
	Offers    []nfttypes.NFTokenOffer `json:"offers"`
}
