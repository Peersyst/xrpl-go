package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	requests "github.com/Peersyst/xrpl-go/xrpl/queries/transactions"
	jsoniter "github.com/json-iterator/go"
)

type Client struct {
	Config *Config
}

type ClientError struct {
	ErrorString string
}

func (e *ClientError) Error() string {
	return e.ErrorString
}

var ErrIncorrectID = errors.New("incorrect id")

func NewClient(cfg *Config) *Client {
	return &Client{
		Config: cfg,
	}
}

// satisfy the Client interface
func (c *Client) SendRequest(reqParams XRPLRequest) (XRPLResponse, error) {

	err := reqParams.Validate()
	if err != nil {
		return nil, err
	}

	body, err := CreateRequest(reqParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.Config.URL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// add timeout context to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req = req.WithContext(ctx)

	req.Header = c.Config.Headers

	var response *http.Response

	response, err = c.Config.HTTPClient.Do(req)
	if err != nil || response == nil {
		return nil, err
	}

	// allow client to reuse persistent connection
	defer response.Body.Close()

	// Check for service unavailable response and retry if so
	if response.StatusCode == 503 {

		maxRetries := 3
		backoffDuration := 1 * time.Second

		for i := 0; i < maxRetries; i++ {
			time.Sleep(backoffDuration)

			// Make request again after waiting
			response, err = c.Config.HTTPClient.Do(req)
			if err != nil {
				return nil, err
			}

			if response.StatusCode != 503 {
				break
			}

			// Increase backoff duration for the next retry
			backoffDuration *= 2
		}

		if response.StatusCode == 503 {
			// Return service unavailable error here after retry 3 times
			return nil, &ClientError{ErrorString: "Server is overloaded, rate limit exceeded"}
		}

	}

	var jr Response
	jr, err = CheckForError(response)
	if err != nil {
		return nil, err
	}

	return &jr, nil
}

func (c *Client) SubmitTransactionBlob(txBlob string, failHard bool) (XRPLResponse, error) {
	submitRequest := &requests.SubmitRequest{
		TxBlob:   txBlob,
		FailHard: failHard,
	}

	response, err := c.SendRequest(submitRequest)

	return response, err
}

// CreateRequest formats the parameters and method name ready for sending request
// Params will have been serialised if required and added to request struct before being passed to this method
func CreateRequest(reqParams XRPLRequest) ([]byte, error) {

	var body Request

	body = Request{
		Method: reqParams.Method(),
		// each param object will have a struct with json serialising tags
		Params: [1]interface{}{reqParams},
	}

	// Omit the Params field if method doesn't require any
	paramBytes, err := jsoniter.Marshal(body.Params)
	if err != nil {
		return nil, err
	}
	paramString := string(paramBytes)
	if strings.Compare(paramString, "[{}]") == 0 {
		// need to remove params field from the body if it is empty
		body = Request{
			Method: reqParams.Method(),
		}

		jsonBytes, err := jsoniter.Marshal(body)
		if err != nil {
			return nil, err
		}

		return jsonBytes, nil
	}

	jsonBytes, err := jsoniter.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON-RPC request for method %s with parameters %+v: %w", reqParams.Method(), reqParams, err)
	}

	return jsonBytes, nil
}

// CheckForError reads the http response and formats the error if it exists
func CheckForError(res *http.Response) (Response, error) {

	var jr Response

	b, err := io.ReadAll(res.Body)
	if err != nil || b == nil {
		return jr, err
	}

	// In case a different error code is returned
	if res.StatusCode != 200 {
		return jr, &ClientError{ErrorString: string(b)}
	}

	jDec := json.NewDecoder(bytes.NewReader(b))
	jDec.UseNumber()
	err = jDec.Decode(&jr)
	if err != nil {
		return jr, err
	}

	// result will have 'error' if error response
	if _, ok := jr.Result["error"]; ok {
		return jr, &ClientError{ErrorString: jr.Result["error"].(string)}
	}

	return jr, nil
}
