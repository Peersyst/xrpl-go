package ledger

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/require"
)

func TestCredential_EntryType(t *testing.T) {
	credential := &Credential{}
	require.Equal(t, credential.EntryType(), CredentialEntry)
}

func TestCredential_SetLsfAccepted(t *testing.T) {
	credential := &Credential{}
	credential.SetLsfAccepted()
	require.Equal(t, credential.Flags, lsfAccepted)
}

func TestCredential_Flatten(t *testing.T) {
	credential := &Credential{
		Index:             types.Hash256("A738A1E6E8505E1FC77BBB9FEF84FF9A9C609F2739E0F9573CDD6367100A0AA9"),
		LedgerEntryType:   CredentialEntry,
		Flags:             lsfAccepted,
		CredentialType:    types.CredentialType("6D795F63726564656E7469616C"),
		Subject:           types.Address("rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8"),
		Issuer:            types.Address("ra5nK24KXen9AHvsdFTKHSANinZseWnPcX"),
		IssuerNode:        "0000000000000000",
		PreviousTxnID:     types.Hash256("8089451B193AAD110ACED3D62BE79BB523658545E6EE8B7BB0BE573FED9BCBFB"),
		PreviousTxnLgrSeq: 234644,
		SubjectNode:       "0000000000000000",
		URI:               "987654321",
	}

	json := `{
	"index": "A738A1E6E8505E1FC77BBB9FEF84FF9A9C609F2739E0F9573CDD6367100A0AA9",
	"LedgerEntryType": "Credential",
	"Flags": 65536,
	"CredentialType": "6D795F63726564656E7469616C",
	"Issuer": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
	"IssuerNode": "0000000000000000",
	"PreviousTxnID": "8089451B193AAD110ACED3D62BE79BB523658545E6EE8B7BB0BE573FED9BCBFB",
	"PreviousTxnLgrSeq": 234644,
	"Subject": "rsUiUMpnrgxQp24dJYZDhmV4bE3aBtQyt8",
	"SubjectNode": "0000000000000000",
	"URI": "987654321"
}`

	if err := testutil.SerializeAndDeserialize(t, credential, json); err != nil {
		t.Error(err)
	}
}
