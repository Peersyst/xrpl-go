package transaction

import (
	"testing"

	ledger "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestIsSigner(t *testing.T) {
	tests := []struct {
		name     string
		input    SignerData
		expected bool
	}{
		{
			name: "Valid Signer object",
			input: SignerData{
				Account:       "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
				TxnSignature:  "0123456789abcdef",
				SigningPubKey: "abcdef0123456789",
			},
			expected: true,
		},
		{
			name: "Signer object with missing fields",
			input: SignerData{
				Account:       "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
				SigningPubKey: "abcdef0123456789",
			},
			expected: false,
		},
		{
			name: "Invalid Signer object with empty XRPL account",
			input: SignerData{
				Account:       "  ",
				SigningPubKey: "abcdef0123456789",
				TxnSignature:  "0123456789abcdef",
			},
			expected: false,
		},
		{
			name: "Invalid Signer object with invalid XRPL account",
			input: SignerData{
				Account:       "invalid",
				SigningPubKey: "abcdef0123456789",
				TxnSignature:  "0123456789abcdef",
			},
			expected: false,
		},
		{
			name: "Invalid Signer object with empty TxnSignature",
			input: SignerData{
				Account:       "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
				TxnSignature:  "  ",
				SigningPubKey: "abcdef0123456789",
			},
			expected: false,
		},
		{
			name: "Invalid Signer object with empty SigningPubKey",
			input: SignerData{
				Account:       "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
				TxnSignature:  "0123456789abcdef",
				SigningPubKey: "  ",
			},
			expected: false,
		},
		{
			name:     "Nil object",
			input:    SignerData{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ok, err := IsSigner(tt.input); ok != tt.expected {
				t.Errorf("Expected IsSigner to return %v, but got %v with error: %v", tt.expected, ok, err)
			}
		})
	}
}
func TestIsIssuedCurrency(t *testing.T) {
	tests := []struct {
		name     string
		input    types.CurrencyAmount
		expected bool
	}{
		{
			name: "Valid IssuedCurrency object",
			input: types.IssuedCurrencyAmount{
				Value:    "100",
				Issuer:   "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
				Currency: "USD",
			},
			expected: true,
		},
		{
			name:     "Invalid IssuedCurrency object",
			input:    types.XRPCurrencyAmount(100), // should be non XRP
			expected: false,
		},
		{
			name: "IssuedCurrency object with missing currency and issuer fields",
			input: types.IssuedCurrencyAmount{
				Value: "100",
			},
			expected: false,
		},
		{
			name: "IssuedCurrency object with missing issuer and value fields",
			input: types.IssuedCurrencyAmount{
				Currency: "USD",
			},
			expected: false,
		},
		{
			name: "IssuedCurrency object with missing currency and value fields",
			input: types.IssuedCurrencyAmount{
				Issuer: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
			},
			expected: false,
		},
		{
			name: "IssuedCurrency object with empty currency",
			input: types.IssuedCurrencyAmount{
				Issuer:   "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
				Currency: "   ",
				Value:    "100",
			},
			expected: false,
		},
		{
			name: "IssuedCurrency object with XRP currency",
			input: types.IssuedCurrencyAmount{
				Issuer:   "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
				Currency: "XRp", // will be uppercased during validation
				Value:    "100",
			},
			expected: false,
		},
		{
			name: "IssuedCurrency object with empty value",
			input: types.IssuedCurrencyAmount{
				Issuer:   "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
				Currency: "USD",
				Value:    "  ",
			},
			expected: false,
		},
		{
			name: "IssuedCurrency object with invalid issuer",
			input: types.IssuedCurrencyAmount{
				Issuer:   "invalid",
				Currency: "USD",
				Value:    "100",
			},
			expected: false,
		},
		{
			name:     "Empty object",
			input:    types.IssuedCurrencyAmount{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ok, err := IsIssuedCurrency(tt.input); ok != tt.expected {
				t.Errorf("Expected IsIssuedCurrency to return %v, but got %v with error: %v", tt.expected, ok, err)
			}
		})
	}
}

