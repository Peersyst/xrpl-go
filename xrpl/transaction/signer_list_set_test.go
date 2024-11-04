package transaction

import (
	"testing"

	ledger "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestSignerListSet_TxType(t *testing.T) {
	entry := &SignerListSet{}
	assert.Equal(t, SignerListSetTx, entry.TxType())
}

func TestSignerListSet_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		entry    *SignerListSet
		expected string
	}{
		{
			name: "With SignerEntries",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					Fee:     types.XRPCurrencyAmount(12),
				},
				SignerQuorum: 3,
				SignerEntries: []ledger.SignerEntryWrapper{
					{
						SignerEntry: ledger.SignerEntry{
							Account:      "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
							SignerWeight: 2,
						},
					},
					{
						SignerEntry: ledger.SignerEntry{
							Account:      "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
							SignerWeight: 1,
						},
					},
					{
						SignerEntry: ledger.SignerEntry{
							Account:      "raKEEVSGnKSD9Zyvxu4z6Pqpm4ABH8FS6n",
							SignerWeight: 1,
						},
					},
				},
			},
			expected: `{
				"TransactionType": "SignerListSet",
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Fee": "12",
				"SignerQuorum": 3,
				"SignerEntries": [
					{
						"SignerEntry": {
							"Account": "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
							"SignerWeight": 2
						}
					},
					{
						"SignerEntry": {
							"Account": "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
							"SignerWeight": 1
						}
					},
					{
						"SignerEntry": {
							"Account": "raKEEVSGnKSD9Zyvxu4z6Pqpm4ABH8FS6n",
							"SignerWeight": 1
						}
					}
				]
			}`,
		},
		{
			name: "Without SignerEntries",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					Fee:     types.XRPCurrencyAmount(12),
				},
				SignerQuorum: 0,
			},
			expected: `{
				"TransactionType": "SignerListSet",
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Fee": "12",
				"SignerQuorum": 0
			}`,
		},
		{
			name: "Without SignerEntries and SignerQuorum",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					Fee:     types.XRPCurrencyAmount(12),
				},
			},
			expected: `{
				"TransactionType": "SignerListSet",
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Fee": "12",
				"SignerQuorum": 0
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
func TestSignerListSet_Validate(t *testing.T) {
	tests := []struct {
		name    string
		entry   *SignerListSet
		wantErr bool
	}{
		{
			name: "Valid SignerListSet",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: SignerListSetTx,
					Fee:             types.XRPCurrencyAmount(12),
				},
				SignerQuorum: 3,
				SignerEntries: []ledger.SignerEntryWrapper{
					{
						SignerEntry: ledger.SignerEntry{
							Account:      "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
							SignerWeight: 2,
						},
					},
					{
						SignerEntry: ledger.SignerEntry{
							Account:      "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
							SignerWeight: 1,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid SignerListSet BaseTx",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					Fee:     types.XRPCurrencyAmount(12),
				},
				SignerQuorum: 3,
				SignerEntries: []ledger.SignerEntryWrapper{
					{
						SignerEntry: ledger.SignerEntry{
							Account:      "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
							SignerWeight: 2,
						},
					},
					{
						SignerEntry: ledger.SignerEntry{
							Account:      "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
							SignerWeight: 1,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid SignerListSet with no SignerEntries and quorum > 0",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: SignerListSetTx,
					Fee:             types.XRPCurrencyAmount(12),
				},
				SignerQuorum: 3,
			},
			wantErr: true,
		},
		{
			name: "Invalid SignerListSet with too many SignerEntries",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: SignerListSetTx,
					Fee:             types.XRPCurrencyAmount(12),
				},
				SignerQuorum: 3,
				SignerEntries: func() []ledger.SignerEntryWrapper {
					entries := make([]ledger.SignerEntryWrapper, 33)
					for i := range entries {
						entries[i] = ledger.SignerEntryWrapper{
							SignerEntry: ledger.SignerEntry{
								Account:      "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
								SignerWeight: 1,
							},
						}
					}
					return entries
				}(),
			},
			wantErr: true,
		},
		{
			name: "Invalid SignerListSet with invalid WalletLocator",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: SignerListSetTx,
					Fee:             types.XRPCurrencyAmount(12),
				},
				SignerQuorum: 3,
				SignerEntries: []ledger.SignerEntryWrapper{
					{
						SignerEntry: ledger.SignerEntry{
							Account:       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
							SignerWeight:  2,
							WalletLocator: "invalid_hex",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid SignerListSet with SignerQuorum greater than sum of SignerWeights",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: SignerListSetTx,
					Fee:             types.XRPCurrencyAmount(12),
				},
				SignerQuorum: 5,
				SignerEntries: []ledger.SignerEntryWrapper{
					{
						SignerEntry: ledger.SignerEntry{
							Account:      "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
							SignerWeight: 2,
						},
					},
					{
						SignerEntry: ledger.SignerEntry{
							Account:      "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
							SignerWeight: 1,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Valid SignerListSet with SignerQuorum 0",
			entry: &SignerListSet{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: SignerListSetTx,
					Fee:             types.XRPCurrencyAmount(12),
				},
				SignerQuorum: 0,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.entry.Validate()
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
