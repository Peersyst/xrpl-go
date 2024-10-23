package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOracle_EntryType(t *testing.T) {
	oracle := &Oracle{}
	assert.Equal(t, OracleEntry, oracle.EntryType())
}
