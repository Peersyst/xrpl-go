package testutil

import (
	"encoding/json"
	"fmt"
)

// Func to compare a flatten object and a JSON string
func CompareFlattenAndExpected(flattened map[string]interface{}, expected []byte) error {
	// Convert flattened to JSON
	flattenedJSON, err := json.Marshal(flattened)
	if err != nil {
		return fmt.Errorf("error marshaling payment flattened, error: %v", err)
	}

	// Normalize expected JSON
	var expectedMap map[string]interface{}
	if err := json.Unmarshal([]byte(expected), &expectedMap); err != nil {
		return fmt.Errorf("error unmarshaling expected, error: %v", err)
	}
	expectedJSON, err := json.Marshal(expectedMap)
	if err != nil {
		return fmt.Errorf("error marshaling expected payment object: %v", err)
	}

	// Compare JSON strings
	if string(flattenedJSON) != string(expectedJSON) {
		return fmt.Errorf("the flattened and expected JSON are not equal.\nGot:      %v\nExpected: %v", string(flattenedJSON), string(expectedJSON))
	}

	return nil
}
