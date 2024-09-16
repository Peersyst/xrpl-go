package transactions

import (
	"errors"
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

const (
	//
	// Account Set Flags
	//
	// Require a destination tag to send transactions to this account.
	asfRequireDest uint = 1
	// Require authorization for users to hold balances issued by this address.
	// Can only be enabled if the address has no trust lines connected to it.
	asfRequireAuth uint = 2
	// XRP should not be sent to this account.
	asfDisallowXRP uint = 3
	// Disallow use of the master key pair. Can only be enabled if the account
	// has configured another way to sign transactions, such as a Regular Key or a
	// Signer List.
	asfDisableMaster uint = 4
	// Track the ID of this account's most recent transaction. Required for
	// AccountTxnID.
	asfAccountTxnID uint = 5
	// Permanently give up the ability to freeze individual trust lines or
	// disable Global Freeze. This flag can never be disabled after being enabled.
	asfNoFreeze uint = 6
	// Freeze all assets issued by this account.
	asfGlobalFreeze uint = 7
	// Enable rippling on this account's trust lines by default.
	asfDefaultRipple uint = 8
	// Enable Deposit Authorization on this account.
	asfDepositAuth uint = 9
	// Allow another account to mint and burn tokens on behalf of this account.
	asfAuthorizedNFTokenMinter uint = 10
	// asf 11 is reserved for Hooks amendment
	// Disallow other accounts from creating incoming NFTOffers
	asfDisallowIncomingNFTokenOffer uint = 12
	// Disallow other accounts from creating incoming Checks
	asfDisallowIncomingCheck uint = 13
	// Disallow other accounts from creating incoming Payment Channels
	asfDisallowIncomingPayChan uint = 14
	// Disallow other accounts from creating incoming TrustLines
	asfDisallowIncomingTrustLine uint = 15
	// Permanently gain the ability to claw back issued IOUs
	asfAllowTrustLineClawback uint = 16

	//
	// Transaction Flags
	//
	// The same as SetFlag: asfRequireDest.
	tfRequireDestTag uint = 65536 // 0x00010000
	// The same as ClearFlag: asfRequireDestTag.
	tfOptionalDestTag uint = 131072 // 0x00020000
	// The same as SetFlag: asfRequireAuth.
	tfRequireAuth uint = 262144 // 0x00040000
	// The same as ClearFlag: asfRequireAuth.
	tfOptionalAuth uint = 524288 // 0x00080000
	// The same as SetFlag: asfDisallowXRP.
	tfDisallowXRP uint = 1048576 // 0x00100000
	// The same as ClearFlag: asfDisallowXRP.
	tfAllowXRP uint = 2097152 // 0x00200000
)

// An AccountSet transaction modifies the properties of an account in the XRP
// Ledger.
type AccountSet struct {
	BaseTx
	// ClearFlag: asfRequireDestTag, asfOptionalDestTag, asfRequireAuth, asfOptionalAuth, asfDisallowXRP, asfAllowXRP
	ClearFlag uint `json:",omitempty"`
	// The domain that owns this account, as a string of hex representing the.
	// ASCII for the domain in lowercase.
	Domain string `json:",omitempty"`
	// Hash of an email address to be used for generating an avatar image.
	EmailHash types.Hash128 `json:",omitempty"`
	//Public key for sending encrypted messages to this account.
	MessageKey string `json:",omitempty"`
	// Sets an alternate account that is allowed to mint NFTokens on this
	// account's behalf using NFTokenMint's `Issuer` field.
	NFTokenMinter string `json:",omitempty"`
	// Integer flag to enable for this account.
	SetFlag uint `json:",omitempty"`
	// The fee to charge when users transfer this account's issued currencies,
	// represented as billionths of a unit. Cannot be more than 2000000000 or less
	// than 1000000000, except for the special case 0 meaning no fee.
	TransferRate uint `json:",omitempty"`
	// Tick size to use for offers involving a currency issued by this address.
	// The exchange rates of those offers is rounded to this many significant
	// digits. Valid values are 3 to 15 inclusive, or 0 to disable.
	TickSize      uint8         `json:",omitempty"`
	WalletLocator types.Hash256 `json:",omitempty"`
	WalletSize    uint          `json:",omitempty"`
}

// TxType returns the type of the transaction (AccountSet).
func (*AccountSet) TxType() TxType {
	return AccountSetTx
}

// Flatten returns the flattened map of the AccountSet transaction.
func (s *AccountSet) Flatten() FlatTransaction {
	flattened := s.BaseTx.Flatten()

	flattened["TransactionType"] = "AccountSet"

	if s.ClearFlag != 0 {
		flattened["ClearFlag"] = int(s.ClearFlag)
	}
	if s.Domain != "" {
		flattened["Domain"] = s.Domain
	}
	if s.EmailHash != "" {
		flattened["EmailHash"] = s.EmailHash.String()
	}
	if s.MessageKey != "" {
		flattened["MessageKey"] = s.MessageKey
	}
	if s.NFTokenMinter != "" {
		flattened["NFTokenMinter"] = s.NFTokenMinter
	}
	if s.SetFlag != 0 {
		flattened["SetFlag"] = int(s.SetFlag)
	}
	if s.TransferRate != 0 {
		flattened["TransferRate"] = int(s.TransferRate)
	}
	if s.TickSize != 0 {
		flattened["TickSize"] = s.TickSize
	}
	if s.WalletLocator != "" {
		flattened["WalletLocator"] = s.WalletLocator.String()
	}
	if s.WalletSize != 0 {
		flattened["WalletSize"] = int(s.WalletSize)
	}

	return flattened
}

// SetRequireDestTag sets the require destination tag flag.
func (s *AccountSet) SetRequireDestTag() {
	s.Flags |= tfRequireDestTag
}

// SetRequireAuth sets the require auth flag.
func (s *AccountSet) SetRequireAuth() {
	s.Flags |= tfRequireAuth
}

// SetDisallowXRP sets the disallow XRP flag.
func (s *AccountSet) SetDisallowXRP() {
	s.Flags |= tfDisallowXRP
}

// SetOptionalDestTag sets the optional destination tag flag.
func (s *AccountSet) SetOptionalDestTag() {
	s.Flags |= tfOptionalDestTag
}

// SetOptionalAuth sets the optional auth flag.
func (s *AccountSet) SetOptionalAuth() {
	s.Flags |= tfOptionalAuth
}

// SetAllowXRP sets the allow XRP flag.
func (s *AccountSet) SetAllowXRP() {
	s.Flags |= tfAllowXRP
}

// SetAsfRequireDest sets the require destination tag flag.
func (s *AccountSet) SetAsfRequireDest() {
	s.SetFlag = asfRequireDest
}

// ClearAsfRequireDest clears the require destination tag flag.
func (s *AccountSet) ClearAsfRequireDest() {
	s.ClearFlag = asfRequireDest
}

// SetAsfRequireAuth sets the require authorization flag.
func (s *AccountSet) SetAsfRequireAuth() {
	s.SetFlag = asfRequireAuth
}

// ClearAsfRequireAuth clears the require authorization flag.
func (s *AccountSet) ClearAsfRequireAuth() {
	s.ClearFlag = asfRequireAuth
}

// SetAsfDisallowXRP sets the disallow XRP flag.
func (s *AccountSet) SetAsfDisallowXRP() {
	s.SetFlag = asfDisallowXRP
}

// ClearAsfDisallowXRP clears the disallow XRP flag.
func (s *AccountSet) ClearAsfDisallowXRP() {
	s.ClearFlag = asfDisallowXRP
}

// SetAsfDisableMaster sets the disable master key flag.
func (s *AccountSet) SetAsfDisableMaster() {
	s.SetFlag = asfDisableMaster
}

// ClearAsfDisableMaster clears the disable master key flag.
func (s *AccountSet) ClearAsfDisableMaster() {
	s.ClearFlag = asfDisableMaster
}

// SetAsfAccountTxnID sets the account transaction ID flag.
func (s *AccountSet) SetAsfAccountTxnID() {
	s.SetFlag = asfAccountTxnID
}

// ClearAsfAccountTxnID clears the account transaction ID flag.
func (s *AccountSet) ClearAsfAccountTxnID() {
	s.ClearFlag = asfAccountTxnID
}

// SetAsfNoFreeze sets the no freeze flag.
func (s *AccountSet) SetAsfNoFreeze() {
	s.SetFlag = asfNoFreeze
}

// ClearAsfNoFreeze clears the no freeze flag.
func (s *AccountSet) ClearAsfNoFreeze() {
	s.ClearFlag = asfNoFreeze
}

// SetAsfGlobalFreeze sets the global freeze flag.
func (s *AccountSet) SetAsfGlobalFreeze() {
	s.SetFlag = asfGlobalFreeze
}

// ClearAsfGlobalFreeze clears the global freeze flag.
func (s *AccountSet) ClearAsfGlobalFreeze() {
	s.ClearFlag = asfGlobalFreeze
}

// SetAsfDefaultRipple sets the default ripple flag.
func (s *AccountSet) SetAsfDefaultRipple() {
	s.SetFlag = asfDefaultRipple
}

// ClearAsfDefaultRipple clears the default ripple flag.
func (s *AccountSet) ClearAsfDefaultRipple() {
	s.ClearFlag = asfDefaultRipple
}

// SetAsfDepositAuth sets the deposit authorization flag.
func (s *AccountSet) SetAsfDepositAuth() {
	s.SetFlag = asfDepositAuth
}

// ClearAsfDepositAuth clears the deposit authorization flag.
func (s *AccountSet) ClearAsfDepositAuth() {
	s.ClearFlag = asfDepositAuth
}

// SetAsfAuthorizedNFTokenMinter sets the authorized NFToken minter flag.
func (s *AccountSet) SetAsfAuthorizedNFTokenMinter() {
	s.SetFlag = asfAuthorizedNFTokenMinter
}

// ClearAsfAuthorizedNFTokenMinter clears the authorized NFToken minter flag.
func (s *AccountSet) ClearAsfAuthorizedNFTokenMinter() {
	s.ClearFlag = asfAuthorizedNFTokenMinter
}

// SetAsfDisallowIncomingNFTokenOffer sets the disallow incoming NFToken offer flag.
func (s *AccountSet) SetAsfDisallowIncomingNFTokenOffer() {
	s.SetFlag = asfDisallowIncomingNFTokenOffer
}

// ClearAsfDisallowIncomingNFTokenOffer clears the disallow incoming NFToken offer flag.
func (s *AccountSet) ClearAsfDisallowIncomingNFTokenOffer() {
	s.ClearFlag = asfDisallowIncomingNFTokenOffer
}

// SetAsfDisallowIncomingCheck sets the disallow incoming check flag.
func (s *AccountSet) SetAsfDisallowIncomingCheck() {
	s.SetFlag = asfDisallowIncomingCheck
}

// ClearAsfDisallowIncomingCheck clears the disallow incoming check flag.
func (s *AccountSet) ClearAsfDisallowIncomingCheck() {
	s.ClearFlag = asfDisallowIncomingCheck
}

// SetAsfDisallowIncomingPayChan sets the disallow incoming payment channel flag.
func (s *AccountSet) SetAsfDisallowIncomingPayChan() {
	s.SetFlag = asfDisallowIncomingPayChan
}

// ClearAsfDisallowIncomingPayChan clears the disallow incoming payment channel flag.
func (s *AccountSet) ClearAsfDisallowIncomingPayChan() {
	s.ClearFlag = asfDisallowIncomingPayChan
}

// SetAsfDisallowIncomingTrustLine sets the disallow incoming trust line flag.
func (s *AccountSet) SetAsfDisallowIncomingTrustLine() {
	s.SetFlag = asfDisallowIncomingTrustLine
}

// ClearAsfDisallowIncomingTrustLine clears the disallow incoming trust line flag.
func (s *AccountSet) ClearAsfDisallowIncomingTrustLine() {
	s.ClearFlag = asfDisallowIncomingTrustLine
}

// SetAsfAllowTrustLineClawback sets the allow trust line clawback flag.
func (s *AccountSet) SetAsfAllowTrustLineClawback() {
	s.SetFlag = asfAllowTrustLineClawback
}

// ClearAsfAllowTrustLineClawback clears the allow trust line clawback flag.
func (s *AccountSet) ClearAsfAllowTrustLineClawback() {
	s.ClearFlag = asfAllowTrustLineClawback
}

const MIN_TICK_SIZE = 3
const MAX_TICK_SIZE = 15

// Validate the AccountSet transaction fields.
func ValidateAccountSet(tx FlatTransaction) error {
	// validate the base transaction
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "ClearFlag", typecheck.IsUint)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "Domain", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "EmailHash", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "MessageKey", typecheck.IsString)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "SetFlat", typecheck.IsUint)
	if err != nil {
		return err
	}

	// check if SetFlag is within the valid range
	if tx["SetFlag"].(uint) < asfRequireDest || tx["SetFlag"].(uint) > asfAllowTrustLineClawback {
		return errors.New("AccountSet: SetFlag must be between asfRequireDest (1) and asfAllowTrustLineClawback (16)")
	}

	err = ValidateOptionalField(tx, "TransferRate", typecheck.IsUint)
	if err != nil {
		return err
	}

	err = ValidateOptionalField(tx, "TickSize", typecheck.IsUint)
	if err != nil {
		return err
	}

	// check if TickSize is within the valid range and different from 0
	if tx["TickSize"].(uint) != 0 && (tx["TickSize"].(uint) < MIN_TICK_SIZE || tx["TickSize"].(uint) > MAX_TICK_SIZE) {
		return fmt.Errorf("AccountSet: TickSize must be between %d and %d", MIN_TICK_SIZE, MAX_TICK_SIZE)
	}

	return nil
}
