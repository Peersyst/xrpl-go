package rpc

type Request struct {
	Method string         `json:"method"`
	Params [1]interface{} `json:"params,omitempty"`
}

type APIVersionRequest interface {
	APIVersion() int
	SetAPIVersion(apiVersion int)
}

type XRPLRequest interface {
	APIVersionRequest
	Method() string
	Validate() error
}
