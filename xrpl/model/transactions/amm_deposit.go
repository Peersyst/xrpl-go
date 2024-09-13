package transactions

import "errors"

type AMMDeposit struct {
	BaseTx
}

func (*AMMDeposit) TxType() TxType {
	return AMMDepositTx
}

// TODO: Implement flatten
func (s *AMMDeposit) Flatten() FlatTransaction {
	return nil
}

func ValidateAMMDeposit(tx FlatTransaction) error {
	// Validate base transaction fields
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	// Check if the field Asset is set
	if _, ok := tx["Asset"]; !ok {
		return errors.New("AMMDeposit: missing field Asset")
	}

	// Check if the field Asset is an Asset
	if !IsAsset(tx["Asset"]) {
		return errors.New("AMMDeposit: Asset must be an issued currency or XRP")
	}

	// Check if the field Asset2 is set
	if _, ok := tx["Asset2"]; !ok {
		return errors.New("AMMDeposit: missing field Asset2")
	}

	// Check if the field Asset2 is an Asset
	if !IsAsset(tx["Asset2"]) {
		return errors.New("AMMDeposit: Asset2 must be an issued currency or XRP")
	}

	if tx["Amount2"] != nil && tx["Amount"] == nil {
		return errors.New("AMMDeposit: must set Amount with Amount2")
	} else if tx["EPrice"] != nil && tx["Amount"] == nil {
		return errors.New("AMMDeposit: must set Amount with EPrice")
	} else if tx["LPTokenOut"] == nil && tx["Amount"] == nil {
		return errors.New("AMMDeposit: must set at least LPTokenOut or Amount")
	}

	// Check if the field LPTokenOut is an IssuedCurrencyAmount
	if tx["LPTokenOut"] != nil && !IsIssuedCurrency(tx["LPTokenOut"]) {
		return errors.New("AMMDeposit: LPTokenOut must be an IssuedCurrencyAmount")
	}

	// Check if the field Amount is an Amount
	if tx["Amount"] != nil && !IsAmount(tx["Amount"]) {
		return errors.New("AMMDeposit: Amount must be an Amount")
	}

	// Check if the field Amount2 is an Amount
	if tx["Amount2"] != nil && !IsAmount(tx["Amount2"]) {
		return errors.New("AMMDeposit: Amount2 must be an Amount")
	}

	// Check if the field EPrice is an Amount
	if tx["EPrice"] != nil && !IsAmount(tx["EPrice"]) {
		return errors.New("AMMDeposit: EPrice must be an Amount")
	}

	return nil
}
