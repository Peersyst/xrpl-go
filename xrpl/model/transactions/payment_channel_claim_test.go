package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/test"
)

func TestPaymentChannelClaimTx(t *testing.T) {
	s := PaymentChannelClaim{
		BaseTx: BaseTx{
			Account:         "abc",
			TransactionType: PaymentChannelClaimTx,
		},
		Channel:   "C1AE6DDDEEC05CF2978C0BAD6FE302948E9533691DC749DCDD3B9E5992CA6198",
		Balance:   types.XRPCurrencyAmount(1000000),
		Amount:    types.XRPCurrencyAmount(1000000),
		Signature: "30440220718D264EF05CAED7C781FF6DE298DCAC68D002562C9BF3A07C1E721B420C0DAB02203A5A4779EF4D2CCC7BC3EF886676D803A9981B928D3B8ACA483B80ECA3CD7B9B",
		PublicKey: "32D2471DB72B27E3310F355BB33E339BF26F8392D5A93D3BC0FC3B566612DA0F0A",
	}

	j := `{
	"Account": "abc",
	"TransactionType": "PaymentChannelClaim",
	"Channel": "C1AE6DDDEEC05CF2978C0BAD6FE302948E9533691DC749DCDD3B9E5992CA6198",
	"Balance": "1000000",
	"Amount": "1000000",
	"Signature": "30440220718D264EF05CAED7C781FF6DE298DCAC68D002562C9BF3A07C1E721B420C0DAB02203A5A4779EF4D2CCC7BC3EF886676D803A9981B928D3B8ACA483B80ECA3CD7B9B",
	"PublicKey": "32D2471DB72B27E3310F355BB33E339BF26F8392D5A93D3BC0FC3B566612DA0F0A"
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
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

func TestPaymentChannelClaimFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*PaymentChannelClaim)
		expected uint
	}{
		{
			name: "SetRenewFlag",
			setter: func(p *PaymentChannelClaim) {
				p.SetRenewFlag()
			},
			expected: tfRenew,
		},
		{
			name: "SetCloseFlag",
			setter: func(p *PaymentChannelClaim) {
				p.SetCloseFlag()
			},
			expected: tfClose,
		},
		{
			name: "SetRenewFlag and SetCloseFlag",
			setter: func(p *PaymentChannelClaim) {
				p.SetRenewFlag()
				p.SetCloseFlag()
			},
			expected: tfRenew | tfClose,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaymentChannelClaim{}
			tt.setter(p)
			if p.Flags != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, p.Flags)
			}
		})
	}
}
