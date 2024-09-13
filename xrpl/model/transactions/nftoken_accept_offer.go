package transactions

import (
	"encoding/json"
	"errors"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type NFTokenAcceptOffer struct {
	BaseTx
	NFTokenSellOffer types.Hash256        `json:",omitempty"`
	NFTokenBuyOffer  types.Hash256        `json:",omitempty"`
	NFTokenBrokerFee types.CurrencyAmount `json:",omitempty"`
}

func (*NFTokenAcceptOffer) TxType() TxType {
	return NFTokenAcceptOfferTx
}

func (n *NFTokenAcceptOffer) UnmarshalJSON(data []byte) error {
	type naoHelper struct {
		BaseTx
		NFTokenSellOffer types.Hash256   `json:",omitempty"`
		NFTokenBuyOffer  types.Hash256   `json:",omitempty"`
		NFTokenBrokerFee json.RawMessage `json:",omitempty"`
	}
	var h naoHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*n = NFTokenAcceptOffer{
		BaseTx:           h.BaseTx,
		NFTokenSellOffer: h.NFTokenSellOffer,
		NFTokenBuyOffer:  h.NFTokenBuyOffer,
	}

	fee, err := types.UnmarshalCurrencyAmount(h.NFTokenBrokerFee)
	if err != nil {
		return err
	}
	n.NFTokenBrokerFee = fee
	return nil
}

// TODO: Implement flatten
func (s *NFTokenAcceptOffer) Flatten() FlatTransaction {
	return nil
}

func ValidateNFTokenAcceptOffer(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	_, hasSellOffer := tx["NFTokenSellOffer"]
	_, hasBuyOffer := tx["NFTokenBuyOffer"]
	_, hasNFTokenBrokerFee := tx["NFTokenBrokerFee"]

	// Validate broker fee
	if hasNFTokenBrokerFee {
		err = validateNFTokenBrokerFee(tx, hasSellOffer, hasBuyOffer)
		if err != nil {
			return err
		}
	}

	// Check either NFTokenSellOffer or NFTokenBuyOffer is present

	if !hasSellOffer && !hasBuyOffer {
		return errors.New("NFTokenAcceptOffer: must set either NFTokenSellOffer or NFTokenBuyOffer")
	}

	return nil
}

func validateNFTokenBrokerFee(tx FlatTransaction, hasSellOffer, hasBuyOffer bool) error {
	value, err := ParseAmountValue(tx["NFTokenBrokerFee"])
	if err != nil {
		return err
	}

	if value <= 0 {
		return errors.New("NFTokenAcceptOffer: NFTokenBrokerFee must be greater than 0; omit if there is no fee")
	}

	// Check if both NFTokenSellOffer and NFTokenBuyOffer are set
	if !hasSellOffer && !hasBuyOffer {
		return errors.New("NFTokenAcceptOffer: both NFTokenSellOffer and NFTokenBuyOffer must be set if using brokered mode")
	}

	return nil
}
