---
sidebar_position: 1
---

# Introduction

This documentation englobes the `xrpl-go` project, a Go SDK for interacting with the XRP Ledger.

## What is the XRP Ledger?

The XRP Ledger (XRPL) is a decentralized, open-source blockchain optimized for fast, low-cost transactions and financial applications. It enables tokenization and seamless cross-border payments without intermediaries.

To learn more about the XRP Ledger, you can visit the [official website](https://xrpl.org/).

## What is xrpl-go?

The [`xrpl-go`](https://github.com/Peersyst/xrpl-go) project is an SDK written in Go for interacting with the XRP Ledger. It provides a set of tools and libraries for building applications on the XRP Ledger.

The SDK can be split into the following packages:

- [`binary-codec`](packages/binary-codec.md): A package for encoding and decoding XRPL binary messages, objects and transactions.
- [`address-codec`](packages/address-codec.md): A package for encoding and decoding XRPL addresses.
- [`keypairs`](packages/keypairs.md): A package for generating and managing cryptographic keypairs.
- [`xrpl`](packages/xrpl.md): The biggest package of the SDK. It contains clients, types, transactions, and utils to interact with the XRP Ledger.
