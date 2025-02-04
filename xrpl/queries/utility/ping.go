package utility

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
)

// ############################################################################
// Request
// ############################################################################

// The ping command returns an acknowledgement, so that clients can test the
// connection status and latency.
type PingRequest struct {
	common.BaseRequest
}

func (*PingRequest) Method() string {
	return "ping"
}

func (*PingRequest) APIVersion() int {
	return version.RippledAPIV2
}

func (*PingRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the ping method.
type PingResponse struct {
	Role      string `json:"role,omitempty"`
	Unlimited bool   `json:"unlimited,omitempty"`
}
