package account

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

type AccountTransactionsResponse struct {
	Account        types.Address        `json:"account"`
	LedgerIndexMin common.LedgerIndex   `json:"ledger_index_min"`
	LedgerIndexMax common.LedgerIndex   `json:"ledger_index_max"`
	Limit          int                  `json:"limit"`
	Marker         any                  `json:"marker,omitempty"`
	Transactions   []AccountTransaction `json:"transactions"`
	Validated      bool                 `json:"validated"`
}
