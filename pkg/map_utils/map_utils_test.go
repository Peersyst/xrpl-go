package maputils

import (
	"reflect"
	"testing"
)

func TestGetKeys(t *testing.T) {
	// Test case 1: Empty map
	m1 := make(map[string]interface{})
	expected1 := []string{}
	if keys1 := GetKeys(m1); !reflect.DeepEqual(keys1, expected1) {
		t.Errorf("Expected %v, but got %v", expected1, keys1)
	}

	// Test case 2: Map with one key-value pair
	m2 := map[string]interface{}{
		"key1": "value1",
	}
	expected2 := []string{"key1"}
	if keys2 := GetKeys(m2); !reflect.DeepEqual(keys2, expected2) {
		t.Errorf("Expected %v, but got %v", expected2, keys2)
	}

	// Test case 3: Map with multiple key-value pairs
	m3 := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	expected3 := []string{"key1", "key2", "key3"}
	if keys3 := GetKeys(m3); !reflect.DeepEqual(keys3, expected3) {
		t.Errorf("Expected %v, but got %v", expected3, keys3)
	}
}
