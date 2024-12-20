package rpc

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/queries/account"
	"github.com/stretchr/testify/assert"
)

func TestGetResult(t *testing.T) {
	t.Run("correctly decodes", func(t *testing.T) {

		jr := Response{
			Result: AnyJSON{
				"account":      "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"ledger_hash":  "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
				"ledger_index": json.Number(strconv.FormatInt(71766343, 10)),
			},
			Warning: "none",
			Warnings: []XRPLResponseWarning{{
				ID:      1,
				Message: "message",
			},
			},
		}

		expected := account.ChannelsResponse{
			Account:     "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			LedgerHash:  "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
			LedgerIndex: 71766343,
		}

		var acr account.ChannelsResponse
		err := jr.GetResult(&acr)

		assert.NoError(t, err)
		assert.Equal(t, expected, acr)
	})
	t.Run("throws error for incorrect mapping", func(t *testing.T) {

		jr := Response{
			Result: AnyJSON{
				"account":      123,
				"ledger_hash":  "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
				"ledger_index": json.Number(strconv.FormatInt(71766343, 10)),
			},
			Warning: "none",
			Warnings: []XRPLResponseWarning{{
				ID:      1,
				Message: "message",
			},
			},
		}

		var acr account.ChannelsResponse
		err := jr.GetResult(&acr)

		assert.Error(t, err)
	})
}
