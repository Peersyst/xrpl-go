package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestAMMCreateFlatten(t *testing.T) {
	s := AMMCreate{
		BaseTx: BaseTx{
			Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType: AMMCreateTx,
			Fee:             types.XRPCurrencyAmount(10),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Amount: types.XRPCurrencyAmount(100),
		Amount2: types.IssuedCurrencyAmount{
			Currency: "USD",
			Value:    "200",
			Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
		},
		TradingFee: 10,
	}

	flattened := s.Flatten()

	expected := `{
		"Account":         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType": "AMMCreate",
		"Fee":             "10",
		"Sequence":        1234,
		"SigningPubKey":   "ghijk",
		"TxnSignature":    "A1B2C3D4E5F6",
		"Amount":          "100",
		"Amount2":         {
			"currency": "USD",
			"value":    "200",
			"issuer":   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ"
		},
		"TradingFee":      10
	}`

	err := testutil.CompareFlattenAndExpected(flattened, []byte(expected))
	if err != nil {
		t.Error(err)
	}
}

func TestAMMCreateValidate(t *testing.T) {
	tests := []struct {
		name    string
		amm     AMMCreate
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid AMMCreate",
			amm: AMMCreate{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMCreateTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Amount: types.XRPCurrencyAmount(100),
				Amount2: types.IssuedCurrencyAmount{
					Currency: "USD",
					Value:    "200",
					Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
				},
				TradingFee: 10,
			},
			wantErr: false,
		},
		{
			name: "Invalid BaseTx for AMMCreate, missing Account",
			amm: AMMCreate{
				BaseTx: BaseTx{
					TransactionType: AMMCreateTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Amount: types.XRPCurrencyAmount(100),
				Amount2: types.IssuedCurrencyAmount{
					Currency: "USD",
					Value:    "200",
					Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
				},
				TradingFee: 10,
			},
			wantErr: true,
		},
		{
			name: "missing Amount",
			amm: AMMCreate{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMCreateTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Amount2: types.IssuedCurrencyAmount{
					Currency: "USD",
					Value:    "200",
					Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
				},
				TradingFee: 10,
			},
			wantErr: true,
		},
		{
			name: "invalid Amount value",
			amm: AMMCreate{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMCreateTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Amount: types.IssuedCurrencyAmount{
					Currency: "USD",
					Value:    "-100",
					Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
				},
				Amount2: types.IssuedCurrencyAmount{
					Currency: "USD",
					Value:    "200",
					Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
				},
				TradingFee: 10,
			},
			wantErr: true,
		},
		{
			name: "missing Amount2",
			amm: AMMCreate{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMCreateTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Amount:     types.XRPCurrencyAmount(100),
				TradingFee: 10,
			},
			wantErr: true,
		},
		{
			name: "invalid Amount2 value",
			amm: AMMCreate{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMCreateTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Amount: types.XRPCurrencyAmount(100),
				Amount2: types.IssuedCurrencyAmount{
					Currency: "USD",
					Value:    "-200",
					Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
				},
				TradingFee: 10,
			},
			wantErr: true,
		},
		{
			name: "trading fee too high",
			amm: AMMCreate{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: AMMCreateTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Amount: types.XRPCurrencyAmount(100),
				Amount2: types.IssuedCurrencyAmount{
					Currency: "USD",
					Value:    "200",
					Issuer:   "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
				},
				TradingFee: 2000,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.amm.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !valid {
				t.Errorf("Expected valid AMMCreate, got invalid")
			}
		})
	}
}

func TestAMMCreate_TxType(t *testing.T) {
	tx := &AMMCreate{}
	assert.Equal(t, AMMCreateTx, tx.TxType())
}
