package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

func TestIsSigner(t *testing.T) {
	t.Run("Valid Signer object", func(t *testing.T) {
		validSigner := SignerData{
			Account:       "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
			TxnSignature:  "0123456789abcdef",
			SigningPubKey: "abcdef0123456789",
		}

		if ok, err := IsSigner(validSigner); !ok {
			t.Errorf("Expected IsSigner to return true, but got false with error: %v", err)
		}
	})

	t.Run("Signer object with missing fields", func(t *testing.T) {
		invalidSigner := SignerData{
			Account:       "r4ES5Mmnz4HGbu2asdicuECBaBWo4knhXW",
			SigningPubKey: "abcdef0123456789",
		}

		if ok, _ := IsSigner(invalidSigner); ok {
			t.Errorf("Expected IsSigner to return false, but got true")
		}
	})

	t.Run("Nil object", func(t *testing.T) {
		invalidSigner := SignerData{}
		if ok, _ := IsSigner(invalidSigner); ok {
			t.Errorf("Expected IsSigner to return false, but got true")
		}
	})
}
func TestIsIssuedCurrency(t *testing.T) {
	t.Run("Valid IssuedCurrency object", func(t *testing.T) {
		obj1 := types.IssuedCurrencyAmount{
			Value:    "100",
			Issuer:   "r1234567890",
			Currency: "USD",
		}
		if ok, err := IsIssuedCurrency(obj1); !ok {
			t.Errorf("Expected IsIssuedCurrency to return true, but got false with error: %v", err)
		}
	})

	t.Run("IssuedCurrency object with missing fields", func(t *testing.T) {
		invalid := types.IssuedCurrencyAmount{
			Value: "100",
		}
		if ok, _ := IsIssuedCurrency(invalid); ok {
			t.Errorf("Expected IsIssuedCurrency to return false, but got true")
		}
	})

	t.Run("Empty object", func(t *testing.T) {
		invalid := types.IssuedCurrencyAmount{}
		if ok, _ := IsIssuedCurrency(invalid); ok {
			t.Errorf("Expected IsIssuedCurrency to return false, but got true")
		}
	})
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

	t.Run("Memo object with non hex values", func(t *testing.T) {
		obj := Memo{
			MemoData:   "bob",
			MemoFormat: "alice",
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
