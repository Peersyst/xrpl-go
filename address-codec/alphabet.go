// Huge thanks to the btcsuite developers for creating this code, which we adapted for our use in compliance with the Copyfree Initiative.
// btcsuite base58 repo: https://github.com/btcsuite/btcd/tree/master/btcutil/base58
//
// AUTOGENERATED by gen_alphabet.go; do not edit.
package addresscodec

const (
	// alphabet is the modified base58 alphabet used by XRP.
	xrpAlphabet  = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
	alphabetIdx0 = 'r'
)

var b58 = [256]byte{
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 50, 33, 7, 21, 41, 40, 27,
	45, 8, 255, 255, 255, 255, 255, 255,
	255, 54, 10, 38, 12, 14, 47, 15,
	16, 255, 17, 18, 19, 20, 13, 255,
	22, 23, 24, 25, 26, 11, 28, 29,
	30, 31, 32, 255, 255, 255, 255, 255,
	255, 5, 34, 35, 36, 37, 6, 39,
	3, 49, 42, 43, 255, 44, 4, 46,
	1, 48, 0, 2, 51, 52, 53, 9,
	55, 56, 57, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
}
