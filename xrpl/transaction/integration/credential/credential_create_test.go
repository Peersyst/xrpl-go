package credential

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/Peersyst/xrpl-go/xrpl/testutil/integration"
	rippleTime "github.com/Peersyst/xrpl-go/xrpl/time"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	e2eIntegration "github.com/Peersyst/xrpl-go/xrpl/transaction/integration"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/results"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
	"github.com/stretchr/testify/require"
)

type CredentialCreateTest struct {
	Name               string
	CredentialCreate   *transaction.CredentialCreate
	ExpectedResultCode results.TxResult
	ExpectedError      string
}

func TestIntegrationCredentialCreateWebsocket(t *testing.T) {
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
	// - pastTime: 1000 seconds (for testing expired credentials)
	futureTime := nowTime + 10000
	pastTime := 1000

	// Define credential types in hex format:
	// - "6D795F63726564656E7469616C" = "my_credential" in ASCII
	// - "6D795F63726564656E7469616C32" = "my_credential2" in ASCII
	// - "6D795F63726564656E7469616C33" = "my_credential3" in ASCII
	credentialType := types.CredentialType("6D795F63726564656E7469616C")
	credentialType2 := types.CredentialType("6D795F63726564656E7469616C32")
	credentialType3 := types.CredentialType("6D795F63726564656E7469616C33")

	tt := []CredentialCreateTest{
		{
			Name: "pass - valid CredentialCreate",
			CredentialCreate: &transaction.CredentialCreate{
				BaseTx: transaction.BaseTx{
					Account:         sender.GetAddress(),
					TransactionType: transaction.CredentialCreateTx,
					Fee:             types.XRPCurrencyAmount(10),
				},
				CredentialType: credentialType,
				Subject:        types.Address(receiver.GetAddress()),
				Expiration:     uint32(futureTime),
				URI:            hex.EncodeToString([]byte("https://example.com")),
			},
			ExpectedResultCode: results.TesSUCCESS,
		},
		{
			Name: "fail - Expiration is in the past	",
			CredentialCreate: &transaction.CredentialCreate{
				BaseTx: transaction.BaseTx{
					Account:         sender.GetAddress(),
					TransactionType: transaction.CredentialCreateTx,
				},
				Subject:        types.Address(receiver.GetAddress()),
				CredentialType: credentialType2,
				Expiration:     uint32(pastTime),
			},
			ExpectedResultCode: results.TecEXPIRED,
		},
		{
			Name: "fail - Subject is missing",
			CredentialCreate: &transaction.CredentialCreate{
				BaseTx: transaction.BaseTx{
					Account:         sender.GetAddress(),
					TransactionType: transaction.CredentialCreateTx,
				},
				CredentialType: credentialType3,
			},
			ExpectedResultCode: results.TesSUCCESS, // not relevant as rippled will throw "invalidTransaction" immediately upon the transaction submission
			ExpectedError:      e2eIntegration.ErrInvalidTransaction,
		},
		{
			Name: "fail - Duplicate CredentialCreate",
			CredentialCreate: &transaction.CredentialCreate{
				BaseTx: transaction.BaseTx{
					Account:         sender.GetAddress(),
					TransactionType: transaction.CredentialCreateTx,
				},
				Subject:        types.Address(receiver.GetAddress()),
				CredentialType: credentialType,
			},
			ExpectedResultCode: results.TecDUPLICATE,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			flatTx := tc.CredentialCreate.Flatten()

			_, err = runner.TestTransaction(&flatTx, sender, tc.ExpectedResultCode.String())
			if tc.ExpectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ExpectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}

	err = runner.Teardown()
	require.NoError(t, err)
}
