package transaction

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

var (
	ErrInvalidRegularKey        = errors.New("invalid xrpl address for the RegularKey field")
	ErrRegularKeyMatchesAccount = errors.New("regular key must not match the account address")
)

// A SetRegularKey transaction assigns, changes, or removes the regular key pair associated with an account.
//
// You can protect your account by assigning a regular key pair to it and using it instead of the master key pair to sign transactions whenever possible.
// If your regular key pair is compromised, but your master key pair is not, you can use a SetRegularKey transaction to regain control of your account.
//
// Example:
//
// ```json
//
//	{
//	    "Flags": 0,
//	    "TransactionType": "SetRegularKey",
//	    "Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
//	    "Fee": "12",
//	    "RegularKey": "rAR8rR8sUkBoCZFawhkWzY4Y5YoyuznwD"
//	}
//
// ```
type SetRegularKey struct {
	BaseTx
	// (Optional) A base-58-encoded Address that indicates the regular key pair to be assigned to the account.
	// If omitted, removes any existing regular key pair from the account. Must not match the master key pair for the address.
	RegularKey types.Address `json:",omitempty"`
}

// TxType returns the transaction type for this transaction (SetRegularKey).
func (*SetRegularKey) TxType() TxType {
	return SetRegularKeyTx
}

// Flatten returns the flattened map of the SetRegularKey transaction.
func (s *SetRegularKey) Flatten() FlatTransaction {
	flattened := s.BaseTx.Flatten()

	flattened["TransactionType"] = "SetRegularKey"

	if s.RegularKey != "" {
		flattened["RegularKey"] = s.RegularKey.String()
	}

	return flattened
}

// Validate checks if the SetRegularKey struct is valid.
func (s *SetRegularKey) Validate() (bool, error) {
	ok, err := s.BaseTx.Validate()
	if err != nil || !ok {
		return false, err
	}

	// Check if the regular key is not the same as the account address
	if s.RegularKey != "" && s.RegularKey == s.Account {
		return false, ErrRegularKeyMatchesAccount
	}

	// Check if the regular key is a valid xrpl address
	if s.RegularKey != "" && !addresscodec.IsValidClassicAddress(s.RegularKey.String()) {
		return false, ErrInvalidRegularKey
	}

	return true, nil
}
