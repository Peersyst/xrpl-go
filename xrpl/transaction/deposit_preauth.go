package transaction

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

var (
	ErrDepositPreauthInvalidAuthorize                     = errors.New("deposit preauth: invalid Authorize")
	ErrDepositPreauthInvalidUnauthorize                   = errors.New("deposit preauth: invalid Unauthorize")
	ErrDepositPreauthCannotSetBothAuthorizeAndUnauthorize = errors.New("deposit preauth: cannot set both Authorize and Unauthorize")
	ErrDepositPreauthMustSetEitherAuthorizeOrUnauthorize  = errors.New("deposit preauth: must set either Authorize or Unauthorize")
	ErrDepositPreauthAuthorizeCannotBeSender              = errors.New("deposit preauth: Authorize cannot be the same as the sender's account")
	ErrDepositPreauthUnauthorizeCannotBeSender            = errors.New("deposit preauth: Unauthorize cannot be the same as the sender's account")
)

// Added by the DepositPreauth amendment.
// A DepositPreauth transaction gives another account pre-approval to deliver payments to the sender
// of this transaction.
// This is only useful if the sender of this transaction is using (or plans to use) Deposit
// Authorization.
//
// ```json
//
//	{
//	  "TransactionType" : "DepositPreauth",
//	  "Account" : "rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8",
//	  "Authorize" : "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
//	  "Fee" : "10",
//	  "Flags" : 2147483648,
//	  "Sequence" : 2
//	}
//
// ```
type DepositPreauth struct {
	BaseTx
	// (Optional) The XRP Ledger address of the sender to preauthorize.
	Authorize types.Address `json:",omitempty"`
	// (Optional) The XRP Ledger address of a sender whose preauthorization should be revoked.
	Unauthorize types.Address `json:",omitempty"`
}

// TxType implements the TxType method for the DepositPreauth struct.
func (*DepositPreauth) TxType() TxType {
	return DepositPreauthTx
}

// Flatten implements the Flatten method for the DepositPreauth struct.
func (s *DepositPreauth) Flatten() FlatTransaction {
	flattened := s.BaseTx.Flatten()

	flattened["TransactionType"] = DepositPreauthTx.String()

	if s.Authorize != "" {
		flattened["Authorize"] = s.Authorize.String()
	}

	if s.Unauthorize != "" {
		flattened["Unauthorize"] = s.Unauthorize.String()
	}

	return flattened
}

// Validate implements the Validate method for the DepositPreauth struct.
func (s *DepositPreauth) Validate() (bool, error) {
	_, err := s.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if s.Authorize != "" && s.Unauthorize != "" {
		return false, ErrDepositPreauthCannotSetBothAuthorizeAndUnauthorize
	}

	if s.Authorize == "" && s.Unauthorize == "" {
		return false, ErrDepositPreauthMustSetEitherAuthorizeOrUnauthorize
	}

	if s.Authorize != "" && s.Authorize.String() == s.Account.String() {
		return false, ErrDepositPreauthAuthorizeCannotBeSender
	}

	if s.Unauthorize != "" && s.Unauthorize.String() == s.Account.String() {
		return false, ErrDepositPreauthUnauthorizeCannotBeSender
	}

	if s.Authorize != "" && !addresscodec.IsValidClassicAddress(s.Authorize.String()) {
		return false, ErrDepositPreauthInvalidAuthorize
	}

	if s.Unauthorize != "" && !addresscodec.IsValidClassicAddress(s.Unauthorize.String()) {
		return false, ErrDepositPreauthInvalidUnauthorize
	}

	return true, nil
}
