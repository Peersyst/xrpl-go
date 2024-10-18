package transaction

import (
	"testing"

	ledger "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestAMMDeposit_TxType(t *testing.T) {
	tx := &AMMDeposit{}
	assert.Equal(t, AMMDepositTx, tx.TxType())
}
func TestAMMDeposit_Flatten(t *testing.T) {
	tx := &AMMDeposit{
		BaseTx: BaseTx{
			Account:  "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
			Fee:      types.XRPCurrencyAmount(10),
			Flags:    1048576,
			Sequence: 7,
		},
		Asset: ledger.Asset{
			Currency: "TST",
			Issuer:   "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd",
		},
		Asset2: ledger.Asset{
			Currency: "XRP",
		},
		Amount: types.IssuedCurrencyAmount{
			Value:    "2.5",
			Currency: "TST",
			Issuer:   "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd",
		},
		Amount2: types.XRPCurrencyAmount(30000000),
		EPrice: types.IssuedCurrencyAmount{
			Value:    "1.5",
			Issuer:   "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd",
			Currency: "TST",
		},
		LPTokenOut: types.IssuedCurrencyAmount{
			Value:    "100",
			Currency: "TST",
			Issuer:   "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd",
		},
		TradingFee: 10,
	}

	flattened := tx.Flatten()

	expected := `{
	"Account":         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
	"Fee":             "10",
	"Flags":           1048576,
	"Sequence":        7,
	"TransactionType": "AMMDeposit",
	"Asset": {
		"currency": "TST",
		"issuer":   "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd"
	},
	"Asset2": {
		"currency": "XRP"
	},
	"Amount": {
		"issuer":   "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd",
		"value":    "2.5",
		"currency": "TST"
	},
	"Amount2": "30000000",
	"EPrice": {
		"value": "1.5",
		"issuer": "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd",
		"currency": "TST"
	},
	"LPTokenOut": {
		"value": "100",
		"currency": "TST",
		"issuer": "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd"
	},
	"TradingFee": 10
}`

	err := testutil.CompareFlattenAndExpected(flattened, []byte(expected))
	if err != nil {
		t.Error(err)
	}
}

func TestAMMDeposit_Flags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*AMMDeposit)
		expected uint
	}{
		{
			name: "SetLPTokentFlag",
			setter: func(a *AMMDeposit) {
				a.SetLPTokentFlag()
			},
			expected: tfLPToken,
		},
		{
			name: "SetSingleAssetFlag",
			setter: func(a *AMMDeposit) {
				a.SetSingleAssetFlag()
			},
			expected: tfSingleAsset,
		},
		{
			name: "SetTwoAssetFlag",
			setter: func(a *AMMDeposit) {
				a.SetTwoAssetFlag()
			},
			expected: tfTwoAsset,
		},
		{
			name: "SetOneAssetLPTokenFlag",
			setter: func(a *AMMDeposit) {
				a.SetOneAssetLPTokenFlag()
			},
			expected: tfOneAssetLPToken,
		},
		{
			name: "SetLimitLPTokenFlag",
			setter: func(a *AMMDeposit) {
				a.SetLimitLPTokenFlag()
			},
			expected: tfLimitLPToken,
		},
		{
			name: "SetTwoAssetIfEmptyFlag",
			setter: func(a *AMMDeposit) {
				a.SetTwoAssetIfEmptyFlag()
			},
			expected: tfTwoAssetIfEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AMMDeposit{}
			tt.setter(a)
			if a.Flags != tt.expected {
				t.Errorf("Expected AMMDeposit Flags to be %d, got %d", tt.expected, a.Flags)
			}
		})
	}
}
