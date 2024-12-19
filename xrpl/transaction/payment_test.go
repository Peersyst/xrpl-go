package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestPayment_TxType(t *testing.T) {
	tx := &Payment{}
	assert.Equal(t, PaymentTx, tx.TxType())
}

func TestPaymentFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*Payment)
		expected uint32
	}{
		{
			name: "pass - SetRippleNotDirectFlag",
			setter: func(p *Payment) {
				p.SetRippleNotDirectFlag()
			},
			expected: tfRippleNotDirect,
		},
		{
			name: "pass - SetPartialPaymentFlag",
			setter: func(p *Payment) {
				p.SetPartialPaymentFlag()
			},
			expected: tfPartialPayment,
		},
		{
			name: "pass - SetLimitQualityFlag",
			setter: func(p *Payment) {
				p.SetLimitQualityFlag()
			},
			expected: tfLimitQuality,
		},
		{
			name: "pass - SetRippleNotDirectFlag and SetPartialPaymentFlag",
			setter: func(p *Payment) {
				p.SetRippleNotDirectFlag()
				p.SetPartialPaymentFlag()
			},
			expected: tfRippleNotDirect | tfPartialPayment,
		},
		{
			name: "pass - SetRippleNotDirectFlag and SetLimitQualityFlag",
			setter: func(p *Payment) {
				p.SetRippleNotDirectFlag()
				p.SetLimitQualityFlag()
			},
			expected: tfRippleNotDirect | tfLimitQuality,
		},
		{
			name: "pass - SetPartialPaymentFlag and SetLimitQualityFlag",
			setter: func(p *Payment) {
				p.SetPartialPaymentFlag()
				p.SetLimitQualityFlag()
			},
			expected: tfPartialPayment | tfLimitQuality,
		},
		{
			name: "pass - all flags",
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
func TestPayment_Validate(t *testing.T) {
	tests := []struct {
		name        string
		payment     Payment
		wantValid   bool
		wantErr     bool
		expectedErr error
	}{
		{
			name: "pass - valid Payment",
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
				Destination:    "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
				DestinationTag: 123,
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "fail - invalid BaseTx Payment, missing TransactionType",
			payment: Payment{
				BaseTx: BaseTx{
					Account: "rJwjoukM94WwKwxM428V7b9npHjpkSvif",
					Fee:     types.XRPCurrencyAmount(1000),
					Flags:   262144,
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					Currency: "USD",
					Value:    "1",
				},
				Destination: "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidTransactionType,
		},
		{
			name: "fail - missing Amount",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           262144,
				},
				Destination: "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrMissingAmount("Amount"),
		},
		{
			name: "fail - invalid Amount",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           262144,
				},
				Destination: "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "invalid",
					Currency: "USD",
					Value:    "1",
				},
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidIssuer,
		},
		{
			name: "fail - invalid SendMax Issuer",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           262144,
				},
				Destination: "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "rNqvZ6vEQ5b8PuhfarQ1aViCEqAWr2JkZ",
					Currency: "USD",
					Value:    "1",
				},
				SendMax: types.IssuedCurrencyAmount{
					Issuer:   "invalid",
					Currency: "USD",
					Value:    "1",
				},
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidIssuer,
		},
		{
			name: "fail - invalid DeliverMax Value",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           262144,
				},
				Destination: "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "rNqvZ6vEQ5b8PuhfarQ1aViCEqAWr2JkZ",
					Currency: "USD",
					Value:    "1",
				},
				DeliverMax: types.IssuedCurrencyAmount{
					Issuer:   "rNqvZ6vEQ5b8PuhfarQ1aViCEqAWr2JkZ",
					Currency: "USD",
					Value:    "invalid",
				},
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidTokenValue,
		},
		{
			name: "fail - invalid DeliverMin Currency",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           262144,
				},
				Destination: "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "rNqvZ6vEQ5b8PuhfarQ1aViCEqAWr2JkZ",
					Currency: "USD",
					Value:    "1",
				},
				DeliverMin: types.IssuedCurrencyAmount{
					Issuer:   "rNqvZ6vEQ5b8PuhfarQ1aViCEqAWr2JkZ",
					Currency: "XRP", // must not be XRP
					Value:    "1",
				},
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidTokenCurrency,
		},
		{
			name: "fail - invalid Destination",
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
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidDestination,
		},
		{
			name: "fail - invalid Paths, both account and currency",
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
				Destination: "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
				Paths: [][]PathStep{
					{
						{Account: "invalid", Currency: "USD"}, // can't have both account and currency
					},
				},
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidPathStepCombination,
		},
		{
			name: "fail - invalid Paths, both Issuer and currency set to XRP",
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
				Destination: "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
				Paths: [][]PathStep{
					{
						{Issuer: "rLs9Pa3CwsoJTnXf4RzzbGsnD9GeCPAUpj", Currency: "XRP"}, // can't have both Issuer and currency set to XRP
					},
				},
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidPathStepCombination,
		},
		{
			name: "fail - invalid Paths, empty array",
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
				Destination: "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
				Paths:       [][]PathStep{},
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrEmptyPath,
		},
		{
			name: "pass - valid Partial Payment",
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
				Destination: "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
				DeliverMin: types.IssuedCurrencyAmount{
					Issuer:   "r3EeETxLb1JwmN2xWuZZdKrrEkqw7qgeYf",
					Currency: "USD",
					Value:    "0.5",
				},
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "fail - invalid Partial Payment without Flag",
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
				Destination: "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
				DeliverMin: types.IssuedCurrencyAmount{
					Issuer:   "ra2ASKcVifxurMgUpTnb59mGDAf7JSVyzh",
					Currency: "USD",
					Value:    "0.5",
				},
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrPartialPaymentFlagRequired,
		},
		{
			name: "fail - invalid Partial Payment with another Flag than PartialPayment",
			payment: Payment{
				BaseTx: BaseTx{
					Account:         "r3EeETxLb1JwmN2xWuZZdKrrEkqw7qgeYf",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           65536,
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "ra2ASKcVifxurMgUpTnb59mGDAf7JSVyzh",
					Currency: "USD",
					Value:    "1",
				},
				Destination: "rDgHn3T2P7eNAaoHh43iRudhAUjAHmDgEP",
				DeliverMin: types.IssuedCurrencyAmount{
					Issuer:   "ra2ASKcVifxurMgUpTnb59mGDAf7JSVyzh",
					Currency: "USD",
					Value:    "0.5",
				},
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrPartialPaymentFlagRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.payment.Validate()
			if valid != tt.wantValid {
				t.Errorf("expected valid to be %v, got %v", tt.wantValid, valid)
			}
			if (err != nil) && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Validate() got error message: %v, want error message: %v", err, tt.expectedErr)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error presence to be %v, got %v, err: %s", tt.wantErr, err != nil, err)
			}
		})
	}
}

