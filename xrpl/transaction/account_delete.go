package transaction

import (
	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

// Added by the DeletableAccounts amendment
// An AccountDelete transaction deletes an account and any objects it owns in the XRP Ledger, if possible,
// sending the account's remaining XRP to a specified destination account. See Deleting Accounts for the requirements to delete an account.
//
// ```json
//
//	{
//	    "TransactionType": "AccountDelete",
//	    "Account": "rWYkbWkCeg8dP6rXALnjgZSjjLyih5NXm",
//	    "Destination": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
//	    "DestinationTag": 13,
//	    "Fee": "2000000",
//	    "Sequence": 2470665,
//	    "Flags": 2147483648
//	}
//
// ```
type AccountDelete struct {
	BaseTx
	// The address of an account to receive any leftover XRP after deleting the sending account.
	// Must be a funded account in the ledger, and must not be the sending account.
	Destination types.Address
	// (Optional) Arbitrary destination tag that identifies a hosted recipient or other information
	// for the recipient of the deleted account's leftover XRP.
	DestinationTag uint32 `json:",omitempty"`
}

// TxType implements the TxType method for the AccountDelete struct.
func (*AccountDelete) TxType() TxType {
	return AccountDeleteTx
}

// Flatten implements the Flatten method for the AccountDelete struct.
func (s *AccountDelete) Flatten() FlatTransaction {
	flatTx := s.BaseTx.Flatten()
	flatTx["TransactionType"] = s.TxType().String()

	if s.Destination != "" {
		flatTx["Destination"] = s.Destination.String()
	}

	if s.DestinationTag != 0 {
		flatTx["DestinationTag"] = s.DestinationTag
	}
	return flatTx
}

// Validate implements the Validate method for the AccountDelete struct.
func (s *AccountDelete) Validate() (bool, error) {
	_, err := s.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if !addresscodec.IsValidClassicAddress(s.Destination.String()) {
		return false, ErrInvalidDestinationAddress
	}

	return true, nil
}
