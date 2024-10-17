package types

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	typeerrors "github.com/Peersyst/xrpl-go/binary-codec/types/errors"
	"github.com/Peersyst/xrpl-go/binary-codec/types/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestXChainBridge_FromJson(t *testing.T) {
	tt := []struct {
		name string
		json any
		want []byte
		err  error
	}{
		{
			name: "valid xchain bridge",
			json: map[string]any{
				"LockingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"LockingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"IssuingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"IssuingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
			},
			want: []byte{83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18},
			err:  nil,
		},
		{
			name: "invalid LockingChainDoor classic address",
			json: map[string]any{
				"LockingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p1",
				"LockingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"IssuingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"IssuingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
			},
			want: nil,
			err:  &ErrDecodeClassicAddress{},
		},
		{
			name: "invalid LockingChainIssue classic address",
			json: map[string]any{
				"LockingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"LockingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p1",
				"IssuingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"IssuingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
			},
			want: nil,
			err:  &ErrDecodeClassicAddress{},
		},
		{
			name: "invalid IssuingChainDoor classic address",
			json: map[string]any{
				"LockingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"LockingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"IssuingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p1",
				"IssuingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
			},
			want: nil,
			err:  &ErrDecodeClassicAddress{},
		},
		{
			name: "invalid IssuingChainIssue classic address",
			json: map[string]any{
				"LockingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"LockingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"IssuingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"IssuingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p1",
			},
			want: nil,
			err:  &ErrDecodeClassicAddress{},
		},
		{
			name: "not a valid json",
			json: "not a valid json",
			want: nil,
			err:  &typeerrors.ErrNotValidJson{},
		},
		{
			name: "invalid xchain bridge",
			json: map[string]any{
				"LockingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"IssuingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"IssuingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
			},
			want: nil,
			err:  &ErrNotValidXChainBridge{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			xcb := &XChainBridge{}
			got, err := xcb.FromJson(tc.json)
			if err != tc.err {
				t.Errorf("FromJson() error = %v, want %v", err.Error(), tc.err.Error())
			}
			if !bytes.Equal(got, tc.want) {
				t.Errorf("FromJson() got = %v, want %v", got, tc.want)
			}

		})
	}
}

func TestXChainBridge_ToJson(t *testing.T) {
	tt := []struct {
		name  string
		input []byte
		opts  []int
		want  map[string]string
		err   error
		setup func(t *testing.T) (*XChainBridge, *testutil.MockBinaryParser)
	}{
		{
			name:  "Valid xchain bridge",
			input: []byte{83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18},
			opts:  []int{80},
			want: map[string]string{
				"LockingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"LockingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"IssuingChainDoor":  "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
				"IssuingChainIssue": "r3e7qTG44Mg8pHXgxPtyRx286Re5Urtx2p",
			},
			err: nil,
			setup: func(t *testing.T) (*XChainBridge, *testutil.MockBinaryParser) {
				ctrl := gomock.NewController(t)
				mock := testutil.NewMockBinaryParser(ctrl)
				mock.EXPECT().ReadBytes(80).Return([]byte{83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18}, nil)
				return &XChainBridge{}, mock
			},
		},
		{
			name:  "No length prefix",
			input: []byte{83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18},
			opts:  nil,
			want:  nil,
			err:   ErrNoLengthPrefix,
			setup: func(t *testing.T) (*XChainBridge, *testutil.MockBinaryParser) {
				return &XChainBridge{}, nil
			},
		},
		{
			name:  "ReadBytes error",
			input: []byte{83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18, 83, 223, 129, 195, 127, 70, 21, 146, 66, 247, 202, 145, 99, 224, 159, 4, 64, 41, 204, 18},
			opts:  []int{80},
			want:  nil,
			err:   &ErrReadBytes{},
			setup: func(t *testing.T) (*XChainBridge, *testutil.MockBinaryParser) {
				ctrl := gomock.NewController(t)
				mock := testutil.NewMockBinaryParser(ctrl)
				mock.EXPECT().ReadBytes(80).Return([]byte{}, errors.New("read bytes error"))
				return &XChainBridge{}, mock
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			xcb, mock := tc.setup(t)
			got, err := xcb.ToJson(mock, tc.opts...)
			if err != tc.err {
				t.Errorf("ToJson() error = %v, want %v", err.Error(), tc.err.Error())
			} else if tc.err == nil && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ToJson() got = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestErrNotValidXChainBridge_Error(t *testing.T) {
	err := &ErrNotValidXChainBridge{}
	require.Equal(t, "not a valid xchain bridge", err.Error())
}

func TestErrNotValidJson_Error(t *testing.T) {
	err := &typeerrors.ErrNotValidJson{}
	require.Equal(t, "not a valid json", err.Error())
}

func TestErrDecodeClassicAddress_Error(t *testing.T) {
	err := &ErrDecodeClassicAddress{}
	require.Equal(t, "decode classic address error", err.Error())
}

func TestErrReadBytes_Error(t *testing.T) {
	err := &ErrReadBytes{}
	require.Equal(t, "read bytes error", err.Error())
}
