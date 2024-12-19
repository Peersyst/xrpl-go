package types

import "github.com/Peersyst/xrpl-go/v1/xrpl/ledger-entry-types"

type State struct {
	Data            string                  `json:"data,omitempty"`
	LedgerEntryType ledger.EntryType        `json:",omitempty"`
	LedgerObject    ledger.FlatLedgerObject `json:"-"`
	Index           string                  `json:"index"`
}
