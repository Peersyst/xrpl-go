package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestEscrowCreate_TxType(t *testing.T) {
	entry := &EscrowCreate{}
	assert.Equal(t, EscrowCreateTx, entry.TxType())
}

func TestEscrowCreate_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		entry    *EscrowCreate
		expected string
	}{
		{
			name: "pass - all fields set",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				},
				Amount:         types.XRPCurrencyAmount(10000),
				Destination:    "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				CancelAfter:    533257958,
				FinishAfter:    533171558,
				Condition:      "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
				DestinationTag: 23480,
			},
			expected: `{
				"TransactionType": "EscrowCreate",
				"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Amount":          "10000",
				"Destination":     "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"CancelAfter":     533257958,
				"FinishAfter":     533171558,
				"Condition":       "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
				"DestinationTag":  23480
			}`,
		},
		{
			name: "pass - optional fields omitted",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
			},
			expected: `{
				"TransactionType": "EscrowCreate",
				"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Amount":          "10000",
				"Destination":     "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.entry.Flatten(), []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestEscrowCreate_Validate(t *testing.T) {
	tests := []struct {
		name      string
		entry     *EscrowCreate
		wantValid bool
		wantErr   bool
	}{

		{
			name: "fail - invalid transaction with only CancelAfter",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				CancelAfter: 533257958,
			},
			wantValid: false,
			wantErr:   true,
		},
		{
			name: "fail - invalid transaction with only Condition",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				Condition:   "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
			},
			wantValid: false,
			wantErr:   true,
		},
		{
			name: "fail - invalid transaction with no Condition and FinishAfter",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				CancelAfter: 533257958,
			},
			wantValid: false,
			wantErr:   true,
		},
		{
			name: "fail - invalid transaction with invalid destination address",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "invalidAddress",
				CancelAfter: 533257958,
			},
			wantValid: false,
			wantErr:   true,
		},
		{
			name: "fail - invalid BaseTx, missing TransactionType",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				CancelAfter: 533257958,
			},
			wantValid: false,
			wantErr:   true,
		},
		{
			name: "pass - valid transaction - Conditional with expiration",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				CancelAfter: 533257958,
				Condition:   "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "pass - valid transaction - Time based",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				FinishAfter: 533171558,
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "pass - valid transaction - Time based with expiration",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				FinishAfter: 533171558,
				CancelAfter: 533257958,
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "pass - valid transaction - Timed conditional",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				FinishAfter: 533171558,
				Condition:   "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "pass - valid transaction - Timed conditional with Expiration",
			entry: &EscrowCreate{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: EscrowCreateTx,
				},
				Amount:      types.XRPCurrencyAmount(10000),
				Destination: "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				FinishAfter: 533171558,
				Condition:   "A0258020E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855810100",
				CancelAfter: 533257958,
			},
			wantValid: true,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.entry.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("escrowCreate.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if valid != tt.wantValid {
				t.Errorf("escrowCreate.Validate() = %v, want %v", valid, tt.wantValid)
			}
		})
	}
}
