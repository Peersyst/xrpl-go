package types

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

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
