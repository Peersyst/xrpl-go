package types

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/binary-codec/definitions"
	"github.com/Peersyst/xrpl-go/v1/binary-codec/serdes"
	"github.com/stretchr/testify/require"
)

func TestGetSerializedType(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected SerializedType
	}{
		{
			name:     "pass - uint8",
			input:    "UInt8",
			expected: &UInt8{},
		},
		{
			name:     "pass - uint16",
			input:    "UInt16",
			expected: &UInt16{},
		},
		{
			name:     "pass - uint32",
			input:    "UInt32",
			expected: &UInt32{},
		},
		{
			name:     "pass - uint64",
			input:    "UInt64",
			expected: &UInt64{},
		},
		{
			name:     "pass - hash128",
			input:    "Hash128",
			expected: NewHash128(),
		},
		{
			name:     "pass - hash160",
			input:    "Hash160",
			expected: NewHash160(),
		},
		{
			name:     "pass - hash256",
			input:    "Hash256",
			expected: NewHash256(),
		},
		{
			name:     "pass - accountid",
			input:    "AccountID",
			expected: &AccountID{},
		},
		{
			name:     "pass - amount",
			input:    "Amount",
			expected: &Amount{},
		},
		{
			name:     "pass - vector256",
			input:    "Vector256",
			expected: &Vector256{},
		},
		{
			name:     "pass - blob",
			input:    "Blob",
			expected: &Blob{},
		},
		{
			name:     "pass - stobject",
			input:    "STObject",
			expected: NewSTObject(serdes.NewBinarySerializer(serdes.NewFieldIDCodec(definitions.Get()))),
		},
		{
			name:     "pass - starray",
			input:    "STArray",
			expected: &STArray{},
		},
		{
			name:     "pass - pathset",
			input:    "PathSet",
			expected: &PathSet{},
		},
		{
			name:     "fail - unknown type",
			input:    "Unknown",
			expected: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, GetSerializedType(tc.input))
		})
	}
}
