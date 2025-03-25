package transaction

import (
	"testing"

	ledger "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/require"
)

func TestAMMClawback_TxType(t *testing.T) {
	tx := &AMMClawback{}
	require.Equal(t, AMMClawbackTx, tx.TxType())
}

func TestAMMClawback_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		tx       *AMMClawback
		expected string
	}{
		{
			name: "pass - AMMClawback with Amount",
			tx: &AMMClawback{
				BaseTx: BaseTx{
					TransactionType: AMMClawbackTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        1234,
					SigningPubKey:   "abcdef",
					TxnSignature:    "ABCDEF123456",
				},
				Account: "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
				Holder: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				Asset: ledger.Asset{
					Currency: "FOO",
					Issuer:   "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
				},
				Asset2: ledger.Asset{
					Currency: "BAR",
					Issuer:   "rHtptZx1yHf6Yv43s1RWffM3XnEYv3XhRg",
				},
				Amount: &types.IssuedCurrencyAmount{
					Currency: "FOO",
					Issuer:   "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
					Value:    "1000",
				},
			},
			expected: `{
				"TransactionType": "AMMClawback",
				"Account": "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
				"Holder": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				"Asset": {"currency": "FOO", "issuer": "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL"},
				"Asset2": {"currency": "BAR", "issuer": "rHtptZx1yHf6Yv43s1RWffM3XnEYv3XhRg"},
				"Amount": {"currency": "FOO", "issuer": "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL", "value": "1000"},
				"Fee": "10",
				"Sequence": 1234,
				"SigningPubKey": "abcdef",
				"TxnSignature": "ABCDEF123456"
			}`,
		},
		{
			name: "pass - AMMClawback without Amount",
			tx: &AMMClawback{
				BaseTx: BaseTx{
					TransactionType: AMMClawbackTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        1234,
					SigningPubKey:   "abcdef",
					TxnSignature:    "ABCDEF123456",
				},
				Account: "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
				Holder:  "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				Asset: ledger.Asset{
					Currency: "FOO",
					Issuer:   "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
				},
				Asset2: ledger.Asset{
					Currency: "BAR",
					Issuer:   "rHtptZx1yHf6Yv43s1RWffM3XnEYv3XhRg",
				},
				Amount: nil,
			},
			expected: `{
				"TransactionType": "AMMClawback",
				"Account": "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
				"Holder": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				"Asset": {"currency": "FOO", "issuer": "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL"},
				"Asset2": {"currency": "BAR", "issuer": "rHtptZx1yHf6Yv43s1RWffM3XnEYv3XhRg"},
				"Fee": "10",
				"Sequence": 1234,
				"SigningPubKey": "abcdef",
				"TxnSignature": "ABCDEF123456"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flattened := tt.tx.Flatten()
			if err := testutil.CompareFlattenAndExpected(flattened, []byte(tt.expected)); err != nil {
				t.Error(err)
			}
		})
	}
}


func TestAMMClawback_Validate(t *testing.T) {


	validTx := &AMMClawback{
		BaseTx: BaseTx{
			TransactionType: AMMClawbackTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "abcdef",
			TxnSignature:    "ABCDEF123456",
		},
		Account: types.Address("rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL"),
		Holder:  types.Address("rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"),
		Asset: ledger.Asset{
			Currency: "FOO",
			Issuer:   "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
		},
		Asset2: ledger.Asset{
			Currency: "BAR",
			Issuer:   "rHtptZx1yHf6Yv43s1RWffM3XnEYv3XhRg",
		},
		Amount: &types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   "rPdYxU9dNkbzC5Y2h4jLbVJ3rMRrk7WVRL",
			Value:    "1000",
		},
	}

	tests := []struct {
		name    string
		tx      *AMMClawback
		wantErr bool
		errMessage     error
	}{
		{
			name:    "pass - valid AMMClawback with Amount",
			tx:      validTx,
			wantErr: false,
		},
		{
			name: "pass - valid AMMClawback without Amount",
			tx: func() *AMMClawback {
				tx := validTx
				tx.Amount = nil
				return tx
			}(),
			wantErr: false,
		},
		{
			name: "fail - Holder equals Asset.Issuer",
			tx: func() *AMMClawback {
				tx := validTx
				tx.Holder = tx.Asset.Issuer
				return tx
			}(),
			wantErr: true,
			errMessage:     ErrHolderMatchesIssuer,
		},
		{
			name: "fail - Account mismatch (Account != Asset.Issuer)",
			tx: func() *AMMClawback {
				tx := validTx
				tx.Account = "rDifferentAccountAddress"
				return tx
			}(),
			wantErr: true,
			errMessage:     ErrAccountMismatch,
		},
		{
			name: "fail - Amount currency mismatch",
			tx: func() *AMMClawback {
				tx := validTx
				if tx.Amount != nil {
					tx.Amount.Currency = "DIFF"
				}
				return tx
			}(),
			wantErr: true,
			errMessage:     ErrAmountCurrencyMismatch,
		},
		{
			name: "fail - Amount issuer mismatch",
			tx: func() *AMMClawback {
				tx := validTx
				if tx.Amount != nil {
					tx.Amount.Issuer = "rDifferentAccountAddress"
				}
				return tx
			}(),
			wantErr: true,
			errMessage:     ErrAmountIssuerMismatch,
		},
		{
			name: "fail - invalid Holder address",
			tx: func() *AMMClawback {
				tx := validTx
				tx.Holder = "invalid"
				return tx
			}(),
			wantErr: true,
		},
	}

		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.tx.Validate()
			if tt.wantErr {
				require.Error(t, err)
				require.Equal(t, tt.errMessage, err)
				require.False(t, valid)
			} else {
				require.NoError(t, err)
				require.True(t, valid)
			}
		})
	}
}

func TestAMMClawback_Flags(t *testing.T) {
	tx := &AMMClawback{}

	tx.SetTwoAssetFlag()
	require.Equal(t, uint32(1), tx.Flags)
}