# hash

## Overview

The `hash` package contains functions and types related to the XRPL hash types. Currently, it only contains the function `SignTxBlob` that hashes a signed transaction blob, which is mainly used for multisigning.

## Usage

To import the package, you can use the following code:

```go
import "github.com/Peersyst/xrpl-go/xrpl/hash"
```

## API

```go
func SignTxBlob(blob []byte, secret string) ([]byte, error)
```