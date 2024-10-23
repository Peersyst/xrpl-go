package types

import (
	"fmt"
	"sort"

	"github.com/Peersyst/xrpl-go/binary-codec/definitions"
	"github.com/Peersyst/xrpl-go/binary-codec/serdes"
	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

// STObject represents a map of serialized field instances, where each key is a field name
// and the associated value is the field's value. This structure allows us to represent nested
// and complex structures of the Ripple protocol.
type STObject struct{}

// FromJson converts a JSON object into a serialized byte slice.
// It works by converting the JSON object into a map of field instances (which include the field definition
// and value), and then serializing each field instance.
// This method returns an error if the JSON input is not a valid object.
func (t *STObject) FromJson(json any) ([]byte, error) {
	s := serdes.NewSerializer()
	if _, ok := json.(map[string]any); !ok {
		return nil, fmt.Errorf("not a valid json node")
	}
	fimap, err := createFieldInstanceMapFromJson(json.(map[string]any))

	if err != nil {
		return nil, err
	}

	sk := getSortedKeys(fimap)

	for _, v := range sk {
		if !v.IsSerialized {
			continue
		}

		st := GetSerializedType(v.Type)
		b, err := st.FromJson(fimap[v])
		if err != nil {
			return nil, err
		}
		err = s.WriteFieldAndValue(v, b)
		if err != nil {
			return nil, err
		}
	}
	return s.GetSink(), nil
}

// ToJson takes a BinaryParser and optional parameters, and converts the serialized byte data
// back to a JSON value. It will continue parsing until it encounters an end marker for an object
// or an array, or until the parser has no more data.
func (t *STObject) ToJson(p interfaces.BinaryParser, opts ...int) (any, error) {
	m := make(map[string]any)

	for p.HasMore() {

		fi, err := p.ReadField()
		if err != nil {
			return nil, err
		}

		if fi.FieldName == "ObjectEndMarker" || fi.FieldName == "ArrayEndMarker" {
			break
		}

		st := GetSerializedType(fi.Type)

		var res any
		if fi.IsVLEncoded {
			vlen, err := p.ReadVariableLength()
			if err != nil {
				return nil, err
			}
			res, err = st.ToJson(p, vlen)
			if err != nil {
				return nil, err
			}

		} else {
			res, err = st.ToJson(p)
			if err != nil {
				return nil, err
			}
		}
		res, err = enumToStr(fi.FieldName, res)
		if err != nil {
			return nil, err
		}

		m[fi.FieldName] = res
	}
	return m, nil
}

// nolint
// createFieldInstanceMapFromJson creates a map of field instances from a JSON object.
// Each key-value pair in the JSON object is converted into a field instance, where the key
// represents the field name and the value is the field's value.
//
//lint:ignore U1000 // ignore this for now
func createFieldInstanceMapFromJson(json map[string]any) (map[definitions.FieldInstance]any, error) {
	m := make(map[definitions.FieldInstance]any, len(json))

	for k, v := range json {
		fi, err := definitions.Get().GetFieldInstanceByFieldName(k)

		if err != nil {
			return nil, err
		}

		m[*fi] = v
	}
	return m, nil
}

// nolint
//
// getSortedKeys is a helper function to sort the keys of a map of field instances based on
// their ordinal values. This is used to ensure that the fields are serialized in the
// correct order.
//
//lint:ignore U1000 // ignore this for now
func getSortedKeys(m map[definitions.FieldInstance]any) []definitions.FieldInstance {
	keys := make([]definitions.FieldInstance, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i].Ordinal < keys[j].Ordinal
	})
	return keys
}

// enumToStr is a helper function that takes a field name and its associated value,
// and returns a string representation of the value if the field is an enumerated type
// (i.e., TransactionType, TransactionResult, LedgerEntryType).
// If the field is not an enumerated type, the original value is returned.
func enumToStr(fieldType string, value any) (any, error) {
	switch fieldType {
	case "TransactionType":
		return definitions.Get().GetTransactionTypeNameByTransactionTypeCode(int32(value.(int)))
	case "TransactionResult":
		return definitions.Get().GetTransactionResultNameByTransactionResultTypeCode(int32(value.(int)))
	case "LedgerEntryType":
		return definitions.Get().GetLedgerEntryTypeNameByLedgerEntryTypeCode(int32(value.(int)))
	default:
		return value, nil
	}
}
