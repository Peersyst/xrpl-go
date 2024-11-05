package transaction

import (
	"errors"
	"fmt"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// TODO: Refactor to use a single interface for all transaction types
type Tx interface {
	TxType() TxType
}

type TxHash string

func (*TxHash) TxType() TxType {
	return HashedTx
}

type Binary struct {
	TxBlob string `json:"tx_blob"`
}

func (tx *Binary) TxType() TxType {
	return BinaryTx
}

type BaseTx struct {
	// The unique address of the transaction sender.
	Account types.Address
	//
	// The type of transaction. Valid types include: `Payment`, `OfferCreate`,
	// `TrustSet`, and many others.
	//
	TransactionType TxType
	//
	// Integer amount of XRP, in drops, to be destroyed as a cost for
	// distributing this transaction to the network. Some transaction types have
	// different minimum requirements.
	//
	Fee types.XRPCurrencyAmount `json:",omitempty"`
	//
	// The sequence number of the account sending the transaction. A transaction
	// is only valid if the Sequence number is exactly 1 greater than the previous
	// transaction from the same account. The special case 0 means the transaction
	// is using a Ticket instead.
	//
	Sequence uint `json:",omitempty"`
	//
	// Hash value identifying another transaction. If provided, this transaction
	// is only valid if the sending account's previously-sent transaction matches
	// the provided hash.
	//
	AccountTxnID types.Hash256 `json:",omitempty"`
	// Set of bit-flags for this transaction.
	Flags uint `json:",omitempty"`
	//
	// Highest ledger index this transaction can appear in. Specifying this field
	// places a strict upper limit on how long the transaction can wait to be
	// validated or rejected.
	//
	LastLedgerSequence uint `json:",omitempty"`
	//
	// Additional arbitrary information used to identify this transaction.
	//
	Memos []MemoWrapper `json:",omitempty"`
	// The network id of the transaction.
	NetworkID uint `json:",omitempty"`
	//
	// Array of objects that represent a multi-signature which authorizes this
	// transaction.
	//
	Signers []Signer `json:",omitempty"`
	//
	// Arbitrary integer used to identify the reason for this payment, or a sender
	// on whose behalf this transaction is made. Conventionally, a refund should
	// specify the initial payment's SourceTag as the refund payment's
	// DestinationTag.
	//
	SourceTag uint `json:",omitempty"`
	//
	// Hex representation of the public key that corresponds to the private key
	// used to sign this transaction. If an empty string, indicates a
	// multi-signature is present in the Signers field instead.
	//
	SigningPubKey string `json:",omitempty"`
	//
	// The sequence number of the ticket to use in place of a Sequence number. If
	// this is provided, Sequence must be 0. Cannot be used with AccountTxnID.
	//
	TicketSequence uint `json:",omitempty"`
	//
	// The signature that verifies this transaction as originating from the
	// account it says it is from.
	//
	TxnSignature string `json:",omitempty"`
}

func (tx *BaseTx) TxType() TxType {
	return tx.TransactionType
}

func (tx *BaseTx) Flatten() FlatTransaction {
	flattened := make(FlatTransaction)

	if tx.Account != "" {
		flattened["Account"] = tx.Account.String()
	}
	if tx.TransactionType != "" {
		flattened["TransactionType"] = tx.TransactionType.String()
	}
	if tx.Fee != 0 {
		flattened["Fee"] = tx.Fee.String()
	}
	if tx.Sequence != 0 {
		flattened["Sequence"] = tx.Sequence
	}
	if tx.AccountTxnID != "" {
		flattened["AccountTxnID"] = tx.AccountTxnID.String()
	}
	if tx.Flags != 0 {
		flattened["Flags"] = int(tx.Flags)
	}
	if tx.LastLedgerSequence != 0 {
		flattened["LastLedgerSequence"] = tx.LastLedgerSequence
	}
	if len(tx.Memos) > 0 {
		flattenedMemos := make([]any, 0)
		for _, memo := range tx.Memos {
			flattenedMemo := memo.Flatten()
			if flattenedMemo != nil {
				flattenedMemos = append(flattenedMemos, flattenedMemo)
			}
		}
		flattened["Memos"] = flattenedMemos
	}
	if tx.NetworkID != 0 {
		flattened["NetworkId"] = tx.NetworkID
	}
	if len(tx.Signers) > 0 {
		flattenedSigners := make([]FlatSigner, 0)
		for _, signer := range tx.Signers {
			flattenedSigner := signer.Flatten()
			if flattenedSigner != nil {
				flattenedSigners = append(flattenedSigners, flattenedSigner)
			}
		}
		flattened["Signers"] = flattenedSigners
	}
	if tx.SourceTag != 0 {
		flattened["SourceTag"] = tx.SourceTag
	}
	if tx.SigningPubKey != "" {
		flattened["SigningPubKey"] = tx.SigningPubKey
	}
	if tx.TicketSequence != 0 {
		flattened["TicketSequence"] = tx.TicketSequence
	}
	if tx.TxnSignature != "" {
		flattened["TxnSignature"] = tx.TxnSignature
	}

	return flattened
}

func (tx *BaseTx) Validate() (bool, error) {
	flattenTx := tx.Flatten()

	if !addresscodec.IsValidClassicAddress(tx.Account.String()) {
		return false, fmt.Errorf("invalid xrpl address for the Account field")
	}

	if tx.TransactionType == "" {
		return false, fmt.Errorf("transaction type is required")
	}

	if !typecheck.IsStringNumericUint(tx.Fee.String()) {
		return false, errors.New("invalid fee amount, not a uint")
	}

	err := ValidateOptionalField(flattenTx, "Sequence", typecheck.IsUint)
	if err != nil {
		return false, err
	}

	err = ValidateOptionalField(flattenTx, "AccountTxnID", typecheck.IsString)
	if err != nil {
		return false, err
	}

	err = ValidateOptionalField(flattenTx, "LastLedgerSequence", typecheck.IsUint)
	if err != nil {
		return false, err
	}

	err = ValidateOptionalField(flattenTx, "SourceTag", typecheck.IsUint)
	if err != nil {
		return false, err
	}

	err = ValidateOptionalField(flattenTx, "SigningPubKey", typecheck.IsString)
	if err != nil {
		return false, err
	}

	err = ValidateOptionalField(flattenTx, "TicketSequence", typecheck.IsUint)
	if err != nil {
		return false, err
	}

	err = ValidateOptionalField(flattenTx, "TxnSignature", typecheck.IsString)
	if err != nil {
		return false, err
	}

	err = ValidateOptionalField(flattenTx, "NetworkID", typecheck.IsUint)
	if err != nil {
		return false, err
	}

	// memos
	err = validateMemos(tx.Memos)
	if err != nil {
		return false, err
	}

	// signers
	err = validateSigners(tx.Signers)
	if err != nil {
		return false, err
	}

	return true, nil
}
