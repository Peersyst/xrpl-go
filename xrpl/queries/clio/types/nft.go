package types

import (
	"github.com/Peersyst/xrpl-go/v1/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

// NFToken is a struct that represents an NFToken.
// It contains the information about the NFToken, such as the ID,
// the ledger index, the owner, if it is burned, the flags, the transfer fee,
// the issuer, the taxon, the sequence, the URI, and the URI JSON.
type NFToken struct {
	NFTokenID       types.NFTokenID    `json:"nft_id"`
	LedgerIndex     common.LedgerIndex `json:"ledger_index"`
	Owner           types.Address      `json:"owner"`
	IsBurned        bool               `json:"is_burned"`
	Flags           uint               `json:"flags"`
	TransferFee     uint               `json:"transfer_fee"`
	Issuer          types.Address      `json:"issuer"`
	NFTokenTaxon    uint               `json:"nft_taxon"`
	NFTokenSequence uint               `json:"nft_sequence"`
	URI             types.NFTokenURI   `json:"uri,omitempty"`
}
