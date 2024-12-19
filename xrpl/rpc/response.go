package rpc

import (
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction"
	"github.com/mitchellh/mapstructure"
)

type Response struct {
	Result    AnyJSON               `json:"result"`
	Warning   string                `json:"warning,omitempty"`
	Warnings  []XRPLResponseWarning `json:"warnings,omitempty"`
	Forwarded bool                  `json:"forwarded,omitempty"`
}

type XRPLResponseWarning struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

type AnyJSON transaction.FlatTransaction

type APIWarning struct {
	ID      int         `json:"id"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (r Response) GetResult(v any) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json",
		Result: &v, DecodeHook: mapstructure.TextUnmarshallerHookFunc()})

	if err != nil {
		return err
	}
	err = dec.Decode(r.Result)
	if err != nil {
		return err
	}
	return nil
}

type XRPLResponse interface {
	GetResult(v any) error
}
