package types

import (
	"errors"
	"testing"

	"github.com/Peersyst/xrpl-go/binary-codec/definitions"
	"github.com/Peersyst/xrpl-go/binary-codec/serdes"
	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
	"github.com/Peersyst/xrpl-go/binary-codec/types/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestPathSet_FromJson(t *testing.T) {
	testcases := []struct {
		name   string
		input  any
		output []byte
		err    error
	}{
		{
			name: "fail - empty path set",
			input: []any{
				map[string]any{},
			},
			err: ErrInvalidPathSet,
		},
		{
			name: "fail - invalid path set",
			input: []any{
				[]any{
					map[string]any{},
				},
			},
			err: ErrInvalidPathSet,
		},
		{
			name: "pass - valid path set",
			input: []any{
				[]any{
					map[string]any{
						"account":  "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
						"currency": "USD",
						"issuer":   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
					},
				},
			},
			output: []byte{0x31, 0xb5, 0xf7, 0x62, 0x79, 0x8a, 0x53, 0xd5, 0x43, 0xa0, 0x14, 0xca, 0xf8, 0xb2, 0x97, 0xcf, 0xf8, 0xf2, 0xf9, 0x37, 0xe8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb5, 0xf7, 0x62, 0x79, 0x8a, 0x53, 0xd5, 0x43, 0xa0, 0x14, 0xca, 0xf8, 0xb2, 0x97, 0xcf, 0xf8, 0xf2, 0xf9, 0x37, 0xe8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			err:    nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			pathset := PathSet{}
			act, err := pathset.FromJSON(tc.input)
			if tc.err != nil {
				require.Error(t, err)
				require.Equal(t, tc.err, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.output, act)
		})
	}
}

func TestPathSet_ToJson(t *testing.T) {
	testcases := []struct {
		name     string
		malleate func(t *testing.T) interfaces.BinaryParser
		output   any
		err      error
	}{
		{
			name: "fail - binary parser peek",
			malleate: func(t *testing.T) interfaces.BinaryParser {
				parser := testutil.NewMockBinaryParser(gomock.NewController(t))
				parser.EXPECT().HasMore().Return(true)
				parser.EXPECT().Peek().Return(uint8(0), errors.New("peek error"))
				return parser
			},
			err: errors.New("peek error"),
		},
		{
			name: "fail - binary parser read byte",
			malleate: func(t *testing.T) interfaces.BinaryParser {
				parser := testutil.NewMockBinaryParser(gomock.NewController(t))
				parser.EXPECT().HasMore().AnyTimes().Return(true)
				parser.EXPECT().Peek().AnyTimes().Return(uint8(pathSeparatorByte), nil)
				parser.EXPECT().ReadByte().Return(uint8(0), errors.New("read byte error"))
				return parser
			},
			err: errors.New("read byte error"),
		},
		{
			name: "pass - valid path set",
			malleate: func(t *testing.T) interfaces.BinaryParser {
				return serdes.NewBinaryParser([]byte{0x31, 0xb5, 0xf7, 0x62, 0x79, 0x8a, 0x53, 0xd5, 0x43, 0xa0, 0x14, 0xca, 0xf8, 0xb2, 0x97, 0xcf, 0xf8, 0xf2, 0xf9, 0x37, 0xe8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb5, 0xf7, 0x62, 0x79, 0x8a, 0x53, 0xd5, 0x43, 0xa0, 0x14, 0xca, 0xf8, 0xb2, 0x97, 0xcf, 0xf8, 0xf2, 0xf9, 0x37, 0xe8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, definitions.Get())
			},
			output: []any{
				[]any{
					map[string]any{
						"account":  "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
						"currency": "USD",
						"issuer":   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
						"type":     16,
						"type_hex": "0000000000000010",
					},
				},
			},
			err: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			pathset := PathSet{}
			act, err := pathset.ToJSON(tc.malleate(t))
			if tc.err != nil {
				require.Error(t, err)
				require.Equal(t, tc.err, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.output, act)
		})
	}
}

func TestIsPathStep(t *testing.T) {

	tt := []struct {
		description string
		input       map[string]any
		expected    bool
	}{
		{
			description: "represents valid path step",
			input: map[string]any{
				"account":  "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				"currency": "USD",
				"issuer":   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
			},
			expected: true,
		},
		{
			description: "represents valid path step",
			input: map[string]any{
				"account":  "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				"currency": "USD",
			},
			expected: true,
		},
		{
			description: "represents invalid path step",
			input:       map[string]any{},
			expected:    false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			require.Equal(t, tc.expected, isPathStep(tc.input))
		})
	}
}

func TestNewPathStep(t *testing.T) {

	tt := []struct {
		description string
		input       map[string]any
		expected    []byte
	}{
		{
			description: "created valid path step",
			input: map[string]any{
				"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
				"currency": "USD",
				"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
			},
			expected: []byte{0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			require.Equal(t, tc.expected, newPathStep(tc.input))
		})
	}
}

func TestNewPath(t *testing.T) {

	tt := []struct {
		description string
		input       []any
		expected    []byte
	}{
		{
			description: "created valid path",
			input: []any{
				map[string]any{
					"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
					"currency": "USD",
					"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
				},
				map[string]any{
					"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
					"currency": "USD",
					"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
				},
			},
			expected: []byte{0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			require.Equal(t, tc.expected, newPath(tc.input))
		})
	}

}

func TestNewPathSet(t *testing.T) {
	tt := []struct {
		description string
		input       []any
		expected    []byte
	}{
		{
			description: "created valid path set with multiple paths",
			input: []any{
				[]any{
					map[string]any{
						"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
						"currency": "USD",
						"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					},
					map[string]any{
						"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
						"currency": "USD",
						"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					},
				},
				[]any{
					map[string]any{
						"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
						"currency": "USD",
						"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					},
					map[string]any{
						"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
						"currency": "USD",
						"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					},
				},
				[]any{
					map[string]any{
						"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
						"currency": "USD",
						"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					},
					map[string]any{
						"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
						"currency": "USD",
						"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					},
				},
			},
			expected: []byte{0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			require.Equal(t, tc.expected, newPathSet(tc.input))
		})
	}
}

func TestParsePathStep(t *testing.T) {
	tt := []struct {
		name        string
		malleate    func(t *testing.T) interfaces.BinaryParser
		expected    map[string]any
		expectedErr error
	}{
		{
			name: "fail - invalid path step",
			malleate: func(t *testing.T) interfaces.BinaryParser {
				parser := testutil.NewMockBinaryParser(gomock.NewController(t))
				parser.EXPECT().ReadByte().Return(uint8(0), errors.New("read byte error"))
				return parser
			},
			expected:    nil,
			expectedErr: errors.New("read byte error"),
		},
		{
			name: "pass - successfully parse path step",
			malleate: func(t *testing.T) interfaces.BinaryParser {
				return serdes.NewBinaryParser([]byte{0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d}, definitions.Get())
			},
			expected: map[string]any{
				"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
				"currency": "USD",
				"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
			},
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			p := tc.malleate(t)
			got, err := parsePathStep(p)
			if tc.expectedErr != nil {
				require.Error(t, err)
				require.Equal(t, tc.expectedErr, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, got)
			}
		})
	}
}

func TestParsePath(t *testing.T) {
	tt := []struct {
		name        string
		malleate    func(t *testing.T) interfaces.BinaryParser
		expected    []any
		expectedErr error
	}{
		{
			name: "fail - binary parser peek error",
			malleate: func(t *testing.T) interfaces.BinaryParser {
				parser := testutil.NewMockBinaryParser(gomock.NewController(t))
				parser.EXPECT().HasMore().Return(true)
				parser.EXPECT().Peek().Return(uint8(0), errors.New("peek error"))
				return parser
			},
			expected:    nil,
			expectedErr: errors.New("peek error"),
		},
		{
			name: "fail - binary parser read byte error",
			malleate: func(t *testing.T) interfaces.BinaryParser {
				parser := testutil.NewMockBinaryParser(gomock.NewController(t))
				parser.EXPECT().HasMore().Return(true)
				parser.EXPECT().Peek().Return(uint8(pathSeparatorByte), nil)
				parser.EXPECT().ReadByte().Return(uint8(0), errors.New("read byte error"))
				return parser
			},
			expected:    nil,
			expectedErr: errors.New("read byte error"),
		},
		{
			name: "pass - successfully parse path",
			malleate: func(t *testing.T) interfaces.BinaryParser {
				return serdes.NewBinaryParser([]byte{0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d}, definitions.Get())
			},
			expected: []any{
				map[string]any{
					"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
					"currency": "USD",
					"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := tc.malleate(t)
			got, err := parsePath(p)
			if tc.expectedErr != nil {
				require.Error(t, err)
				require.Equal(t, tc.expectedErr, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, got)
			}
		})
	}
}
