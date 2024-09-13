package transactions

import (
	"testing"
)

func TestParseAmountValue(t *testing.T) {
	// Test valid XRP amount as float
	xrpAmountStr := "10.5"
	expectedXrpAmount := 10.5
	parsedXrpAmount, err := ParseAmountValue(xrpAmountStr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if parsedXrpAmount != expectedXrpAmount {
		t.Errorf("Expected XRP amount: %f, got: %f", expectedXrpAmount, parsedXrpAmount)
	}

	// Test valid XRP amount as int
	xrpAmountInt := "20"
	expectedXrpAmount = 20.0
	parsedXrpAmount, err = ParseAmountValue(xrpAmountInt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if parsedXrpAmount != expectedXrpAmount {
		t.Errorf("Expected XRP amount: %f, got: %f", expectedXrpAmount, parsedXrpAmount)
	}

	// Test valid Issued Currency amount as float
	issuedAmount := map[string]interface{}{
		"value":    "30.75",
		"issuer":   "rB3gZey7VWHYRqJHLoHDEJXJ2pEPNieKiS",
		"currency": "USD",
	}
	expectedIssuedAmount := 30.75
	parsedIssuedAmount, err := ParseAmountValue(issuedAmount)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if parsedIssuedAmount != expectedIssuedAmount {
		t.Errorf("Expected Issued Currency amount: %f, got: %f", expectedIssuedAmount, parsedIssuedAmount)
	}

	// Test valid Issued Currency amount as int
	issuedAmountAsInt := map[string]interface{}{
		"value":    "30",
		"issuer":   "rB3gZey7VWHYRqJHLoHDEJXJ2pEPNieKiS",
		"currency": "USD",
	}
	expectedIssuedAmount2 := 30.0
	parsedIssuedAmount2, err := ParseAmountValue(issuedAmountAsInt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if parsedIssuedAmount2 != expectedIssuedAmount2 {
		t.Errorf("Expected Issued Currency amount: %f, got: %f", expectedIssuedAmount2, parsedIssuedAmount2)
	}

	// Test invalid amount format
	invalidAmount := "invalid"
	_, err = ParseAmountValue(invalidAmount)
	if err == nil {
		t.Error("Expected error for invalid amount format, but got nil")
	}

	// Test unsupported amount type
	unsupportedAmount := []int{1, 2, 3}
	_, err = ParseAmountValue(unsupportedAmount)
	if err == nil {
		t.Error("Expected error for unsupported amount type, but got nil")
	}
}
