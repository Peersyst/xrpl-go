package transaction

import (
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

func TestTrustSetFlatten(t *testing.T) {
	s := TrustSet{
		BaseTx: BaseTx{
			Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType:    TrustSetTx,
			Fee:                types.XRPCurrencyAmount(12),
			Flags:              262144,
			Sequence:           12,
			LastLedgerSequence: 8007750,
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Issuer:   "rsP3mgGb2tcYUrxiLFiHJiQXhsziegtwBc",
			Currency: "USD",
			Value:    "100",
		},
	}

	flattened := s.Flatten()

	expected := FlatTransaction{
		"Account":            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType":    "TrustSet",
		"Fee":                "12",
		"Flags":              uint32(262144),
		"Sequence":           uint32(12),
		"LastLedgerSequence": uint32(8007750),
		"LimitAmount": map[string]interface{}{
			"issuer":   "rsP3mgGb2tcYUrxiLFiHJiQXhsziegtwBc",
			"currency": "USD",
			"value":    "100",
		},
	}

	// Existing DeepEqual check
	if !reflect.DeepEqual(flattened, expected) {
		t.Errorf("Flatten result differs from expected: %v, %v", flattened, expected)
	}
}

func TestTrustSetFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*TrustSet)
		expected uint32
	}{
		{
			name: "pass - SetSetAuthFlag",
			setter: func(ts *TrustSet) {
				ts.SetSetAuthFlag()
			},
			expected: tfSetAuth,
		},
		{
			name: "pass - SetSetNoRippleFlag",
			setter: func(ts *TrustSet) {
				ts.SetSetNoRippleFlag()
			},
			expected: tfSetNoRipple,
		},
		{
			name: "pass - SetClearNoRippleFlag",
			setter: func(ts *TrustSet) {
				ts.SetClearNoRippleFlag()
			},
			expected: tfClearNoRipple,
		},
		{
			name: "pass - SetSetfAuthFlag and SetSetNoRippleFlag",
			setter: func(ts *TrustSet) {
				ts.SetSetAuthFlag()
				ts.SetSetNoRippleFlag()
			},
			expected: tfSetAuth | tfSetNoRipple,
		},
		{
			name: "pass - SetSetfAuthFlag and SetClearNoRippleFlag",
			setter: func(ts *TrustSet) {
				ts.SetSetAuthFlag()
				ts.SetClearNoRippleFlag()
			},
			expected: tfSetAuth | tfClearNoRipple,
		},
		{
			name: "pass - All flags",
			setter: func(ts *TrustSet) {
				ts.SetSetAuthFlag()
				ts.SetSetNoRippleFlag()
				ts.SetClearNoRippleFlag()
				ts.SetSetFreezeFlag()
				ts.SetClearFreezeFlag()
			},
			expected: tfSetAuth | tfSetNoRipple | tfClearNoRipple | tfSetFreeze | tfClearFreeze,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TrustSet{}
			tt.setter(ts)
			if ts.Flags != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, ts.Flags)
			}
		})
	}
}
func TestTrustSetValidate(t *testing.T) {
	tests := []struct {
		name     string
		trustSet *TrustSet
		valid    bool
		err      error
	}{
		{
			name: "pass - valid TrustSet",
			trustSet: &TrustSet{
				BaseTx: BaseTx{
					Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType:    TrustSetTx,
					Fee:                types.XRPCurrencyAmount(12),
					Flags:              262144,
					Sequence:           12,
					LastLedgerSequence: 8007750,
				},
				LimitAmount: types.IssuedCurrencyAmount{
					Issuer:   "rsP3mgGb2tcYUrxiLFiHJiQXhsziegtwBc",
					Currency: "USD",
					Value:    "100",
				},
				QualityIn:  100,
				QualityOut: 200,
			},
			valid: true,
		},
		{
			name: "fail - missing LimitAmount",
			trustSet: &TrustSet{
				BaseTx: BaseTx{
					Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType:    TrustSetTx,
					Fee:                types.XRPCurrencyAmount(12),
					Flags:              262144,
					Sequence:           12,
					LastLedgerSequence: 8007750,
				},
				QualityIn:  100,
				QualityOut: 200,
			},
			valid: false,
		},
		{
			name: "fail - invalid LimitAmount",
			trustSet: &TrustSet{
				BaseTx: BaseTx{
					Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType:    TrustSetTx,
					Fee:                types.XRPCurrencyAmount(12),
					Flags:              262144,
					Sequence:           12,
					LastLedgerSequence: 8007750,
				},
				LimitAmount: types.IssuedCurrencyAmount{
					Issuer:   "r123",
					Currency: "USD",
				},
				QualityIn:  100,
				QualityOut: 200,
			},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.trustSet.Validate()
			if valid != tt.valid {
				t.Errorf("Expected valid to be %v, got %v", tt.valid, valid)
			}
			if (err != nil && tt.valid) || (err == nil && !tt.valid) {
				t.Errorf("Got error: %v", err)
			}

		})
	}
}
