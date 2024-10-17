package serdes

import (
	"errors"
	"testing"

	"github.com/Peersyst/xrpl-go/binary-codec/definitions"
	"github.com/stretchr/testify/require"
)

func TestBinaryParser_ReadVariableLength(t *testing.T) {
	tt := []struct {
		description string
		input       []byte
		output      int
		expectedErr error
	}{
		{
			description: "Length less than 193",
			input:       []byte{190, 230, 131},
			output:      190,
		},
		{
			description: "length > 192 & length < 241",
			input:       []byte{195, 230, 112, 234, 98},
			output:      935,
		},
		{
			description: "length > 192 & length < 241 missing bytes",
			input:       []byte{195},
			output:      0,
			expectedErr: ErrParserOutOfBound,
		},
		{
			description: "length > 240 & length < 255",
			input:       []byte{242, 112, 78, 95, 115},
			output:      106767,
		},
		{
			description: "length > 240 & length < 255 missing bytes",
			input:       []byte{242},
			output:      0,
			expectedErr: ErrParserOutOfBound,
		},
		{
			description: "length > 240 & length < 255 missing bytes",
			input:       []byte{242, 112},
			output:      0,
			expectedErr: ErrParserOutOfBound,
		},
		{
			description: "empty input",
			input:       []byte{},
			output:      0,
			expectedErr: ErrParserOutOfBound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			p := NewBinaryParser(tc.input)
			actual, err := p.ReadVariableLength()
			require.Equal(t, tc.output, actual)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestBinaryParser_HasMore(t *testing.T) {
	tt := []struct {
		description string
		input       []byte
		output      bool
	}{
		{
			description: "has more",
			input:       []byte{190, 230, 131},
			output:      true,
		},
		{
			description: "has no more",
			input:       []byte{},
			output:      false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			p := NewBinaryParser(tc.input)
			actual := p.HasMore()
			require.Equal(t, tc.output, actual)
		})
	}
}

func TestBinaryParser_Peek(t *testing.T) {
	tt := []struct {
		description string
		input       []byte
		output      byte
		expectedErr error
	}{
		{
			description: "peek byte",
			input:       []byte{190, 230, 131},
			output:      190,
		},
		{
			description: "peek byte with no data",
			input:       []byte{},
			output:      0,
			expectedErr: ErrParserOutOfBound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			p := NewBinaryParser(tc.input)
			actual, err := p.Peek()
			require.Equal(t, tc.output, actual)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestBinaryParser_ReadByte(t *testing.T) {
	tt := []struct {
		description string
		input       []byte
		output      byte
		expectedErr error
	}{
		{
			description: "read byte",
			input:       []byte{190, 230, 131},
			output:      190,
		},
		{
			description: "read byte with no data",
			input:       []byte{},
			output:      0,
			expectedErr: ErrParserOutOfBound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			p := NewBinaryParser(tc.input)
			actual, err := p.ReadByte()
			require.Equal(t, tc.output, actual)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestBinaryParser_ReadBytes(t *testing.T) {
	tt := []struct {
		description  string
		input        []byte
		n            int
		output       []byte
		expectedData []byte
		expectedErr  error
	}{
		{
			description:  "read bytes",
			input:        []byte{190, 230, 131},
			n:            2,
			output:       []byte{190, 230},
			expectedData: []byte{131},
		},
		{
			description:  "read bytes with no data",
			input:        []byte{},
			n:            2,
			output:       []byte(nil),
			expectedErr:  ErrParserOutOfBound,
			expectedData: []byte{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			p := NewBinaryParser(tc.input)
			actual, err := p.ReadBytes(tc.n)
			require.Equal(t, tc.output, actual)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestBinaryParser_readFieldHeader(t *testing.T) {
	tt := []struct {
		description string
		input       []byte
		output      *definitions.FieldHeader
		expectedErr error
	}{
		{
			description: "read field header",
			input:       []byte{190, 230, 131},
			output: &definitions.FieldHeader{
				TypeCode:  11,
				FieldCode: 14,
			},
		},
		{
			description: "read field header with one byte",
			input:       []byte{0},
			output:      nil,
			expectedErr: ErrParserOutOfBound,
		},
		{
			description: "read field header with no data",
			input:       []byte{0, 0},
			output:      nil,
			expectedErr: errors.New("invalid typecode"),
		},
		{
			description: "read field header with two bytes",
			input:       []byte{0, 16},
			output:      nil,
			expectedErr: ErrParserOutOfBound,
		},
		{
			description: "invalid fieldcode",
			input:       []byte{0, 16, 0},
			output:      nil,
			expectedErr: errors.New("invalid fieldcode"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			p := NewBinaryParser(tc.input)
			actual, err := p.readFieldHeader()
			require.Equal(t, tc.output, actual)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestBinaryParser_ReadField(t *testing.T) {
	tt := []struct {
		description string
		input       []byte
		output      *definitions.FieldInstance
		expectedErr error
	}{
		// {
		// 	description: "read field",
		// 	input:       []byte{190, 230, 131},
		// 	output:      &definitions.FieldInstance{
		// 		FieldHeader: &definitions.FieldHeader{
		// 			TypeCode:  11,
		// 			FieldCode: 14,
		// 		},
		// 	},
		// },
		{
			description: "read field with no data",
			input:       []byte{},
			output:      nil,
			expectedErr: ErrParserOutOfBound,
		},
		{
			description: "read field with invalid typecode",
			input:       []byte{0, 0},
			output:      nil,
			expectedErr: errors.New("invalid typecode"),
		},
		{
			description: "read field with invalid fieldcode",
			input:       []byte{0, 16, 0},
			output:      nil,
			expectedErr: errors.New("invalid fieldcode"),
		},
		// {
		// 	description: "not found error field header",
		// 	input:       []byte{190, 230, 131},
		// 	output:      nil,
		// 	expectedErr: errors.New("not found error field header"),
		// },
		// {
		// 	description: "not found error field name",
		// 	input:       []byte{190, 230, 131},
		// 	output:      nil,
		// 	expectedErr: errors.New("not found error field name"),
		// },
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			p := NewBinaryParser(tc.input)
			actual, err := p.ReadField()
			require.Equal(t, tc.output, actual)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}
