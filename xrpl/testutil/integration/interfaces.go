package integration

import (
	"github.com/Peersyst/xrpl-go/xrpl/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

type FaucetProvider interface {
	common.FaucetProvider
}

type Client interface {
	FaucetProvider() common.FaucetProvider

	FundWallet(wallet *wallet.Wallet) error
	Autofill(tx *transaction.FlatTransaction) error
	Submit(blob string, validate bool) (*transactions.SubmitResponse, error)
	SubmitMultisigned(blob string, validate bool) (*transactions.SubmitMultisignedResponse, error)
}

type Connectable interface {
	Connect() error
	Disconnect() error
}
