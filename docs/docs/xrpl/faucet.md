# faucet

## Overview

`faucet` is a package that allows the user to get XRP for testing purposes on **testnet** and **devnet** ledgers and even from custom chains. To be able to fund your accounts programmatically, you can initialize the desired `FaucetProvider` for the ledger you want to use.

The package already exposes the `TestnetFaucetProvider` and `DevnetFaucetProvider` providers. If you want to use a custom chain, you can implement the `FaucetProvider` interface and use your own provider.

## Usage

To import the package, you can use the following code:

```go
import "github.com/Peersyst/xrpl-go/xrpl/faucet"
```

For devnet, you can use the following code:

```go
devnetFaucet := faucet.NewDevnetFaucetProvider()

err := devnetFaucet.FundWallet("rJ96831v5JXxna35JYvsW9VRmENwq23ib9")
if err != nil {
    // ...
}
```

for testnet, you can use the following code:

```go
testnetFaucet := faucet.NewTestnetFaucetProvider()

err := testnetFaucet.FundWallet("rJ96831v5JXxna35JYvsW9VRmENwq23ib9")
if err != nil {
    // ...
}
```