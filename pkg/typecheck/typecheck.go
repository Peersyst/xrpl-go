package typecheck

import (
	"reflect"
	"regexp"
)

// IsString checks if the given interface is a string.
func IsString(str interface{}) bool {
	_, ok := str.(string)
	return ok
}

// IsUint32 checks if the given interface is a uint32.
func IsUint32(num interface{}) bool {
	_, ok := num.(uint32)
	return ok
}

// IsUint64 checks if the given interface is a uint64.
func IsUint64(num interface{}) bool {
	_, ok := num.(uint64)
	return ok
}

// IsUint checks if the given interface is a uint.
func IsUint(num interface{}) bool {
	_, ok := num.(uint)
	return ok
}

// IsInt checks if the given interface is an int.
func IsInt(num interface{}) bool {
	_, ok := num.(int)
	return ok
}

// IsBool checks if the given interface is a bool.
func IsBool(b interface{}) bool {
	_, ok := b.(bool)
	return ok
}

// IsHex checks if the given string is a valid hexadecimal string.
func IsHex(s string) bool {
	// Define a regular expression for a valid hexadecimal string
	var validHexPattern = regexp.MustCompile(`^[0-9a-fA-F]+$`)
	return validHexPattern.MatchString(s)
}

// IsMap checks if the given interface is a map.
func IsMap(m interface{}) bool {
	_, ok := m.(map[string]interface{})
	return ok
}

// IsArrayOrSlice checks if the given value is an array or a slice.
func IsArrayOrSlice(value interface{}) bool {
	v := reflect.ValueOf(value)
	return v.Kind() == reflect.Array || v.Kind() == reflect.Slice
}
