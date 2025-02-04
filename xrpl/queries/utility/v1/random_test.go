package v1

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
)

func TestRandomResponse(t *testing.T) {
	s := RandomResponse{
		Random: "8ED765AEBBD6767603C2C9375B2679AEC76E6A8133EF59F04F9FC1AAA70E41AF",
	}

	j := `{
	"random": "8ED765AEBBD6767603C2C9375B2679AEC76E6A8133EF59F04F9FC1AAA70E41AF"
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
