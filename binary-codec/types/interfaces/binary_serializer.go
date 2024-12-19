package interfaces

import "github.com/Peersyst/xrpl-go/v1/binary-codec/definitions"

type BinarySerializer interface {
	WriteFieldAndValue(fieldInstance definitions.FieldInstance, value []byte) error
	GetSink() []byte
}
