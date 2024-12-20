package path

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type BookOffersRequest struct {
	TakerGets   types.IssuedCurrencyAmount `json:"taker_gets"`
	TakerPays   types.IssuedCurrencyAmount `json:"taker_pays"`
	LedgerHash  common.LedgerHash          `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier     `json:"ledger_index,omitempty"`
	Limit       int                        `json:"limit,omitempty"`
	Taker       types.Address              `json:"taker,omitempty"`
}

func (*BookOffersRequest) Method() string {
	return "book_offers"
}

func (r *BookOffersRequest) UnmarshalJSON(data []byte) error {
	type borHelper struct {
		TakerGets   types.IssuedCurrencyAmount `json:"taker_gets"`
		TakerPays   types.IssuedCurrencyAmount `json:"taker_pays"`
		LedgerHash  common.LedgerHash          `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage            `json:"ledger_index,omitempty"`
		Limit       int                        `json:"limit,omitempty"`
		Taker       types.Address              `json:"taker,omitempty"`
	}
	var h borHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*r = BookOffersRequest{
		TakerGets:  h.TakerGets,
		TakerPays:  h.TakerPays,
		LedgerHash: h.LedgerHash,
		Limit:      h.Limit,
		Taker:      h.Taker,
	}
	var i common.LedgerSpecifier
	i, err = common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}
