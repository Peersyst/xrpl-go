package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestOfferCreate_TxType(t *testing.T) {
	tx := &OfferCreate{}
	assert.Equal(t, OfferCreateTx, tx.TxType())
}

func TestOfferCreateFlatten(t *testing.T) {
	tests := []struct {
		name     string
		input    OfferCreate
		expected string
	}{
		{
			name: "pass - with Expiration and OfferSequence",
			input: OfferCreate{
				BaseTx: BaseTx{
					Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType:    OfferCreateTx,
					Fee:                types.XRPCurrencyAmount(12),
					Sequence:           8,
					LastLedgerSequence: 7108682,
				},
				Expiration:    6000000,
				OfferSequence: 10,
				TakerGets:     types.XRPCurrencyAmount(6000000),
				TakerPays: types.IssuedCurrencyAmount{
					Issuer:   "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
					Currency: "GKO",
					Value:    "2",
				},
			},
			expected: `{
				"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
				"TransactionType": "OfferCreate",
				"Fee": "12",
				"Sequence": 8,
				"LastLedgerSequence": 7108682,
				"Expiration": 6000000,
				"OfferSequence": 10,
				"TakerGets": "6000000",
				"TakerPays": {
					"issuer": "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
					"currency": "GKO",
					"value": "2"
				}
			}`,
		},
		{
			name: "pass - without Expiration and OfferSequence",
			input: OfferCreate{
				BaseTx: BaseTx{
					Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType:    OfferCreateTx,
					Fee:                types.XRPCurrencyAmount(12),
					Sequence:           8,
					LastLedgerSequence: 7108682,
				},
				TakerGets: types.XRPCurrencyAmount(6000000),
				TakerPays: types.IssuedCurrencyAmount{
					Issuer:   "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
					Currency: "GKO",
					Value:    "2",
				},
			},
			expected: `{
				"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
				"TransactionType": "OfferCreate",
				"Fee": "12",
				"Sequence": 8,
				"LastLedgerSequence": 7108682,
				"TakerGets": "6000000",
				"TakerPays": {
					"issuer": "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
					"currency": "GKO",
					"value": "2"
				}
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Flatten()
			err := testutil.CompareFlattenAndExpected(result, []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestOfferCreate_Validate(t *testing.T) {
	tests := []struct {
		name     string
		input    OfferCreate
		expected bool
	}{
		{
			name: "pass - valid OfferCreate",
			input: OfferCreate{
				BaseTx: BaseTx{
					Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType:    OfferCreateTx,
					Fee:                types.XRPCurrencyAmount(12),
					Sequence:           8,
					LastLedgerSequence: 7108682,
				},
				TakerGets: types.XRPCurrencyAmount(6000000),
				TakerPays: types.IssuedCurrencyAmount{
					Issuer:   "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
					Currency: "GKO",
					Value:    "2",
				},
			},
			expected: true,
		},
		{
			name: "fail - invalid OfferCreate, missing Account",
			input: OfferCreate{
				BaseTx: BaseTx{
					TransactionType:    OfferCreateTx,
					Fee:                types.XRPCurrencyAmount(12),
					Sequence:           8,
					LastLedgerSequence: 7108682,
				},
				TakerGets: types.XRPCurrencyAmount(6000000),
				TakerPays: types.IssuedCurrencyAmount{
					Issuer:   "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
					Currency: "GKO",
					Value:    "2",
				},
			},
			expected: false,
		},
		{
			name: "fail - invalid TakerGets",
			input: OfferCreate{
				BaseTx: BaseTx{
					Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType:    OfferCreateTx,
					Fee:                types.XRPCurrencyAmount(12),
					Sequence:           8,
					LastLedgerSequence: 7108682,
				},
				TakerGets: types.IssuedCurrencyAmount{
					Issuer:   "",
					Currency: "",
					Value:    "",
				},
				TakerPays: types.IssuedCurrencyAmount{
					Issuer:   "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
					Currency: "GKO",
					Value:    "2",
				},
			},
			expected: false,
		},
		{
			name: "fail - invalid TakerPays",
			input: OfferCreate{
				BaseTx: BaseTx{
					Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType:    OfferCreateTx,
					Fee:                types.XRPCurrencyAmount(12),
					Sequence:           8,
					LastLedgerSequence: 7108682,
				},
				TakerGets: types.XRPCurrencyAmount(6000000),
				TakerPays: types.IssuedCurrencyAmount{
					Issuer:   "",
					Currency: "",
					Value:    "",
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.input.Validate()
			if valid != tt.expected {
				t.Errorf("expected %v, got %v, error: %v", tt.expected, valid, err)
			}
		})
	}
}

func TestOfferCreate_Flags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*OfferCreate)
		expected uint32
	}{
		{
			name: "pass - SetPassiveFlag",
			setter: func(a *OfferCreate) {
				a.SetPassiveFlag()
			},
			expected: tfPassive,
		},
		{
			name: "pass - SetImmediateOrCancelFlag",
			setter: func(a *OfferCreate) {
				a.SetImmediateOrCancelFlag()
			},
			expected: tfImmediateOrCancel,
		},
		{
			name: "pass - SetFillOrKillFlag",
			setter: func(a *OfferCreate) {
				a.SetFillOrKillFlag()
			},
			expected: tfFillOrKill,
		},
		{
			name: "pass - SetSellFlag",
			setter: func(a *OfferCreate) {
				a.SetSellFlag()
			},
			expected: tfSell,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OfferCreate{}
			tt.setter(o)
			if o.Flags != tt.expected {
				t.Errorf("Expected OfferCreate Flags to be %d, got %d", tt.expected, o.Flags)
			}
		})
	}
}
