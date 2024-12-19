package testutil

import (
	"testing"

	definitions "github.com/Peersyst/xrpl-go/v1/binary-codec/definitions"
)

func GetFieldInstance(t *testing.T, fieldName string) definitions.FieldInstance {
	t.Helper()
	fi, err := definitions.Get().GetFieldInstanceByFieldName(fieldName)
	if err != nil {
		t.Fatalf("FieldInstance with FieldName %v", fieldName)
	}
	return *fi
}
