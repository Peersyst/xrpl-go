package websocket

import (
	"github.com/mitchellh/mapstructure"
)

type ResponseWarning struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

type ErrorWebsocketClientXrplResponse struct {
	Type    string
	Request map[string]any
}

func (e *ErrorWebsocketClientXrplResponse) Error() string {
	return e.Type
}

type ClientResponse struct {
	ID        int               `json:"id"`
	Status    string            `json:"status"`
	Type      string            `json:"type"`
	Error     string            `json:"error,omitempty"`
	Result    map[string]any    `json:"result,omitempty"`
	Value     map[string]any    `json:"value,omitempty"`
	Warning   string            `json:"warning,omitempty"`
	Warnings  []ResponseWarning `json:"warnings,omitempty"`
	Forwarded bool              `json:"forwarded,omitempty"`
}

func (r *ClientResponse) GetResult(v any) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &v, DecodeHook: mapstructure.TextUnmarshallerHookFunc()})
	if err != nil {
		return err
	}
	err = dec.Decode(r.Result)
	if err != nil {
		return err
	}
	return nil
}

func (r *ClientResponse) CheckError() error {
	if r.Error != "" {
		return &ErrorWebsocketClientXrplResponse{
			Type:    r.Error,
			Request: r.Value,
		}
	}
	return nil
}
