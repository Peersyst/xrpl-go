package websocket

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/client"
	"github.com/mitchellh/mapstructure"
)

var _ client.XRPLResponse = (*WebSocketClientXrplResponse)(nil)

type ErrorWebsocketClientXrplResponse struct {
	Type    string
	Request map[string]any
}

func (e *ErrorWebsocketClientXrplResponse) Error() string {
	return e.Type
}

type WebSocketClientXrplResponse struct {
	ID        int                          `json:"id"`
	Status    string                       `json:"status"`
	Type      string                       `json:"type"`
	Error     string                       `json:"error,omitempty"`
	Result    map[string]any               `json:"result,omitempty"`
	Value     map[string]any               `json:"value,omitempty"`
	Warning   string                       `json:"warning,omitempty"`
	Warnings  []client.XRPLResponseWarning `json:"warnings,omitempty"`
	Forwarded bool                         `json:"forwarded,omitempty"`
}

func (r *WebSocketClientXrplResponse) GetResult(v any) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &v, DecodeHook: mapstructure.TextUnmarshallerHookFunc()})
	if err != nil {
		return err
	}
	err = dec.Decode(r.Result)
	if err != nil {
		fmt.Println("aqui", err)
		return err
	}
	return nil
}

func (r *WebSocketClientXrplResponse) CheckError() error {
	if r.Error != "" {
		return &ErrorWebsocketClientXrplResponse{
			Type:    r.Error,
			Request: r.Value,
		}
	}
	return nil
}