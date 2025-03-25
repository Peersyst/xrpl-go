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

	holder := types.Address("rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B")
	account := types.Address("rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh")
	issuer := account

	tests := []struct {
		name     string
		tx       *AMMClawback
		expected string
	}{
		{
			name: "pass - AMMClawback with Amount",
			tx: &AMMClawback{
				BaseTx: BaseTx{
					Account: account,
					TransactionType: AMMClawbackTx,
					
				},
				Holder:holder,
				Asset: ledger.Asset{
					Currency: "FOO",
					Issuer: issuer,
				},
				Asset2: ledger.Asset{
					Currency: "BAR",
					Issuer:"rHtptZx1yHf6Yv43s1RWffM3XnEYv3XhRg",
				},
				Amount: types.IssuedCurrencyAmount{
					Currency: "FOO",
					Issuer:issuer,
					Value:    "1000",
				},
			},
			expected: `{
				"TransactionType": "AMMClawback",
				"Account": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				"Holder": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				"Asset": {"currency": "FOO", "issuer": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh"},
				"Asset2": {"currency": "BAR", "issuer": "rHtptZx1yHf6Yv43s1RWffM3XnEYv3XhRg"},
				"Amount": {"currency": "FOO", "issuer": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh", "value": "1000"}
			}`,
		},
		{
			name: "pass - AMMClawback without Amount",
			tx: &AMMClawback{
				BaseTx: BaseTx{
					Account: account,
					TransactionType: AMMClawbackTx,
				},
				Holder:holder,
				Asset: ledger.Asset{
					Currency: "FOO",
					Issuer:issuer,
				},
				Asset2: ledger.Asset{
					Currency: "BAR",
					Issuer:"rHtptZx1yHf6Yv43s1RWffM3XnEYv3XhRg",
				},
			},
			expected: `{
				"TransactionType": "AMMClawback",
				"Account": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				"Holder": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
				"Asset": {"currency": "FOO", "issuer": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh"},
				"Asset2": {"currency": "BAR", "issuer": "rHtptZx1yHf6Yv43s1RWffM3XnEYv3XhRg"}
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

	holder := types.Address("rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B")
	account := types.Address("rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh")
	issuer := account


	validTx := &AMMClawback{
		BaseTx: BaseTx{
			Account: account,
			TransactionType: AMMClawbackTx,
		},
		Holder:  holder,
		Asset: ledger.Asset{
			Currency: "FOO",
			Issuer:   issuer,
		},
		Asset2: ledger.Asset{
			Currency: "XRP",
		},
		Amount: types.IssuedCurrencyAmount{
			Currency: "FOO",            // must match Asset.Currency
			Issuer:   issuer, // must match Asset.Issuer (here, account)
			Value:    "5",
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
				// Create a copy and set Amount to its zero value.
				tx := *validTx
				tx.Amount = types.IssuedCurrencyAmount{}
				return &tx
			}(),
			wantErr: false,
		},
		{
			name: "fail - Holder equals Asset.Issuer",
			tx: func() *AMMClawback {
				tx := *validTx
				// Set Holder equal to Asset.Issuer.
				tx.Holder = tx.Asset.Issuer
				return &tx
			}(),
			wantErr:    true,
			errMessage: ErrHolderMatchesIssuer,
		},
		{
			name: "fail - Account mismatch (Account != Asset.Issuer)",
			tx: func() *AMMClawback {
				tx := *validTx
				tx.Account = "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd"
				return &tx
			}(),
			wantErr:    true,
			errMessage: ErrAccountMismatch,
		},
		{
			name: "fail - Amount currency mismatch",
			tx: func() *AMMClawback {
				tx := *validTx
				tx.Amount.Currency = "DIFF"
				return &tx
			}(),
			wantErr:    true,
			errMessage: ErrAmountCurrencyMismatch,
		},
		{
			name: "fail - Amount issuer mismatch",
			tx: func() *AMMClawback {
				tx := *validTx
				tx.Amount.Issuer = "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd"
				return &tx
			}(),
			wantErr:    true,
			errMessage: ErrAmountIssuerMismatch,
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