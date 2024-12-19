# XRPL-GO

[![Go Reference](https://pkg.go.dev/badge/github.com/Peersyst/xrpl-go/v1.svg)](https://pkg.go.dev/github.com/Peersyst/xrpl-go/v1)
[![Go Report Card](https://goreportcard.com/badge/github.com/Peersyst/xrpl-go/v1)](https://goreportcard.com/report/github.com/Peersyst/xrpl-go/v1)
[![Release Card](https://img.shields.io/github/v/release/Peersyst/xrpl-go?include_prereleases)](https://github.com/Peersyst/xrpl-go/v1/releases)


The `xrpl-go` library provides a Go implementation for interacting with the XRP Ledger. From serialization to signing transactions, the library allows users to work with the most
complex elements of the XRP Ledger. A full library of models for all transactions and core server rippled API objects are provided.

## Disclaimer
This library is still in development and not yet intended for production use.

## Requirements

Requiring Go version `1.22.0` and later.
[Download latest Go version](https://go.dev/dl/)

## Packages

| Name | Description |
|---------|-------------|
| addresscodec | Provides functions for encoding and decoding XRP Ledger addresses |
| binarycodec | Implements binary serialization and deserialization of XRP Ledger objects |
| keypairs | Handles generation and management of cryptographic key pairs for XRP Ledger accounts |
| xrpl | Core package containing the main functionality for interacting with the XRP Ledger |
| examples | Contains example code demonstrating usage of the xrpl-go library |

## Report an issue

If you find any issues, please report them to the [XRPL-GO GitHub repository](https://github.com/Peersyst/xrpl-go/issues). 

## License
The `xrpl-go` library is licensed under the MIT License.
