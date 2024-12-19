package types

import (
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

const (
	Burnable     NFTokenFlag = 0x0001
	OnlyXRP      NFTokenFlag = 0x0002
	Transferable NFTokenFlag = 0x0008
	ReservedFlag NFTokenFlag = 0x8000
)

type NFTokenFlag uint32

type NFT struct {
	Flags        NFTokenFlag `json:",omitempty"`
	Issuer       types.Address
	NFTokenID    types.NFTokenID
	NFTokenTaxon uint
	URI          types.NFTokenURI `json:",omitempty"`
	NFTSerial    uint             `json:"nft_serial"`
}
