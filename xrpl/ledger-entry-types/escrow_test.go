package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
	"github.com/stretchr/testify/require"
)

func TestEscrow(t *testing.T) {
	var s Object = &Escrow{
		LedgerEntryType:   EscrowEntry,
		Flags:             0,
		Account:           "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		Amount:            types.XRPCurrencyAmount(10000),
		CancelAfter:       545440232,
		Condition:         "A0258020A82A88B2DF843A54F58772E4A3861866ECDB4157645DD9AE528C1D3AEEDABAB6810120",
		Destination:       "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		DestinationNode:   "0000000000000000",
		DestinationTag:    23480,
		FinishAfter:       545354132,
		OwnerNode:         "0000000000000000",
		PreviousTxnID:     "C44F2EB84196B9AD820313DBEBA6316A15C9A2D35787579ED172B87A30131DA7",
		PreviousTxnLgrSeq: 28991004,
		SourceTag:         11747,
	}

	j := `{
	"LedgerEntryType": "Escrow",
	"Flags": 0,
	"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
	"Amount": "10000",
	"CancelAfter": 545440232,
	"Condition": "A0258020A82A88B2DF843A54F58772E4A3861866ECDB4157645DD9AE528C1D3AEEDABAB6810120",
	"Destination": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
	"DestinationNode": "0000000000000000",
	"DestinationTag": 23480,
	"FinishAfter": 545354132,
	"OwnerNode": "0000000000000000",
	"PreviousTxnID": "C44F2EB84196B9AD820313DBEBA6316A15C9A2D35787579ED172B87A30131DA7",
	"PreviousTxnLgrSeq": 28991004,
	"SourceTag": 11747
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestEscrow_EntryType(t *testing.T) {
	s := &Escrow{}
	require.Equal(t, s.EntryType(), EscrowEntry)
}
