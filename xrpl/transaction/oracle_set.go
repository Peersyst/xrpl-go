package transaction

import (
	"fmt"

	ledger "github.com/Peersyst/xrpl-go/v1/xrpl/ledger-entry-types"
)

const (
	// The maximum number of PriceData objects that can be included in a PriceDataSeries array.
	OracleSetMaxPriceDataSeriesItems int = 10
	OracleSetProviderMaxLength       int = 256
)

var (
	ErrProviderLength       = fmt.Errorf("provider length must be less than %d bytes", OracleSetProviderMaxLength)
	ErrPriceDataSeriesItems = fmt.Errorf("price data series items must be less than %d", OracleSetMaxPriceDataSeriesItems)
)

// Creates a new Oracle ledger entry or updates the fields of an existing one, using the Oracle ID.
//
// The oracle provider must complete these steps before submitting this transaction:
// 1. Create or own the XRPL account in the Owner field and have enough XRP to meet the reserve and transaction fee requirements.
// 2. Publish the XRPL account public key, so it can be used for verification by dApps.
// 3. Publish a registry of available price oracles with their unique OracleDocumentID.
//
// ```json
//
//	{
//	  "TransactionType": "OracleSet",
//	  "Account": "rNZ9m6AP9K7z3EVg6GhPMx36V4QmZKeWds",
//	  "OracleDocumentID": 34,
//	  "Provider": "70726F7669646572",
//	  "LastUpdateTime": 1724871860,
//	  "AssetClass": "63757272656E6379",
//	  "PriceDataSeries": [
//	    {
//	      "PriceData": {
//	        "BaseAsset": "XRP",
//	        "QuoteAsset": "USD",
//	        "AssetPrice": 740,
//	        "Scale": 3
//	      }
//	    }
//	  ]
//	}
//
// ```
type OracleSet struct {
	BaseTx
	// A unique identifier of the price oracle for the Account. It is 0 by default.
	OracleDocumentID uint32
	// The time the data was last updated, in seconds since the UNIX Epoch.
	// It is 0 by default.
	LastUpdatedTime uint32
	// (Variable) An arbitrary value that identifies an oracle provider, such as Chainlink, Band, or DIA. This field is a string, up to 256 ASCII hex encoded characters (0x20-0x7E).
	// This field is required when creating a new Oracle ledger entry, but is optional for updates.
	Provider string `json:",omitempty"`
	// (Optional) An optional Universal Resource Identifier to reference price data off-chain. This field is limited to 256 bytes.
	URI string `json:",omitempty"`
	// (Variable) Describes the type of asset, such as "currency", "commodity", or "index". This field is a string, up to 16 ASCII hex encoded characters (0x20-0x7E).
	// This field is required when creating a new Oracle ledger entry, but is optional for updates.
	AssetClass string `json:",omitempty"`
	// An array of up to 10 PriceData objects, each representing the price information for a token pair. More than five PriceData objects require two owner reserves.
	PriceDataSeries []ledger.PriceData
}

// Returns the type of the transaction.
func (tx *OracleSet) TxType() TxType {
	return OracleSetTx
}

// Returns a flattened transaction.
func (tx *OracleSet) Flatten() map[string]interface{} {
	flattened := tx.BaseTx.Flatten()

	flattened["TransactionType"] = tx.TxType()

	if tx.Account != "" {
		flattened["Account"] = tx.Account.String()
	}

	flattened["OracleDocumentID"] = tx.OracleDocumentID

	if tx.Provider != "" {
		flattened["Provider"] = tx.Provider
	}
	if tx.URI != "" {
		flattened["URI"] = tx.URI
	}

	flattened["LastUpdatedTime"] = tx.LastUpdatedTime

	if tx.AssetClass != "" {
		flattened["AssetClass"] = tx.AssetClass
	}

	if len(tx.PriceDataSeries) > 0 {
		flattenedPriceDataSeries := make([]map[string]interface{}, 0, len(tx.PriceDataSeries))
		for _, priceData := range tx.PriceDataSeries {
			flattenedPriceDataSeries = append(flattenedPriceDataSeries, priceData.Flatten())
		}
		flattened["PriceDataSeries"] = flattenedPriceDataSeries
	}

	return flattened
}

// Validates the transaction.
func (tx *OracleSet) Validate() (bool, error) {
	if ok, err := tx.BaseTx.Validate(); !ok {
		return false, err
	}

	if len([]byte(tx.Provider)) > OracleSetProviderMaxLength {
		return false, ErrProviderLength
	}

	if len(tx.PriceDataSeries) > OracleSetMaxPriceDataSeriesItems {
		return false, ErrPriceDataSeriesItems
	}

	for _, priceData := range tx.PriceDataSeries {
		if err := priceData.Validate(); err != nil {
			return false, err
		}
	}

	return true, nil
}
