package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXChainOwnedCreateAccountClaimID_EntryType(t *testing.T) {
	entry := &XChainOwnedCreateAccountClaimID{}
	assert.Equal(t, XChainOwnedCreateAccountClaimIDEntry, entry.EntryType())
}
