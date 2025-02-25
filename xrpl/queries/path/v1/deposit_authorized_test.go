package v1

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
)

func TestDepositAuthorizedRequest(t *testing.T) {
	s := DepositAuthorizedRequest{
		SourceAccount:      "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
		DestinationAccount: "rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8",
		LedgerIndex:        common.Validated,
	}

	j := `{
	"source_account": "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
	"destination_account": "rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8",
	"ledger_index": "validated"
}`

	if err := testutil.Serialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestDepositAuthorizedResponse(t *testing.T) {
	s := DepositAuthorizedResponse{
		DepositAuthorized:  true,
		DestinationAccount: "rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8",
		LedgerHash:         "BD03A10653ED9D77DCA859B7A735BF0580088A8F287FA2C5403E0A19C58EF322",
		LedgerIndex:        8,
		SourceAccount:      "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
		Validated:          true,
	}

	j := `{
	"deposit_authorized": true,
	"destination_account": "rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8",
	"ledger_hash": "BD03A10653ED9D77DCA859B7A735BF0580088A8F287FA2C5403E0A19C58EF322",
	"ledger_index": 8,
	"source_account": "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
	"validated": true
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
