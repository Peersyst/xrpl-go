package v1

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	accounttypes "github.com/Peersyst/xrpl-go/xrpl/queries/account/types"
	typesv1 "github.com/Peersyst/xrpl-go/xrpl/queries/account/v1/types"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestAccountInfoRequest(t *testing.T) {
	s := InfoRequest{
		Account:     "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
		LedgerIndex: common.Closed,
		Queue:       true,
		SignerLists: false,
		Strict:      true,
	}

	// SignerLists assigned to default, omitted due to omitempty
	j := `{
	"account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
	"ledger_index": "closed",
	"queue": true,
	"strict": true
}`
	if err := testutil.Serialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestAccountInfoResponse(t *testing.T) {
	s := InfoResponse{
		AccountData: typesv1.AccountData{
			AccountRoot: ledger.AccountRoot{
				Account:           "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
				Balance:           types.XRPCurrencyAmount(999999999960),
				Flags:             8388608,
				LedgerEntryType:   ledger.AccountRootEntry,
				OwnerCount:        0,
				PreviousTxnID:     "4294BEBE5B569A18C0A2702387C9B1E7146DC3A5850C1E87204951C6FDAA4C42",
				PreviousTxnLgrSeq: 3,
				Sequence:          6,
			},
			SignerLists: []ledger.SignerList{
				{
					LedgerEntryType:   ledger.SignerListEntry,
					Flags:             0,
					PreviousTxnID:     "4294BEBE5B569A18C0A2702387C9B1E7146DC3A5850C1E87204951C6FDAA4C42",
					PreviousTxnLgrSeq: 3,
					OwnerNode:         "4294BEBE5B569A18C0A2702387C9B1E7146DC3A5850C1E87204951C6FDAA4C42",
					SignerListID:      1,
					SignerQuorum:      3,
					SignerEntries: []ledger.SignerEntryWrapper{
						{
							SignerEntry: ledger.SignerEntry{
								Account:      "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
								SignerWeight: 1,
							},
						},
					},
				},
			},
		},
		LedgerCurrentIndex: 4,
		QueueData: accounttypes.QueueData{
			TxnCount:           5,
			AuthChangeQueued:   true,
			LowestSequence:     6,
			HighestSequence:    10,
			MaxSpendDropsTotal: types.XRPCurrencyAmount(500),
			Transactions: []accounttypes.QueueTransaction{
				{
					AuthChange:    false,
					Fee:           types.XRPCurrencyAmount(100),
					FeeLevel:      types.XRPCurrencyAmount(2560),
					MaxSpendDrops: types.XRPCurrencyAmount(100),
					Seq:           6,
				},
				{
					AuthChange:    true,
					Fee:           types.XRPCurrencyAmount(100),
					FeeLevel:      types.XRPCurrencyAmount(2560),
					MaxSpendDrops: types.XRPCurrencyAmount(100),
					Seq:           10,
				},
			},
		},
		Validated: false,
	}

	j := `{
	"account_data": {
		"Flags": 8388608,
		"LedgerEntryType": "AccountRoot",
		"Account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
		"Balance": "999999999960",
		"OwnerCount": 0,
		"PreviousTxnID": "4294BEBE5B569A18C0A2702387C9B1E7146DC3A5850C1E87204951C6FDAA4C42",
		"PreviousTxnLgrSeq": 3,
		"Sequence": 6,
		"signer_lists": [
			{
				"LedgerEntryType": "SignerList",
				"PreviousTxnID": "4294BEBE5B569A18C0A2702387C9B1E7146DC3A5850C1E87204951C6FDAA4C42",
				"PreviousTxnLgrSeq": 3,
				"OwnerNode": "4294BEBE5B569A18C0A2702387C9B1E7146DC3A5850C1E87204951C6FDAA4C42",
				"SignerEntries": [
					{
						"SignerEntry": {
							"Account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
							"SignerWeight": 1
						}
					}
				],
				"SignerListID": 1,
				"SignerQuorum": 3,
				"Flags": 0
			}
		]
	},
	"ledger_current_index": 4,
	"queue_data": {
		"txn_count": 5,
		"auth_change_queued": true,
		"lowest_sequence": 6,
		"highest_sequence": 10,
		"max_spend_drops_total": "500",
		"transactions": [
			{
				"auth_change": false,
				"fee": "100",
				"fee_level": "2560",
				"max_spend_drops": "100",
				"seq": 6
			},
			{
				"auth_change": true,
				"fee": "100",
				"fee_level": "2560",
				"max_spend_drops": "100",
				"seq": 10
			}
		]
	},
	"validated": false
}`
	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
