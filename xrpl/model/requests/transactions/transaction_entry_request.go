package transactions

import (
	"encoding/json"

	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
)

type TransactionEntryRequest struct {
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	TxHash      string                 `json:"tx_hash"`
}

func (*TransactionEntryRequest) Method() string {
	return "transaction_entry"
}

func (t *TransactionEntryRequest) UnmarshalJSON(data []byte) error {
	type terHelper struct {
		LedgerHash  common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage   `json:"ledger_index,omitempty"`
		TxHash      string            `json:"tx_hash"`
	}
	var h terHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*t = TransactionEntryRequest{
		LedgerHash: h.LedgerHash,
		TxHash:     h.TxHash,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	t.LedgerIndex = i
	return nil
}
