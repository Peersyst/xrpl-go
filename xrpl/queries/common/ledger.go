package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type LedgerSpecifier interface {
	Ledger() string
}

func UnmarshalLedgerSpecifier(data []byte) (LedgerSpecifier, error) {
	if len(data) == 0 {
		return nil, nil
	}
	switch data[0] {
	case '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return nil, err
		}
		switch s {
		case Current.Ledger():
			return Current, nil
		case Validated.Ledger():
			return Validated, nil
		case Closed.Ledger():
			return Closed, nil
		}
		return nil, fmt.Errorf("decoding LedgerTitle: invalid string %s", s)
	default:
		var i LedgerIndex
		if err := json.Unmarshal(data, &i); err != nil {
			return nil, err
		}
		return i, nil
	}
}

type LedgerIndex uint32

func (l LedgerIndex) Ledger() string {
	return strconv.FormatUint(uint64(l), 10)
}

func (l LedgerIndex) Uint32() uint32 {
	return uint32(l)
}

func (l LedgerIndex) Int() int {
	return int(l)
}

type LedgerTitle string

const (
	Current   LedgerTitle = "current"
	Validated LedgerTitle = "validated"
	Closed    LedgerTitle = "closed"
)

func (l LedgerTitle) Ledger() string {
	return string(l)
}

type LedgerHash string
