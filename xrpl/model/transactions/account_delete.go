package transactions

import (
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type AccountDelete struct {
	BaseTx
	Destination    types.Address
	DestinationTag uint `json:",omitempty"`
}

func (*AccountDelete) TxType() TxType {
	return AccountDeleteTx
}

// TODO: Implement flatten
func (s *AccountDelete) Flatten() FlatTransaction {
	return nil
}

func ValidateAccountDelete(tx FlatTransaction) error {
	// Validate base transaction fields
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	// TODO: Replace IsString by a check for valid xrpl address
	ValidateRequiredField(tx, "Destination", typecheck.IsString)

	ValidateOptionalField(tx, "DestinationTag", typecheck.IsInt)

	return nil
}
