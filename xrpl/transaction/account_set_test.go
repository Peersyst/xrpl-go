package transaction

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestAccountSetTransaction(t *testing.T) {
	s := AccountSet{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: AccountSetTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Domain:  "A5B21758D2318FA2C",
		SetFlag: 10,
	}
	j := `{
	"Account": "abcdef",
	"TransactionType": "AccountSet",
	"Fee": "1",
	"Sequence": 1234,
	"SigningPubKey": "ghijk",
	"TxnSignature": "A1B2C3D4E5F6",
	"Domain": "A5B21758D2318FA2C",
	"SetFlag": 10
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

func TestAccountSetFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*AccountSet)
		expected uint
	}{
		{
			name: "SetRequireDestTag",
			setter: func(s *AccountSet) {
				s.SetRequireDestTag()
			},
			expected: tfRequireDestTag,
		},
		{
			name: "SetRequireAuth",
			setter: func(s *AccountSet) {
				s.SetRequireAuth()
			},
			expected: tfRequireAuth,
		},
		{
			name: "SetDisallowXRP",
			setter: func(s *AccountSet) {
				s.SetDisallowXRP()
			},
			expected: tfDisallowXRP,
		},
		{
			name: "SetOptionalDestTag",
			setter: func(s *AccountSet) {
				s.SetOptionalDestTag()
			},
			expected: tfOptionalDestTag,
		},
		{
			name: "SetRequireDestTag and SetRequireAuth",
			setter: func(s *AccountSet) {
				s.SetRequireDestTag()
				s.SetRequireAuth()
			},
			expected: tfRequireDestTag | tfRequireAuth,
		},
		{
			name: "SetDisallowXRP and SetOptionalDestTag",
			setter: func(s *AccountSet) {
				s.SetDisallowXRP()
				s.SetOptionalDestTag()
			},
			expected: tfDisallowXRP | tfOptionalDestTag,
		},
		{
			name: "SetRequireDestTag, SetRequireAuth, and SetDisallowXRP",
			setter: func(s *AccountSet) {
				s.SetRequireDestTag()
				s.SetRequireAuth()
				s.SetDisallowXRP()
			},
			expected: tfRequireDestTag | tfRequireAuth | tfDisallowXRP,
		},
		{
			name: "All flags",
			setter: func(s *AccountSet) {
				s.SetRequireDestTag()
				s.SetRequireAuth()
				s.SetDisallowXRP()
				s.SetOptionalDestTag()
			},
			expected: tfRequireDestTag | tfRequireAuth | tfDisallowXRP | tfOptionalDestTag,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AccountSet{}
			tt.setter(s)
			if s.Flags != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, s.Flags)
			}
		})
	}
}
func TestAccountSet_Validate(t *testing.T) {
	testCases := []struct {
		name       string
		accountSet AccountSet
		valid      bool
	}{
		{
			name: "Valid AccountSet",
			accountSet: AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				ClearFlag:    1,
				Domain:       "A5B21758D2318FA2C",
				EmailHash:    "1234567890abcdef",
				MessageKey:   "messagekey",
				SetFlag:      2,
				TransferRate: 1000000000,
				TickSize:     5,
			},
			valid: true,
		},
		{
			name: "Valid AccountSet without options, just the commons fields",
			accountSet: AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
			},
			valid: true,
		},
		{
			name: "Invalid AccountSet with high SetFlag",
			accountSet: AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				SetFlag: 18, // too high
			},
			valid: false,
		},
		{
			name: "Invalid AccountSet with low TickSize",
			accountSet: AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				TickSize: 2, // too low
			},
			valid: false,
		},
		{
			name: "Invalid AccountSet with high TickSize",
			accountSet: AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				TickSize: 16, // too high
			},
			valid: false,
		},
		{
			name: "Valid AccountSet TickSize set to 0 to disable it",
			accountSet: AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				TickSize: 0,
			},
			valid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			valid, err := tc.accountSet.Validate()
			if valid != tc.valid {
				t.Errorf("Validation result for %s is incorrect. Expected: %v, Got: %v", tc.name, tc.valid, valid)
			}
			if err != nil && tc.valid {
				t.Errorf("Validation failed for %s: %s", tc.name, err)
			}
			if err == nil && !tc.valid {
				t.Errorf("Validation should have failed for %s", tc.name)
			}
		})
	}
}
