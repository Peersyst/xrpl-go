package types

import (
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

type NFTokenOffer struct {
	Amount            any           `json:"amount"`
	Flags             uint          `json:"flags"`
	NFTokenOfferIndex string        `json:"nft_offer_index"`
	Owner             types.Address `json:"owner"`
	Destination       types.Address `json:"destination,omitempty"`
	Expiration        int           `json:"expiration,omitempty"`
}
