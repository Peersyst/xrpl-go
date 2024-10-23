package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBridge_EntryType(t *testing.T) {
	entry := &Bridge{}
	assert.Equal(t, BridgeEntry, entry.EntryType())
}
