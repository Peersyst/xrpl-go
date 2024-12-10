package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/utility"
	"github.com/Peersyst/xrpl-go/xrpl/rpc/testutil"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestJsonRpcClientCreation(t *testing.T) {

	t.Run("Set config with valid port + ip", func(t *testing.T) {

		cfg, _ := NewClientConfig("url")

		jsonRpcClient := NewClient(cfg)

		assert.Equal(t, &Client{Config: cfg}, jsonRpcClient)
	})
}

func TestCheckForError(t *testing.T) {

	t.Run("Error Response", func(t *testing.T) {

		json := `{
			"result": {
				"error": "ledgerIndexMalformed",
				"request": {
					"account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
					"command": "account_info",
					"ledger_index": "-",
					"strict": true
				},
				"status": "error"
			}
		}`

		b := io.NopCloser(bytes.NewReader([]byte(json)))
		res := &http.Response{
			StatusCode: 200, // error response still returns a 200
			Body:       b,
		}

		bodyBytes, err := CheckForError(res)
		assert.NotNil(t, bodyBytes)
		expError := &ClientError{ErrorString: "ledgerIndexMalformed"}
		assert.Equal(t, expError, err)
	})

	t.Run("Error Response with error code", func(t *testing.T) {

		json := "Null Method" // https://xrpl.org/error-formatting.html#universal-errors

		b := io.NopCloser(bytes.NewReader([]byte(json)))
		res := &http.Response{
			StatusCode: 400,
			Body:       b,
		}

		bodyBytes, err := CheckForError(res)
		assert.NotNil(t, bodyBytes)
		expErrpr := &ClientError{ErrorString: "Null Method"}
		assert.Equal(t, expErrpr, err)
	})

	t.Run("No error Response", func(t *testing.T) {

		json := `{
			"result": {
			  "account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			  "channels": [
				{
				  "account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				  "amount": "1000",
				  "balance": "0",
				  "channel_id": "C7F634794B79DB40E87179A9D1BF05D05797AE7E92DF8E93FD6656E8C4BE3AE7",
				  "destination_account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
				  "public_key": "aBR7mdD75Ycs8DRhMgQ4EMUEmBArF8SEh1hfjrT2V9DQTLNbJVqw",
				  "public_key_hex": "03CFD18E689434F032A4E84C63E2A3A6472D684EAF4FD52CA67742F3E24BAE81B2",
				  "settle_delay": 60
				}
			  ],
			  "ledger_hash": "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
			  "ledger_index": 71766343,
			  "status": "success",
			  "validated": true
			}
		  }`

		b := io.NopCloser(bytes.NewReader([]byte(json)))
		res := &http.Response{
			StatusCode: 200,
			Body:       b,
		}

		bodyBytes, err := CheckForError(res)

		assert.Nil(t, err)
		assert.NotNil(t, bodyBytes)
	})
}

func TestCreateRequest(t *testing.T) {
	t.Run("Create request", func(t *testing.T) {

		req := &account.ChannelsRequest{
			Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
			DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
			LedgerIndex:        common.Validated,
		}

		expetedBody := Request{
			Method: "account_channels",
			Params: [1]interface{}{req},
		}
		expectedRequestBytes, _ := jsoniter.Marshal(expetedBody)

		byteRequest, err := CreateRequest(req)

		assert.NoError(t, err)
		// assert bytes equal
		assert.Equal(t, expectedRequestBytes, byteRequest)
		// assert json equal
		assert.Equal(t, string(expectedRequestBytes), string(byteRequest))
	})
	t.Run("Create request - no parameters with using pointer declaration", func(t *testing.T) {

		var req *utility.PingRequest // params sent in as zero value struct

		expetedBody := Request{
			Method: "ping",
		}
		expectedRequestBytes, _ := jsoniter.Marshal(expetedBody)

		byteRequest, err := CreateRequest(req)

		assert.NoError(t, err)
		// assert bytes equal
		assert.Equal(t, expectedRequestBytes, byteRequest)
		// assert json equal
		assert.Equal(t, string(expectedRequestBytes), string(byteRequest))
	})

	t.Run("Create request - no parameters with struct initialisation", func(t *testing.T) {

		req := &utility.PingRequest{} // means params get set an empty object

		expetedBody := Request{
			Method: "ping",
		}
		expectedRequestBytes, _ := jsoniter.Marshal(expetedBody)

		byteRequest, err := CreateRequest(req)

		assert.NoError(t, err)
		// assert bytes equal
		assert.Equal(t, expectedRequestBytes, byteRequest)
		// assert json equal
		assert.Equal(t, string(expectedRequestBytes), string(byteRequest))
	})
}

