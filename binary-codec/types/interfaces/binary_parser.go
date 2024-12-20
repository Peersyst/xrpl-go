package interfaces

import "github.com/Peersyst/xrpl-go/v1/binary-codec/definitions"

type BinaryParser interface {
	ReadByte() (byte, error)
	ReadField() (*definitions.FieldInstance, error)
	Peek() (byte, error)
	ReadBytes(n int) ([]byte, error)
	HasMore() bool
	ReadVariableLength() (int, error)
}
