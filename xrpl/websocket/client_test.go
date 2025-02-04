package websocket

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	commonconstants "github.com/Peersyst/xrpl-go/xrpl/common"
	account "github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/websocket/interfaces"
	"github.com/Peersyst/xrpl-go/xrpl/websocket/testutil"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

func TestClient_Connect(t *testing.T) {
	ws := NewClient(NewClientConfig().WithHost("wss://s.altnet.rippletest.net"))
	err := ws.Connect()
	require.NoError(t, err)
	require.True(t, ws.IsConnected())
	ws.Disconnect()
	require.False(t, ws.IsConnected())
}

func TestClient_Disconnect(t *testing.T) {
	ws := NewClient(NewClientConfig().WithHost("wss://s.altnet.rippletest.net"))
	ws.Connect()
	require.True(t, ws.IsConnected())
	ws.Disconnect()
	require.False(t, ws.IsConnected())
}

func TestClient_IsConnected(t *testing.T) {
	ws := NewClient(NewClientConfig().WithHost("wss://s.altnet.rippletest.net"))
	require.False(t, ws.IsConnected())
	ws.Connect()
	require.True(t, ws.IsConnected())
	ws.Disconnect()
	require.False(t, ws.IsConnected())
}

func TestClient_SendRequest(t *testing.T) {
	tt := []struct {
		description    string
		req            interfaces.Request
		res            *ClientResponse
		expectedErr    error
		serverMessages []map[string]any
	}{
		{
			description: "successful request",
			req: &account.ChannelsRequest{
				Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
			},
			res: &ClientResponse{
				ID: 1,
				Result: map[string]any{
					"account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					"channels": []any{
						map[string]any{
							"account":             "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"amount":              "1000",
							"balance":             "0",
							"channel_id":          "C7F634794B79DB40E87179A9D1BF05D05797AE7E92DF8E93FD6656E8C4BE3AE7",
							"destination_account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
							"public_key":          "aBR7mdD75Ycs8DRhMgQ4EMUEmBArF8SEh1hfjrT2V9DQTLNbJVqw",
							"public_key_hex":      "03CFD18E689434F032A4E84C63E2A3A6472D684EAF4FD52CA67742F3E24BAE81B2",
							"settle_delay":        json.Number("60"),
						},
					},
					"ledger_hash":  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
					"ledger_index": json.Number("71766314"),
					"validated":    true,
				},
			},
			expectedErr: nil,
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
						"account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
						"channels": []any{
							map[string]any{
								"account":             "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
								"amount":              "1000",
								"balance":             "0",
								"channel_id":          "C7F634794B79DB40E87179A9D1BF05D05797AE7E92DF8E93FD6656E8C4BE3AE7",
								"destination_account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
								"public_key":          "aBR7mdD75Ycs8DRhMgQ4EMUEmBArF8SEh1hfjrT2V9DQTLNbJVqw",
								"public_key_hex":      "03CFD18E689434F032A4E84C63E2A3A6472D684EAF4FD52CA67742F3E24BAE81B2",
								"settle_delay":        60,
							},
						},
						"ledger_hash":  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
						"ledger_index": 71766314,
						"validated":    true,
					},
				},
			},
		},
		{
			description: "Invalid ID",
			req: &account.ChannelsRequest{
				Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
			},
			res: &ClientResponse{
				Result: map[string]any{
					"account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					"channels": []any{
						map[string]any{
							"account":             "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"amount":              "1000",
							"balance":             "0",
							"channel_id":          "C7F634794B79DB40E87179A9D1BF05D05797AE7E92DF8E93FD6656E8C4BE3AE7",
							"destination_account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
							"public_key":          "aBR7mdD75Ycs8DRhMgQ4EMUEmBArF8SEh1hfjrT2V9DQTLNbJVqw",
							"public_key_hex":      "03CFD18E689434F032A4E84C63E2A3A6472D684EAF4FD52CA67742F3E24BAE81B2",
							"settle_delay":        json.Number("60"),
						},
					},
					"ledger_hash":  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
					"ledger_index": json.Number("71766314"),
					"validated":    true,
				},
			},
			expectedErr: ErrIncorrectID,
			serverMessages: []map[string]any{
				{
					"id": 2,
					"result": map[string]any{
						"account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
						"channels": []any{
							map[string]any{
								"account":             "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
								"amount":              "1000",
								"balance":             "0",
								"channel_id":          "C7F634794B79DB40E87179A9D1BF05D05797AE7E92DF8E93FD6656E8C4BE3AE7",
								"destination_account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
								"public_key":          "aBR7mdD75Ycs8DRhMgQ4EMUEmBArF8SEh1hfjrT2V9DQTLNbJVqw",
								"public_key_hex":      "03CFD18E689434F032A4E84C63E2A3A6472D684EAF4FD52CA67742F3E24BAE81B2",
								"settle_delay":        60,
							},
						},
						"ledger_hash":  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
						"ledger_index": 71766314,
						"validated":    true,
					},
				},
			},
		},
		{
			description: "Error response",
			req: &account.ChannelsRequest{
				Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
			},
			res: &ClientResponse{
				ID: 1,
				Result: map[string]any{
					"account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					"channels": []any{
						map[string]any{
							"account":             "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"amount":              "1000",
							"balance":             "0",
							"channel_id":          "C7F634794B79DB40E87179A9D1BF05D05797AE7E92DF8E93FD6656E8C4BE3AE7",
							"destination_account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
							"public_key":          "aBR7mdD75Ycs8DRhMgQ4EMUEmBArF8SEh1hfjrT2V9DQTLNbJVqw",
							"public_key_hex":      "03CFD18E689434F032A4E84C63E2A3A6472D684EAF4FD52CA67742F3E24BAE81B2",
							"settle_delay":        json.Number("60"),
						},
					},
					"ledger_hash":  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
					"ledger_index": json.Number("71766314"),
					"validated":    true,
				},
			},
			expectedErr: &ErrorWebsocketClientXrplResponse{
				Type: "invalidParams",
				Request: map[string]any{
					"account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				},
			},
			serverMessages: []map[string]any{
				{
					"id":    1,
					"error": "invalidParams",
					"value": map[string]any{
						"account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			ws := &testutil.MockWebSocketServer{Msgs: tc.serverMessages}
			s := ws.TestWebSocketServer(func(c *websocket.Conn) {
				for _, m := range tc.serverMessages {
					err := c.WriteJSON(m)
					if err != nil {
						println("error writing message")
					}
				}
			})
			defer s.Close()
			url, _ := testutil.ConvertHTTPToWS(s.URL)
			cl := &Client{cfg: ClientConfig{
				host: url,
			}}
			if err := cl.Connect(); err != nil {
				t.Errorf("Error connecting to server: %v", err)
			}

			res, err := cl.Request(tc.req)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.EqualValues(t, tc.res, res)
			}

			cl.Disconnect()
		})
	}
}

