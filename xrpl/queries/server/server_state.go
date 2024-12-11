package server

import (
	servertypes "github.com/Peersyst/xrpl-go/xrpl/queries/server/types"
)

// ############################################################################
// Request
// ############################################################################

// The server_state command asks the server for various machine-readable
// information about the rippled server's current state. The response is almost
// the same as the server_info method, but uses units that are easier to process
// instead of easier to read.
type StateRequest struct {
}

func (*StateRequest) Method() string {
	return "server_state"
}

// TODO: Implement V2
func (*StateRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the server_state method.
type StateResponse struct {
	State servertypes.State `json:"state"`
}