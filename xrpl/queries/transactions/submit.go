package transaction

import (
	"errors"

	"github.com/Peersyst/xrpl-go/v1/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction"
)

var (
	ErrNoTxBlob = errors.New("no TxBlob defined")
)

// ############################################################################
// Request
// ############################################################################

// The submit method applies a transaction and sends it to the network to be
// confirmed and included in future ledgers.
type SubmitRequest struct {
	TxBlob   string `json:"tx_blob"`
	FailHard bool   `json:"fail_hard,omitempty"`
}

func (*SubmitRequest) Method() string {
	return "submit"
}

func (req *SubmitRequest) Validate() error {
	if req.TxBlob == "" {
		return ErrNoTxBlob
	}
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the submit method.
type SubmitResponse struct {
	EngineResult             string                      `json:"engine_result"`
	EngineResultCode         int                         `json:"engine_result_code"`
	EngineResultMessage      string                      `json:"engine_result_message"`
	TxBlob                   string                      `json:"tx_blob"`
	Tx                       transaction.FlatTransaction `json:"tx_json"`
	Accepted                 bool                        `json:"accepted"`
	AccountSequenceAvailable uint                        `json:"account_sequence_available"`
	AccountSequenceNext      uint                        `json:"account_sequence_next"`
	Applied                  bool                        `json:"applied"`
	Broadcast                bool                        `json:"broadcast"`
	Kept                     bool                        `json:"kept"`
	Queued                   bool                        `json:"queued"`
	OpenLedgerCost           string                      `json:"open_ledger_cost"`
	ValidatedLedgerIndex     common.LedgerIndex          `json:"validated_ledger_index"`
}
