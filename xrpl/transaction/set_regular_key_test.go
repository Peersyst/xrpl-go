package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/stretchr/testify/assert"
)

func TestSetRegularKey_TxType(t *testing.T) {
	entry := &SetRegularKey{}
	assert.Equal(t, SetRegularKeyTx, entry.TxType())
}

func TestSetRegularKey_Flatten(t *testing.T) {
	tests := []struct {
		name       string
		regularKey *SetRegularKey
		want       string
	}{
		{
			name: "Valid SetRegularKey",
			regularKey: &SetRegularKey{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: SetRegularKeyTx,
				},
				RegularKey: "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
			},
			want: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "SetRegularKey",
				"RegularKey":      "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v"
			}`,
		},
		{
			name: "Without RegularKey",
			regularKey: &SetRegularKey{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: SetRegularKeyTx,
				},
			},
			want: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "SetRegularKey"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.regularKey.Flatten(), []byte(tt.want))
			if err != nil {
				t.Error(err)
			}
		})
	}
}
func TestSetRegularKey_Validate(t *testing.T) {
	tests := []struct {
		name       string
		regularKey *SetRegularKey
		wantErr    bool
	}{
		{
			name: "Valid SetRegularKey",
			regularKey: &SetRegularKey{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: SetRegularKeyTx,
				},
				RegularKey: "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
			},
			wantErr: false,
		},
		{
			name: "Invalid SetRegularKey BaseTx",
			regularKey: &SetRegularKey{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				},
				RegularKey: "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
			},
			wantErr: true,
		},
		{
			name: "RegularKey same as Account",
			regularKey: &SetRegularKey{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: SetRegularKeyTx,
				},
				RegularKey: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			},
			wantErr: true,
		},
		{
			name: "Invalid RegularKey address",
			regularKey: &SetRegularKey{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: SetRegularKeyTx,
				},
				RegularKey: "invalidAddress",
			},
			wantErr: true,
		},
		{
			name: "Without RegularKey",
			regularKey: &SetRegularKey{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: SetRegularKeyTx,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.regularKey.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !valid && !tt.wantErr {
				t.Errorf("Validate() = %v, want %v", valid, !tt.wantErr)
			}
		})
	}
}
