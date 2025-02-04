package account

import (
	accounttypes "github.com/Peersyst/xrpl-go/xrpl/queries/account/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The `account_nfts` method retrieves all of the NFTs currently owned by the
// specified account.
type NFTsRequest struct {
	common.BaseRequest
	Account     types.Address          `json:"account"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
}

func (*NFTsRequest) Method() string {
	return "account_nfts"
}

func (*NFTsRequest) APIVersion() int {
	return version.RippledAPIV2
}

// TODO: Implement (V2)
func (*NFTsRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the account_nfts method.
type NFTsResponse struct {
	Account            types.Address      `json:"account"`
	AccountNFTs        []accounttypes.NFT `json:"account_nfts"`
	LedgerIndex        common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash  `json:"ledger_hash,omitempty"`
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index,omitempty"`
	Validated          bool               `json:"validated"`
	Marker             any                `json:"marker,omitempty"`
	Limit              int                `json:"limit,omitempty"`
}
