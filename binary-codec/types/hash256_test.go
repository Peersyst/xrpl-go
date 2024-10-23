package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHash256(t *testing.T) {
	hash := NewHash256()
	require.Equal(t, 32, hash.getLength())
}
