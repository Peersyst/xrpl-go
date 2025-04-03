package integration

import (
	"reflect"
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

// AMMClawbackTest defines a test case for AMMClawback transactions.
type AMMClawbackTest struct {
	Name          string
	Clawback      *transaction.AMMClawback
	ExpectedError string
}

// logFlattenedTx logs each field in the flattened transaction with its type and value.
func logFlattenedTx(t *testing.T, label string, tx map[string]interface{}) {
	t.Logf("Logging flattened transaction (%s):", label)
	for key, val := range tx {
		t.Logf("  Field: %s, Type: %T, Value: %+v", key, val, val)
	}
}


// SetupAMMClawbackTestEnv sets up the environment for testing AMMClawback:
//   1. Enables clawback on the issuer (via AccountSet).
//   2. Establishes a trust line from the holder to the issuer for "FOO".
//   3. Issues 1000 FOO tokens from the issuer to the holder (via Payment).
//   4. Creates an AMM pool (via AMMCreate) depositing 500 FOO and some XRP.
// It returns the issuer and holder wallets.
func SetupAMMClawbackTestEnv(t *testing.T, runner *integration.Runner) (issuer, holder *wallet.Wallet) {
	t.Log("Setting up AMMClawback test environment...")
	issuer = runner.GetWallet(0)
	holder = runner.GetWallet(1)

	// 1. Enable Clawback on the Issuer:
	accountSet := &transaction.AccountSet{
		BaseTx: transaction.BaseTx{
			Account: issuer.GetAddress(),
		},
	}
	accountSet.SetAsfAllowTrustLineClawback()
	flatTx := accountSet.Flatten()
	t.Log("AccountSet flattened transaction:")
	logFlattenedTx(t, "IssuerSet", flatTx)
	_,_, err := runner.ProcessTransactionAndWait(&flatTx, issuer)
	require.NoError(t, err, "Issuer should enable clawback successfully")


	accountSet2 := &transaction.AccountSet{
		BaseTx: transaction.BaseTx{
			Account: issuer.GetAddress(),
		},
	}
	accountSet2.SetAsfDefaultRipple()
	flatTx2 := accountSet2.Flatten()
	t.Log("AccountSet flattened transaction:")
	logFlattenedTx(t, "IssuerSet", flatTx)
	_,_, err2 := runner.ProcessTransactionAndWait(&flatTx2, issuer)
	require.NoError(t, err2, "Issuer should enable clawback successfully")

	// Enable DefaultRipple on the holder.
	holderSet := &transaction.AccountSet{
		BaseTx: transaction.BaseTx{
			Account: holder.GetAddress(),
		},
	}
	holderSet.SetAsfRequireAuth()
	flatTx = holderSet.Flatten()
	t.Log("Holder DefaultRipple AccountSet flattened transaction:")
	logFlattenedTx(t, "HolderSet", flatTx)
	_, _, err = runner.ProcessTransactionAndWait(&flatTx, holder)
	require.NoError(t, err, "Holder should enable DefaultRipple successfully")

	// 2. Establish a Trust Line from Holder to Issuer for asset "FOO":
	trustSet := &transaction.TrustSet{
		BaseTx: transaction.BaseTx{
			Account: holder.GetAddress(),
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   issuer.GetAddress(),
			Value:    "1000000",
		},
	}
	flatTx = trustSet.Flatten()
	t.Log("TrustSet flattened transaction:")
	logFlattenedTx(t, "TrustSet", flatTx)
	_,_, err = runner.ProcessTransactionAndWait(&flatTx, holder)
	require.NoError(t, err, "Holder should establish a trust line successfully")

	// 3. Issue Tokens: Issuer sends 1000 FOO to Holder.
	payment := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account: issuer.GetAddress(),
		},
		Destination: holder.GetAddress(),
		Amount: types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   issuer.GetAddress(),
			Value:    "1000",
		},
	}
	flatTx = payment.Flatten()
	t.Log("Payment flattened transaction:")
	logFlattenedTx(t, "Payment", flatTx)
	_,_, err = runner.ProcessTransactionAndWait(&flatTx, issuer)
	require.NoError(t, err, "Payment should succeed, issuing tokens to the holder")

	// 4. Create an AMM Pool: Holder creates a pool with 500 FOO and 50 XRP.
	ammCreate := &transaction.AMMCreate{
		BaseTx: transaction.BaseTx{
			Account: holder.GetAddress(),
			Fee:     200000, // set fee explicitly
		},
		Amount: types.IssuedCurrencyAmount{
			Currency: "FOO",
			Issuer:   issuer.GetAddress(),
			Value:    "5",
		},
		Amount2:    types.XRPCurrencyAmount(5),
		TradingFee: 100,
	}
	flatTx = ammCreate.Flatten()
	flatTx["TradingFee"] = int(ammCreate.TradingFee)
	t.Log("AMMCreate flattened transaction:")
	logFlattenedTx(t, "AMMCreate", flatTx)
	_,_, err = runner.ProcessTransactionAndWait(&flatTx, holder)
	require.NoError(t, err, "AMM pool creation should succeed")

	return issuer, holder
}

