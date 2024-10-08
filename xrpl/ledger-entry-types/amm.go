package ledger

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type AMM struct {
	// The address of the special account that holds this AMM's assets.
	Account types.Address
	// The definition for one of the two assets this AMM holds. In JSON, this is an object with currency and issuer fields.
	Asset Asset
	// The definition for the other asset this AMM holds. In JSON, this is an object with currency and issuer fields.
	Asset2 Asset
	// Details of the current owner of the auction slot, as an Auction Slot object.
	AuctionSlot AuctionSlot `json:",omitempty"`
	// The total outstanding balance of liquidity provider tokens from this AMM instance.
	// The holders of these tokens can vote on the AMM's trading fee in proportion to their holdings, or redeem the tokens for a share of the AMM's assets which grows with the trading fees collected.
	LPTokenBalance types.CurrencyAmount
	// The percentage fee to be charged for trades against this AMM instance, in units of 1/100,000. The maximum value is 1000, for a 1% fee.
	TradingFee uint16
	// A list of vote objects, representing votes on the pool's trading fee.
	VoteSlots []VoteEntry `json:",omitempty"`
}

func (a *AMM) UnmarshalJSON(data []byte) error {
	type ammHelper struct {
		Asset          Asset
		Asset2         Asset
		Account        types.Address
		AuctionSlot    AuctionSlot
		LPTokenBalance json.RawMessage
		TradingFee     uint16
		VoteSlots      []VoteEntry
	}
	var h ammHelper
	var err error
	if err = json.Unmarshal(data, &h); err != nil {
		return err
	}
	*a = AMM{
		Asset:       h.Asset,
		Asset2:      h.Asset2,
		Account:     h.Account,
		AuctionSlot: h.AuctionSlot,
		TradingFee:  h.TradingFee,
		VoteSlots:   h.VoteSlots,
	}

	a.LPTokenBalance, err = types.UnmarshalCurrencyAmount(h.LPTokenBalance)
	if err != nil {
		return err
	}

	return nil
}

// ---------------------------------------------
// Asset Object
// ---------------------------------------------

type Asset struct {
	Currency string
	Issuer   types.Address
}

func (a *Asset) Flatten() map[string]interface{} {
	var flattened = make(map[string]interface{})

	if a.Currency != "" {
		flattened["Currency"] = a.Currency
	}

	if a.Issuer != "" {
		flattened["Issuer"] = a.Issuer.String()
	}

	return flattened
}

// ---------------------------------------------
// Auction Slot Object
// ---------------------------------------------

type AuctionSlot struct {
	// The current owner of this auction slot.
	Account types.Address
	// A list of at most 4 additional accounts that are authorized to trade at the discounted fee for this AMM instance.
	AuthAccounts []AuthAccounts `json:",omitempty"`
	// The trading fee to be charged to the auction owner, in the same format as TradingFee. Normally, this is 1/10 of the normal fee for this AMM.
	DiscountedFee uint32
	// The amount the auction owner paid to win this slot, in LP Tokens.
	Price types.CurrencyAmount
	// The time when this slot expires, in seconds since the Ripple Epoch. https://xrpl.org/docs/references/protocol/data-types/basic-data-types#specifying-time.
	Expiration uint32
}

// UnmarshalJSON is a custom JSON unmarshaler for the AuctionSlot struct.
// It decodes JSON data into an AuctionSlot instance, handling the nested
// structure and converting the Price field using a custom unmarshal function.
// If the JSON data cannot be unmarshaled into the expected structure,
// or if the Price field cannot be converted, an error is returned.
func (s *AuctionSlot) UnmarshalJSON(data []byte) error {
	type aasHelper struct {
		Account       types.Address
		AuthAccounts  []AuthAccounts
		DiscountedFee uint32
		Price         json.RawMessage
		Expiration    uint32
	}
	var h aasHelper
	var err error
	if err = json.Unmarshal(data, &h); err != nil {
		return err
	}
	*s = AuctionSlot{
		Account:       h.Account,
		AuthAccounts:  h.AuthAccounts,
		DiscountedFee: h.DiscountedFee,
		Expiration:    h.Expiration,
	}

	s.Price, err = types.UnmarshalCurrencyAmount(h.Price)
	if err != nil {
		return err
	}
	return nil
}

// Flatten converts the AuctionSlot struct into a map with string keys and interface{} values.
// It includes non-zero and non-nil fields from the AuctionSlot struct.
func (a AuctionSlot) Flatten() map[string]interface{} {
	var flattened = make(map[string]interface{})

	if a.Account != "" {
		flattened["Account"] = a.Account.String()
	}

	if a.DiscountedFee != 0 {
		flattened["DiscountedFee"] = a.DiscountedFee
	}

	if a.Price != nil {
		flattened["Price"] = a.Price.Flatten()
	}

	if a.Expiration != 0 {
		flattened["Expiration"] = a.Expiration
	}

	if len(a.AuthAccounts) > 0 {
		flattenedAuthAccounts := make([]interface{}, 0)
		for _, authAccount := range a.AuthAccounts {
			flattenedAuthAccount := authAccount.Flatten()
			if flattenedAuthAccount != nil {
				flattenedAuthAccounts = append(flattenedAuthAccounts, flattenedAuthAccount)
			}
		}
		flattened["AuthAccounts"] = flattenedAuthAccounts
	}

	return flattened
}

// ---------------------------------------------
// AuthAccounts Object
// ---------------------------------------------

type AuthAccounts struct {
	AuthAccount AuthAccount
}

type AuthAccount struct {
	// Authorized account to trade at the discounted fee for this AMM instance.
	Account types.Address
}

func (a AuthAccounts) Flatten() map[string]interface{} {
	var flattened = make(map[string]interface{})

	if a.AuthAccount.Account != "" {
		flattened["Account"] = a.AuthAccount.Account.String()
	}

	return flattened
}

type VoteEntry struct {
	Account     types.Address
	TradingFee  uint
	VoteWeither uint
}
