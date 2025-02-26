package integration

import (
	"testing"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/queries/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
	"github.com/stretchr/testify/require"
)

type Runner struct {
	t      *testing.T
	config *RunnerConfig

	client  Client
	wallets []*wallet.Wallet
}

// NewRunner creates a new runner. It doesn't connect to the websocket or generate wallets until Setup is called.
// A testing.T is required to use the require package.
func NewRunner(t *testing.T, config *RunnerConfig) *Runner {
	return &Runner{
		t:      t,
		config: config,
	}
}

// Setup creates a new websocket client and generates the required number of wallets.
// It also connects to the websocket and starts the client.
// For every wallet, it will create a new account and fund it with the faucet.
func (r *Runner) Setup() error {
	r.client = r.config.Client

	if connectable, ok := r.client.(Connectable); ok {
		err := connectable.Connect()
		if err != nil {
			return err
		}
	}

	for i := 0; i < r.config.WalletCount; i++ {
		w, err := wallet.New(crypto.ED25519())
		if err != nil {
			return err
		}
		err = r.client.FundWallet(&w)
		if err != nil {
			return err
		}
		r.wallets = append(r.wallets, &w)
	}
	return nil
}

// Teardown closes the websocket client.
func (r *Runner) Teardown() error {
	if connectable, ok := r.client.(Connectable); ok {
		err := connectable.Disconnect()
		if err != nil {
			return err
		}
	}

	return nil
}

// TestTransaction submits a signed transaction and validates the result.
// If validate is nil, the transaction is not validated.
func (r *Runner) TestTransaction(flatTx *transaction.FlatTransaction, signer *wallet.Wallet) (*transactions.SubmitResponse, error) {
	tx, hash, err := r.processTransaction(flatTx, signer)
	if err != nil {
		return nil, err
	}

	require.NoError(r.t, err)
	require.Equal(r.t, tx.EngineResult, "tesSUCCESS")
	require.Equal(r.t, hash, tx.Tx["hash"].(string))

	return tx, nil
}

// TestMultisigTransaction submits a multisigned transaction and validates the result.
// If validate is nil, the transaction is not validated.
func (r *Runner) TestMultisigTransaction(blob string) (*transactions.SubmitMultisignedResponse, error) {
	tx, err := r.client.SubmitMultisigned(blob, true)
	if err != nil {
		return nil, err
	}

	require.NoError(r.t, err)
	require.Equal(r.t, tx.EngineResult, "tesSUCCESS")

	return tx, nil
}

// GetWallet returns a wallet by index.
func (r *Runner) GetWallet(index int) *wallet.Wallet {
	if index < 0 || index >= len(r.wallets) {
		return nil
	}
	return r.wallets[index]
}

// GetWallets returns all wallets.
func (r *Runner) GetWallets() []*wallet.Wallet {
	return r.wallets
}

// GetClient returns the websocket client.
func (r *Runner) GetClient() Client {
	return r.client
}

func (r *Runner) processTransaction(flatTx *transaction.FlatTransaction, signer *wallet.Wallet) (*transactions.SubmitResponse, string, error) {
	attempts := 0

	for {
		err := r.client.Autofill(flatTx)
		if err != nil {
			return nil, "", err
		}

		blob, hash, err := signer.Sign(*flatTx)
		if err != nil {
			return nil, hash, err
		}

		tx, err := r.client.Submit(blob, true)
		if err != nil {
			return nil, hash, err
		}

		if tx.EngineResult != "tefPAST_SEQ" || attempts >= r.config.MaxRetries {
			return tx, hash, nil
		}
		attempts++
	}
}