func TestIsMemo(t *testing.T) {
	t.Run("Valid Memo object with all fields", func(t *testing.T) {
		obj := Memo{
			MemoData:   "0123456789abcdef",
			MemoFormat: "abcdef0123456789",
			MemoType:   "abcdef0123456789",
		}

		ok, _ := IsMemo(obj)

		if !(ok) {
			t.Errorf("Expected IsMemo to return true, but got false")
		}
	})

	t.Run("Valid memo object with missing fields", func(t *testing.T) {
		obj := Memo{
			MemoData: "0123456789abcdef",
		}

		ok, err := IsMemo(obj)

		if !ok {
			t.Errorf("Expected IsMemo to return true, but got false with error: %v", err)
		}
	})

	t.Run("Memo object with MemoData non hex value", func(t *testing.T) {
		obj := Memo{
			MemoData: "bob",
		}

		if ok, _ := IsMemo(obj); ok {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})

	t.Run("Memo object with MemoFormat non hex value", func(t *testing.T) {
		obj := Memo{
			MemoData:   "0123456789abcdef",
			MemoFormat: "non-hex",
		}

		if ok, _ := IsMemo(obj); ok {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})

	t.Run("Memo object with MemoType non hex value", func(t *testing.T) {
		obj := Memo{
			MemoData:   "0123456789abcdef",
			MemoFormat: "0123456789abcdef",
			MemoType:   "non-hex",
		}

		if ok, _ := IsMemo(obj); ok {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})

	t.Run("Empty object", func(t *testing.T) {
		obj := Memo{}
		if ok, _ := IsMemo(obj); ok {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})
}
func TestIsAsset(t *testing.T) {
	t.Run("Valid Asset object with currency XRP only", func(t *testing.T) {
		obj := ledger.Asset{
			Currency: "xrP", // will be converted to XRP in the Validate function
		}

		ok, err := IsAsset(obj)

		if !ok {
			t.Errorf("Expected IsAsset to return true, but got false with error: %v", err)
		}
	})

	t.Run("Invalid Asset object with currency only and different than XRP", func(t *testing.T) {
		obj := ledger.Asset{
			Currency: "USD", // missing issuer
		}

		ok, err := IsAsset(obj)

		if ok {
			t.Errorf("Expected IsAsset to return true, but got false with error: %v", err)
		}
	})

	t.Run("Valid Asset object with currency and issuer", func(t *testing.T) {
		obj := ledger.Asset{
			Currency: "USD",
			Issuer:   "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
		}

		ok, err := IsAsset(obj)

		if !ok {
			t.Errorf("Expected IsAsset to return true, but got false with error: %v", err)
		}
	})

	t.Run("Asset object with missing currency", func(t *testing.T) {
		obj := ledger.Asset{
			Issuer: "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
		}

		ok, err := IsAsset(obj)

		if ok {
			t.Errorf("Expected IsAsset to return false, but got true")
		} else if err == nil {
			t.Errorf("Expected an error, but got nil")
		}
	})

	t.Run("Empty Asset object", func(t *testing.T) {
		obj := ledger.Asset{}

		ok, err := IsAsset(obj)

		if ok {
			t.Errorf("Expected IsAsset to return false, but got true")
		} else if err == nil {
			t.Errorf("Expected an error, but got nil")
		}
	})
}
func TestIsPath(t *testing.T) {
	tests := []struct {
		name     string
		input    []PathStep
		expected bool
	}{
		{
			name: "Valid path with account only",
			input: []PathStep{
				{Account: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW"},
			},
			expected: true,
		},
		{
			name: "Valid path with currency only",
			input: []PathStep{
				{Currency: "USD"},
			},
			expected: true,
		},
		{
			name: "Valid path with issuer only",
			input: []PathStep{
				{Issuer: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW"},
			},
			expected: true,
		},
		{
			name: "Valid path with currency and issuer",
			input: []PathStep{
				{Currency: "USD", Issuer: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW"},
			},
			expected: true,
		},
		{
			name: "Invalid path with account and currency",
			input: []PathStep{
				{Account: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW", Currency: "USD"},
			},
			expected: false,
		},
		{
			name: "Invalid path with account and issuer",
			input: []PathStep{
				{Account: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW", Issuer: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW"},
			},
			expected: false,
		},
		{
			name: "Invalid path with currency XRP and issuer",
			input: []PathStep{
				{Currency: "XRP", Issuer: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW"},
			},
			expected: false,
		},
		{
			name:     "Empty path",
			input:    []PathStep{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ok, err := IsPath(tt.input); ok != tt.expected {
				t.Errorf("Expected IsPath to return %v, but got %v with error: %v", tt.expected, ok, err)
			}
		})
	}
}
func TestIsPaths(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]PathStep
		expected bool
	}{
		{
			name: "Valid paths with single path and single step",
			input: [][]PathStep{
				{
					{Account: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW"},
				},
			},
			expected: true,
		},
		{
			name: "Valid paths with multiple paths and steps",
			input: [][]PathStep{
				{
					{Account: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW"},
					{Currency: "USD"},
				},
				{
					{Issuer: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW"},
					{Currency: "EUR", Issuer: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW"},
				},
			},
			expected: true,
		},
		{
			name: "Invalid paths with empty path",
			input: [][]PathStep{
				{},
			},
			expected: false,
		},
		{
			name: "Invalid paths with empty path step",
			input: [][]PathStep{
				{
					{},
				},
			},
			expected: false,
		},
		{
			name: "Invalid paths with invalid path step, account and currency cannot be together",
			input: [][]PathStep{
				{
					{Account: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW", Currency: "USD"},
				},
			},
			expected: false,
		},
		{
			name: "Invalid paths with invalid path step having currency XRP and issuer",
			input: [][]PathStep{
				{
					{Currency: "XRP", Issuer: "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW"},
				},
			},
			expected: false,
		},
		{
			name:     "Empty paths",
			input:    [][]PathStep{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ok, err := IsPaths(tt.input); ok != tt.expected {
				t.Errorf("Expected IsPaths to return %v, but got %v with error: %v", tt.expected, ok, err)
			}
		})
	}
}
