package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/v1/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestAMMBidFlatten(t *testing.T) {
	s := AMMBid{
		BaseTx: BaseTx{
			Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType: AMMCreateTx,
			Fee:             types.XRPCurrencyAmount(10),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Asset:  ledger.Asset{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"},
		Asset2: ledger.Asset{Currency: "XRP"},
		BidMin: types.XRPCurrencyAmount(100),
		BidMax: types.IssuedCurrencyAmount{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ", Value: "200"},
		AuthAccounts: []ledger.AuthAccounts{
			{
				AuthAccount: ledger.AuthAccount{
					Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
				},
			},
			{
				AuthAccount: ledger.AuthAccount{
					Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
				}},
		},
	}

	flattened := s.Flatten()

	expected := `{
	"Account":         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
	"TransactionType": "AMMBid",
	"Fee":             "10",
	"Sequence":        1234,
	"SigningPubKey":   "ghijk",
	"TxnSignature":    "A1B2C3D4E5F6",
	"Asset": {
		"currency": "USD",
		"issuer": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"
	},
	"Asset2": {
		"currency": "XRP"
	},
	"BidMin": "100",
	"BidMax": {
		"currency": "USD",
		"issuer": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
		"value": "200"
	},
	"AuthAccounts": [
		{
			"AuthAccount": {
				"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ"
			}
		},
		{
			"AuthAccount": {
				"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ"
			}
		}
	]
}`

	err := testutil.CompareFlattenAndExpected(flattened, []byte(expected))
	if err != nil {
		t.Error(err)
	}
}

func TestAMMBid_Validate(t *testing.T) {
	tests := []struct {
		name    string
		ammBid  AMMBid
		wantErr bool
	}{
		{
			name: "pass - valid AMMBid",
			ammBid: AMMBid{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMBidTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Asset:  ledger.Asset{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"},
				Asset2: ledger.Asset{Currency: "XRP"},
				BidMin: types.XRPCurrencyAmount(100),
				BidMax: types.IssuedCurrencyAmount{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ", Value: "200"},
				AuthAccounts: []ledger.AuthAccounts{
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcE",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "fail - invalid BaseTx AMMBid, Account missing",
			ammBid: AMMBid{
				BaseTx: BaseTx{
					TransactionType: AMMBidTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Asset:  ledger.Asset{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"},
				Asset2: ledger.Asset{Currency: "XRP"},
				BidMin: types.XRPCurrencyAmount(100),
				BidMax: types.IssuedCurrencyAmount{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ", Value: "200"},
				AuthAccounts: []ledger.AuthAccounts{
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcE",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail - invalid AMMBid with more than 4 AuthAccounts",
			ammBid: AMMBid{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMBidTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Asset:  ledger.Asset{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"},
				Asset2: ledger.Asset{Currency: "XRP"},
				BidMin: types.XRPCurrencyAmount(100),
				BidMax: types.IssuedCurrencyAmount{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ", Value: "200"},
				AuthAccounts: []ledger.AuthAccounts{
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcE",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcA",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcB",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcC",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail - invalid AMMBid with more an AuthAccount with invalid address",
			ammBid: AMMBid{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMBidTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Asset:  ledger.Asset{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"},
				Asset2: ledger.Asset{Currency: "XRP"},
				BidMin: types.XRPCurrencyAmount(100),
				BidMax: types.IssuedCurrencyAmount{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ", Value: "200"},
				AuthAccounts: []ledger.AuthAccounts{
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcE",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcA",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "invalid",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail - invalid AMMBid with invalid Asset, currency empty",
			ammBid: AMMBid{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMBidTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Asset:  ledger.Asset{Currency: "", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"},
				Asset2: ledger.Asset{Currency: "XRP"},
				BidMin: types.XRPCurrencyAmount(100),
				BidMax: types.IssuedCurrencyAmount{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ", Value: "200"},
				AuthAccounts: []ledger.AuthAccounts{
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcE",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail - invalid AMMBid with invalid Asset2, issuer empty with currency non empty",
			ammBid: AMMBid{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMBidTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Asset:  ledger.Asset{Currency: "XRP"},
				Asset2: ledger.Asset{Currency: "USD", Issuer: ""},
				BidMin: types.XRPCurrencyAmount(100),
				BidMax: types.IssuedCurrencyAmount{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ", Value: "200"},
				AuthAccounts: []ledger.AuthAccounts{
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcE",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail - invalid AMMBid with invalid Asset and Asset2, two XRP assets",
			ammBid: AMMBid{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMBidTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Asset:  ledger.Asset{Currency: "XRP"},
				Asset2: ledger.Asset{Currency: "XRP"},
				BidMin: types.XRPCurrencyAmount(100),
				BidMax: types.IssuedCurrencyAmount{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ", Value: "200"},
				AuthAccounts: []ledger.AuthAccounts{
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcE",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail - invalid AMMBid with invalid BidMin, missing value and issuer",
			ammBid: AMMBid{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMBidTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Asset:  ledger.Asset{Currency: "XRP"},
				Asset2: ledger.Asset{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"},
				BidMin: types.IssuedCurrencyAmount{Currency: "USD"}, // missing value and issuer
				BidMax: types.IssuedCurrencyAmount{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ", Value: "200"},
				AuthAccounts: []ledger.AuthAccounts{
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcE",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail - invalid AMMBid with invalid BidMax, missing value and issuer",
			ammBid: AMMBid{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMBidTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Asset:  ledger.Asset{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"},
				Asset2: ledger.Asset{Currency: "XRP"},
				BidMin: types.IssuedCurrencyAmount{Currency: "USD", Issuer: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ", Value: "200"},
				BidMax: types.IssuedCurrencyAmount{Currency: "USD"}, // missing value and issuer
				AuthAccounts: []ledger.AuthAccounts{
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcZ",
						},
					},
					{
						AuthAccount: ledger.AuthAccount{
							Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcE",
						},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.ammBid.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("AMMBid.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if valid != !tt.wantErr {
				t.Errorf("AMMBid.Validate() = %v, want %v", valid, !tt.wantErr)
			}
		})
	}
}

func TestAMMBid_TxType(t *testing.T) {
	entry := &AMMBid{}
	assert.Equal(t, AMMBidTx, entry.TxType())
}
