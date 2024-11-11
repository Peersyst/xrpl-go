package transaction

import (
	"encoding/hex"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestTx_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		tx        *BaseTx
		wantValid bool
		wantErr   bool
	}{
		{
			name: "Valid transaction",
			tx: &BaseTx{
				Account:            "rhbi7TGHknHCsRrVYmW57tQHmHjmFgjEpU",
				TransactionType:    PaymentTx,
				Fee:                types.XRPCurrencyAmount(10),
				Sequence:           1,
				AccountTxnID:       "abcdef123456",
				LastLedgerSequence: 100,
				SourceTag:          123,
				SigningPubKey:      "abcdefg",
				TicketSequence:     2,
				TxnSignature:       "xyz123",
				NetworkID:          1,
				Memos: []MemoWrapper{
					{
						Memo: Memo{
							MemoType:   hex.EncodeToString([]byte("text")),
							MemoData:   hex.EncodeToString([]byte("Hello, world!")),
							MemoFormat: hex.EncodeToString([]byte("plain")),
						},
					},
					{
						Memo: Memo{
							MemoType:   hex.EncodeToString([]byte("text")),
							MemoData:   hex.EncodeToString([]byte("Hello, world 2!")),
							MemoFormat: hex.EncodeToString([]byte("plain")),
						},
					},
				},
				Signers: []Signer{
					{
						SignerData{
							Account:       "rDqbKhee18wUCnvjPjZA5Kgpe4zeubLQUC",
							TxnSignature:  "abc123",
							SigningPubKey: "def456",
						},
					},
				},
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "Missing required Account field",
			tx: &BaseTx{
				TransactionType: PaymentTx,
			},
			wantValid: false,
			wantErr:   true,
		},
		{
			name: "Missing required TransactionType field",
			tx: &BaseTx{
				Account: "rhbi7TGHknHCsRrVYmW57tQHmHjmFgjEpU",
			},
			wantValid: false,
			wantErr:   true,
		},
		{
			name: "Invalid memos",
			tx: &BaseTx{
				Account:            "rhbi7TGHknHCsRrVYmW57tQHmHjmFgjEpU",
				TransactionType:    PaymentTx,
				Fee:                types.XRPCurrencyAmount(10),
				Sequence:           1,
				AccountTxnID:       "abcdef123456",
				LastLedgerSequence: 100,
				SourceTag:          123,
				SigningPubKey:      "abcdefg",
				TicketSequence:     2,
				TxnSignature:       "xyz123",
				Memos: []MemoWrapper{
					{
						Memo: Memo{
							MemoType:   "invalid",
							MemoData:   "Hello, world!",
							MemoFormat: "plain",
						},
					},
				},
			},
			wantValid: false,
			wantErr:   true,
		},
		{
			name: "Invalid signers",
			tx: &BaseTx{
				Account:            "rhbi7TGHknHCsRrVYmW57tQHmHjmFgjEpU",
				TransactionType:    PaymentTx,
				Fee:                types.XRPCurrencyAmount(10),
				Sequence:           1,
				AccountTxnID:       "abcdef123456",
				LastLedgerSequence: 100,
				SourceTag:          123,
				SigningPubKey:      "abcdefg",
				TicketSequence:     2,
				TxnSignature:       "xyz123",
				Signers: []Signer{
					{
						SignerData{
							Account: "rDqbKhee18wUCnvjPjZA5Kgpe4zeubLQUC",
						},
					},
				},
			},
			wantValid: false,
			wantErr:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			valid, err := tc.tx.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if valid != tc.wantValid {
				t.Errorf("Validate() = %v, want %v", valid, !tc.wantErr)
			}
		})
	}
}

func TestBinary_TxType(t *testing.T) {
	tx := &Binary{}
	assert.Equal(t, BinaryTx, tx.TxType())
}

func TestTxHash_TxType(t *testing.T) {
	var tx TxHash = "abcdef123456"
	assert.Equal(t, HashedTx, tx.TxType())
}

func TestBaseTx_TxType(t *testing.T) {
	tx := &BaseTx{
		TransactionType: PaymentTx,
	}
	assert.Equal(t, PaymentTx, tx.TxType())
}

func TestBaseTx_Flatten(t *testing.T) {
	testCases := []struct {
		name     string
		tx       *BaseTx
		expected string
	}{
		{
			name: "All fields populated",
			tx: &BaseTx{
				Account:            "rhbi7TGHknHCsRrVYmW57tQHmHjmFgjEpU",
				TransactionType:    PaymentTx,
				Fee:                types.XRPCurrencyAmount(10),
				Sequence:           1,
				AccountTxnID:       "abcdef123456",
				Flags:              2147483648,
				LastLedgerSequence: 100,
				Memos: []MemoWrapper{
					{
						Memo: Memo{
							MemoType:   hex.EncodeToString([]byte("text")),
							MemoData:   hex.EncodeToString([]byte("Hello, world!")),
							MemoFormat: hex.EncodeToString([]byte("plain")),
						},
					},
				},
				NetworkID:      1,
				Signers:        []Signer{{SignerData{Account: "rDqbKhee18wUCnvjPjZA5Kgpe4zeubLQUC", TxnSignature: "abc123", SigningPubKey: "def456"}}},
				SourceTag:      123,
				SigningPubKey:  "abcdefg",
				TicketSequence: 2,
				TxnSignature:   "xyz123",
			},
			expected: `{
				"Account": "rhbi7TGHknHCsRrVYmW57tQHmHjmFgjEpU",
				"TransactionType": "Payment",
				"Fee": "10",
				"Sequence": 1,
				"AccountTxnID": "abcdef123456",
				"Flags": 2147483648,
				"LastLedgerSequence": 100,
				"Memos": [
					{
						"Memo": {
							"MemoType": "74657874",
							"MemoData": "48656c6c6f2c20776f726c6421",
							"MemoFormat": "706c61696e"
						}
					}
				],
				"NetworkId": 1,
				"Signers": [
					{
						"SignerData": {
							"Account": "rDqbKhee18wUCnvjPjZA5Kgpe4zeubLQUC", 
							"TxnSignature": "abc123", 
							"SigningPubKey": "def456"
						}
					}
				],
				"SourceTag": 123,
				"SigningPubKey": "abcdefg",
				"TicketSequence": 2,
				"TxnSignature": "xyz123"
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tc.tx.Flatten(), []byte(tc.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}
