package path

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/testutil"
)

func TestPathFindCloseRequest(t *testing.T) {
	s := FindCloseRequest{
		Subcommand: Close,
	}

	j := `{
	"subcommand": "close"
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestPathFindStatusRequest(t *testing.T) {
	s := FindStatusRequest{
		Subcommand: Status,
	}

	j := `{
	"subcommand": "status"
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
