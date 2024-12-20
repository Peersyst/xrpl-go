package clio

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type NFTInfoRequest struct {
	NFTokenID   types.NFTokenID        `json:"nft_id"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
}

func (*NFTInfoRequest) Method() string {
	return "nft_info"
}

func (r *NFTInfoRequest) UnmarshalJSON(data []byte) error {
	type nirHelper struct {
		NFTokenID   types.NFTokenID   `json:"nft_id"`
		LedgerHash  common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage   `json:"ledger_index,omitempty"`
	}
	var h nirHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*r = NFTInfoRequest{
		NFTokenID:  h.NFTokenID,
		LedgerHash: h.LedgerHash,
	}
	r.LedgerIndex, err = common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	return err
}
