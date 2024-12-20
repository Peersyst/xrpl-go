package faucet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
)

const (
	TestnetFaucetHost = "faucet.altnet.rippletest.net"
	TestnetFaucetPath = "/accounts"
)

// TestnetFaucetProvider implements the FaucetProvider interface for the XRPL Testnet.
// It provides functionality to interact with the Testnet faucet for funding wallets.
type TestnetFaucetProvider struct {
	host        string
	accountPath string
}

// NewTestnetFaucetProvider creates and returns a new instance of TestnetFaucetProvider
// with predefined Testnet faucet host and account path.
func NewTestnetFaucetProvider() *TestnetFaucetProvider {
	return &TestnetFaucetProvider{
		host:        TestnetFaucetHost,
		accountPath: TestnetFaucetPath,
	}
}

// FundWallet sends a request to the Testnet faucet to fund the specified wallet address.
// It returns an error if the funding request fails.
func (fp *TestnetFaucetProvider) FundWallet(address types.Address) error {
	url := fmt.Sprintf("https://%s%s", fp.host, fp.accountPath)
	payload := map[string]string{"destination": address.String(), "userAgent": UserAgent}
	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
