package utils

import (
	"bytes"
	"math/big"
	"testing"
)

func TestHexToNumber(t *testing.T) {
	t.Run("Empty hex string should return 0", func(t *testing.T) {
		result, err := HexToNumber("")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := big.NewInt(0)
		if result.Cmp(expected) != 0 {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Valid hex string should return corresponding number", func(t *testing.T) {
		result, err := HexToNumber("123456")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := big.NewInt(0x123456)
		if result.Cmp(expected) != 0 {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Invalid hex string should return error", func(t *testing.T) {
		_, err := HexToNumber("hello")
		if err == nil {
			t.Error("Expected error, but got nil")
		}
	})
}
func TestHexToBytes(t *testing.T) {
	t.Run("Empty hex string should return empty byte slice", func(t *testing.T) {
		result, err := HexToBytes("")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := []byte{}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("String should return corresponding byte slice", func(t *testing.T) {
		result, err := HexToBytes("48656c6c6f20576f726c64") // "Hello World" in hex
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Invalid hex string should return error", func(t *testing.T) {
		_, err := HexToBytes("hello")
		if err == nil {
			t.Error("Expected error, but got nil")
		}
	})

	t.Run("tring should return corresponding byte slice", func(t *testing.T) {
		result, err := HexToBytes("DEADBEEF") // "Hello World" in hex
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := []byte{222, 173, 190, 239}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
func TestBytesToNumberBE(t *testing.T) {
	t.Run("Empty byte slice should return 0", func(t *testing.T) {
		result, err := BytesToNumberBE([]byte{})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := big.NewInt(0)
		if result.Cmp(expected) != 0 {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Byte slice with single byte should return corresponding number", func(t *testing.T) {
		result, err := BytesToNumberBE([]byte{0x12})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := big.NewInt(0x12)
		if result.Cmp(expected) != 0 {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Byte slice with multiple bytes should return corresponding number", func(t *testing.T) {
		result, err := BytesToNumberBE([]byte{0x12, 0x34, 0x56})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := big.NewInt(0x123456)
		if result.Cmp(expected) != 0 {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
func TestBytesToNumberLE(t *testing.T) {
	t.Run("Empty byte slice should return 0", func(t *testing.T) {
		result, err := BytesToNumberLE([]byte{})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := big.NewInt(0)
		if result.Cmp(expected) != 0 {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Byte slice with single byte should return corresponding number", func(t *testing.T) {
		result, err := BytesToNumberLE([]byte{0x12})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := big.NewInt(0x12)
		if result.Cmp(expected) != 0 {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Byte slice with multiple bytes should return corresponding number", func(t *testing.T) {
		result, err := BytesToNumberLE([]byte{0x12, 0x34, 0x56})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := big.NewInt(0x563412)
		if result.Cmp(expected) != 0 {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
func TestNumberToBytesBE(t *testing.T) {
	t.Run("Number 0 should return byte slice with all zeros", func(t *testing.T) {
		num := big.NewInt(0)
		result, err := NumberToBytesBE(num, 4)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := []byte{0x00, 0x00, 0x00, 0x00}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Number 255 should return byte slice with all 0xFF", func(t *testing.T) {
		num := big.NewInt(255)
		result, err := NumberToBytesBE(num, 4)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := []byte{0x00, 0x00, 0x00, 0xFF}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Number 65535 should return byte slice with 0xFF in the last two bytes", func(t *testing.T) {
		num := big.NewInt(65535)
		result, err := NumberToBytesBE(num, 4)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := []byte{0x00, 0x00, 0xFF, 0xFF}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Number 4294967295 should return byte slice with all 0xFF", func(t *testing.T) {
		num := big.NewInt(4294967295)
		result, err := NumberToBytesBE(num, 4)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := []byte{0xFF, 0xFF, 0xFF, 0xFF}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
func TestNumberToBytesLE(t *testing.T) {
	t.Run("Number 0 should return byte slice with all zeros", func(t *testing.T) {
		num := big.NewInt(0)
		result, err := NumberToBytesLE(num, 4)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := []byte{0x00, 0x00, 0x00, 0x00}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Number 255 should return byte slice with all 0xFF", func(t *testing.T) {
		num := big.NewInt(255)
		result, err := NumberToBytesLE(num, 4)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := []byte{0xFF, 0x00, 0x00, 0x00}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Number 65535 should return byte slice with 0xFF in the first two bytes", func(t *testing.T) {
		num := big.NewInt(65535)
		result, err := NumberToBytesLE(num, 4)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := []byte{0xFF, 0xFF, 0x00, 0x00}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Number 4294967295 should return byte slice with all 0xFF", func(t *testing.T) {
		num := big.NewInt(4294967295)
		result, err := NumberToBytesLE(num, 4)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := []byte{0xFF, 0xFF, 0xFF, 0xFF}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
