package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestOfferDirectoryNode(t *testing.T) {
	var s LedgerObject = &DirectoryNode{
		Flags: 0,
		Indexes: []types.Hash256{
			"AD7EAE148287EF12D213A251015F86E6D4BD34B3C4A0A1ED9A17198373F908AD",
		},
		LedgerEntryType:   DirectoryNodeEntry,
		RootIndex:         "1BBEF97EDE88D40CEE2ADE6FEF121166AFE80D99EBADB01A4F069BA8FF484000",
		TakerGetsCurrency: "0000000000000000000000000000000000000000",
		TakerGetsIssuer:   "0000000000000000000000000000000000000000",
		TakerPaysCurrency: "0000000000000000000000004A50590000000000",
		TakerPaysIssuer:   "5BBC0F22F61D9224A110650CFE21CC0C4BE13098",
	}

	j := `{
	"Flags": 0,
	"Indexes": [
		"AD7EAE148287EF12D213A251015F86E6D4BD34B3C4A0A1ED9A17198373F908AD"
	],
	"LedgerEntryType": "DirectoryNode",
	"RootIndex": "1BBEF97EDE88D40CEE2ADE6FEF121166AFE80D99EBADB01A4F069BA8FF484000",
	"TakerGetsCurrency": "0000000000000000000000000000000000000000",
	"TakerGetsIssuer": "0000000000000000000000000000000000000000",
	"TakerPaysCurrency": "0000000000000000000000004A50590000000000",
	"TakerPaysIssuer": "5BBC0F22F61D9224A110650CFE21CC0C4BE13098"
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestOwnerDirectoryNode(t *testing.T) {
	var s LedgerObject = &DirectoryNode{
		Flags: 0,
		Indexes: []types.Hash256{
			"AD7EAE148287EF12D213A251015F86E6D4BD34B3C4A0A1ED9A17198373F908AD",
			"E83BBB58949A8303DF07172B16FB8EFBA66B9191F3836EC27A4568ED5997BAC5",
		},
		LedgerEntryType: DirectoryNodeEntry,
		Owner:           "rpR95n1iFkTqpoy1e878f4Z1pVHVtWKMNQ",
		RootIndex:       "193C591BF62482468422313F9D3274B5927CA80B4DD3707E42015DD609E39C94",
	}

	j := `{
	"Flags": 0,
	"Indexes": [
		"AD7EAE148287EF12D213A251015F86E6D4BD34B3C4A0A1ED9A17198373F908AD",
		"E83BBB58949A8303DF07172B16FB8EFBA66B9191F3836EC27A4568ED5997BAC5"
	],
	"LedgerEntryType": "DirectoryNode",
	"Owner": "rpR95n1iFkTqpoy1e878f4Z1pVHVtWKMNQ",
	"RootIndex": "193C591BF62482468422313F9D3274B5927CA80B4DD3707E42015DD609E39C94"
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
