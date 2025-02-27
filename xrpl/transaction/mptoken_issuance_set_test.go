package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/require"
)

func TestMPTokenIssuanceSet_TxType(t *testing.T) {
	tx := &MPTokenIssuanceSet{}
	require.Equal(t, MPTokenIssuanceSetTx, tx.TxType())
}

func TestMPTokenIssuanceSet_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		tx       *MPTokenIssuanceSet
		expected FlatTransaction
	}{
		{
			name: "pass - with holder",
			tx: &MPTokenIssuanceSet{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					Flags:   1,
				},
				MPTokenIssuanceID: "00070C4495F14B0E44F78A264E41713C64B5F89242540EE255534400000000000000",
				Holder:            types.Holder("rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn"),
			},
			expected: FlatTransaction{
				"TransactionType":   "MPTokenIssuanceSet",
				"Account":           "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Flags":             uint32(1),
				"MPTokenIssuanceID": "00070C4495F14B0E44F78A264E41713C64B5F89242540EE255534400000000000000",
				"Holder":            "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			},
		},
		{
			name: "pass - without holder",
			tx: &MPTokenIssuanceSet{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					Flags:   1,
				},
				MPTokenIssuanceID: "00070C4495F14B0E44F78A264E41713C64B5F89242540EE255534400000000000000",
			},
			expected: FlatTransaction{
				"TransactionType":   "MPTokenIssuanceSet",
				"Account":           "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"Flags":             uint32(1),
				"MPTokenIssuanceID": "00070C4495F14B0E44F78A264E41713C64B5F89242540EE255534400000000000000",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flattened := tt.tx.Flatten()
			require.Equal(t, tt.expected, flattened)
		})
	}
}

func TestMPTokenIssuanceSet_Validate(t *testing.T) {
	tests := []struct {
		name    string
		tx      *MPTokenIssuanceSet
		wantOk  bool
		wantErr error
	}{
		{
			name: "pass - valid transaction",
			tx: &MPTokenIssuanceSet{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: MPTokenIssuanceSetTx,
				},
				MPTokenIssuanceID: "00070C4495F14B0E44F78A264E41713C64B5F89242540EE255534400000000000000",
				Holder:            types.Holder("rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn"),
			},
			wantOk:  true,
			wantErr: nil,
		},
		{
			name: "fail - invalid holder address",
			tx: &MPTokenIssuanceSet{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: MPTokenIssuanceSetTx,
				},
				MPTokenIssuanceID: "00070C4495F14B0E44F78A264E41713C64B5F89242540EE255534400000000000000",
				Holder:            types.Holder("invalid"),
			},
			wantOk:  false,
			wantErr: ErrInvalidAccount,
		},
		{
			name: "fail - conflicting flags",
			tx: &MPTokenIssuanceSet{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: MPTokenIssuanceSetTx,
					Flags:           tfMPTLock | tfMPTUnlock,
				},
				MPTokenIssuanceID: "00070C4495F14B0E44F78A264E41713C64B5F89242540EE255534400000000000000",
			},
			wantOk:  false,
			wantErr: ErrMPTokenIssuanceSetFlags,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := tt.tx.Validate()
			require.Equal(t, tt.wantOk, ok)
			require.Equal(t, tt.wantErr, err)
		})
	}
}

func TestMPTokenIssuanceSet_Flags(t *testing.T) {
	tests := []struct {
		name     string
		setFlags func(*MPTokenIssuanceSet)
		want     uint32
	}{
		{
			name: "pass - set MPTLock flag",
			setFlags: func(tx *MPTokenIssuanceSet) {
				tx.SetMPTLockFlag()
			},
			want: tfMPTLock,
		},
		{
			name: "pass - set MPTUnlock flag",
			setFlags: func(tx *MPTokenIssuanceSet) {
				tx.SetMPTUnlockFlag()
			},
			want: tfMPTUnlock,
		},
		{
			name: "pass - set both flags",
			setFlags: func(tx *MPTokenIssuanceSet) {
				tx.SetMPTLockFlag()
				tx.SetMPTUnlockFlag()
			},
			want: tfMPTLock | tfMPTUnlock,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &MPTokenIssuanceSet{}
			tt.setFlags(tx)
			require.Equal(t, tt.want, tx.Flags)
		})
	}
}