func TestPayment_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		payment  *Payment
		expected string
	}{
		{
			name: "pass - flatten with all fields",
			payment: &Payment{
				BaseTx: BaseTx{
					Account:         "rJwjoukM94WwKwxM428V7b9npHjpkSvif",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
					Flags:           tfRippleNotDirect | tfPartialPayment,
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					Currency: "USD",
					Value:    "1",
				},
				DeliverMax: types.IssuedCurrencyAmount{
					Issuer:   "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					Currency: "USD",
					Value:    "2",
				},
				DeliverMin: types.IssuedCurrencyAmount{
					Issuer:   "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					Currency: "USD",
					Value:    "0.5",
				},
				Destination:    "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
				DestinationTag: 12345,
				InvoiceID:      "ABC123",
				Paths: [][]PathStep{
					{
						{
							Account:  "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
							Currency: "USD",
							Issuer:   "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
						},
						{
							Account:  "r4D6ptkGYmpNpUWTtc3MpKcdcEtsonrbVf",
							Currency: "USD",
							Issuer:   "rJwrc4W71kVUNTJX77qGHySRJj7BxSgQqt",
						},
					},
				},
				SendMax: types.IssuedCurrencyAmount{
					Issuer:   "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					Currency: "USD",
					Value:    "3",
				},
			},
			expected: `{
				"Account": "rJwjoukM94WwKwxM428V7b9npHjpkSvif",
				"TransactionType": "Payment",
				"Fee": "1000",
				"Flags": 196608,
				"Amount": {
					"issuer": "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					"currency": "USD",
					"value": "1"
				},
				"DeliverMax": {
					"issuer": "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					"currency": "USD",
					"value": "2"
				},
				"DeliverMin": {
					"issuer": "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					"currency": "USD",
					"value": "0.5"
				},
				"Destination": "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
				"DestinationTag": 12345,
				"InvoiceID": "ABC123",
				"Paths": [
					[
						{
							"account": "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
							"currency": "USD",
							"issuer": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn"
						},
						{
							"account": "r4D6ptkGYmpNpUWTtc3MpKcdcEtsonrbVf",
							"currency": "USD",
							"issuer": "rJwrc4W71kVUNTJX77qGHySRJj7BxSgQqt"
						}
					]
				],
				"SendMax": {
					"issuer": "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					"currency": "USD",
					"value": "3"
				}
			}`,
		},
		{
			name: "pass - flatten with minimal fields",
			payment: &Payment{
				BaseTx: BaseTx{
					Account:         "rJwjoukM94WwKwxM428V7b9npHjpkSvif",
					TransactionType: PaymentTx,
					Fee:             types.XRPCurrencyAmount(1000),
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
					Currency: "USD",
					Value:    "1",
				},
				Destination: "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn",
			},
			expected: `{
				"Account": "rJwjoukM94WwKwxM428V7b9npHjpkSvif",
				"TransactionType": "Payment",
				"Fee": "1000",
				"Amount": {
					"issuer": "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn", 
					"currency": "USD", 
					"value": "1"
				},
				"Destination": "r3dFAtNXwRFCyBGz5BcWhMj9a4cm7qkzzn"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.payment.Flatten(), []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}
