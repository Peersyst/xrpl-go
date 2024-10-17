package interfaces

import "github.com/Peersyst/xrpl-go/binary-codec/definitions"

type BinaryParser interface {
	ReadByte() (byte, error)
	ReadField() (*definitions.FieldInstance, error)
	Peek() (byte, error)
	ReadBytes(n int) ([]byte, error)
	HasMore() bool
	ReadVariableLength() (int, error)
}
