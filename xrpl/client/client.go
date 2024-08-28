package client

type Client interface {
	SendRequest(req XRPLRequest) (XRPLResponse, error)
	SubmitTransactionBlob(txBlob string, failHard bool) (XRPLResponse, error)
}

type XRPLClient struct {
	client  Client
	Account Account
}

type XRPLRequest interface {
	Method() string
	Validate() error
}

type XRPLResponse interface {
	GetResult(v any) error
}

type XRPLResponseWarning struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func NewXRPLClient(cl Client) *XRPLClient {
	return &XRPLClient{
		client:  cl,
		Account: &accountImpl{client: cl},
	}
}

func (c *XRPLClient) Client() Client {
	return c.client
}
