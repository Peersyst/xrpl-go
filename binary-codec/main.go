package binarycodec

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
)

func Encode(json map[string]interface{}) (string, error) {

	fimap, err := createFieldInstanceMapFromJson(json)

	if err != nil {
		return "", err
	}

	sk := getSortedKeys(fimap)

	var sink []byte

	for _, v := range sk {

		h, err := EncodeFieldID(v.FieldName)

		if err != nil {
			return "", err
		}

		sink = append(sink, h...)
		fmt.Println(hex.EncodeToString(sink))

		// fmt.Println(fimap[v])
		// val := uint32(fimap[v].(int))

		// need to write bytes to new buffers
		// amount, uint, hash all big endian
		buf := new(bytes.Buffer)
		err = binary.Write(buf, binary.BigEndian, uint32(fimap[v].(int)))

		if err != nil {
			return "", err
		}

		fmt.Println(buf.Bytes())
		sink = append(sink, buf.Bytes()...)
		// fmt.Println(hex.EncodeToString(sink))
	}

	// Loop through and create map of map[FieldInstance]interface{}
	// Sort by Ordinal
	// Start serializing
	//	optimize encode from field id codec, making same call twice

	// fmt.Println(string(sink))

	return hex.EncodeToString(sink), nil
}

// func Serialize(json string) (string, error) {
// 	return "120007220008000024001ABED82A2380BF2C2019001ABED764D55920AC9391400000000000000000000000000055534400000000000A20B3C85F482532A9578DBB3950B85CA06594D165400000037E11D60068400000000000000A732103EE83BB432547885C219634A1BC407A9DB0474145D69737D09CCDC63E1DEE7FE3744630440220143759437C04F7B61F012563AFE90D8DAFC46E86035E1D965A9CED282C97D4CE02204CFD241E86F17E011298FC1A39B63386C74306A5DE047E213B0F29EFA4571C2C8114DD76483FACDEE26E60D8A586BB58D09F27045C46", nil
// }

//lint:ignore U1000 // ignore this for now
//nolint
func createFieldInstanceMapFromJson(json map[string]interface{}) (map[definitions.FieldInstance]any, error) {

	m := make(map[definitions.FieldInstance]interface{}, len(json))

	for k, v := range json {
		fi, err := definitions.Get().GetFieldInstanceByFieldName(k)

		if err != nil {
			return nil, err
		}

		m[*fi] = v

	}

	return m, nil
}

//lint:ignore U1000 // ignore this for now
//nolint
func getSortedKeys(m map[definitions.FieldInstance]interface{}) []definitions.FieldInstance {
	keys := make([]definitions.FieldInstance, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i].Ordinal < keys[j].Ordinal
	})

	return keys
}
