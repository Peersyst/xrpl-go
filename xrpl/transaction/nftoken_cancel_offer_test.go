package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestNFTokenCancelOffer_TxType(t *testing.T) {
	tx := &NFTokenCancelOffer{}
	assert.Equal(t, NFTokenCancelOfferTx, tx.TxType())
}

func TestNFTokenCancelOffer_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		tx       *NFTokenCancelOffer
		expected string
	}{
		{
			name: "Empty NFTokenOffers",
			tx: &NFTokenCancelOffer{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: NFTokenCancelOfferTx,
				},
				NFTokenOffers: []types.NFTokenID{},
			},
			expected: `{
				"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
				"TransactionType": "NFTokenCancelOffer"
			}`,
		},
		{
			name: "With NFTokenOffers",
			tx: &NFTokenCancelOffer{
				BaseTx: BaseTx{
					Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
				},
				NFTokenOffers: []types.NFTokenID{
					"9C92E061381C1EF37A8CDE0E8FC35188BFC30B1883825042A64309AC09F4C36D",
				},
			},
			expected: `{
				"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
				"TransactionType": "NFTokenCancelOffer",
				"NFTokenOffers": [
					"9C92E061381C1EF37A8CDE0E8FC35188BFC30B1883825042A64309AC09F4C36D"
				]
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.tx.Flatten(), []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestNFTokenCancelOffer_Validate(t *testing.T) {
	tests := []struct {
		name      string
		tx        *NFTokenCancelOffer
		wantValid bool
		wantErr   bool
	}{
		{
			name: "pass - Valid NFTokenCancelOffer",
			tx: &NFTokenCancelOffer{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: NFTokenCancelOfferTx,
				},
				NFTokenOffers: []types.NFTokenID{
					"9C92E061381C1EF37A8CDE0E8FC35188BFC30B1883825042A64309AC09F4C36D",
				},
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "fail - Invalid NFTokenCancelOffer - Empty NFTokenOffers",
			tx: &NFTokenCancelOffer{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: NFTokenCancelOfferTx,
				},
				NFTokenOffers: []types.NFTokenID{},
			},
			wantValid: false,
			wantErr:   true,
		},
		{
			name: "fail - Invalid NFTokenCancelOffer BaseTx - Missing TransactionType",
			tx: &NFTokenCancelOffer{
				BaseTx: BaseTx{
					Account: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
				},
				NFTokenOffers: []types.NFTokenID{
					"9C92E061381C1EF37A8CDE0E8FC35188BFC30B1883825042A64309AC09F4C36D",
				},
			},
			wantValid: false,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.tx.Validate()
			if valid != tt.wantValid {
				t.Errorf("Validate() valid = %v, want %v", valid, tt.wantValid)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
