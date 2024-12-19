package addresscodec

import (
	"bytes"
)

var (
	MainnetXAddressPrefix = []byte{0x05, 0x44}
	TestnetXAddressPrefix = []byte{0x04, 0x93}
	// X-address length - value is 35
	XAddressLength = 35
)

// IsValidXAddress returns true if the x-address is valid. Otherwise, it returns false.
func IsValidXAddress(xAddress string) bool {
	_, _, _, err := DecodeXAddress(xAddress)
	return err == nil
}

// EncodeXAddress returns the x-address encoding of the accountId, tag, and testnet boolean.
// If the accountId is not 20 bytes long, it returns an error.
func EncodeXAddress(accountID []byte, tag uint32, tagFlag, testnetFlag bool) (string, error) {
	if len(accountID) != AccountAddressLength {
		return "", ErrInvalidAccountID
	}

	xAddressBytes := make([]byte, 0, XAddressLength)

	if testnetFlag {
		xAddressBytes = append(xAddressBytes, TestnetXAddressPrefix...)
	} else {
		xAddressBytes = append(xAddressBytes, MainnetXAddressPrefix...)
	}

	xAddressBytes = append(xAddressBytes, accountID...)

	if tagFlag {
		xAddressBytes = append(xAddressBytes, byte(1))
	} else {
		xAddressBytes = append(xAddressBytes, byte(0))
	}

	xAddressBytes = append(
		xAddressBytes,
		byte(tag&0xff),
		byte((tag>>8)&0xff),
		byte((tag>>16)&0xff),
		byte((tag>>24)&0xff),
		0,
		0,
		0,
		0,
	)

	cksum := checksum(xAddressBytes)
	xAddressBytes = append(xAddressBytes, cksum[:]...)

	return EncodeBase58(xAddressBytes), nil
}

// DecodeXAddress returns the accountId, tag, and testnet boolean decoding of the x-address.
// If the x-address is invalid, it returns an error.
func DecodeXAddress(xAddress string) (accountID []byte, tag uint32, testnet bool, err error) {
	xAddressBytes := DecodeBase58(xAddress)

	switch {
	case bytes.HasPrefix(xAddressBytes, MainnetXAddressPrefix):
		testnet = false
	case bytes.HasPrefix(xAddressBytes, TestnetXAddressPrefix):
		testnet = true
	default:
		return nil, 0, false, ErrInvalidXAddress
	}

	tag, err = decodeTag(xAddressBytes)
	if err != nil {
		return nil, 0, false, err
	}

	return xAddressBytes[2:22], tag, testnet, nil
}

// XAddressToClassicAddress converts the x-address to a classic address.
// It returns the classic address, tag and testnet boolean.
// If the x-address is invalid, it returns an error.
func XAddressToClassicAddress(xAddress string) (classicAddress string, tag uint32, testnet bool, err error) {
	accountID, tag, testnet, err := DecodeXAddress(xAddress)
	if err != nil {
		return "", 0, false, err
	}

	classicAddress, err = EncodeAccountIDToClassicAddress(accountID)
	if err != nil {
		return "", 0, false, err
	}

	return classicAddress, tag, testnet, nil
}

// ClassicAddressToXAddress converts the classic address to an x-address.
// It returns the x-address.
// If the classic address is invalid, it returns an error.
func ClassicAddressToXAddress(address string, tag uint32, tagFlag, testnetFlag bool) (string, error) {
	_, accountID, err := DecodeClassicAddressToAccountID(address)
	if err != nil {
		return "", err
	}

	return EncodeXAddress(accountID, tag, tagFlag, testnetFlag)
}

// decodeTag returns the tag from the x-address.
// If the tag is invalid, it returns an error.
func decodeTag(xAddressBytes []byte) (uint32, error) {
	switch {
	case xAddressBytes[22] > 1:
		return 0, ErrInvalidTag
	case xAddressBytes[22] == 1:
		return uint32(xAddressBytes[23]) + uint32(xAddressBytes[24])*256 + uint32(xAddressBytes[25])*65536, nil
	default:
		return 0, nil
	}
}
