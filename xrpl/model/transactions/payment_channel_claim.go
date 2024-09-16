package transactions

import (
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

const (
	// Clear the channel's Expiration time. (Expiration is different from the
	// channel's immutable CancelAfter time.) Only the source address of the
	// payment channel can use this flag.
	tfRenew uint = 65536 // 0x00010000
	// Request to close the channel. Only the channel source and destination
	// addresses can use this flag. This flag closes the channel immediately if it
	// has no more XRP allocated to it after processing the current claim, or if
	// the destination address uses it. If the source address uses this flag when
	// the channel still holds XRP, this schedules the channel to close after
	// SettleDelay seconds have passed. (Specifically, this sets the Expiration of
	// the channel to the close time of the previous ledger plus the channel's
	// SettleDelay time, unless the channel already has an earlier Expiration
	// time.) If the destination address uses this flag when the channel still
	// holds XRP, any XRP that remains after processing the claim is returned to
	// the source address.
	tfClose uint = 131072 // 0x00020000
)

// Claim XRP from a payment channel, adjust the payment channel's expiration,
// or both.
type PaymentChannelClaim struct {
	BaseTx
	Channel   types.Hash256
	Balance   types.XRPCurrencyAmount `json:",omitempty"`
	Amount    types.XRPCurrencyAmount `json:",omitempty"`
	Signature string                  `json:",omitempty"`
	PublicKey string                  `json:",omitempty"`
}

// TxType returns the type of the transaction (PaymentChannelClaim).
func (*PaymentChannelClaim) TxType() TxType {
	return PaymentChannelClaimTx
}

// Flatten returns a flattened map of the PaymentChannelClaim transaction.
func (s *PaymentChannelClaim) Flatten() FlatTransaction {
	flattened := s.BaseTx.Flatten()

	flattened["TransactionType"] = "PaymentChannelClaim"

	if s.Channel != "" {
		flattened["Channel"] = s.Channel.String()
	}
	if s.Balance != 0 {
		flattened["Balance"] = s.Balance.Flatten()
	}
	if s.Amount != 0 {
		flattened["Amount"] = s.Amount.Flatten()
	}
	if s.Signature != "" {
		flattened["Signature"] = s.Signature
	}
	if s.PublicKey != "" {
		flattened["PublicKey"] = s.PublicKey
	}
	return flattened
}

// SetRenewFlag sets the Renew flag.
//
// Renew: Clear the channel's Expiration time. (Expiration is different from the
// channel's immutable CancelAfter time.) Only the source address of the
// payment channel can use this flag.
func (s *PaymentChannelClaim) SetRenewFlag() {
	s.Flags |= tfRenew
}

// SetCloseFlag sets the Close flag.
//
// Close: Request to close the channel. Only the channel source and destination
// addresses can use this flag. This flag closes the channel immediately if it
// has no more XRP allocated to it after processing the current claim, or if
// the destination address uses it. If the source address uses this flag when
// the channel still holds XRP, this schedules the channel to close after
// SettleDelay seconds have passed. (Specifically, this sets the Expiration of
// the channel to the close time of the previous ledger plus the channel's
// SettleDelay time, unless the channel already has an earlier Expiration
// time.) If the destination address uses this flag when the channel still
// holds XRP, any XRP that remains after processing the claim is returned to
// the source address.
func (s *PaymentChannelClaim) SetCloseFlag() {
	s.Flags |= tfClose
}

func ValidatePaymentChannelClaim(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	ValidateRequiredField(tx, "Channel", typecheck.IsString)

	ValidateOptionalField(tx, "Balance", typecheck.IsString)
	ValidateOptionalField(tx, "Amount", typecheck.IsString)
	ValidateOptionalField(tx, "Signature", typecheck.IsString)
	ValidateOptionalField(tx, "PublicKey", typecheck.IsString)

	return nil
}
