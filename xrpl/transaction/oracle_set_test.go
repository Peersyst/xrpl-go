package transaction

import (
	"strings"
	"testing"

	ledger "github.com/Peersyst/xrpl-go/v1/xrpl/ledger-entry-types"
	"github.com/stretchr/testify/assert"
)

func TestOracleSet_TxType(t *testing.T) {
	tx := &OracleSet{}
	assert.Equal(t, tx.TxType(), OracleSetTx)
}

func TestOracleSet_Flatten(t *testing.T) {
	testcases := []struct {
		name     string
		tx       *OracleSet
		expected map[string]interface{}
	}{
		{
			name: "pass - empty",
			tx:   &OracleSet{},
			expected: map[string]interface{}{
				"TransactionType":  OracleSetTx,
				"OracleDocumentID": uint32(0),
				"LastUpdatedTime":  uint32(0),
			},
		},
		{
			name: "pass - complete",
			tx: &OracleSet{
				BaseTx: BaseTx{
					Account:            "r9cZA1mTh4KVPD5PXPBGVdqw9XRybCz6z",
					Fee:                1000000,
					Sequence:           1,
					LastLedgerSequence: 3000000,
				},
				OracleDocumentID: 1,
				Provider:         "Chainlink",
				URI:              "https://example.com",
				LastUpdatedTime:  1715702400,
				AssetClass:       "currency",
				PriceDataSeries: []ledger.PriceData{
					{
						BaseAsset:  "XRP",
						QuoteAsset: "USD",
						AssetPrice: 740,
						Scale:      3,
					},
				},
			},
			expected: map[string]interface{}{
				"TransactionType":    OracleSetTx,
				"Account":            "r9cZA1mTh4KVPD5PXPBGVdqw9XRybCz6z",
				"Fee":                "1000000",
				"Sequence":           uint32(1),
				"LastLedgerSequence": uint32(3000000),
				"OracleDocumentID":   uint32(1),
				"Provider":           "Chainlink",
				"URI":                "https://example.com",
				"LastUpdatedTime":    uint32(1715702400),
				"AssetClass":         "currency",
				"PriceDataSeries": []map[string]interface{}{
					{
						"BaseAsset":  "XRP",
						"QuoteAsset": "USD",
						"AssetPrice": uint64(740),
						"Scale":      uint8(3),
					},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			assert.Equal(t, testcase.tx.Flatten(), testcase.expected)
		})
	}
}

func TestOracleSet_Validate(t *testing.T) {
	testcases := []struct {
		name     string
		tx       *OracleSet
		expected error
	}{
		{
			name: "fail - base tx invalid",
			tx: &OracleSet{
				BaseTx: BaseTx{
					TransactionType: OracleSetTx,
				},
			},
			expected: ErrInvalidAccount,
		},
		{
			name: "fail - provider length",
			tx: &OracleSet{
				BaseTx: BaseTx{
					Account:         "r9cZA1mTh4KVPD5PXPBGVdqw9XRybCz6z",
					TransactionType: OracleSetTx,
				},
				Provider: strings.Repeat("a", 257),
			},
			expected: ErrProviderLength,
		},
		{
			name: "fail - price data series items",
			tx: &OracleSet{
				BaseTx: BaseTx{
					Account:         "r9cZA1mTh4KVPD5PXPBGVdqw9XRybCz6z",
					TransactionType: OracleSetTx,
				},
				PriceDataSeries: make([]ledger.PriceData, 100),
			},
			expected: ErrPriceDataSeriesItems,
		},
		{
			name: "fail - price data series item invalid",
			tx: &OracleSet{
				BaseTx: BaseTx{
					Account:         "r9cZA1mTh4KVPD5PXPBGVdqw9XRybCz6z",
					TransactionType: OracleSetTx,
				},
				PriceDataSeries: []ledger.PriceData{
					{
						BaseAsset: "XRP",
					},
				},
			},
			expected: ledger.ErrPriceDataQuoteAsset,
		},
		{
			name: "fail - price data series item scale",
			tx: &OracleSet{
				BaseTx: BaseTx{
					Account:         "r9cZA1mTh4KVPD5PXPBGVdqw9XRybCz6z",
					TransactionType: OracleSetTx,
				},
				PriceDataSeries: []ledger.PriceData{
					{
						BaseAsset:  "XRP",
						QuoteAsset: "USD",
						Scale:      11,
					},
				},
			},
			expected: ledger.ErrPriceDataScale,
		},
		{
			name: "fail - price data series item asset price and scale",
			tx: &OracleSet{
				BaseTx: BaseTx{
					Account:         "r9cZA1mTh4KVPD5PXPBGVdqw9XRybCz6z",
					TransactionType: OracleSetTx,
				},
				PriceDataSeries: []ledger.PriceData{
					{
						BaseAsset:  "XRP",
						QuoteAsset: "USD",
						Scale:      10,
					},
				},
			},
			expected: ledger.ErrPriceDataAssetPriceAndScale,
		},
		{
			name: "pass - complete",
			tx: &OracleSet{
				BaseTx: BaseTx{
					Account:         "r9cZA1mTh4KVPD5PXPBGVdqw9XRybCz6z",
					TransactionType: OracleSetTx,
				},
				OracleDocumentID: 1,
				Provider:         "Chainlink",
				URI:              "https://example.com",
				LastUpdatedTime:  1715702400,
				AssetClass:       "currency",
				PriceDataSeries: []ledger.PriceData{
					{
						BaseAsset:  "XRP",
						QuoteAsset: "USD",
						AssetPrice: 740,
						Scale:      3,
					},
				},
			},
			expected: nil,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			ok, err := testcase.tx.Validate()
			assert.Equal(t, ok, testcase.expected == nil)
			assert.Equal(t, err, testcase.expected)
		})
	}
}
