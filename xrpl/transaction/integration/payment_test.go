package integration

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/xrpl/testutil/integration"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
	"github.com/stretchr/testify/require"
)

type PaymentTest struct {
	Name          string
	Payment       *transaction.Payment
	ExpectedError string
}

func TestIntegrationPayment_Websocket(t *testing.T) {
	env := integration.GetWebsocketEnv(t)
	client := websocket.NewClient(websocket.NewClientConfig().WithHost(env.Host).WithFaucetProvider(env.FaucetProvider))

	runner := integration.NewRunner(t, client, &integration.RunnerConfig{
		WalletCount: 2,
	})

	err := runner.Setup()
	require.NoError(t, err)
	defer runner.Teardown()

	sender := runner.GetWallet(0)
	receiver := runner.GetWallet(1)

	tt := []PaymentTest{
		{
			Name: "pass - XRP to XRP",
			Payment: &transaction.Payment{
				BaseTx: transaction.BaseTx{
					Account: sender.GetAddress(),
				},
				Amount:      types.XRPCurrencyAmount(1),
				Destination: receiver.GetAddress(),
			},
		},
		{
			Name: "fail - missing Destination",
			Payment: &transaction.Payment{
				BaseTx: transaction.BaseTx{
					Account: sender.GetAddress(),
				},
				Amount: types.XRPCurrencyAmount(1),
			},
			ExpectedError: ErrInvalidTransaction,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			flatTx := tc.Payment.Flatten()
			_, err := runner.TestTransaction(&flatTx, sender, "tesSUCCESS")
			if tc.ExpectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ExpectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIntegrationPayment_RPCClient(t *testing.T) {
	env := integration.GetRPCEnv(t)
	clientCfg, err := rpc.NewClientConfig(env.Host, rpc.WithFaucetProvider(env.FaucetProvider))
	require.NoError(t, err)
	client := rpc.NewClient(clientCfg)

	runner := integration.NewRunner(t, client, integration.NewRunnerConfig(
		integration.WithWallets(2),
	))

	err = runner.Setup()
	require.NoError(t, err)
	defer runner.Teardown()

	sender := runner.GetWallet(0)
	receiver := runner.GetWallet(1)

	tt := []PaymentTest{
		{
			Name: "pass - XRP to XRP",
			Payment: &transaction.Payment{
				BaseTx: transaction.BaseTx{
					Account: sender.GetAddress(),
				},
				Amount:      types.XRPCurrencyAmount(1),
				Destination: receiver.GetAddress(),
			},
		},
		{
			Name: "fail - missing Destination",
			Payment: &transaction.Payment{
				BaseTx: transaction.BaseTx{
					Account: sender.GetAddress(),
				},
				Amount: types.XRPCurrencyAmount(1),
			},
			ExpectedError: ErrInvalidTransaction,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			flatTx := tc.Payment.Flatten()
			_, err := runner.TestTransaction(&flatTx, sender, "tesSUCCESS")
			if tc.ExpectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ExpectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
