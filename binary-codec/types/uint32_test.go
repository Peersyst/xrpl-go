package types

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/Peersyst/xrpl-go/v1/binary-codec/definitions"
	"github.com/Peersyst/xrpl-go/v1/binary-codec/serdes"
	"github.com/Peersyst/xrpl-go/v1/binary-codec/types/interfaces"
	"github.com/Peersyst/xrpl-go/v1/binary-codec/types/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUint32_FromJson(t *testing.T) {

	tt := []struct {
		name        string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			name:        "Valid uint32",
			input:       uint32(1),
			expected:    []byte{0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			name:        "Valid uint32 (2)",
			input:       uint32(100),
			expected:    []byte{0, 0, 0, 100},
			expectedErr: nil,
		},
		{
			name:        "Valid uint32 (3)",
			input:       uint32(255),
			expected:    []byte{0, 0, 0, 255},
			expectedErr: nil,
		},
		// TODO: Add test for overflow case
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			uint32 := &UInt32{}
			actual, err := uint32.FromJSON(tc.input)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if !bytes.Equal(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestUint32_ToJson(t *testing.T) {
	defs := definitions.Get()

	tt := []struct {
		name        string
		input       []byte
		malleate    func(t *testing.T) interfaces.BinaryParser
		expected    uint32
		expectedErr error
	}{
		{
			name:  "fail - invalid uint32",
			input: []byte{0, 0, 0, 1},
			malleate: func(t *testing.T) interfaces.BinaryParser {
				parserMock := testutil.NewMockBinaryParser(gomock.NewController(t))
				parserMock.EXPECT().ReadBytes(gomock.Any()).Return([]byte{}, errors.New("binary parser has no data"))
				return parserMock
			},
			expected:    0,
			expectedErr: fmt.Errorf("binary parser has no data"),
		},
		{
			name:  "pass - valid uint32",
			input: []byte{0, 0, 0, 1},
			malleate: func(t *testing.T) interfaces.BinaryParser {
				return serdes.NewBinaryParser([]byte{0, 0, 0, 1}, defs)
			},
			expected:    1,
			expectedErr: nil,
		},
		{
			name:  "pass - valid uint32 (2)",
			input: []byte{0, 0, 0, 100},
			malleate: func(t *testing.T) interfaces.BinaryParser {
				return serdes.NewBinaryParser([]byte{0, 0, 0, 100}, defs)
			},
			expected:    100,
			expectedErr: nil,
		},
		{
			name:  "pass - valid uint32 (3)",
			input: []byte{0, 0, 0, 255},
			malleate: func(t *testing.T) interfaces.BinaryParser {
				return serdes.NewBinaryParser([]byte{0, 0, 0, 255}, defs)
			},
			expected:    255,
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			class := &UInt32{}
			parser := tc.malleate(t)
			actual, err := class.ToJSON(parser)
			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.Equal(t, tc.expected, actual)
			}
		})
	}

}
