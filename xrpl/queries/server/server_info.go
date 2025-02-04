package server

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	servertypes "github.com/Peersyst/xrpl-go/xrpl/queries/server/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
)

// ############################################################################
// Request
// ############################################################################

// The server_info command asks the server for a human-readable version of
// various information about the rippled server being queried.
type InfoRequest struct {
	common.BaseRequest
}

func (*InfoRequest) Method() string {
	return "server_info"
}

func (*InfoRequest) APIVersion() int {
	return version.RippledAPIV2
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
