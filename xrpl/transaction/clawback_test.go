package transaction

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestClawbackTransaction(t *testing.T) {
	clawbackTx := Clawback{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: ClawbackTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   "def",
			Currency: "USD",
			Value:    "1",
		},
	}

	clawbackJSON := `{
	"Account": "abcdef",
	"TransactionType": "Clawback",
	"Fee": "1",
	"Sequence": 1234,
	"SigningPubKey": "ghijk",
	"TxnSignature": "A1B2C3D4E5F6",
	"Amount": {
		"issuer": "def",
		"currency": "USD",
		"value": "1"
	}
}`

	if err := testutil.SerializeAndDeserialize(t, clawbackTx, clawbackJSON); err != nil {
		t.Error(err)
	}

	tx, err := UnmarshalTx(json.RawMessage(clawbackJSON))
	if err != nil {
		t.Errorf("UnmarshalTx error: %s", err.Error())
	}
	if !reflect.DeepEqual(tx, &clawbackTx) {
		fmt.Println("Expected: ", clawbackTx)
		fmt.Println("Actual: ", tx)
		t.Error("UnmarshalTx result differs from expected")
	}
}

func TestClawbackFlatten(t *testing.T) {
	s := Clawback{
		BaseTx: BaseTx{
			Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType: ClawbackTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			Currency: "USD",
			Value:    "1",
		},
	}

	flattened := s.Flatten()

	expected := FlatTransaction{
		"Account":         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType": "Clawback",
		"Fee":             "1",
		"Sequence":        uint(1234),
		"Amount": map[string]interface{}{
			"issuer":   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			"currency": "USD",
			"value":    "1",
		},
	}

	if !reflect.DeepEqual(flattened, expected) {
		t.Errorf("Flatten result differs from expected: %v, %v", flattened, expected)
	}
}
func TestClawbackValidate(t *testing.T) {
	tests := []struct {
		name       string
		clawback   Clawback
		shouldPass bool
	}{
		{
			name: "Valid Clawback transaction",
			clawback: Clawback{
				BaseTx: BaseTx{
					Account:         "rnLYcEcYw2r3w6BDsFDSScoFmvZXbwa6EQ",
					TransactionType: ClawbackTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "rhsTg7mm7v3oEGrF85n1KdB3JjCk5KPT4M",
					Currency: "USD",
					Value:    "1",
				},
			},
			shouldPass: true,
		},
		{
			name: "Clawback transaction with missing Amount field",
			clawback: Clawback{
				BaseTx: BaseTx{
					Account:         "rnLYcEcYw2r3w6BDsFDSScoFmvZXbwa6EQ",
					TransactionType: ClawbackTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
			},
			shouldPass: false,
		},
		{
			name: "Clawback transaction with invalid Amount",
			clawback: Clawback{
				BaseTx: BaseTx{
					Account:         "rL2ek7KyeTk6NiyJcxYFrfiPcv8PFVoqgR",
					TransactionType: ClawbackTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "rnLYcEcYw2r3w6BDsFDSScoFmvZXbwa6EQ",
					Currency: "USD",
					Value:    "invalid",
				},
			},
			shouldPass: false,
		},
		{
			name: "Clawback transaction with Account same as the issuer",
			clawback: Clawback{
				BaseTx: BaseTx{
					Account:         "rL2ek7KyeTk6NiyJcxYFrfiPcv8PFVoqgR",
					TransactionType: ClawbackTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "rL2ek7KyeTk6NiyJcxYFrfiPcv8PFVoqgR",
					Currency: "USD",
					Value:    "1",
				},
			},
			shouldPass: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.clawback.Validate()
			if tt.shouldPass {
				if err != nil {
					t.Errorf("Validation failed for valid Clawback transaction with error: %s", err.Error())
				}
				if !valid {
					t.Error("Validation should pass for valid Clawback transaction")
				}
			} else {
				if err == nil || valid {
					t.Error("Validation should fail for invalid Clawback transaction")
				}
			}
		})
	}
}
