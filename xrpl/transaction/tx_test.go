package transaction

import (
	"encoding/hex"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		name   string
		tx     *BaseTx
		valid  bool
		errMsg string
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
				NetworkId:          1,
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
			valid: true,
		},
		{
			name: "Missing required Account field",
			tx: &BaseTx{
				TransactionType: PaymentTx,
			},
			valid: false,
		},
		{
			name: "Missing required TransactionType field",
			tx: &BaseTx{
				Account: "rhbi7TGHknHCsRrVYmW57tQHmHjmFgjEpU",
			},
			valid: false,
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
			valid: false,
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
			valid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			valid, err := tc.tx.Validate()
			if valid != tc.valid || (err == nil && !tc.valid) {
				t.Errorf("Test case %s failed: expected valid=%v, errMsg=%s", tc.name, tc.valid, err)
			}
		})
	}
}
