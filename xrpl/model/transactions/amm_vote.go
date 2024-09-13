package transactions

import (
	"errors"
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
)

type AMMVote struct {
	BaseTx
}

func (*AMMVote) TxType() TxType {
	return AMMVoteTx
}

// TODO: Implement flatten
func (s *AMMVote) Flatten() FlatTransaction {
	return nil
}

func ValidateAMMVote(tx FlatTransaction) error {
	// Validate base transaction fields
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	// Check if the Asset field is set
	if _, ok := tx["Asset"]; !ok {
		return errors.New("AMMVote: missing field Asset")
	}

	// Check if Asset is an Asset
	if !IsAsset(tx["Asset"]) {
		return errors.New("AMMVote: Asset must be a Currency")
	}

	// Check if the Asset2 field is set
	if _, ok := tx["Asset2"]; !ok {
		return errors.New("AMMVote: missing field Asset2")
	}

	// Check if Asset2 is an Asset
	if !IsAsset(tx["Asset2"]) {
		return errors.New("AMMVote: Asset2 must be a Currency")
	}

	// Check if the TradingFee field is set
	if _, ok := tx["TradingFee"]; !ok {
		return errors.New("AMMVote: missing field TradingFee")
	}

	// Check if TradingFee is a number TODO: check later if IsInt is the right check
	if !typecheck.IsInt(tx["TradingFee"]) {
		return errors.New("AMMVote: TradingFee must be a number")
	}

	// Check if TradingFee is between 0 and AMM_MAX_TRADING_FEE
	if tx["TradingFee"].(int) < 0 || tx["TradingFee"].(int) > AMM_MAX_TRADING_FEE {
		return fmt.Errorf("AMMVote: TradingFee must be between 0 and %v", AMM_MAX_TRADING_FEE)
	}

	return nil
}
