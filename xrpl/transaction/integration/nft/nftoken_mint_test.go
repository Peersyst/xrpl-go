package integration

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil/integration"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
	"github.com/stretchr/testify/require"
)

type NFTokenMintTest struct {
	Name          string
	NFTokenMint   *transaction.NFTokenMint
	ExpectedError string
}

func TestIntegrationNFTokenMint_Websocket(t *testing.T) {
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

	tt := []NFTokenMintTest{
		{
			Name: "pass - Minting an NFT with a Destination wallet address",
			NFTokenMint: &transaction.NFTokenMint{
				BaseTx: transaction.BaseTx{
					Account: sender.GetAddress(),
				},
				Amount:       types.XRPCurrencyAmount(1),
				NFTokenTaxon: 0,
				URI:          types.NFTokenURI("68747470733A2F2F676F6F676C652E636F6D"), // https://google.com
				Destination:  receiver.GetAddress(),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			flatTx := tc.NFTokenMint.Flatten()
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
