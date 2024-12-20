package types

import (
	"errors"
	"testing"

	addresscodec "github.com/Peersyst/xrpl-go/v1/address-codec"
	"github.com/Peersyst/xrpl-go/v1/binary-codec/types/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestIssue_FromJson(t *testing.T) {
	tt := []struct {
		name        string
		input       string
		expected    []byte
		expectedErr error
	}{
		{
			name:  "Valid Issue",
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
			name:        "Invalid Issue",
			input:       "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p2",
			expected:    nil,
			expectedErr: addresscodec.ErrInvalidClassicAddress,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			issue := &Issue{}
			actual, err := issue.FromJSON(tc.input)
			require.Equal(t, tc.expected, actual)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestIssue_ToJson(t *testing.T) {
	tt := []struct {
		name     string
		input    []byte
		expected string
		opts     []int
		err      error
		setup    func(t *testing.T) (*Issue, *testutil.MockBinaryParser)
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
			setup: func(t *testing.T) (*Issue, *testutil.MockBinaryParser) {
				ctrl := gomock.NewController(t)
				mock := testutil.NewMockBinaryParser(ctrl)
				mock.EXPECT().ReadBytes(20).Return([]byte{
					83, 223, 129, 195, 127, 70,
					21, 146, 66, 247, 202, 145,
					99, 224, 159, 4, 64, 41,
					204, 18,
				}, nil)
				return &Issue{}, mock
			},
		},
		{
			name:     "No length prefix",
			input:    []byte{},
			expected: "",
			opts:     nil,
			err:      ErrNoLengthPrefix,
			setup: func(t *testing.T) (*Issue, *testutil.MockBinaryParser) {
				ctrl := gomock.NewController(t)
				mock := testutil.NewMockBinaryParser(ctrl)
				return &Issue{}, mock
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
			err:      errors.New("errReadBytes"),
			setup: func(t *testing.T) (*Issue, *testutil.MockBinaryParser) {
				ctrl := gomock.NewController(t)
				mock := testutil.NewMockBinaryParser(ctrl)
				mock.EXPECT().ReadBytes(20).Return([]byte{}, errors.New("errReadBytes"))
				return &Issue{}, mock
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			issue, parser := tc.setup(t)
			actual, err := issue.ToJSON(parser, tc.opts...)

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
