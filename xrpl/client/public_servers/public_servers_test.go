package publicservers

import (
	"testing"
)

func TestNewWSPublicServersUrls(t *testing.T) {
	publicServers := NewWSPublicServersUrls()

	// Test mainnet URLs
	mainnet := publicServers.Mainnet()
	if mainnet.XRPLedgerFoundation() != xrpLedgerFoundation {
		t.Errorf("Expected XRPLedgerFoundation URL to be %s, but got %s", xrpLedgerFoundation, mainnet.XRPLedgerFoundation())
	}
	if mainnet.RippleS1() != rippleS1 {
		t.Errorf("Expected RippleS1 URL to be %s, but got %s", rippleS1, mainnet.RippleS1())
	}
	if mainnet.RippleS2() != rippleS2 {
		t.Errorf("Expected RippleS2 URL to be %s, but got %s", rippleS2, mainnet.RippleS2())
	}

	// Test testnet URLs
	testnet := publicServers.Testnet()
	if testnet.Ripple() != rippleTestnet {
		t.Errorf("Expected Ripple URL to be %s, but got %s", rippleTestnet, testnet.Ripple())
	}
	if testnet.XRPLLabs() != xrplLabs {
		t.Errorf("Expected XRPLLabs URL to be %s, but got %s", xrplLabs, testnet.XRPLLabs())
	}
	if testnet.Clio() != clio {
		t.Errorf("Expected Clio URL to be %s, but got %s", clio, testnet.Clio())
	}

	// Test devnet URLs
	devnet := publicServers.Devnet()
	if devnet.RippleDevnet() != rippleDevnet {
		t.Errorf("Expected RippleDevnet URL to be %s, but got %s", rippleDevnet, devnet.RippleDevnet())
	}
	if devnet.RippleDevnetClio() != rippleDevnetClio {
		t.Errorf("Expected RippleDevnetClio URL to be %s, but got %s", rippleDevnetClio, devnet.RippleDevnetClio())
	}
}
