package path

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	pathtypes "github.com/Peersyst/xrpl-go/xrpl/queries/path/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type SubCommand string

const (
	Create SubCommand = "create"
	Close  SubCommand = "close"
	Status SubCommand = "status"
)

// ############################################################################
// Create Request
// ############################################################################

// Start sending pathfinding information.
type FindCreateRequest struct {
	common.BaseRequest
	Subcommand         SubCommand             `json:"subcommand"`
	SourceAccount      types.Address          `json:"source_account,omitempty"`
	DestinationAccount types.Address          `json:"destination_account,omitempty"`
	DestinationAmount  types.CurrencyAmount   `json:"destination_amount,omitempty"`
	SendMax            types.CurrencyAmount   `json:"send_max,omitempty"`
	Paths              []transaction.PathStep `json:"paths,omitempty"`
}

func (*FindCreateRequest) Method() string {
	return "path_find"
}

func (*FindCreateRequest) APIVersion() int {
	return version.RippledAPIV1
}

// TODO: Implement V2
func (*FindCreateRequest) Validate() error {
	return nil
}

// ############################################################################
// Close Request
// ############################################################################

// Stop sending pathfinding information.
type FindCloseRequest struct {
	common.BaseRequest
	Subcommand SubCommand `json:"subcommand"`
}

func (*FindCloseRequest) Method() string {
	return "path_find"
}

func (*FindCloseRequest) APIVersion() int {
	return version.RippledAPIV2
}

// TODO: Implement V2
func (*FindCloseRequest) Validate() error {
	return nil
}

// ############################################################################
// Status Request
// ############################################################################

// Get the information of the currently-open pathfinding request.
type FindStatusRequest struct {
	common.BaseRequest
	Subcommand SubCommand `json:"subcommand"`
}

func (*FindStatusRequest) Method() string {
	return "path_find"
}

func (*FindStatusRequest) APIVersion() int {
	return version.RippledAPIV2
}

// TODO: Implement V2
func (*FindStatusRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// TODO: Add ID handling (v2)

// The expected response from the path_find method.
type FindResponse struct {
	Alternatives       []pathtypes.Alternative `json:"alternatives"`
	DestinationAccount types.Address           `json:"destination_account"`
	// DestinationAmount  types.CurrencyAmount    `json:"destination_amount"`
	DestinationAmount any           `json:"destination_amount"`
	SourceAccount     types.Address `json:"source_account"`
	FullReply         bool          `json:"full_reply"`
	Closed            bool          `json:"closed,omitempty"`
	Status            bool          `json:"status,omitempty"`
}
