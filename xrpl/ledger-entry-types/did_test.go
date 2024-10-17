package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDID_EntryType(t *testing.T) {
	did := &DID{}
	assert.Equal(t, did.EntryType(), DIDEntry)
}