func TestSendRequest(t *testing.T) {

	t.Run("SendRequest - Check headers and URL", func(t *testing.T) {

		req := &account.ChannelsRequest{
			Account: "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		}
		var capturedRequest *http.Request

		mc := &testutil.JSONRPCMockClient{}
		mc.DoFunc = func(req *http.Request) (*http.Response, error) {
			capturedRequest = req
			return testutil.MockResponse(`{}`, 200, mc)(req)
		}

		cfg, err := NewClientConfig("http://testnode/", WithHTTPClient(mc))
		assert.NoError(t, err)

		jsonRpcClient := NewClient(cfg)

		_, err = jsonRpcClient.SendRequest(req)

		assert.NotNil(t, capturedRequest)
		assert.NoError(t, err)
		assert.Equal(t, "POST", capturedRequest.Method)
		assert.Equal(t, "http://testnode/", capturedRequest.URL.String())
		assert.Equal(t, "application/json", capturedRequest.Header.Get("Content-Type"))
	})

	t.Run("SendRequest - sucessful response", func(t *testing.T) {

		req := &account.ChannelsRequest{
			Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
			DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
			LedgerIndex:        common.Validated,
		}

		response := `{
			"result": {
			  "account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			  "channels": [
				{
					"account":             "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					"amount":              "1000",
					"balance":             "0",
					"channel_id":          "C7F634794B79DB40E87179A9D1BF05D05797AE7E92DF8E93FD6656E8C4BE3AE7",
					"destination_account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					"public_key":          "aBR7mdD75Ycs8DRhMgQ4EMUEmBArF8SEh1hfjrT2V9DQTLNbJVqw",
					"public_key_hex":      "03CFD18E689434F032A4E84C63E2A3A6472D684EAF4FD52CA67742F3E24BAE81B2",
					"settle_delay":        60
				}
			  ],
			  "ledger_hash": "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
			  "ledger_index": 71766314,
			  "validated": true
			},
			"warning": "none",
			"warnings":
			[{
				"id": 1,
				"message": "message"
			}]
		  }`

		mc := &testutil.JSONRPCMockClient{}
		mc.DoFunc = testutil.MockResponse(response, 200, mc)

		cfg, err := NewClientConfig("http://testnode/", WithHTTPClient(mc))
		assert.NoError(t, err)

		jsonRpcClient := NewClient(cfg)

		xrplResponse, err := jsonRpcClient.SendRequest(req)

		expectedXrplResponse := &Response{
			Result: AnyJSON{
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
			Warning: "none",
			Warnings: []XRPLResponseWarning{{
				ID:      1,
				Message: "message",
			},
			},
		}

		var channelsResponse account.ChannelsResponse
		_ = xrplResponse.GetResult(&channelsResponse)

		expected := &account.ChannelsResponse{
			Account:     "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			LedgerIndex: 71766314,
			LedgerHash:  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
		}

		assert.NoError(t, err)

		assert.Equal(t, expectedXrplResponse, xrplResponse)

		assert.Equal(t, expected.Account, channelsResponse.Account)
		assert.Equal(t, expected.LedgerIndex, channelsResponse.LedgerIndex)
		assert.Equal(t, expected.LedgerHash, channelsResponse.LedgerHash)
	})

	t.Run("SendRequest - error response", func(t *testing.T) {

		req := &account.ChannelsRequest{
			Account: "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		}
		response := `{
			"result": {
				"error": "ledgerIndexMalformed",
				"request": {
					"account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
					"command": "account_info",
					"ledger_index": "-",
					"strict": true
				},
				"status": "error"
			}
		}`

		mc := &testutil.JSONRPCMockClient{}
		mc.DoFunc = testutil.MockResponse(response, 200, mc)

		cfg, err := NewClientConfig("http://testnode/", WithHTTPClient(mc))
		assert.NoError(t, err)

		jsonRpcClient := NewClient(cfg)

		_, err = jsonRpcClient.SendRequest(req)

		assert.EqualError(t, err, "ledgerIndexMalformed")
	})

	t.Run("SendRequest - 503 response", func(t *testing.T) {

		req := &account.ChannelsRequest{
			Account: "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		}
		response := `Service Unavailable`

		mc := &testutil.JSONRPCMockClient{}
		mc.DoFunc = func(req *http.Request) (*http.Response, error) {
			mc.RequestCount++
			return testutil.MockResponse(response, 503, mc)(req)
		}

		cfg, err := NewClientConfig("http://testnode/", WithHTTPClient(mc))
		assert.NoError(t, err)

		jsonRpcClient := NewClient(cfg)

		_, err = jsonRpcClient.SendRequest(req)

		// Check that 3 extra requests were made
		assert.Equal(t, 4, mc.RequestCount)
		assert.EqualError(t, err, "Server is overloaded, rate limit exceeded")

	})

	t.Run("SendRequest - 503 response sucessfully resolves", func(t *testing.T) {

		req := &account.ChannelsRequest{
			Account: "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		}
		sucessResponse := `{
			"result": {
			  "account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			  "ledger_hash": "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
			  "ledger_index": 71766343
				}
			}`

		mc := &testutil.JSONRPCMockClient{}
		mc.DoFunc = func(req *http.Request) (*http.Response, error) {
			if mc.RequestCount < 3 {
				// Return 503 response for the first three requests
				mc.RequestCount++
				return testutil.MockResponse(`Service Unavailable`, 503, mc)(req)
			}
			// Return 200 response for the fourth request
			return testutil.MockResponse(sucessResponse, 200, mc)(req)
		}

		cfg, err := NewClientConfig("http://testnode/", WithHTTPClient(mc))
		assert.NoError(t, err)

		jsonRpcClient := NewClient(cfg)

		xrplResponse, err := jsonRpcClient.SendRequest(req)

		var channelsResponse account.ChannelsResponse
		_ = xrplResponse.GetResult(&channelsResponse)

		expected := &account.ChannelsResponse{
			Account:     "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			LedgerIndex: 71766343,
			LedgerHash:  "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
		}

		// Check that only 2 extra requests were made
		assert.Equal(t, 3, mc.RequestCount)

		assert.NoError(t, err)
		assert.Equal(t, expected.Account, channelsResponse.Account)
		assert.Equal(t, expected.LedgerIndex, channelsResponse.LedgerIndex)
		assert.Equal(t, expected.LedgerHash, channelsResponse.LedgerHash)
	})

	t.Run("SendRequest - timeout", func(t *testing.T) {
		req := &account.ChannelsRequest{
			Account: "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		}

		mc := &testutil.JSONRPCMockClient{}
		mc.DoFunc = func(req *http.Request) (*http.Response, error) {
			// hit the timeout by not responding
			time.Sleep(time.Second * 5)
			return nil, errors.New("timeout")
		}

		cfg, err := NewClientConfig("http://testnode/", WithHTTPClient(mc))
		assert.NoError(t, err)

		jsonRpcClient := NewClient(cfg)

		_, err = jsonRpcClient.SendRequest(req)

		// Check that the expected timeout error occurred
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "timeout")
	})
}
