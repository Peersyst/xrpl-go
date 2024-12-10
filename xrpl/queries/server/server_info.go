package server

import (
	servertypes "github.com/Peersyst/xrpl-go/xrpl/queries/server/types"
)

// ############################################################################
// Request
// ############################################################################

// The server_info command asks the server for a human-readable version of
// various information about the rippled server being queried.
type InfoRequest struct {
}

func (*InfoRequest) Method() string {
	return "server_info"
}

// TODO: Implement V2
func (*InfoRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the server_info method.
type InfoResponse struct {
	Info servertypes.Info `json:"info"`
}
