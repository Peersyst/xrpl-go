package types

import "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"

type AccountData struct {
	ledger.AccountRoot
	SignerLists []ledger.SignerList `json:"signer_lists,omitempty"`
}
