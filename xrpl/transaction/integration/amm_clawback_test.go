package integration

import (
	"testing"

	ledger "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/xrpl/testutil/integration"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
	"github.com/stretchr/testify/require"
)

type AMMClawbackTest struct {
	Name          string
	Clawback      *transaction.AMMClawback
	ExpectedError string
}

// SetupAMMClawbackTestEnv sets up the environment for testing AMMClawback:
//   1. Enables clawback on the issuer (via AccountSet).
//   2. Establishes a trust line from the holder to the issuer for "FOO".
//   3. Issues 1000 FOO tokens from the issuer to the holder (via Payment).
//   4. Creates an AMM pool (via AMMCreate) depositing 500 FOO and some XRP.
// It returns the issuer and holder wallets for further testing.
func SetupAMMClawbackTestEnv(t *testing.T, runner *integration.Runner) (issuer, holder *wallet.Wallet) {
	issuer = runner.GetWallet(0)
	holder = runner.GetWallet(1)

	// 1. Enable Clawback on the Issuer:
	accountSet := &transaction.AccountSet{
		BaseTx: transaction.BaseTx{
			Account:         issuer.GetAddress(),
			// TransactionType: transaction.AccountSetTx,
			// Fee:             types.XRPCurrencyAmount(10),
			// Sequence:        1,
		},
	}
	accountSet.SetAsfAllowTrustLineClawback()
	flatTx := accountSet.Flatten()
	_, err := runner.TestTransaction(&flatTx, issuer, "tesSUCCESS")
	require.NoError(t, err, "Issuer should enable clawback successfully")

	// 2. Establish a Trust Line from Holder to Issuer for asset "FOO":
	trustSet := &transaction.TrustSet{
		BaseTx: transaction.BaseTx{
			Account:         holder.GetAddress(),
			// TransactionType: transaction.TrustSetTx,
			// Fee:             types.XRPCurrencyAmount(10),
			// Sequence:        1,
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   issuer.GetAddress(),
			Value:    "1000000",
		},
	}
	flatTx = trustSet.Flatten()
	_, err = runner.TestTransaction(&flatTx, holder, "tesSUCCESS")
	require.NoError(t, err, "Holder should establish a trust line successfully")

	// 3. Issue Tokens: Issuer sends 1000 FOO to Holder.
	payment := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account:         issuer.GetAddress(),
			// TransactionType: transaction.PaymentTx,
			// Fee:             types.XRPCurrencyAmount(10),
			// Sequence:        2,
		},
		Destination: holder.GetAddress(),
		Amount: types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   issuer.GetAddress(),
			Value:    "1000",
		},
	}
	flatTx = payment.Flatten()
	_, err = runner.TestTransaction(&flatTx, issuer, "tesSUCCESS")
	require.NoError(t, err, "Payment should succeed, issuing tokens to the holder")

	// 4. Create an AMM Pool: Holder creates a pool with 500 FOO and 50 XRP.
	ammCreate := &transaction.AMMCreate{
		BaseTx: transaction.BaseTx{
			Account:         holder.GetAddress(),
			// TransactionType: transaction.AMMCreateTx,
			// Fee:             types.XRPCurrencyAmount(10),
			// Sequence:        2,
		},
		Amount: types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   issuer.GetAddress(),
			Value:    "500",
		},
		Amount2:    types.XRPCurrencyAmount(50000000),
		// TradingFee: 500,
	}
	flatTx = ammCreate.Flatten()
	_, err = runner.TestTransaction(&flatTx, holder, "tesSUCCESS")
	require.NoError(t, err, "AMM pool creation should succeed")

	return issuer, holder
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

	// info, err := client.GetAccountInfo(&account.InfoRequest{Account: runner.GetWallet(0).GetAddress()})
	// require.NoError(t, err)
	// // t.Logf("Issuer current sequence: %d", info.AccountData.Sequence)

	issuer, holder := SetupAMMClawbackTestEnv(t, runner)

	tests := []AMMClawbackTest{
		{
			Name: "pass - valid AMMClawback",
			Clawback: &transaction.AMMClawback{
				BaseTx: transaction.BaseTx{
					Account:         issuer.GetAddress(),
					// TransactionType: transaction.AMMClawbackTx,
					// Fee:             types.XRPCurrencyAmount(10),
					// Sequence:        3, 
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
					// TransactionType: transaction.AMMClawbackTx,
					// Fee:             types.XRPCurrencyAmount(10),
					// // Sequence:        4, // next sequence for issuer
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

	issuer, holder := SetupAMMClawbackTestEnv(t, runner)

	tests := []AMMClawbackTest{
		{
			Name: "pass - valid AMMClawback",
			Clawback: &transaction.AMMClawback{
				BaseTx: transaction.BaseTx{
					Account:         issuer.GetAddress(),
					// TransactionType: transaction.AMMClawbackTx,
					// Fee:             types.XRPCurrencyAmount(10),
					// Sequence:        3,
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
					// TransactionType: transaction.AMMClawbackTx,
					// Fee:             types.XRPCurrencyAmount(10),
					// Sequence:        4,
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