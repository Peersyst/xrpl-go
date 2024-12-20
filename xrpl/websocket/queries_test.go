package websocket

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/server"
	"github.com/Peersyst/xrpl-go/xrpl/websocket/testutil"
	"github.com/gorilla/websocket"
)

func TestWebsocketClient_GetServerInfo(t *testing.T) {
	tests := []struct {
		name           string
		serverMessages []map[string]any
		expected       *server.InfoResponse
		expectedErr    error
	}{
		{
			name: "Valid server info",
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
						"info": map[string]any{
							"build_version":     "1.9.4",
							"complete_ledgers":  "32570-62964740",
							"hostid":            "MIST",
							"load_factor":       float64(1),
							"peers":             float64(96),
							"pubkey_node":       "n9KUjqxCr5FKThSNXdzb7oqN8rYwScB2dUnNqxQxbEA17JkaWy5x",
							"server_state":      "full",
							"validation_quorum": float64(28),
						},
					},
				},
			},
			expected: &server.InfoResponse{
				Info: server.Info{
					BuildVersion:     "1.9.4",
					CompleteLedgers:  "32570-62964740",
					HostID:           "MIST",
					LoadFactor:       1,
					Peers:            96,
					PubkeyNode:       "n9KUjqxCr5FKThSNXdzb7oqN8rYwScB2dUnNqxQxbEA17JkaWy5x",
					ServerState:      "full",
					ValidationQuorum: 28,
				},
			},
			expectedErr: nil,
		},
		{
			name:           "Error response",
			serverMessages: []map[string]any{{"error": "Server error"}},
			expected:       nil,
			expectedErr:    errors.New("incorrect id"),
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

			result, err := cl.GetServerInfo(&server.InfoRequest{})

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("Expected %+v, but got %+v", tt.expected, result)
			}

			cl.Disconnect()
		})
	}
}

func TestGetAccountInfo(t *testing.T) {
	tests := []struct {
		name           string
		serverMessages []map[string]any
		expected       *account.InfoResponse
		expectedErr    error
	}{
		{
			name: "Successful response",
			serverMessages: []map[string]any{
				{

					"id": 1,
					"result": map[string]any{
						"account_data": map[string]any{
							"Account":           "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
							"Flags":             0,
							"LedgerEntryType":   "AccountRoot",
							"OwnerCount":        0,
							"PreviousTxnID":     "4294BEBE5B569A18C0A2702387C9B1E7146DC3A5850C1E87204951C6FDAA4C42",
							"PreviousTxnLgrSeq": 3,
							"Sequence":          6,
						},
						"validated": false,
					},
				},
			},
			expected: &account.InfoResponse{
				AccountData: ledger.AccountRoot{
					Account:           "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
					Flags:             0,
					LedgerEntryType:   "AccountRoot",
					OwnerCount:        0,
					PreviousTxnID:     "4294BEBE5B569A18C0A2702387C9B1E7146DC3A5850C1E87204951C6FDAA4C42",
					PreviousTxnLgrSeq: 3,
					Sequence:          6,
				},
				Validated: false,
			},
			expectedErr: nil,
		},
		{
			name:           "Error response",
			serverMessages: []map[string]any{{"error": "Account not found"}},
			expected:       nil,
			expectedErr:    errors.New("incorrect id"),
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

			result, err := cl.GetAccountInfo(&account.InfoRequest{
				Account: "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
			})

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("Expected %+v, but got %+v", tt.expected, result)
			}

			cl.Disconnect()
		})
	}
}

