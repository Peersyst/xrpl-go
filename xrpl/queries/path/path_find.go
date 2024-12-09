package path

import (
	pathtypes "github.com/Peersyst/xrpl-go/xrpl/queries/path/types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type SubCommand string

const (
	Create SubCommand = "create"
	Close  SubCommand = "close"
	Status SubCommand = "status"
)

type FindCreateRequest struct {
	Subcommand         SubCommand             `json:"subcommand"`
	SourceAccount      types.Address          `json:"source_account,omitempty"`
	DestinationAccount types.Address          `json:"destination_account,omitempty"`
	DestinationAmount  types.CurrencyAmount   `json:"destination_amount,omitempty"`
	SendMax            types.CurrencyAmount   `json:"send_max,omitempty"`
	Paths              []transaction.PathStep `json:"paths,omitempty"`
}

type FindCloseRequest struct {
	Subcommand SubCommand `json:"subcommand"`
}

type FindStatusRequest struct {
	Subcommand SubCommand `json:"subcommand"`
}

func (*FindCreateRequest) Method() string {
	return "path_find"
}

func (*FindCloseRequest) Method() string {
	return "path_find"
}

func (*FindStatusRequest) Method() string {
	return "path_find"
}

// TODO: Add ID handling (v2)
type FindResponse struct {
	Alternatives       []pathtypes.Alternative        `json:"alternatives"`
	DestinationAccount types.Address        `json:"destination_account"`
	DestinationAmount  types.CurrencyAmount `json:"destination_amount"`
	SourceAccount      types.Address        `json:"source_account"`
	FullReply          bool                 `json:"full_reply"`
	Closed             bool                 `json:"closed,omitempty"`
	Status             bool                 `json:"status,omitempty"`
}
