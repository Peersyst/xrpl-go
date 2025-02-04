package transactions

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
)

// ############################################################################
// Request
// ############################################################################

// The submit_multisigned command applies a multi-signed transaction and sends
// it to the network to be included in future ledgers.
type SubmitMultisignedRequest struct {
	common.BaseRequest
	Tx       transaction.FlatTransaction `json:"tx_json"`
	FailHard bool                        `json:"fail_hard"`
}

func (*SubmitMultisignedRequest) Method() string {
	return "submit_multisigned"
}

func (*SubmitMultisignedRequest) APIVersion() int {
	return version.RippledAPIV2
}

func (*SubmitMultisignedRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the submit_multisigned method.
type SubmitMultisignedResponse struct {
	EngineResult        string                      `json:"engine_result"`
	EngineResultCode    int                         `json:"engine_result_code"`
	EngineResultMessage string                      `json:"engine_result_message"`
	TxBlob              string                      `json:"tx_blob"`
	Tx                  transaction.FlatTransaction `json:"tx_json"`
}
