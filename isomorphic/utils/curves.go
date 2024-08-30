package utils

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

func HexToNumber(hexStr string) (*big.Int, error) {
	if hexStr == "" {
		hexStr = "00"
	}
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(bytes), nil
}

func HexToBytes(hexStr string) ([]byte, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

/**
 * @example bytesToHex(Uint8Array.from([0xca, 0xfe, 0x01, 0x23])) // 'cafe0123'
 */
func bytesToHex(bytes []byte) string {
	hex := ""
	for i := 0; i < len(bytes); i++ {
		hex += fmt.Sprintf("%02x", bytes[i])
	}
	return hex
}

// BE: Big Endian, LE: Little Endian
func BytesToNumberBE(bytes []byte) (*big.Int, error) {
	return HexToNumber(bytesToHex(bytes))
}

func BytesToNumberLE(bytes []byte) (*big.Int, error) {
	reversedBytes := make([]byte, len(bytes))
	copy(reversedBytes, bytes)
	for i, j := 0, len(reversedBytes)-1; i < j; i, j = i+1, j-1 {
		reversedBytes[i], reversedBytes[j] = reversedBytes[j], reversedBytes[i]
	}
	return HexToNumber(bytesToHex(reversedBytes))
}

func NumberToBytesBE(n *big.Int, length int) ([]byte, error) {
	hexStr := n.Text(16)
	hexStr = strings.Repeat("0", length*2-len(hexStr)) + hexStr
	return HexToBytes(hexStr)
}

func NumberToBytesLE(n *big.Int, len int) ([]byte, error) {
	bytesBE, err := NumberToBytesBE(n, len)
	if err != nil {
		return nil, err
	}
	reversedBytes := make([]byte, len)
	for i := 0; i < len; i++ {
		reversedBytes[i] = bytesBE[len-i-1]
	}
	return reversedBytes, nil
}
