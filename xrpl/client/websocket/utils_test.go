package websocket

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/client"
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/account"
	"github.com/stretchr/testify/require"
)

func TestFormatRequest(t *testing.T) {
	ws := &WebsocketClient{cfg: &WebsocketConfig{}}
	tt := []struct {
		description string
		req         client.XRPLRequest
		id          int
		marker      any
		expected    string
		expectedErr error
	}{
		{
			description: "valid request",
			req: &account.AccountChannelsRequest{
				Account:            "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				DestinationAccount: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				Limit:              70,
			},
			id:     1,
			marker: nil,
			expected: `{
				"id": 1,
				"command":"account_channels",
				"account":"r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"destination_account":"r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"limit":70
			}`,
			expectedErr: nil,
		},
		{
			description: "valid request with marker",
			req: &account.AccountChannelsRequest{
				Account:            "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				DestinationAccount: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				Limit:              70,
			},
			id:     1,
			marker: "hdsohdaoidhadasd",
			expected: `{
				"id": 1,
				"command":"account_channels",
				"account":"r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"destination_account":"r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"limit":70,
				"marker":"hdsohdaoidhadasd"
			}`,
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			a, err := ws.formatRequest(tc.req, tc.id, tc.marker)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.JSONEq(t, tc.expected, string(a))
			}
		})
	}
}
