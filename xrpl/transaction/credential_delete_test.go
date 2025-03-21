package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/require"
)

func TestCredentialDelete_TxType(t *testing.T) {
	tx := &CredentialDelete{}
	require.Equal(t, CredentialDeleteTx, tx.TxType())
}

func TestCredentialDelete_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		input    *CredentialDelete
		expected FlatTransaction
	}{
		{
			name: "pass - valid CredentialDelete",
			input: &CredentialDelete{
				BaseTx: BaseTx{
					Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TransactionType: CredentialDeleteTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        10,
				},
				CredentialType: "6D795F63726564656E7469616C",
				Subject:        "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
				Issuer:         "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
			},
			expected: FlatTransaction{
				"Account":         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
				"TransactionType": "CredentialDelete",
				"Fee":             "10",
				"Sequence":        uint32(10),
				"CredentialType":  "6D795F63726564656E7469616C",
				"Subject":         "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
				"Issuer":          "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			flattened := test.input.Flatten()
			require.Equal(t, test.expected, flattened)
		})
	}
}

func TestCredentialDelete_Validate(t *testing.T) {
	tests := []struct {
		name     string
		input    *CredentialDelete
		expected bool
	}{
		{
			name: "pass - valid CredentialDelete",
			input: &CredentialDelete{
				BaseTx: BaseTx{
					Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TransactionType: CredentialDeleteTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        10,
				},
				CredentialType: "6D795F63726564656E7469616C",
				Subject:        "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
				Issuer:         "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
			},
			expected: true,
		},
		{
			name: "fail - invalid CredentialType",
			input: &CredentialDelete{
				BaseTx: BaseTx{
					Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TransactionType: CredentialDeleteTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        10,
				},
				CredentialType: "invalid",
			},
			expected: false,
		},
		{
			name: "fail - invalid Subject",
			input: &CredentialDelete{
				BaseTx: BaseTx{
					Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TransactionType: CredentialDeleteTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        10,
				},
				CredentialType: "6D795F63726564656E7469616C",
				Subject:        "invalid",
				Issuer:         "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
			},
			expected: false,
		},
		{
			name: "fail - invalid Issuer",
			input: &CredentialDelete{
				BaseTx: BaseTx{
					Account:         "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
					TransactionType: CredentialDeleteTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        10,
				},
				CredentialType: "6D795F63726564656E7469616C",
				Subject:        "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
				Issuer:         "invalid",
			},
			expected: false,
		},
		{
			name: "fail - invalid BaseTx",
			input: &CredentialDelete{
				BaseTx: BaseTx{
					Account:         "invalid",
					TransactionType: CredentialDeleteTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        10,
				},
				CredentialType: "6D795F63726564656E7469616C",
				Subject:        "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
				Issuer:         "rJZdUoJnJb5q8tHb9cYfYh5vZg9G6z2v1d",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			valid, err := test.input.Validate()
			require.Equal(t, test.expected, valid)
			if test.expected {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
