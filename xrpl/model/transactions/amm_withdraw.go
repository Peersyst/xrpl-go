package transactions

import "errors"

type AMMWithdraw struct {
	BaseTx
}

func (*AMMWithdraw) TxType() TxType {
	return AMMWithdrawTx
}

// TODO: Implement flatten
func (s *AMMWithdraw) Flatten() FlatTransaction {
	return nil
}

func ValidateAMMWithdraw(tx FlatTransaction) error {
	// Validate base transaction fields
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	// Check if the Asset field is set
	if _, ok := tx["Asset"]; !ok {
		return errors.New("AMMWithdraw: missing field Asset")
	}

	// Chck if Asset is an Asset
	if !IsAsset(tx["Asset"]) {
		return errors.New("AMMWithdraw: Asset must be a Currency")
	}

	// Check if the Asset2 field is set
	if _, ok := tx["Asset2"]; !ok {
		return errors.New("AMMWithdraw: missing field Asset2")
	}

	// Check if Asset2 is an Asset
	if !IsAsset(tx["Asset2"]) {
		return errors.New("AMMWithdraw: Asset2 must be a Currency")
	}

	if tx["Amount2"] != nil && tx["Amount"] == nil {
		return errors.New("AMMWithdraw: must set Amount with Amount2")
	} else if tx["EPrice"] != nil && tx["Amount"] == nil {
		return errors.New("AMMWithdraw: must set Amount with EPrice")
	}

	if tx["LPTokenIn"] != nil && !IsIssuedCurrency(tx["LPTokenIn"]) {
		return errors.New("AMMWithdraw: must set at least LPTokenOut or Amount")
	}

	// Check if the field Amount is an Amount
	if tx["Amount"] != nil && !IsAmount(tx["Amount"]) {
		return errors.New("AMMWithdraw: Amount must be an Amount")
	}

	// Check if the field Amount2 is an Amount
	if tx["Amount2"] != nil && !IsAmount(tx["Amount2"]) {
		return errors.New("AMMWithdraw: Amount2 must be an Amount")
	}

	// Check if the field EPrice is an Amount
	if tx["EPrice"] != nil && !IsAmount(tx["EPrice"]) {
		return errors.New("AMMWithdraw: EPrice must be an Amount")
	}

	return nil
}
