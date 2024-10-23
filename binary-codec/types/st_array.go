package types

import (
	"errors"

	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

const (
	ArrayEndMarker  = 0xF1
	ObjectEndMarker = 0xE1
)

// STArray represents an array of STObject instances.
type STArray struct{}

var ErrNotSTObjectInSTArray = errors.New("not STObject in STArray. Array fields must be STObjects")

// FromJson is a method that takes a JSON value (which should be a slice of JSON objects),
// and converts it to a byte slice, representing the serialized form of the STArray.
// It loops through the JSON slice, and for each element, calls the FromJson method
// of an STObject, appending the resulting byte slice to a "sink" slice.
// The method returns an error if the JSON value is not a slice.
func (t *STArray) FromJson(json any) ([]byte, error) {
	if _, ok := json.([]any); !ok {
		return nil, ErrNotSTObjectInSTArray
	}

	var sink []byte
	for _, v := range json.([]any) {
		st := &STObject{}
		b, err := st.FromJson(v)
		if err != nil {
			return nil, err
		}
		sink = append(sink, b...)
	}
	sink = append(sink, ArrayEndMarker)

	return sink, nil
}

// ToJson is a method that takes a BinaryParser and optional parameters, and converts
// the serialized byte data back to a JSON value.
// The method loops until the BinaryParser has no more data, and for each loop,
// it calls the ToJson method of an STObject, appending the resulting JSON value to a "value" slice.
func (t *STArray) ToJson(p interfaces.BinaryParser, opts ...int) (any, error) {
	var value []any
	count := 0

	for p.HasMore() {

		stObj := make(map[string]any)
		fi, err := p.ReadField()
		if err != nil {
			return nil, err
		}
		if count == 0 && fi.Type != "STObject" {
			return nil, ErrNotSTObjectInSTArray
		} else if fi.FieldName == "ArrayEndMarker" {
			break
		}
		fn := fi.FieldName
		st := GetSerializedType(fi.Type)
		res, err := st.ToJson(p)
		if err != nil {
			return nil, err
		}
		stObj[fn] = res
		value = append(value, stObj)
		count++
	}
	return value, nil
}
