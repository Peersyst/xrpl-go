package path

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type NFTokenSellOffersRequest struct {
	NFTokenID   types.NFTokenID        `json:"nft_id"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
}

func (*NFTokenSellOffersRequest) Method() string {
	return "nft_sell_offers"
}

func (r *NFTokenSellOffersRequest) UnmarshalJSON(data []byte) error {
	type borHelper struct {
		NFTokenID   types.NFTokenID   `json:"nft_id"`
		LedgerHash  common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage   `json:"ledger_index,omitempty"`
		Limit       int               `json:"limit,omitempty"`
		Marker      any               `json:"marker,omitempty"`
	}
	var h borHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*r = NFTokenSellOffersRequest{
		NFTokenID:  h.NFTokenID,
		LedgerHash: h.LedgerHash,
		Limit:      h.Limit,
		Marker:     h.Marker,
	}
	var i common.LedgerSpecifier
	i, err = common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
