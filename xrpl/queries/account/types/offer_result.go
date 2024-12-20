package types

type OfferResultFlags uint

type OfferResult struct {
	Flags    OfferResultFlags `json:"flags"`
	Sequence uint             `json:"seq"`
	// TakerGets  types.CurrencyAmount `json:"taker_gets"`
	TakerGets any `json:"taker_gets"`
	// TakerPays  types.CurrencyAmount `json:"taker_pays"`
	TakerPays  any    `json:"taker_pays"`
	Quality    string `json:"quality"`
	Expiration uint   `json:"expiration,omitempty"`
}
