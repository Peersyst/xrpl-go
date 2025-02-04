package v1

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ############################################################################
// Request
// ############################################################################

// The random command provides a random number to be used as a source of
// entropy for random number generation by clients.
type RandomRequest struct {
	common.BaseRequest
}

func (*RandomRequest) Method() string {
	return "random"
}

func (*RandomRequest) APIVersion() int {
	return version.RippledAPIV1
}

func (*RandomRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the random method.
type RandomResponse struct {
	Random types.Hash256 `json:"random"`
}
