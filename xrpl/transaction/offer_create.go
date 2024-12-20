package transaction

import (
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

// An OfferCreate transaction places an Offer in the decentralized exchange.
//
// Example:
//
// ```json
//
//	{
//	    "TransactionType": "OfferCreate",
//	    "Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
//	    "Fee": "12",
//	    "Flags": 0,
//	    "LastLedgerSequence": 7108682,
//	    "Sequence": 8,
//	    "TakerGets": "6000000",
//	    "TakerPays": {
//	      "currency": "GKO",
//	      "issuer": "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
//	      "value": "2"
//	    }
//	}
//
// ```
type OfferCreate struct {
	BaseTx
	// (Optional) Time after which the Offer is no longer active, in seconds since the Ripple Epoch.
	Expiration uint32 `json:",omitempty"`
	// (Optional) An Offer to delete first, specified in the same way as OfferCancel.
	OfferSequence uint32 `json:",omitempty"`
	// The amount and type of currency being sold.
	TakerGets types.CurrencyAmount
	// The amount and type of currency being bought.
	TakerPays types.CurrencyAmount
}

// **********************************
// OfferCreate Flags
// **********************************

const (
	// tfPassive indicates that the offer is passive, meaning it does not consume offers that exactly match it, and instead waits to be consumed by an offer that exactly matches it.
	tfPassive uint32 = 65536
	// Treat the Offer as an Immediate or Cancel order. The Offer never creates an Offer object in the ledger: it only trades as much as it can by consuming existing Offers at the time the transaction is processed. If no Offers match, it executes "successfully" without trading anything. In this case, the transaction still uses the result code tesSUCCESS.
	tfImmediateOrCancel uint32 = 131072
	// Treat the offer as a Fill or Kill order. The Offer never creates an Offer object in the ledger, and is canceled if it cannot be fully filled at the time of execution. By default, this means that the owner must receive the full TakerPays amount; if the tfSell flag is enabled, the owner must be able to spend the entire TakerGets amount instead.
	tfFillOrKill uint32 = 262144
	// tfSell indicates that the offer is selling, not buying.
	tfSell uint32 = 524288
)

// tfPassive indicates that the offer is passive, meaning it does not consume offers that exactly match it, and instead waits to be consumed by an offer that exactly matches it.
func (o *OfferCreate) SetPassiveFlag() {
	o.Flags |= tfPassive
}

// Treat the Offer as an Immediate or Cancel order. The Offer never creates an Offer object in the ledger: it only trades as much as it can by consuming existing Offers at the time the transaction is processed. If no Offers match, it executes "successfully" without trading anything. In this case, the transaction still uses the result code tesSUCCESS.
func (o *OfferCreate) SetImmediateOrCancelFlag() {
	o.Flags |= tfImmediateOrCancel
}

// Treat the offer as a Fill or Kill order. The Offer never creates an Offer object in the ledger, and is canceled if it cannot be fully filled at the time of execution. By default, this means that the owner must receive the full TakerPays amount; if the tfSell flag is enabled, the owner must be able to spend the entire TakerGets amount instead.
func (o *OfferCreate) SetFillOrKillFlag() {
	o.Flags |= tfFillOrKill
}

// tfSell indicates that the offer is selling, not buying.
func (o *OfferCreate) SetSellFlag() {
	o.Flags |= tfSell
}

// TxType returns the type of the transaction (OfferCreate).
func (*OfferCreate) TxType() TxType {
	return OfferCreateTx
}

// Flatten returns a map of the OfferCreate transaction fields.
func (o *OfferCreate) Flatten() FlatTransaction {
	flattened := o.BaseTx.Flatten()

	if o.Expiration != 0 {
		flattened["Expiration"] = o.Expiration
	}
	if o.OfferSequence != 0 {
		flattened["OfferSequence"] = o.OfferSequence
	}
	flattened["TakerGets"] = o.TakerGets.Flatten()
	flattened["TakerPays"] = o.TakerPays.Flatten()

	return flattened
}

// Validates the OfferCreate transaction.
func (o *OfferCreate) Validate() (bool, error) {
	_, err := o.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if ok, err := IsAmount(o.TakerGets, "TakerGets", true); !ok {
		return false, err
	}

	if ok, err := IsAmount(o.TakerPays, "TakerPays", true); !ok {
		return false, err
	}

	return true, nil
}
