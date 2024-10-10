package transaction

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestAMMCreateTransaction(t *testing.T) {
	s := AMMCreate{
		BaseTx: BaseTx{
			Account:         "abcdef",
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
	}

	j := `{
	"Account": "abcdef",
	"TransactionType": "AMMCreate",
	"Fee": "1",
	"Sequence": 1234,
	"SigningPubKey": "ghijk",
	"TxnSignature": "A1B2C3D4E5F6",
	"Amount": "100",
	"Amount2": {
		"issuer": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcQ",
		"currency": "USD",
		"value": "200"
	},
	"TradingFee": 10
}`

	if err := testutil.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

	tx, err := UnmarshalTx(json.RawMessage(j))
	if err != nil {
		t.Errorf("UnmarshalTx error: %s", err.Error())
	}
	if !reflect.DeepEqual(tx, &s) {
		t.Error("UnmarshalTx result differs from expected")
	}
}
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
					Account:         "abcdef",
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
			name: "missing Amount",
			amm: AMMCreate{
				BaseTx: BaseTx{
					Account:         "abcdef",
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
			errMsg:  "missing field Amount",
		},
		{
			name: "invalid Amount value",
			amm: AMMCreate{
				BaseTx: BaseTx{
					Account:         "abcdef",
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
			errMsg:  "invalid field Amount",
		},
		{
			name: "missing Amount2",
			amm: AMMCreate{
				BaseTx: BaseTx{
					Account:         "abcdef",
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
			errMsg:  "missing field Amount2",
		},
		{
			name: "invalid Amount2 value",
			amm: AMMCreate{
				BaseTx: BaseTx{
					Account:         "abcdef",
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
			errMsg:  "invalid field Amount2",
		},
		{
			name: "trading fee too high",
			amm: AMMCreate{
				BaseTx: BaseTx{
					Account:         "abcdef",
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
			errMsg:  fmt.Sprintf("trading fee is too high, max value is %d", AMM_MAX_TRADING_FEE),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.amm.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("Validate() error = %v, expected error message %v", err, tt.errMsg)
			}
			if !tt.wantErr && !valid {
				t.Errorf("Expected valid AMMCreate, got invalid")
			}
		})
	}
}
