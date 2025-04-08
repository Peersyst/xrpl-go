package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHash192(t *testing.T) {
	hash := NewHash192()
	require.Equal(t, 24, hash.getLength())
}
