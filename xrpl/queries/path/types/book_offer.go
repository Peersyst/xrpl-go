package types

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type BookOffer struct {
	ledger.Offer
	OwnerFunds      string               `json:"owner_funds,omitempty"`
	TakerGetsFunded types.CurrencyAmount `json:"taker_gets_funded,omitempty"`
	TakerPaysFunded types.CurrencyAmount `json:"taker_pays_funded,omitempty"`
	Quality         string               `json:"quality,omitempty"`
}

type BookOfferCurrency struct {
	Currency string `json:"currency"`
	Issuer   string `json:"issuer,omitempty"`
}
