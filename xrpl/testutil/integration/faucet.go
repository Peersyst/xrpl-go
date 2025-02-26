package integration

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

const (
	LocalGenesisAddress types.Address = "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh"
	LocalGenesisSeed    string        = "snoPBrXtMeMyMHUVTgbuqAfg1SUTb"
)

// FundWallet funds a wallet with the client's faucet provider.
// If the faucet provider is nil, it will fund the wallet with the local genesis wallet.
func (f *Runner) FundWallet(wallet *wallet.Wallet) error {
	attempts := 0
	for {
		err := f.client.FundWallet(wallet)
		if err == nil {
			return nil
		}
		if attempts >= f.config.MaxRetries {
			break
		}
		attempts++
	}

	return f.fundWalletWithGenesis(wallet)
}

// fundWalletWithGenesis funds a wallet with the local genesis wallet.
func (f *Runner) fundWalletWithGenesis(w *wallet.Wallet) error {
	genesisWallet, err := wallet.FromSeed(LocalGenesisSeed, "")
	if err != nil {
		return err
	}
	genesisAddress := genesisWallet.GetAddress()

	payment := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account: genesisAddress,
		},
		Amount:      types.XRPCurrencyAmount(400000000),
		Destination: w.GetAddress(),
	}

	flatTx := payment.Flatten()
	_, err = f.TestTransaction(&flatTx, &genesisWallet, "tesSUCCESS")
	if err != nil {
		return err
	}

	return nil
}
