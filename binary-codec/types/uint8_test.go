package types

import (
	"bytes"
	"testing"

	"github.com/Peersyst/xrpl-go/binary-codec/serdes"
)

func TestUint8_FromJson(t *testing.T) {

	tt := []struct {
		name        string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			name:        "Valid uint8",
			input:       1,
			expected:    []byte{1},
			expectedErr: nil,
		},
		{
			name:        "Valid uint8 (2)",
			input:       100,
			expected:    []byte{100},
			expectedErr: nil,
		},
		{
			name:        "Valid uint8 (3)",
			input:       255,
			expected:    []byte{255},
			expectedErr: nil,
		},
		// TODO: Add test for overflow case
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			uint8 := &UInt8{}
			actual, err := uint8.FromJson(tc.input)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if !bytes.Equal(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestUint8_ToJson(t *testing.T) {
	tt := []struct {
		name        string
		input       []byte
		expected    int
		expectedErr error
	}{
		{
			name:        "Valid uint8",
			input:       []byte{1},
			expected:    1,
			expectedErr: nil,
		},
		{
			name:        "Valid uint8 (2)",
			input:       []byte{100},
			expected:    100,
			expectedErr: nil,
		},
		{
			name:        "Valid uint8 (3)",
			input:       []byte{255},
			expected:    255,
			expectedErr: nil,
		},
		// TODO: Add test for overflow case
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			uint8 := &UInt8{}
			parser := serdes.NewBinaryParser(tc.input)
			actual, err := uint8.ToJson(parser)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}

}
