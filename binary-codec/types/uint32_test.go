package types

import (
	"bytes"
	"testing"

	"github.com/Peersyst/xrpl-go/binary-codec/serdes"
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
			input:       1,
			expected:    []byte{0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			name:        "Valid uint32 (2)",
			input:       100,
			expected:    []byte{0, 0, 0, 100},
			expectedErr: nil,
		},
		{
			name:        "Valid uint32 (3)",
			input:       255,
			expected:    []byte{0, 0, 0, 255},
			expectedErr: nil,
		},
		// TODO: Add test for overflow case
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			uint32 := &UInt32{}
			actual, err := uint32.FromJson(tc.input)
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
	tt := []struct {
		name        string
		input       []byte
		expected    int
		expectedErr error
	}{
		{
			name:        "Valid uint32",
			input:       []byte{0, 0, 0, 1},
			expected:    1,
			expectedErr: nil,
		},
		{
			name:        "Valid uint32 (2)",
			input:       []byte{0, 0, 0, 100},
			expected:    100,
			expectedErr: nil,
		},
		{
			name:        "Valid uint32 (3)",
			input:       []byte{0, 0, 0, 255},
			expected:    255,
			expectedErr: nil,
		},
		// TODO: Add test for overflow case
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			uint32 := &UInt32{}
			parser := serdes.NewBinaryParser(tc.input)
			actual, err := uint32.ToJson(parser)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}

}
