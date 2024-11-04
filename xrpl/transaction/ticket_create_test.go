package transaction

import (
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestTicketCreateFlatten(t *testing.T) {
	s := TicketCreate{
		BaseTx: BaseTx{
			Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType: TicketCreateTx,
			Fee:             types.XRPCurrencyAmount(10),
			Sequence:        50,
		},
		TicketCount: 5,
	}

	flattened := s.Flatten()

	expected := FlatTransaction{
		"Account":         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		"TransactionType": "TicketCreate",
		"Fee":             "10",
		"Sequence":        uint(50),
		"TicketCount":     uint32(5),
	}

	if !reflect.DeepEqual(flattened, expected) {
		t.Errorf("Flatten result differs from expected: %v, %v", flattened, expected)
	}
}

func TestTicketCreate_TxType(t *testing.T) {
	tx := &TicketCreate{}
	assert.Equal(t, TicketCreateTx, tx.TxType())
}

func TestTicketCreate_Validate(t *testing.T) {
	tests := []struct {
		name      string
		ticket    TicketCreate
		wantValid bool
	}{
		{
			name: "valid ticket count",
			ticket: TicketCreate{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: TicketCreateTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        50,
				},
				TicketCount: 5,
			},
			wantValid: true,
		},
		{
			name: "Invalid BaseTx",
			ticket: TicketCreate{
				BaseTx: BaseTx{
					Account:         "",
					TransactionType: TicketCreateTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        50,
				},
				TicketCount: 5,
			},
			wantValid: false,
		},
		{
			name: "ticket count zero",
			ticket: TicketCreate{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: TicketCreateTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        50,
				},
				TicketCount: 0,
			},
			wantValid: false,
		},
		{
			name: "ticket count exceeds limit",
			ticket: TicketCreate{
				BaseTx: BaseTx{
					Account:         "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
					TransactionType: TicketCreateTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        50,
				},
				TicketCount: 251,
			},
			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.ticket.Validate()
			if valid != tt.wantValid {
				t.Errorf("Validate() valid = %v, want %v, err: %v", valid, tt.wantValid, err)
			}
		})
	}
}