func TestClient_formatRequest(t *testing.T) {
	ws := &Client{}
	tt := []struct {
		description string
		req         interfaces.Request
		id          int
		marker      any
		expected    string
		expectedErr error
	}{
		{
			description: "valid request",
			req: &account.ChannelsRequest{
				Account:            "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				DestinationAccount: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				Limit:              70,
			},
			id:     1,
			marker: nil,
			expected: `{
				"id": 1,
				"BaseRequest": {},
				"account":"r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"api_version":2,
				"command":"account_channels",
				"destination_account":"r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"limit":70
			}`,
			expectedErr: nil,
		},
		{
			description: "valid request with marker",
			req: &account.ChannelsRequest{
				Account:            "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				DestinationAccount: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				Limit:              70,
			},
			id:     1,
			marker: "hdsohdaoidhadasd",
			expected: `{
				"id": 1,
				"BaseRequest": {},
				"account":"r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"api_version": 2,
				"command":"account_channels",
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

func TestClient_convertTransactionAddressToClassicAddress(t *testing.T) {
	ws := &Client{}
	tests := []struct {
		name      string
		tx        transaction.FlatTransaction
		fieldName string
		expected  transaction.FlatTransaction
	}{
		{
			name: "No conversion for classic address",
			tx: transaction.FlatTransaction{
				"Destination": "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
			fieldName: "Destination",
			expected: transaction.FlatTransaction{
				"Destination": "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
		},
		{
			name: "Field not present in transaction",
			tx: transaction.FlatTransaction{
				"Amount": "1000000",
			},
			fieldName: "Destination",
			expected: transaction.FlatTransaction{
				"Amount": "1000000",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws.convertTransactionAddressToClassicAddress(&tt.tx, tt.fieldName)
			if reflect.DeepEqual(tt.expected, &tt.tx) {
				t.Errorf("expected %+v, result %+v", tt.expected, &tt.tx)
			}
		})
	}
}

func TestClient_validateTransactionAddress(t *testing.T) {
	ws := &Client{}
	tests := []struct {
		name         string
		tx           transaction.FlatTransaction
		addressField string
		tagField     string
		expected     transaction.FlatTransaction
		expectedErr  error
	}{
		{
			name: "Valid classic address without tag",
			tx: transaction.FlatTransaction{
				"Account": "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
			addressField: "Account",
			tagField:     "SourceTag",
			expected: transaction.FlatTransaction{
				"Account": "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
			expectedErr: nil,
		},
		{
			name: "Valid classic address with tag",
			tx: transaction.FlatTransaction{
				"Destination":    "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"DestinationTag": uint32(12345),
			},
			addressField: "Destination",
			tagField:     "DestinationTag",
			expected: transaction.FlatTransaction{
				"Destination":    "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"DestinationTag": uint32(12345),
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ws.validateTransactionAddress(&tt.tx, tt.addressField, tt.tagField)

			if tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(tt.expected, tt.tx) {
				t.Errorf("Expected %v, but got %v", tt.expected, tt.tx)
			}
		})
	}
}

func TestClient_setValidTransactionAddresses(t *testing.T) {
	tests := []struct {
		name        string
		tx          transaction.FlatTransaction
		expected    transaction.FlatTransaction
		expectedErr error
	}{
		{
			name: "Valid transaction with classic addresses",
			tx: transaction.FlatTransaction{
				"Account":     "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"Destination": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
			},
			expected: transaction.FlatTransaction{
				"Account":     "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"Destination": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
			},
			expectedErr: nil,
		},
		{
			name: "Transaction with additional address fields",
			tx: transaction.FlatTransaction{
				"Account":     "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"Destination": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
				"Owner":       "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"RegularKey":  "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
			expected: transaction.FlatTransaction{
				"Account":     "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"Destination": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
				"Owner":       "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"RegularKey":  "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
			expectedErr: nil,
		},
	}

	ws := &Client{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ws.setValidTransactionAddresses(&tt.tx)

			if tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(tt.expected, tt.tx) {
				t.Errorf("Expected %v, but got %v", tt.expected, tt.tx)
			}
		})
	}
}

func TestClient_setTransactionNextValidSequenceNumber(t *testing.T) {
	tests := []struct {
		name           string
		tx             transaction.FlatTransaction
		serverMessages []map[string]any
		expected       transaction.FlatTransaction
		expectedErr    error
	}{
		{
			name: "Valid transaction",
			tx: transaction.FlatTransaction{
				"Account": "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
			},
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
						"account_data": map[string]any{
							"Sequence": uint32(42),
						},
						"ledger_current_index": uint32(100),
					},
				},
			},
			expected: transaction.FlatTransaction{
				"Account":  "rGWrZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf",
				"Sequence": uint32(42),
			},
			expectedErr: nil,
		},
		{
			name:           "Missing Account",
			tx:             transaction.FlatTransaction{},
			serverMessages: []map[string]any{},
			expected:       transaction.FlatTransaction{},
			expectedErr:    errors.New("missing Account in transaction"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &testutil.MockWebSocketServer{Msgs: tt.serverMessages}
			s := ws.TestWebSocketServer(func(c *websocket.Conn) {
				for _, m := range tt.serverMessages {
					err := c.WriteJSON(m)
					if err != nil {
						t.Errorf("error writing message: %v", err)
					}
				}
			})
			defer s.Close()

			url, _ := testutil.ConvertHTTPToWS(s.URL)
			cl := &Client{
				cfg: ClientConfig{
					host: url,
				},
			}

			if err := cl.Connect(); err != nil {
				t.Errorf("Error connecting to server: %v", err)
			}

			err := cl.setTransactionNextValidSequenceNumber(&tt.tx)

			if tt.expectedErr != nil {
				if !reflect.DeepEqual(err.Error(), tt.expectedErr.Error()) {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if !reflect.DeepEqual(tt.expected, tt.tx) {
				t.Logf("Expected:")
				for k, v := range tt.expected {
					t.Logf("  %s: %v (type: %T)", k, v, v)
				}
				t.Logf("Got:")
				for k, v := range tt.tx {
					t.Logf("  %s: %v (type: %T)", k, v, v)
				}
				t.Errorf("Expected %v but got %v", tt.expected, tt.tx)
			}

			cl.Disconnect()
		})
	}
}

func TestClient_calculateFeePerTransactionType(t *testing.T) {
	tests := []struct {
		name           string
		tx             transaction.FlatTransaction
		serverMessages []map[string]any
		expectedFee    string
		expectedErr    error
		feeCushion     float32
	}{
		{
			name: "Basic fee calculation",
			tx: transaction.FlatTransaction{
				"TransactionType": transaction.PaymentTx,
			},
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
						"info": map[string]any{
							"validated_ledger": map[string]any{
								"base_fee_xrp": float32(0.00001),
							},
							"load_factor": float32(1),
						},
					},
				},
			},
			expectedFee: "10",
			expectedErr: nil,
			feeCushion:  1,
		},
		{
			name: "Fee calculation with high load factor",
			tx: transaction.FlatTransaction{
				"TransactionType": transaction.PaymentTx,
			},
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
						"info": map[string]any{
							"validated_ledger": map[string]any{
								"base_fee_xrp": float32(0.00001),
							},
							"load_factor": float32(1000),
						},
					},
				},
			},
			expectedFee: "10000",
			expectedErr: nil,
			feeCushion:  1,
		},
		{
			name: "Fee calculation with max fee limit",
			tx: transaction.FlatTransaction{
				"TransactionType": transaction.PaymentTx,
			},
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
						"info": map[string]any{
							"validated_ledger": map[string]any{
								"base_fee_xrp": float32(1),
							},
							"load_factor": float32(1000),
						},
					},
				},
			},
			expectedFee: "2000000",
			expectedErr: nil,
			feeCushion:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &testutil.MockWebSocketServer{Msgs: tt.serverMessages}
			s := ws.TestWebSocketServer(func(c *websocket.Conn) {
				for _, m := range tt.serverMessages {
					err := c.WriteJSON(m)
					if err != nil {
						t.Errorf("error writing message: %v", err)
					}
				}
			})
			defer s.Close()

			url, _ := testutil.ConvertHTTPToWS(s.URL)
			cl := &Client{
				cfg: ClientConfig{
					host:       url,
					feeCushion: tt.feeCushion,
					maxFeeXRP:  DefaultMaxFeeXRP,
				},
			}

			if err := cl.Connect(); err != nil {
				t.Errorf("Error connecting to server: %v", err)
			}

			err := cl.calculateFeePerTransactionType(&tt.tx, 0)

			if tt.expectedErr != nil {
				if !reflect.DeepEqual(err.Error(), tt.expectedErr.Error()) {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(tt.expectedFee, tt.tx["Fee"]) {
					t.Errorf("Expected fee %v, but got %v", tt.expectedFee, tt.tx["Fee"])
				}
			}

			cl.Disconnect()
		})
	}
}

func TestClient_setLastLedgerSequence(t *testing.T) {
	tests := []struct {
		name           string
		serverMessages []map[string]any
		tx             transaction.FlatTransaction
		expectedTx     transaction.FlatTransaction
		expectedErr    error
	}{
		{
			name: "Successfully set LastLedgerSequence",
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": transaction.FlatTransaction{
						"ledger_index": 1000,
					},
				},
			},
			tx:          transaction.FlatTransaction{},
			expectedTx:  transaction.FlatTransaction{"LastLedgerSequence": uint32(1000 + commonconstants.LedgerOffset)},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &testutil.MockWebSocketServer{Msgs: tt.serverMessages}
			s := ws.TestWebSocketServer(func(c *websocket.Conn) {
				for _, m := range tt.serverMessages {
					err := c.WriteJSON(m)
					if err != nil {
						t.Errorf("error writing message: %v", err)
					}
				}
			})
			defer s.Close()

			url, _ := testutil.ConvertHTTPToWS(s.URL)
			cl := &Client{
				cfg: ClientConfig{
					host: url,
				},
			}

			if err := cl.Connect(); err != nil {
				t.Errorf("Error connecting to server: %v", err)
			}

			err := cl.setLastLedgerSequence(&tt.tx)

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(tt.expectedTx, tt.tx) {
					t.Errorf("Expected tx %v, but got %v", tt.expectedTx, tt.tx)
				}
			}

			cl.Disconnect()
		})
	}
}

func TestClient_checkAccountDeleteBlockers(t *testing.T) {
	tests := []struct {
		name           string
		address        types.Address
		serverMessages []map[string]any
		expectedErr    error
	}{
		{
			name:    "No blockers",
			address: "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
						"account":         "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
						"account_objects": []any{},
						"ledger_hash":     "4BC50C9B0D8515D3EAAE1E74B29A95804346C491EE1A95BF25E4AAB854A6A651",
						"ledger_index":    30,
						"validated":       true,
					},
				},
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &testutil.MockWebSocketServer{Msgs: tt.serverMessages}
			s := ws.TestWebSocketServer(func(c *websocket.Conn) {
				for _, m := range tt.serverMessages {
					err := c.WriteJSON(m)
					if err != nil {
						t.Errorf("error writing message: %v", err)
					}
				}
			})
			defer s.Close()

			url, _ := testutil.ConvertHTTPToWS(s.URL)
			cl := &Client{
				cfg: ClientConfig{
					host: url,
				},
			}

			if err := cl.Connect(); err != nil {
				t.Errorf("Error connecting to server: %v", err)
			}

			err := cl.checkAccountDeleteBlockers(tt.address)

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			cl.Disconnect()
		})
	}
}

func TestClient_setTransactionFlags(t *testing.T) {
	tests := []struct {
		name     string
		tx       transaction.FlatTransaction
		expected uint32
		wantErr  bool
	}{
		{
			name: "No flags set",
			tx: transaction.FlatTransaction{
				"TransactionType": string(transaction.PaymentTx),
			},
			expected: uint32(0),
			wantErr:  false,
		},
		{
			name: "Flags already set",
			tx: transaction.FlatTransaction{
				"TransactionType": string(transaction.PaymentTx),
				"Flags":           uint32(1),
			},
			expected: 1,
			wantErr:  false,
		},
		{
			name: "Missing TransactionType",
			tx: transaction.FlatTransaction{
				"Flags": uint32(1),
			},
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{}
			err := c.setTransactionFlags(&tt.tx)

			if (err != nil) != tt.wantErr {

				t.Errorf("setTransactionFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				flags, ok := tt.tx["Flags"]
				if !ok && tt.expected != 0 {
					t.Errorf("setTransactionFlags() got = %v (type %T), want %v (type %T)", flags, flags, tt.expected, tt.expected)
				}
			}
		})
	}
}
