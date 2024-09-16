package transactions

import (
	"errors"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type NFTokenCancelOffer struct {
	BaseTx

	// An array of identifiers of NFTokenOffer objects that should be cancelled
	// by this transaction.
	//
	// It is an error if an entry in this list points to an
	// object that is not an NFTokenOffer object. It is not an
	// error if an entry in this list points to an object that
	// does not exist. This field is required.
	NFTokenOffers []types.Hash256
}

func (*NFTokenCancelOffer) TxType() TxType {
	return NFTokenCancelOfferTx
}

// TODO: Implement flatten
func (s *NFTokenCancelOffer) Flatten() FlatTransaction {
	return nil
}

func ValidateNFTokenCancelOffer(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	if !typecheck.IsArrayOrSlice(tx["NFTokenOffers"]) {
		return errors.New("field 'NFTokenOffers' must be an array")
	}

	if len(tx["NFTokenOffers"].([]types.Hash256)) == 0 {
		return errors.New("field 'NFTokenOffers' must not be empty")
	}

	return nil
}
