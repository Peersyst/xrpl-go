package credential

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/Peersyst/xrpl-go/xrpl/testutil/integration"
	rippleTime "github.com/Peersyst/xrpl-go/xrpl/time"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/results"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
	"github.com/stretchr/testify/require"
)

func TestIntegrationCredentialWorkflowWebsocket(t *testing.T) {
	// Get the test environment configuration for websocket connection
	env := integration.GetWebsocketEnv(t)

	// Create a new websocket client with the test environment host and faucet provider
	client := websocket.NewClient(websocket.NewClientConfig().WithHost(env.Host).WithFaucetProvider(env.FaucetProvider))

	// Initialize a test runner with 2 wallets (one for sender, one for receiver)
	runner := integration.NewRunner(t, client, &integration.RunnerConfig{
		WalletCount: 2,
	})

	// Set up the test environment and create the wallets
	err := runner.Setup()
	require.NoError(t, err)

	// Get the wallet instances for sender and receiver
	sender := runner.GetWallet(0)
	receiver := runner.GetWallet(1)

	// Convert current time to Ripple time (seconds since Ripple Epoch)
	nowTime, err := rippleTime.IsoTimeToRippleTime(time.Now().Format(time.RFC3339))
	require.NoError(t, err)

	// Set up time values for testing:
	// - futureTime: 10000 seconds in the future (for valid expiration)
	futureTime := nowTime + 10000

	// Define credential types in hex format:
	// - "6D795F73757065725F63726564656E7469616C" = "my_super_credential" in ASCII
	credentialType := types.CredentialType("6D795F73757065725F63726564656E7469616C")

	// 1. Start with a valid CredentialCreate transaction
	createTx := &transaction.CredentialCreate{
		BaseTx: transaction.BaseTx{
			Account:         sender.GetAddress(),
			TransactionType: transaction.CredentialCreateTx,
			Fee:             types.XRPCurrencyAmount(10),
		},
		CredentialType: credentialType,
		Subject:        types.Address(receiver.GetAddress()),
		Expiration:     uint32(futureTime),
		URI:            hex.EncodeToString([]byte("https://example.com")),
	}

	createFlatTx := createTx.Flatten()

	_, err = runner.TestTransaction(&createFlatTx, sender, results.TesSUCCESS.String())
	require.NoError(t, err)

	// 2. Create a CredentialAccept transaction
	acceptTx := &transaction.CredentialAccept{
		BaseTx: transaction.BaseTx{
			Account:         receiver.GetAddress(),
			TransactionType: transaction.CredentialAcceptTx,
			Fee:             types.XRPCurrencyAmount(10),
		},
		CredentialType: credentialType,
		Issuer:         types.Address(sender.GetAddress()),
	}

	acceptFlatTx := acceptTx.Flatten()

	_, err = runner.TestTransaction(&acceptFlatTx, receiver, results.TesSUCCESS.String())
	require.NoError(t, err)

	// 3. Delete the Credential
	deleteTx := &transaction.CredentialDelete{
		BaseTx: transaction.BaseTx{
			Account:         sender.GetAddress(),
			TransactionType: transaction.CredentialDeleteTx,
		},
		CredentialType: credentialType,
		Issuer:         types.Address(sender.GetAddress()),
		Subject:        types.Address(receiver.GetAddress()),
	}

	deleteFlatTx := deleteTx.Flatten()

	_, err = runner.TestTransaction(&deleteFlatTx, sender, results.TesSUCCESS.String())
	require.NoError(t, err)

	// Teardown the test environment
	err = runner.Teardown()
	require.NoError(t, err)
}
