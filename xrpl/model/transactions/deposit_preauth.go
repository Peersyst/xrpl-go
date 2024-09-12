package transactions

import (
	"errors"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type DepositPreauth struct {
	BaseTx

	// The XRP Ledger address of the sender to preauthorize.
	Authorize types.Address `json:",omitempty"`

	//   The XRP Ledger address of a sender whose preauthorization should be.
	//   revoked.
	Unauthorize types.Address `json:",omitempty"`
}

func (*DepositPreauth) TxType() TxType {
	return DepositPreauthTx
}

// TODO: Implement flatten
func (s *DepositPreauth) Flatten() FlatTransaction {
	return nil
}

func ValidateDepositPreauth(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	_, hasAuthorize := tx["Authorize"]
	_, hasUnauthorize := tx["Unauthorize"]

	// Check if both Authorize and Unauthorize are set
	if hasAuthorize && hasUnauthorize {
		return errors.New("DepositPreauth: can't provide both Authorize and Unauthorize fields")
	}

	// Check if neither Authorize nor Unauthorize are set
	if !hasAuthorize && !hasUnauthorize {
		return errors.New("DepositPreauth: must provide either Authorize or Unauthorize field")
	}

	// Check if the field Authorize is set
	if _, ok := tx["Authorize"]; !ok {
		return errors.New("DepositPreauth: missing field Authorize")
	}

	if hasAuthorize {
		// Check if the field Authorize is a string
		if !typecheck.IsString(tx["Authorize"]) {
			return errors.New("DepositPreauth: Authorize must be a string")
		}

		// Check if the field Authorize is not the same as the tx.Account
		if tx["Authorize"] == tx["Account"] {
			return errors.New("DepositPreauth: Account can't preauthorize its own address")
		}
	}

	if hasUnauthorize {
		// Check if the field Unauthorize is a string
		if !typecheck.IsString(tx["Unauthorize"]) {
			return errors.New("DepositPreauth: Unauthorize must be a string")
		}

		// Check if the field Unauthorize is not the same as the tx.Account
		if tx["Unauthorize"] == tx["Account"] {
			return errors.New("DepositPreauth: Account can't unauthorize its own address")
		}
	}

	return nil
}
