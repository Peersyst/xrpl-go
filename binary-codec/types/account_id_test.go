package types

import (
	"errors"
	"testing"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/binary-codec/types/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAccountID_FromJson(t *testing.T) {
	tt := []struct {
		name        string
		input       string
		expected    []byte
		expectedErr error
	}{
		{
			name:  "Valid AccountID",
			input: "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
			expected: []byte{
				83, 223, 129, 195, 127, 70,
				21, 146, 66, 247, 202, 145,
				99, 224, 159, 4, 64, 41,
				204, 18,
			},
			expectedErr: nil,
		},
		{
			name:        "Invalid AccountID",
			input:       "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p2",
			expected:    nil,
			expectedErr: &addresscodec.InvalidClassicAddressError{Input: "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p2"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			accountID := &AccountID{}
			actual, err := accountID.FromJson(tc.input)
			require.Equal(t, tc.expected, actual)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestAccountID_ToJson(t *testing.T) {
	tt := []struct {
		name     string
		input    []byte
		expected string
		opts     []int
		err      error
		setup    func(t *testing.T) (*AccountID, *testutil.MockBinaryParser)
	}{
		{
			name: "Valid AccountID",
			input: []byte{
				83, 223, 129, 195, 127, 70,
				21, 146, 66, 247, 202, 145,
				99, 224, 159, 4, 64, 41,
				204, 18,
			},
			expected: "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
			opts:     []int{20},
			err:      nil,
			setup: func(t *testing.T) (*AccountID, *testutil.MockBinaryParser) {
				ctrl := gomock.NewController(t)
				mock := testutil.NewMockBinaryParser(ctrl)
				mock.EXPECT().ReadBytes(20).Return([]byte{
					83, 223, 129, 195, 127, 70,
					21, 146, 66, 247, 202, 145,
					99, 224, 159, 4, 64, 41,
					204, 18,
				}, nil)
				return &AccountID{}, mock
			},
		},
		{
			name:     "No length prefix",
			input:    []byte{},
			expected: "",
			opts:     nil,
			err:      ErrNoLengthPrefix,
			setup: func(t *testing.T) (*AccountID, *testutil.MockBinaryParser) {
				ctrl := gomock.NewController(t)
				mock := testutil.NewMockBinaryParser(ctrl)
				return &AccountID{}, mock
			},
		},
		{
			name: "ReadBytes error",
			input: []byte{
				83, 223, 129, 195, 127, 70,
				21, 146, 66, 247, 202, 145,
				99, 224, 159, 4, 64, 41,
				204, 18,
			},
			expected: "",
			opts:     []int{20},
			err:      errors.New("read bytes error"),
			setup: func(t *testing.T) (*AccountID, *testutil.MockBinaryParser) {
				ctrl := gomock.NewController(t)
				mock := testutil.NewMockBinaryParser(ctrl)
				mock.EXPECT().ReadBytes(20).Return([]byte{}, errors.New("read bytes error"))
				return &AccountID{}, mock
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			accountID, parser := tc.setup(t)
			actual, err := accountID.ToJson(parser, tc.opts...)

			if tc.err != nil {
				require.Error(t, err)
				require.Equal(t, tc.err, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, actual)
			}
		})
	}
}
