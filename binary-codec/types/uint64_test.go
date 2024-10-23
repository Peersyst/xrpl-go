package types

import (
	"bytes"
	"testing"

	"github.com/Peersyst/xrpl-go/binary-codec/serdes"
)

func TestUint64_FromJson(t *testing.T) {

	tt := []struct {
		name        string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			name:        "Valid uint64",
			input:       "1",
			expected:    []byte{0, 0, 0, 0, 0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			name:        "Valid uint64 (2)",
			input:       "100",
			expected:    []byte{0, 0, 0, 0, 0, 0, 1, 0},
			expectedErr: nil,
		},
		{
			name:        "Valid uint64 (3)",
			input:       "255",
			expected:    []byte{0, 0, 0, 0, 0, 0, 2, 85},
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			uint64 := &UInt64{}
			actual, err := uint64.FromJson(tc.input)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if !bytes.Equal(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestUint64_ToJson(t *testing.T) {
	tt := []struct {
		name        string
		input       []byte
		expected    string
		expectedErr error
	}{
		{
			name:        "Valid uint64",
			input:       []byte{0, 0, 0, 0, 0, 0, 0, 1},
			expected:    "0000000000000001",
			expectedErr: nil,
		},
		{
			name:        "Valid uint64 (2)",
			input:       []byte{0, 0, 0, 0, 0, 0, 0, 100},
			expected:    "0000000000000064",
			expectedErr: nil,
		},
		{
			name:        "Valid uint64 (3)",
			input:       []byte{0, 0, 0, 0, 0, 0, 0, 255},
			expected:    "00000000000000FF",
			expectedErr: nil,
		},
		{
			name:        "Valid uint64 (large number)",
			input:       []byte{255, 255, 255, 255, 255, 255, 255, 255},
			expected:    "FFFFFFFFFFFFFFFF", // Max uint64 value
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			uint64 := &UInt64{}
			parser := serdes.NewBinaryParser(tc.input)
			actual, err := uint64.ToJson(parser)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}

}
