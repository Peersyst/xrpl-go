package types

import (
	"bytes"
	"encoding/hex"

	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

type Currency struct{}

func (c *Currency) FromJSON(json any) ([]byte, error) {
	if str, ok := json.(string); ok {
		return c.fromString(str)
	}
	return nil, ErrInvalidCurrency
}

func (c *Currency) ToJSON(p interfaces.BinaryParser, opts ...int) (any, error) {
	currencyBytes, err := p.ReadBytes(opts[0])
	if err != nil {
		return nil, err
	}

	if bytes.Equal(currencyBytes, XRPBytes) {
		return "XRP", nil
	}

	// Check if bytes has exactly 3 non-zero bytes at positions 12-14
	nonZeroCount := 0
	var currencyStr string
	for i := 0; i < len(currencyBytes); i++ {
		if currencyBytes[i] != 0 {
			if i >= 12 && i <= 14 {
				nonZeroCount++
				currencyStr += string(currencyBytes[i])
			} else {
				nonZeroCount = 0
				break
			}
		}
	}

	if nonZeroCount == 3 {
		return currencyStr, nil
	}

	return hex.EncodeToString(currencyBytes), nil
}

func (c *Currency) fromString(str string) ([]byte, error) {
	if len(str) == 3 {
		var bytes [20]byte
		if str != "XRP" {
			isoBytes := []byte(str)
			copy(bytes[12:], isoBytes)
		}
		return bytes[:], nil
	}

	bytes, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
