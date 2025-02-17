# rpc

## Overview

The `rpc` package provides the RPC client for interacting with the XRPL network via its RPC API. This client handles the communication with XRPL nodes, allowing you to:

- Send requests to query the ledger state
- Submit transactions to the network 
- Receive responses and handle errors
- Manage the connections configuration 

## Client

The `rpc` package provides a `Client` type which and communication with XRPL nodes. This client is configurable and let the user submit transactions and make queries. 

In order to create a new `Client`, you can use the `NewClient` function:

```go
cfg, err := rpc.NewClientConfig("<url>")
if err != nil {
	// ...
}
client := rpc.NewClient(cfg)
```

Every time you create a new `Client`, you need to pass a `Config` struct as argument. You can initialize a `Config` struct using the `NewClientConfig` function.

`Config` struct follows the options pattern, so you can pass different options to the `NewClientConfig` function:

```go
// RPC Client config options
func WithHTTPClient(cl HTTPClient) ConfigOpt
func WithMaxFeeXRP(maxFeeXRP float32) ConfigOpt
func WithFeeCushion(feeCushion float32) ConfigOpt
func WithFaucetProvider(fp common.FaucetProvider) ConfigOpt
```

So, for example, if you want to set a custom `FaucetProvider` and `FeeCushion`, you can do it this way:

```go
cfg, err := rpc.NewClientConfig("https://s.altnet.rippletest.net:51234/",
	rpc.WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	rpc.WithFeeCushion(1.5),
)
if err != nil {
	// ...
}
client := rpc.NewClient(cfg)
```

## Methods

`Client` offers different methods to interact with the XRPL network.

### Request

The `Request` method is used to make queries to the XRPL network. It returns a `XRPLResponse` interface. This method is used in the client's queries requests.

```go
// Client methods
func (c *Client) Request(reqParams XRPLRequest) (XRPLResponse, error)
```

### Submit/SubmitMultisigned

The `Submit` method is used to submit a transaction to the XRPL network. It returns a `TxResponse` struct containing the transaction result for the blob submitted. `txBlob` must be signed. There's also a `SubmitMultisigned` method that works the same way but for multisigned transactions.

```go
func (c *Client) Submit(txBlob string, failHard bool) (*requests.TxResponse, error)
func (c *Client) SubmitMultisigned(txBlob string, failHard bool) (*requests.SubmitMultisignedResponse, error)
```

### Autofill/AutofillMultisigned

The `Autofill` method is used to autofill some fields in a flat transaction. This method is useful for adding dynamic fields like `LastLedgerSequence` or `Fee`. It returns an error if the transaction is not valid or some internall call fails. There's also a `AutofillMultisigned` method that works the same way but for multisigned transactions.

```go
func (c *Client) Autofill(tx *transaction.FlatTransaction) error
func (c *Client) AutofillMultisigned(tx *transaction.FlatTransaction, nSigners uint64) error
```

### SubmitAndWait

The `SubmitAndWait` method is used to submit a transaction to the XRPL network and wait for it to be included in a ledger. It returns a `TxResponse` struct containing the transaction result for the blob submitted. 

```go
func (c *Client) SubmitAndWait(txBlob string, failHard bool) (*requests.TxResponse, error)
```

## Queries

`Client` also exposes methods to make queries to the XRPL network. These methods are wrappers of the queries requests exposed by the [`queries`](/xrpl/queries/) package.

## Usage

To use the `rpc` package, you need to import it in your project:

```go
import "github.com/Peersyst/xrpl-go/xrpl/rpc"
```

## Examples

### How to send a payment transaction

This example shows how to send a payment transaction to the XRPL testnet with the `rpc` package.

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/Peersyst/xrpl-go/xrpl/currency"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

func main() {

	// Create a new rpc client config with a testnet faucet provider
	cfg, err := rpc.NewClientConfig(
		"https://s.altnet.rippletest.net:51234/",
		rpc.WithMaxFeeXRP(5.0),
		rpc.WithFeeCushion(1.5),
		rpc.WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}

	// Create a new rpc client with the config
	client := rpc.NewClient(cfg)

	// Create a new wallet with the ed25519 algorithm
	w, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fund the wallet with the testnet faucet
	if err := client.FundWallet(&w); err != nil {
		fmt.Println(err)
		return
	}

	// Convert the amount to drops
	xrpAmount, err := currency.XrpToDrops("1")
	if err != nil {
		fmt.Println(err)
		return
	}

	xrpAmountInt, err := strconv.ParseInt(xrpAmount, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account: types.Address(w.GetAddress()),
		},
		Destination: "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		Amount:      types.XRPCurrencyAmount(xrpAmountInt),
		DeliverMax:  types.XRPCurrencyAmount(xrpAmountInt),
	}

	flattenedTx := p.Flatten()

	// Autofill the transaction with the client's config
	if err := client.Autofill(&flattenedTx); err != nil {
		fmt.Println(err)
		return
	}

	// Sign the transaction with the wallet
	txBlob, _, err := w.Sign(flattenedTx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Submit the transaction to the network and wait for it to be included in a ledger
	res, err := client.SubmitAndWait(txBlob, false)
	if err != nil {
		fmt.Println(err)
		return
	}
}
```