package transaction

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestPaymentTx(t *testing.T) {
	s := Payment{
		BaseTx: BaseTx{
			Account:         "abc",
			TransactionType: PaymentTx,
			Fee:             types.XRPCurrencyAmount(1000),
			Flags:           262144,
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   "def",
			Currency: "USD",
			Value:    "1",
		},
		Destination: "hij",
	}

	j := `{
	"Account": "abc",
	"TransactionType": "Payment",
	"Fee": "1000",
	"Flags": 262144,
	"Amount": {
		"issuer": "def",
		"currency": "USD",
		"value": "1"
	},
	"Destination": "hij"
}`
	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

	tx, err := UnmarshalTx(json.RawMessage(j))
	if err != nil {
		t.Errorf("UnmarshalTx error: %s", err.Error())
	}
	if !reflect.DeepEqual(tx, &s) {
		t.Error("UnmarshalTx result differs from expected")
	}
}

func TestPaymentFlatten(t *testing.T) {
	s := Payment{
		BaseTx: BaseTx{
			Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType: PaymentTx,
			Fee:             types.XRPCurrencyAmount(1000),
			Flags:           262144,
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			Currency: "USD",
			Value:    "1",
		},
		Destination: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		Paths: [][]PathStep{
			{
				{
					Account: "rLs9Pa3CwsoJTnXf4RzzbGsnD9GeCPAUpj",
				},
				{
					Account: "ra8mAnaRqoxijPayDcyneRjcD45Bo2DNnM",
				},
			},
			{
				{
					Currency: "USD",
					Issuer:   "rEFowZFH6y4A6PwuzmAx6cFXsAkr8JiHiS",
				},
			},
		},
	}

	flattened := s.Flatten()

	expected := `{
		"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType": "Payment",
		"Fee": "1000",
		"Flags": 262144,
		"Amount": {
			"issuer": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			"currency": "USD",
			"value": "1"
		},
		"Destination": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"Paths": [
			[
				{
					"account": "rLs9Pa3CwsoJTnXf4RzzbGsnD9GeCPAUpj"
				},
				{
					"account": "ra8mAnaRqoxijPayDcyneRjcD45Bo2DNnM"
				}
			],
			[
				{
					"currency": "USD",
					"issuer": "rEFowZFH6y4A6PwuzmAx6cFXsAkr8JiHiS"
				}
			]
		]
	}`

	// Convert flattened to JSON
	flattenedJSON, err := json.Marshal(flattened)
	if err != nil {
		t.Errorf("Error marshaling payment flattened, error: %v", err)
	}

	// Normalize expected JSON
	var expectedMap map[string]interface{}
	if err := json.Unmarshal([]byte(expected), &expectedMap); err != nil {
		t.Errorf("Error unmarshaling expected, error: %v", err)
	}
	expectedJSON, err := json.Marshal(expectedMap)
	if err != nil {
		t.Errorf("Error marshaling expected payment object: %v", err)
	}

	// Compare JSON strings
	if string(flattenedJSON) != string(expectedJSON) {
		t.Errorf("The flattened and expected Payment JSON are not equal.\nGot: %v\nExpected: %v", string(flattenedJSON), string(expectedJSON))
	}
}

func TestPaymentFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*Payment)
		expected uint
	}{
		{
			name: "SetRippleNotDirectFlag",
			setter: func(p *Payment) {
				p.SetRippleNotDirectFlag()
			},
			expected: tfRippleNotDirect,
		},
		{
			name: "SetPartialPaymentFlag",
			setter: func(p *Payment) {
				p.SetPartialPaymentFlag()
			},
			expected: tfPartialPayment,
		},
		{
			name: "SetLimitQualityFlag",
			setter: func(p *Payment) {
				p.SetLimitQualityFlag()
			},
			expected: tfLimitQuality,
		},
		{
			name: "SetRippleNotDirectFlag and SetPartialPaymentFlag",
			setter: func(p *Payment) {
				p.SetRippleNotDirectFlag()
				p.SetPartialPaymentFlag()
			},
			expected: tfRippleNotDirect | tfPartialPayment,
		},
		{
			name: "SetRippleNotDirectFlag and SetLimitQualityFlag",
			setter: func(p *Payment) {
				p.SetRippleNotDirectFlag()
				p.SetLimitQualityFlag()
			},
			expected: tfRippleNotDirect | tfLimitQuality,
		},
		{
			name: "SetPartialPaymentFlag and SetLimitQualityFlag",
			setter: func(p *Payment) {
				p.SetPartialPaymentFlag()
				p.SetLimitQualityFlag()
			},
			expected: tfPartialPayment | tfLimitQuality,
		},
		{
			name: "All flags",
			setter: func(p *Payment) {
				p.SetRippleNotDirectFlag()
				p.SetPartialPaymentFlag()
				p.SetLimitQualityFlag()
			},
			expected: tfRippleNotDirect | tfPartialPayment | tfLimitQuality,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Payment{}
			tt.setter(p)
			if p.Flags != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, p.Flags)
			}
		})
	}
}
func TestPaymentValidate(t *testing.T) {
	tests := []struct {
		name    string
		payment Payment
		isValid bool
	}{
		{
			name: "Valid Payment",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "rJwjoukM94WwKwxM428V7b9npHjpkSvif",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           262144,
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					Currency: "USD",
					Value:    "1",
				},
				Destination: "hij",
			},
			isValid: true,
		},
		{
			name: "Missing Amount",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           262144,
				},
				Destination: "hij",
			},
			isValid: false,
		},
		{
			name: "Invalid Destination",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           262144,
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "rQLnYrZARjqMhrFhY5Z8Fv1tiRYvHFBXws",
					Currency: "USD",
					Value:    "1",
				},
				Destination: "",
			},
			isValid: false,
		},
		{
			name: "Invalid Paths, both account and currency",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "rQLnYrZARjqMhrFhY5Z8Fv1tiRYvHFBXws",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           262144,
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "rLs9Pa3CwsoJTnXf4RzzbGsnD9GeCPAUpj",
					Currency: "USD",
					Value:    "1",
				},
				Destination: "hij",
				Paths: [][]PathStep{
					{
						{Account: "invalid", Currency: "USD"}, // can't have both account and currency
					},
				},
			},
			isValid: false,
		},
		{
			name: "Invalid Paths, both Issuer and currency set to XRP",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "rQLnYrZARjqMhrFhY5Z8Fv1tiRYvHFBXws",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           262144,
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "rLs9Pa3CwsoJTnXf4RzzbGsnD9GeCPAUpj",
					Currency: "USD",
					Value:    "1",
				},
				Destination: "hij",
				Paths: [][]PathStep{
					{
						{Issuer: "rLs9Pa3CwsoJTnXf4RzzbGsnD9GeCPAUpj", Currency: "XRP"}, // can't have both Issuer and currency set to XRP
					},
				},
			},
			isValid: false,
		},
		{
			name: "Invalid Paths, empty array",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "rQLnYrZARjqMhrFhY5Z8Fv1tiRYvHFBXws",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           262144,
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "rLs9Pa3CwsoJTnXf4RzzbGsnD9GeCPAUpj",
					Currency: "USD",
					Value:    "1",
				},
				Destination: "hij",
				Paths:       [][]PathStep{},
			},
			isValid: false,
		},
		{
			name: "Valid Partial Payment",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "rLs9Pa3CwsoJTnXf4RzzbGsnD9GeCPAUpj",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           tfPartialPayment,
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "r3EeETxLb1JwmN2xWuZZdKrrEkqw7qgeYf",
					Currency: "USD",
					Value:    "1",
				},
				Destination: "hij",
				DeliverMin: types.IssuedCurrencyAmount{
					Issuer:   "r3EeETxLb1JwmN2xWuZZdKrrEkqw7qgeYf",
					Currency: "USD",
					Value:    "0.5",
				},
			},
			isValid: true,
		},
		{
			name: "Invalid Partial Payment without Flag",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "r3EeETxLb1JwmN2xWuZZdKrrEkqw7qgeYf",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           0,
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "ra2ASKcVifxurMgUpTnb59mGDAf7JSVyzh",
					Currency: "USD",
					Value:    "1",
				},
				Destination: "hij",
				DeliverMin: types.IssuedCurrencyAmount{
					Issuer:   "ra2ASKcVifxurMgUpTnb59mGDAf7JSVyzh",
					Currency: "USD",
					Value:    "0.5",
				},
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := tt.payment.Validate()
			if ok != tt.isValid {
				t.Errorf("Expected valid=%v, got %v, with error: %s", tt.isValid, ok, err)
			}
		})
	}
}
