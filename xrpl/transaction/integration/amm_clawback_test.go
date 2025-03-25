package integration

import (
	"testing"

	ledger "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/xrpl/testutil/integration"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
	"github.com/stretchr/testify/require"
)

type AMMClawbackTest struct {
	Name          string
	Clawback      *transaction.AMMClawback
	ExpectedError string
}

func TestIntegrationAMMClawback_Websocket(t *testing.T) {
	env := integration.GetWebsocketEnv(t)
	client := websocket.NewClient(websocket.NewClientConfig().
		WithHost(env.Host).
		WithFaucetProvider(env.FaucetProvider))
	runner := integration.NewRunner(t, client, &integration.RunnerConfig{
		WalletCount: 2,
	})
	err := runner.Setup()
	require.NoError(t, err)
	defer runner.Teardown()

	issuer := runner.GetWallet(0)
	holder := runner.GetWallet(1)

	tests := []AMMClawbackTest{
		{
			Name: "pass - valid AMMClawback",
			Clawback: &transaction.AMMClawback{
				BaseTx: transaction.BaseTx{
					Account: issuer.GetAddress(),
				},
				Holder: holder.GetAddress(),
				Asset: ledger.Asset{
					Currency: "FOO",
					Issuer:   issuer.GetAddress(),
				},
				Asset2: ledger.Asset{
					// Using native XRP for Asset2; issuer is not required.
					Currency: "XRP",
				},
				Amount: types.IssuedCurrencyAmount{
					Currency: "FOO",
					Issuer:   issuer.GetAddress(),
					Value:    "1000",
				},
			},
			ExpectedError: "",
		},
		{
			Name: "fail - missing Holder",
			Clawback: &transaction.AMMClawback{
				BaseTx: transaction.BaseTx{
					Account:         issuer.GetAddress(),
					TransactionType: transaction.AMMClawbackTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        2,
				},
				Asset: ledger.Asset{
					Currency: "FOO",
					Issuer:   issuer.GetAddress(),
				},
				Asset2: ledger.Asset{
					Currency: "XRP",
				},
				Amount: types.IssuedCurrencyAmount{
					Currency: "FOO",
					Issuer:   issuer.GetAddress(),
					Value:    "1000",
				},
			},
			ExpectedError: ErrInvalidTransaction,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			flatTx := tc.Clawback.Flatten()
			_, err := runner.TestTransaction(&flatTx, issuer, "tesSUCCESS")
			if tc.ExpectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ExpectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIntegrationAMMClawback_RPC(t *testing.T) {
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

	issuer := runner.GetWallet(0)
	holder := runner.GetWallet(1)

	tests := []AMMClawbackTest{
		{
			Name: "pass - valid AMMClawback",
			Clawback: &transaction.AMMClawback{
				BaseTx: transaction.BaseTx{
					Account:         issuer.GetAddress(),
				},
				Holder: holder.GetAddress(),
				Asset: ledger.Asset{
					Currency: "FOO",
					Issuer:   issuer.GetAddress(),
				},
				Asset2: ledger.Asset{
					Currency: "XRP",
				},
				Amount: types.IssuedCurrencyAmount{
					Currency: "FOO",
					Issuer:   issuer.GetAddress(),
					Value:    "1000",
				},
			},
			ExpectedError: "",
		},
		{
			Name: "fail - missing Holder",
			Clawback: &transaction.AMMClawback{
				BaseTx: transaction.BaseTx{
					Account:         issuer.GetAddress(),
					TransactionType: transaction.AMMClawbackTx,
					Fee:             types.XRPCurrencyAmount(10),
					Sequence:        2,
				},
				Asset: ledger.Asset{
					Currency: "FOO",
					Issuer:   issuer.GetAddress(),
				},
				Asset2: ledger.Asset{
					Currency: "XRP",
				},
				Amount: types.IssuedCurrencyAmount{
					Currency: "FOO",
					Issuer:   issuer.GetAddress(),
					Value:    "1000",
				},
			},
			ExpectedError: "missing Holder",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			flatTx := tc.Clawback.Flatten()
			_, err := runner.TestTransaction(&flatTx, issuer, "tesSUCCESS")
			if tc.ExpectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ExpectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
