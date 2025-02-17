# time

## Overview

This package contains functions to handle with XRPL time conversions. It enables conversions between RippleTime and UnixTime. To learn more about RippleTime and UnixTime, you can read the 
[official documentation](https://xrpl.org/docs/references/protocol/data-types/basic-data-types#specifying-time).

## Usage

To import the package, you can use the following code:

```go
import "github.com/Peersyst/xrpl-go/xrpl/time"
```

## API

The following functions are available:

```go
func RippleTimeToUnixTime(rpepoch int64) int64
func UnixTimeToRippleTime(timestamp int64) int64
func RippleTimeToISOTime(rippleTime int64) string
func IsoTimeToRippleTime(isoTime string) (int64, error)
```