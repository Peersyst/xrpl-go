package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXChainOwnedClaimID_EntryType(t *testing.T) {
	entry := &XChainOwnedClaimID{}
	assert.Equal(t, XChainOwnedClaimIDEntry, entry.EntryType())
}
