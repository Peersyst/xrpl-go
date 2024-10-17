package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHash128(t *testing.T) {
	hash := NewHash128()
	require.Equal(t, 16, hash.getLength())
}
