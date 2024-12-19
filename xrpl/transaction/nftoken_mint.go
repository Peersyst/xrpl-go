package transaction

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	"github.com/Peersyst/xrpl-go/v1/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

var (
	// ErrInvalidTransferFee is returned when the transferFee is not between 0 and 50000 inclusive.
	ErrInvalidTransferFee = errors.New("transferFee must be between 0 and 50000 inclusive")
	// ErrInvalidURI is returned when the URI is not a valid hexadecimal string.
	ErrInvalidURI = errors.New("invalid URI, must be a valid hexadecimal string")
	// ErrIssuerAccountConflict is returned when the issuer is the same as the account.
	ErrIssuerAccountConflict = errors.New("issuer cannot be the same as the account")
	// ErrTransferFeeRequiresTransferableFlag is returned when the transferFee is set without the tfTransferable flag.
	ErrTransferFeeRequiresTransferableFlag = errors.New("transferFee can only be set if the tfTransferable flag is enabled")
)

// The NFTokenMint transaction creates a non-fungible token and adds it to the relevant NFTokenPage object of the NFTokenMinter as an NFToken object.
// This transaction is the only opportunity the NFTokenMinter has to specify any token fields that are defined as immutable (for example, the TokenFlags).
//
// Example:
//
// ```json
//
//	{
//		"TransactionType": "NFTokenMint",
//		"Account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
//		"TransferFee": 314,
//		"NFTokenTaxon": 0,
//		"Flags": 8,
//		"Fee": "10",
//		"URI": "697066733A2F2F62616679626569676479727A74357366703775646D37687537367568377932366E6634646675796C71616266336F636C67747179353566627A6469",
//		"Memos": [
//			  {
//				  "Memo": {
//					  "MemoType":
//						"687474703A2F2F6578616D706C652E636F6D2F6D656D6F2F67656E65726963",
//					  "MemoData": "72656E74"
//				  }
//			  }
//		  ]
//	  }
//
// ```
type NFTokenMint struct {
	BaseTx
	// An arbitrary taxon, or shared identifier, for a series or collection of related NFTs. To mint a series of NFTs, give them all the same taxon.
	NFTokenTaxon uint32
	// (Optional) The issuer of the token, if the sender of the account is issuing it on behalf of another account.
	// This field must be omitted if the account sending the transaction is the issuer of the NFToken.
	// If provided, the issuer's AccountRoot object must have the NFTokenMinter field set to the sender of this transaction (this transaction's Account field).
	Issuer types.Address `json:",omitempty"`
	// (Optional) The value specifies the fee charged by the issuer for secondary sales of the NFToken, if such sales are allowed.
	// Valid values for this field are between 0 and 50000 inclusive, allowing transfer rates of between 0.00% and 50.00% in increments of 0.001.
	// If this field is provided, the transaction MUST have the tfTransferable flag enabled.
	TransferFee uint16 `json:",omitempty"`
	// (Optional) Up to 256 bytes of arbitrary data. In JSON, this should be encoded as a string of hexadecimal.
	// You can use the xrpl.convertStringToHex utility to convert a URI to its hexadecimal equivalent.
	// This is intended to be a URI that points to the data or metadata associated with the NFT.
	// The contents could decode to an HTTP or HTTPS URL, an IPFS URI, a magnet link, immediate data encoded as an RFC 2379 "data" URL, or even an issuer-specific encoding.
	// The URI is NOT checked for validity.
	URI types.NFTokenURI `json:",omitempty"`
}

// **********************************
// NFTokenMint Flags
// **********************************

const (
	// Allow the issuer (or an entity authorized by the issuer) to destroy the minted NFToken. (The NFToken's owner can always do so.)
	tfBurnable uint32 = 1
	// The minted NFToken can only be bought or sold for XRP. This can be desirable if the token has a transfer fee and the issuer does not want to receive fees in non-XRP currencies.
	tfOnlyXRP uint32 = 2
	// DEPRECATED Automatically create trust lines from the issuer to hold transfer fees received from transferring the minted NFToken. The fixRemoveNFTokenAutoTrustLine amendment makes it invalid to set this flag.
	tfTrustLine uint32 = 4
	// The minted NFToken can be transferred to others. If this flag is not enabled, the token can still be transferred from or to the issuer, but a transfer to the issuer must be made based on a buy offer from the issuer and not a sell offer from the NFT holder.
	tfTransferable uint32 = 8
)

// Allow the issuer (or an entity authorized by the issuer) to destroy the minted NFToken. (The NFToken's owner can always do so.)
func (n *NFTokenMint) SetBurnableFlag() {
	n.Flags |= tfBurnable
}

// The minted NFToken can only be bought or sold for XRP. This can be desirable if the token has a transfer fee and the issuer does not want to receive fees in non-XRP currencies.
func (n *NFTokenMint) SetOnlyXRPFlag() {
	n.Flags |= tfOnlyXRP
}

// DEPRECATED Automatically create trust lines from the issuer to hold transfer fees received from transferring the minted NFToken. The fixRemoveNFTokenAutoTrustLine amendment makes it invalid to set this flag.
func (n *NFTokenMint) SetTrustlineFlag() {
	n.Flags |= tfTrustLine
}

// The minted NFToken can be transferred to others. If this flag is not enabled, the token can still be transferred from or to the issuer, but a transfer to the issuer must be made based on a buy offer from the issuer and not a sell offer from the NFT holder.
func (n *NFTokenMint) SetTransferableFlag() {
	n.Flags |= tfTransferable
}

// TxType returns the type of the transaction (NFTokenMint).
func (*NFTokenMint) TxType() TxType {
	return NFTokenMintTx
}

// Flatten returns a map of the NFTokenMint transaction fields.
func (n *NFTokenMint) Flatten() FlatTransaction {
	flattened := n.BaseTx.Flatten()

	flattened["TransactionType"] = "NFTokenMint"
	flattened["NFTokenTaxon"] = n.NFTokenTaxon

	if n.Issuer != "" {
		flattened["Issuer"] = n.Issuer
	}

	if n.TransferFee != 0 {
		flattened["TransferFee"] = n.TransferFee
	}

	if n.URI != "" {
		flattened["URI"] = n.URI
	}

	return flattened
}

const (
	// Allowing a transfer fee of up to 50%.
	MaxTransferFee = 50000
)

// Validate checks the validity of the NFTokenMint fields.
func (n *NFTokenMint) Validate() (bool, error) {
	ok, err := n.BaseTx.Validate()
	if err != nil || !ok {
		return false, err
	}

	// check transfer fee is between 0 and 50000
	if n.TransferFee > MaxTransferFee {
		return false, ErrInvalidTransferFee
	}

	// check issuer is not the same as the account
	if n.Issuer == n.Account {
		return false, ErrIssuerAccountConflict
	}

	// check issuer is a valid xrpl address
	if n.Issuer != "" && !addresscodec.IsValidClassicAddress(n.Issuer.String()) {
		return false, ErrInvalidIssuer
	}

	// check URI is a valid hexadecimal string
	if n.URI != "" && !typecheck.IsHex(n.URI.String()) {
		return false, ErrInvalidURI
	}

	// check transfer fee can only be set if the tfTransferable flag is enabled
	if n.TransferFee > 0 && !IsFlagEnabled(n.Flags, tfTransferable) {
		return false, ErrTransferFeeRequiresTransferableFlag
	}

	return true, nil
}