func TestIntegrationAMMClawback_Websocket(t *testing.T) {
	t.Log("Starting TestIntegrationAMMClawback_Websocket")
	env := integration.GetWebsocketEnv(t)
	wsClient := websocket.NewClient(websocket.NewClientConfig().
		WithHost(env.Host).
		WithFaucetProvider(env.FaucetProvider))
	runner := integration.NewRunner(t, wsClient, &integration.RunnerConfig{
		WalletCount: 2,
	})

	err := runner.Setup()
	require.NoError(t, err)
	defer runner.Teardown()

	issuer, holder := SetupAMMClawbackTestEnv(t, runner)

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
					Currency: "XRP",
				},
				Amount: types.IssuedCurrencyAmount{
					Currency: "FOO",
					Issuer:   issuer.GetAddress(),
					Value:    "1",
				},
			},
			ExpectedError: "",
		},
		{
			Name: "fail - missing Holder",
			Clawback: &transaction.AMMClawback{
				BaseTx: transaction.BaseTx{
					Account: issuer.GetAddress(),
				},
				// Holder is intentionally missing.
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
					Value:    "1",
				},
			},
			ExpectedError: "missing Holder",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Logf("Starting test: %s", tc.Name)
			flatTx := tc.Clawback.Flatten()
			t.Logf("Flattened AMMClawback transaction for test '%s':", tc.Name)
			logFlattenedTx(t, tc.Name, flatTx)
			for key, val := range flatTx {
				t.Logf("Test %s - Key: %s, Type: %s", tc.Name, key, reflect.TypeOf(val))
			}
			if tc.ExpectedError == "" {
				_,_, err := runner.ProcessTransactionAndWait(&flatTx, issuer)
				require.NoError(t, err)
			} else {
				txBlob, _, err := issuer.Sign(flatTx)
				require.NoError(t, err, "failed to sign transaction")
				_, err = wsClient.SubmitAndWait(txBlob, false)
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ExpectedError)
			}
		})
	}
}

func TestIntegrationAMMClawback_RPC(t *testing.T) {
	t.Log("Starting TestIntegrationAMMClawback_RPC")
	env := integration.GetRPCEnv(t)
	clientCfg, err := rpc.NewClientConfig(env.Host, rpc.WithFaucetProvider(env.FaucetProvider))
	require.NoError(t, err)
	rpcClient := rpc.NewClient(clientCfg)
	runner := integration.NewRunner(t, rpcClient, integration.NewRunnerConfig(
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
					Account: issuer.GetAddress(),
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
					Value:    "1",
				},
			},
			ExpectedError: "",
		},
		{
			Name: "fail - missing Holder",
			Clawback: &transaction.AMMClawback{
				BaseTx: transaction.BaseTx{
					Account: issuer.GetAddress(),
				},
				// Holder is intentionally missing.
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
					Value:    "1",
				},
			},
			ExpectedError: "missing Holder",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			t.Logf("Starting test: %s", tc.Name)
			flatTx := tc.Clawback.Flatten()
			t.Logf("Flattened AMMClawback transaction for test '%s':", tc.Name)
			logFlattenedTx(t, tc.Name, flatTx)
			for key, val := range flatTx {
				t.Logf("Test %s - Key: %s, Type: %s", tc.Name, key, reflect.TypeOf(val))
			}
			if tc.ExpectedError == "" {
				_,_, err := runner.ProcessTransactionAndWait(&flatTx, issuer)
				require.NoError(t, err)
			} else {
				txBlob, _, err := issuer.Sign(flatTx)
				require.NoError(t, err, "failed to sign transaction")
				_, err = rpcClient.SubmitAndWait(txBlob, false)
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ExpectedError)
			}
		})
	}
}