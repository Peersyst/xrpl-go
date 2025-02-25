package types

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/queries/subscription/types"
	"github.com/stretchr/testify/require"
)

func TestMessage_IsRequest(t *testing.T) {
	tests := []struct {
		name    string
		message Message
		want    bool
	}{
		{
			name: "message with ID is request",
			message: Message{
				ID: 1,
			},
			want: true,
		},
		{
			name: "message without ID is not request",
			message: Message{
				Type: types.LedgerStreamType,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.message.IsRequest())
		})
	}
}

func TestMessage_IsStream(t *testing.T) {
	tests := []struct {
		name    string
		message Message
		want    bool
	}{
		{
			name: "message with stream type is stream",
			message: Message{
				Type: types.LedgerStreamType,
			},
			want: true,
		},
		{
			name: "message without stream type is not stream",
			message: Message{
				ID: 1,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.message.IsStream())
		})
	}
}
