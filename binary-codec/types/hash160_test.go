package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHash160(t *testing.T) {
	hash := NewHash160()
	require.Equal(t, 20, hash.getLength())
}
