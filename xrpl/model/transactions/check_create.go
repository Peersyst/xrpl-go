package transactions

import (
	"encoding/json"
	"errors"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type CheckCreate struct {
	BaseTx
	// The unique address of the account that can cash the Check.
	Destination types.Address
	// Maximum amount of source currency the Check is allowed to debit the
	// sender, including transfer fees on non-XRP currencies. The Check can only
	// credit the destination with the same currency (from the same issuer, for
	// non-XRP currencies). For non-XRP amounts, the nested field names MUST be.
	// lower-case.
	SendMax types.CurrencyAmount
	// Arbitrary tag that identifies the reason for the Check, or a hosted.
	// recipient to pay.
	DestinationTag uint `json:",omitempty"`
	// Time after which the Check is no longer valid, in seconds since the Ripple Epoch.
	Expiration uint `json:",omitempty"`
	//  Arbitrary 256-bit hash representing a specific reason or identifier for this Check.
	InvoiceID types.Hash256 `json:",omitempty"`
}

func (*CheckCreate) TxType() TxType {
	return CheckCreateTx
}

// TODO: Implement flatten
func (s *CheckCreate) Flatten() FlatTransaction {
	return nil
}

func (c *CheckCreate) UnmarshalJSON(data []byte) error {
	type ccHelper struct {
		BaseTx
		Destination    types.Address
		SendMax        json.RawMessage
		DestinationTag uint          `json:",omitempty"`
		Expiration     uint          `json:",omitempty"`
		InvoiceID      types.Hash256 `json:",omitempty"`
	}
	var h ccHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*c = CheckCreate{
		BaseTx:         h.BaseTx,
		Destination:    h.Destination,
		DestinationTag: h.DestinationTag,
		Expiration:     h.Expiration,
		InvoiceID:      h.InvoiceID,
	}

	max, err := types.UnmarshalCurrencyAmount(h.SendMax)
	if err != nil {
		return err
	}
	c.SendMax = max

	return nil
}

func ValidateCheckCreate(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	// TODO: Update IsString by IsAccount when that function exists for the Destination check
	ValidateRequiredField(tx, "Destination", typecheck.IsString)
	ValidateOptionalField(tx, "DestinationTag", typecheck.IsUint)

	if !typecheck.IsString(tx["SendMax"]) && IsIssuedCurrency(tx["SendMax"]) {
		return errors.New("CheckCreate: invalid SendMax, must be an issued currency")
	}

	ValidateOptionalField(tx, "Expiration", typecheck.IsUint)
	ValidateOptionalField(tx, "InvoiceID", typecheck.IsString)

	return nil
}