func TestGetAccountObjects(t *testing.T) {
	tests := []struct {
		name           string
		serverMessages []map[string]any
		expected       *account.ObjectsResponse
		expectedErr    error
	}{
		{
			name: "Successful response",
			serverMessages: []map[string]any{{
				"id": 1,
				"result": map[string]any{
					"account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
					"account_objects": []map[string]any{
						{
							"LedgerEntryType": "RippleState",
							"Balance": map[string]any{
								"currency": "USD",
								"issuer":   "rrrrrrrrrrrrrrrrrrrrBZbvji",
								"value":    "100",
							},
						},
					},
					"ledger_hash":  "4C99E5F63C0D0B1C2283B4F5DCE2239F80CE92E8B1A6AED1E110C198FC96E659",
					"ledger_index": 14380380,
					"validated":    true,
				},
			}},
			expected: &account.ObjectsResponse{
				Account: "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
				AccountObjects: []ledger.FlatLedgerObject{
					{
						"LedgerEntryType": "RippleState",
						"Balance": map[string]any{
							"currency": "USD",
							"issuer":   "rrrrrrrrrrrrrrrrrrrrBZbvji",
							"value":    "100",
						},
					},
				},
				LedgerHash:  "4C99E5F63C0D0B1C2283B4F5DCE2239F80CE92E8B1A6AED1E110C198FC96E659",
				LedgerIndex: 14380380,
				Validated:   true,
			},
			expectedErr: nil,
		},
		{
			name:           "Error response",
			serverMessages: []map[string]any{{"error": "Account not found"}},
			expected:       nil,
			expectedErr:    errors.New("incorrect id"),
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

			result, err := cl.GetAccountObjects(&account.ObjectsRequest{
				Account: "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
			})

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("Expected %+v, but got %+v", tt.expected, result)
			}

			cl.Disconnect()
		})
	}
}

func TestGetXrpBalance(t *testing.T) {
	tests := []struct {
		name           string
		serverMessages []map[string]any
		expected       string
		expectedErr    error
	}{
		{
			name: "Successful response",
			serverMessages: []map[string]any{{
				"id": 1,
				"result": map[string]any{
					"account_data": map[string]any{
						"Account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
						"Balance": "1000000000",
					},
					"validated": true,
				},
			}},
			expected:    "1000",
			expectedErr: nil,
		},
		{
			name:           "Error response",
			serverMessages: []map[string]any{{"error": "Account not found"}},
			expected:       "",
			expectedErr:    errors.New("incorrect id"),
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

			result, err := cl.GetXrpBalance("rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn")

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if tt.expected != result {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}

			cl.Disconnect()
		})
	}
}

func TestGetAccountLines(t *testing.T) {
	tests := []struct {
		name           string
		serverMessages []map[string]any
		expected       *account.LinesResponse
		expectedErr    error
	}{
		{
			name: "Successful response",
			serverMessages: []map[string]any{{
				"id": 1,
				"result": map[string]any{
					"account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
					"lines": []map[string]any{
						{
							"account":  "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
							"balance":  "10",
							"currency": "USD",
						},
					},
					"ledger_current_index": 14380380,
					"validated":            true,
				},
			}},
			expected: &account.LinesResponse{
				Account: "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
				Lines: []account.TrustLine{
					{
						Account:  "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
						Balance:  "10",
						Currency: "USD",
					},
				},
				LedgerCurrentIndex: 14380380,
			},
			expectedErr: nil,
		},
		{
			name:           "Error response",
			serverMessages: []map[string]any{{"error": "Account not found"}},
			expected:       nil,
			expectedErr:    errors.New("incorrect id"),
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

			result, err := cl.GetAccountLines(&account.LinesRequest{
				Account: "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
			})

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("Expected %+v, but got %+v", tt.expected, result)
			}

			cl.Disconnect()
		})
	}
}

func TestGetLedgerIndex(t *testing.T) {
	tests := []struct {
		name           string
		serverMessages []map[string]any
		expected       common.LedgerIndex
		expectedErr    error
	}{
		{
			name: "Successful response",
			serverMessages: []map[string]any{{
				"id": 1,
				"result": map[string]any{
					"ledger_index": 14380380,
					"validated":    true,
				},
			}},
			expected:    common.LedgerIndex(14380380),
			expectedErr: nil,
		},
		{
			name:           "Error response",
			serverMessages: []map[string]any{{"error": "Server error"}},
			expected:       common.LedgerIndex(0),
			expectedErr:    errors.New("incorrect id"),
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

			result, err := cl.GetLedgerIndex()

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("Expected %+v, but got %+v", tt.expected, result)
			}

			cl.Disconnect()
		})
	}
}
