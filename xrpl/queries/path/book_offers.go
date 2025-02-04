package path

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	pathtypes "github.com/Peersyst/xrpl-go/xrpl/queries/path/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The book_offers method retrieves a list of offers, also known as the order
// book, between two currencies.
type BookOffersRequest struct {
	common.BaseRequest
	TakerGets   pathtypes.BookOfferCurrency `json:"taker_gets"`
	TakerPays   pathtypes.BookOfferCurrency `json:"taker_pays"`
	Taker       types.Address               `json:"taker,omitempty"`
	LedgerHash  common.LedgerHash           `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerIndex          `json:"ledger_index,omitempty"`
	Limit       int                         `json:"limit,omitempty"`
}

func (*BookOffersRequest) Method() string {
	return "book_offers"
}

func (*BookOffersRequest) APIVersion() int {
	return version.RippledAPIV2
}

// TODO: Implement V2
func (*BookOffersRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the book_offers method.
type BookOffersResponse struct {
	LedgerCurrentIndex common.LedgerIndex    `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex    `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash     `json:"ledger_hash,omitempty"`
	Offers             []pathtypes.BookOffer `json:"offers"`
	Validated          bool                  `json:"validated,omitempty"`
}
