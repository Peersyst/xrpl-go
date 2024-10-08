package typecheck

import (
	"regexp"
	"strconv"
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

// Checks if the given string is a valid number (Float32).
func IsFloat32(s string) bool {
	_, err := strconv.ParseFloat(s, 32)
	return err == nil
}

// Checks if the given string is a valid number (Float64).
func IsFloat64(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
