package websocket

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnection_Connect(t *testing.T) {
	conn := NewConnection("wss://s.altnet.rippletest.net")
	err := conn.Connect()
	require.NoError(t, err)
	require.True(t, conn.IsConnected())
}

func TestConnection_Disconnect(t *testing.T) {
	conn := NewConnection("wss://s.altnet.rippletest.net")
	err := conn.Connect()
	require.NoError(t, err)
	require.True(t, conn.IsConnected())
	err = conn.Disconnect()
	require.NoError(t, err)
	require.False(t, conn.IsConnected())
}

func TestConnection_IsConnected(t *testing.T) {
	conn := NewConnection("wss://s.altnet.rippletest.net")
	require.False(t, conn.IsConnected())
	err := conn.Connect()
	require.NoError(t, err)
	require.True(t, conn.IsConnected())
	err = conn.Disconnect()
	require.NoError(t, err)
	require.False(t, conn.IsConnected())
}
