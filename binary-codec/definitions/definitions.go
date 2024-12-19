package definitions

import (
	_ "embed"

	"github.com/ugorji/go/codec"
)

var (
	//go:embed definitions.json
	docBytes []byte

	// definitions is the singleton instance of the Definitions struct.
	definitions *Definitions
)

type Definitions struct {
	Types              map[string]int32
	LedgerEntryTypes   map[string]int32
	Fields             fieldInstanceMap
	TransactionResults map[string]int32
	TransactionTypes   map[string]int32
	FieldIDNameMap     map[FieldHeader]string
}

func Get() *Definitions {
	return definitions
}

type definitionsDoc struct {
	Types              map[string]int32 `json:"TYPES"`
	LedgerEntryTypes   map[string]int32 `json:"LEDGER_ENTRY_TYPES"`
	Fields             fieldInstanceMap `json:"FIELDS"`
	TransactionResults map[string]int32 `json:"TRANSACTION_RESULTS"`
	TransactionTypes   map[string]int32 `json:"TRANSACTION_TYPES"`
}

// Loads JSON from the definitions file and converts it to a preferred format.
// The definitions file contains information required for the XRP Ledger's
// canonical binary serialization format:
// `Serialization <https://xrpl.org/serialization.html>`_
func loadDefinitions() {

	var jh codec.JsonHandle

	jh.MapKeyAsString = true
	jh.SignedInteger = true

	dec := codec.NewDecoderBytes(docBytes, &jh)
	var data definitionsDoc
	dec.MustDecode(&data)

	definitions = &Definitions{
		Types:              data.Types,
		Fields:             data.Fields,
		LedgerEntryTypes:   data.LedgerEntryTypes,
		TransactionResults: data.TransactionResults,
		TransactionTypes:   data.TransactionTypes,
	}

	addFieldHeadersAndOrdinals()
	createFieldIDNameMap()
}

func convertToFieldInstanceMap(m [][]interface{}) map[string]*FieldInstance {
	nm := make(map[string]*FieldInstance, len(m))

	for _, j := range m {
		k := j[0].(string)
		fi, _ := castFieldInfo(j[1])
		nm[k] = &FieldInstance{
			FieldName: k,
			FieldInfo: &fi,
			Ordinal:   fi.Nth,
		}
	}
	return nm
}

func castFieldInfo(v interface{}) (FieldInfo, error) {
	if fi, ok := v.(map[string]interface{}); ok {
		return FieldInfo{
			// TODO: Check if this is still needed
			//nolint:gosec // G115: Potential hardcoded credentials (gosec)
			Nth:            int32(fi["nth"].(int64)),
			IsVLEncoded:    fi["isVLEncoded"].(bool),
			IsSerialized:   fi["isSerialized"].(bool),
			IsSigningField: fi["isSigningField"].(bool),
			Type:           fi["type"].(string),
		}, nil
	}
	return FieldInfo{}, ErrUnableToCastFieldInfo
}

func addFieldHeadersAndOrdinals() {
	for k := range definitions.Fields {
		t, _ := definitions.GetTypeCodeByTypeName(definitions.Fields[k].Type)

		if fi, ok := definitions.Fields[k]; ok {
			fi.FieldHeader = &FieldHeader{
				TypeCode:  t,
				FieldCode: definitions.Fields[k].Nth,
			}
			fi.Ordinal = (t<<16 | definitions.Fields[k].Nth)
		}
	}
}

func createFieldIDNameMap() {
	definitions.FieldIDNameMap = make(map[FieldHeader]string, len(definitions.Fields))
	for k := range definitions.Fields {
		fh, _ := definitions.GetFieldHeaderByFieldName(k)

		definitions.FieldIDNameMap[*fh] = k
	}
}
