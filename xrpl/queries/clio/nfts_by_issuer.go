package clio

import (
	cliotypes "github.com/Peersyst/xrpl-go/xrpl/queries/clio/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The nfts_by_issuer method returns a list of NFTokens issued by the account.
// The order of the NFTs is not associated with the date the NFTs were minted.
type NFTsByIssuerRequest struct {
	common.BaseRequest
	Issuer   types.Address `json:"issuer"`
	Marker   any           `json:"marker,omitempty"`
	Limit    int           `json:"limit,omitempty"`
	NftTaxon uint32        `json:"nft_taxon,omitempty"`
}

func (*NFTsByIssuerRequest) Method() string {
	return "nfts_by_issuer"
}

func (*NFTsByIssuerRequest) APIVersion() int {
	return version.RippledAPIV2
}

// TODO: Implement V2
func (*NFTsByIssuerRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the nfts_by_issuer method.
type NFTsByIssuerResponse struct {
	Issuer       types.Address       `json:"issuer"`
	NFTs         []cliotypes.NFToken `json:"nfts"`
	Marker       any                 `json:"marker,omitempty"`
	Limit        int                 `json:"limit,omitempty"`
	NFTokenTaxon uint32              `json:"nft_taxon,omitempty"`
}
