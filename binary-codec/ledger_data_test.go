package binarycodec

import (
	"encoding/hex"
	"testing"

	"github.com/Peersyst/xrpl-go/binary-codec/serdes"
	"github.com/stretchr/testify/require"
)

func TestDecodeLedgerData(t *testing.T) {
	testcases := []struct {
		name        string
		input       string
		expected    LedgerData
		expectedErr error
	}{
		{
			name:        "fail - invalid hex string",
			input:       "invalid",
			expectedErr: hex.InvalidByteError(0x69),
		},
		{
			name:        "fail - invalid ledger index",
			input:       "01E914",
			expectedErr: serdes.ErrParserOutOfBound,
		},
		{
			name:        "fail - invalid total coins",
			input:       "01E91435016340767BF1",
			expectedErr: serdes.ErrParserOutOfBound,
		},
		{
			name:        "fail - invalid parent hash",
			input:       "01E91435016340767BF1C4A3EACEB081770D8ADE216C85445DD6FB002C6B5A2930F2DE",
			expectedErr: serdes.ErrParserOutOfBound,
		},
		{
			name:        "fail - invalid transaction hash",
			input:       "01E91435016340767BF1C4A3EACEB081770D8ADE216C85445DD6FB002C6B5A2930F2DECE006DA18150CB18F6DD33",
			expectedErr: serdes.ErrParserOutOfBound,
		},
		{
			name:        "fail - invalid account hash",
			input:       "01E91435016340767BF1C4A3EACEB081770D8ADE216C85445DD6FB002C6B5A2930F2DECE006DA18150CB18F6DD33F6F0990754C962A7CCE62F332FF9C13939B03B864117F0BDA86B6E9B4F873B5C3E520634D343EF5D",
			expectedErr: serdes.ErrParserOutOfBound,
		},
		{
			name:        "fail - invalid parent close time",
			input:       "01E91435016340767BF1C4A3EACEB081770D8ADE216C85445DD6FB002C6B5A2930F2DECE006DA18150CB18F6DD33F6F0990754C962A7CCE62F332FF9C13939B03B864117F0BDA86B6E9B4F873B5C3E520634D343EF5D9D9A4246643D64DAD278BA95DC0EAC6EB5350CF970D521276C",
			expectedErr: serdes.ErrParserOutOfBound,
		},
		{
			name:        "fail - invalid close time",
			input:       "01E91435016340767BF1C4A3EACEB081770D8ADE216C85445DD6FB002C6B5A2930F2DECE006DA18150CB18F6DD33F6F0990754C962A7CCE62F332FF9C13939B03B864117F0BDA86B6E9B4F873B5C3E520634D343EF5D9D9A4246643D64DAD278BA95DC0EAC6EB5350CF970D521276CDE21276C",
			expectedErr: serdes.ErrParserOutOfBound,
		},
		{
			name:        "fail - invalid close time resolution",
			input:       "01E91435016340767BF1C4A3EACEB081770D8ADE216C85445DD6FB002C6B5A2930F2DECE006DA18150CB18F6DD33F6F0990754C962A7CCE62F332FF9C13939B03B864117F0BDA86B6E9B4F873B5C3E520634D343EF5D9D9A4246643D64DAD278BA95DC0EAC6EB5350CF970D521276CDE21276CE6",
			expectedErr: serdes.ErrParserOutOfBound,
		},
		{
			name:        "fail - invalid close flags",
			input:       "01E91435016340767BF1C4A3EACEB081770D8ADE216C85445DD6FB002C6B5A2930F2DECE006DA18150CB18F6DD33F6F0990754C962A7CCE62F332FF9C13939B03B864117F0BDA86B6E9B4F873B5C3E520634D343EF5D9D9A4246643D64DAD278BA95DC0EAC6EB5350CF970D521276CDE21276CE60A",
			expectedErr: serdes.ErrParserOutOfBound,
		},
		{
			name:  "pass - valid encoded ledger data",
			input: "01E91435016340767BF1C4A3EACEB081770D8ADE216C85445DD6FB002C6B5A2930F2DECE006DA18150CB18F6DD33F6F0990754C962A7CCE62F332FF9C13939B03B864117F0BDA86B6E9B4F873B5C3E520634D343EF5D9D9A4246643D64DAD278BA95DC0EAC6EB5350CF970D521276CDE21276CE60A00",
			expected: LedgerData{
				AccountHash:         "3B5C3E520634D343EF5D9D9A4246643D64DAD278BA95DC0EAC6EB5350CF970D5",
				CloseFlags:          0,
				CloseTime:           556231910,
				CloseTimeResolution: 10,
				LedgerIndex:         32052277,
				ParentCloseTime:     556231902,
				ParentHash:          "EACEB081770D8ADE216C85445DD6FB002C6B5A2930F2DECE006DA18150CB18F6",
				TotalCoins:          "99994494362043555",
				TransactionHash:     "DD33F6F0990754C962A7CCE62F332FF9C13939B03B864117F0BDA86B6E9B4F87",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ledgerData, err := DecodeLedgerData(tc.input)
			if tc.expectedErr != nil {
				require.Error(t, err)
				require.Equal(t, tc.expectedErr, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, ledgerData)
			}
		})
	}
}
