package transaction

import (
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestCredentialCreate_TxType(t *testing.T) {
	tx := &CredentialCreate{}
	assert.Equal(t, CredentialCreateTx, tx.TxType())
}

func TestCredentialCreate_Flatten(t *testing.T) {
	s := CredentialCreate{
		BaseTx: BaseTx{
			Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType: CredentialCreateTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
		},
		Subject:        "rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8",
		Expiration:     123456,
		CredentialType: "6D795F63726564656E7469616C",                                   // "my_credential" in hex
		URI:            "687474703A2F2F636F6D70616E792E636F6D2F63726564656E7469616C73", // "http://company.com/credentials" in hex
	}

	flattened := s.Flatten()

	expected := FlatTransaction{
		"Account":         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType": "CredentialCreate",
		"Fee":             "1",
		"Sequence":        uint32(1234),
		"Subject":         "rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8",
		"Expiration":      uint32(123456),
		"CredentialType":  "6D795F63726564656E7469616C",
		"URI":             "687474703A2F2F636F6D70616E792E636F6D2F63726564656E7469616C73",
	}

	if !reflect.DeepEqual(flattened, expected) {
		t.Errorf("Flatten result differs from expected: %v, %v", flattened, expected)
	}
}

func TestCredentialCreate_Validate(t *testing.T) {
	tests := []struct {
		name     string
		input    *CredentialCreate
		expected bool
	}{
		{
			name: "pass - valid CredentialCreate",
			input: &CredentialCreate{
				BaseTx: BaseTx{
					Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TransactionType: "AMMWithdraw",
					Fee:             types.XRPCurrencyAmount(10),
					Flags:           1048576,
					Sequence:        10,
				},
				Subject:        "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
				CredentialType: "6D795F63726564656E7469616C",
				Expiration:     123456,
				URI:            "687474703A2F2F636F6D70616E792E636F6D2F63726564656E7469616C73",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.input.Validate()
			if valid != tt.expected {
				t.Errorf("Expected validation result to be %v, got %v", tt.expected, valid)
			}
			if (err != nil) != !tt.expected {
				t.Errorf("Expected error presence to be %v, got %v", !tt.expected, err != nil)
			}
		})
	}
}
