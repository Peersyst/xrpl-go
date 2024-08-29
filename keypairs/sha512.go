package keypairs

import (
	"crypto/sha512"
	"encoding/binary"
	"hash"
	"math/big"
)

type Sha512 struct {
	hash hash.Hash
}

// NewSha512 creates a new Sha512 instance
func NewSha512() *Sha512 {
	return &Sha512{
		hash: sha512.New(),
	}
}

// Half creates a new Sha512 instance, adds the input, and returns the first 256 bits of the hash
func Half(input []byte) []byte {
	sha := NewSha512()
	sha.Add(input)
	return sha.First256()
}

// Add updates the hash with the given input
func (s *Sha512) Add(bytes []byte) *Sha512 {
	s.hash.Write(bytes)
	return s
}

// AddU32 adds a 32-bit unsigned integer to the hash
func (s *Sha512) AddU32(i uint32) *Sha512 {
	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, i)
	return s.Add(buffer)
}

// Finish finalizes the hash and returns the digest
func (s *Sha512) Finish() []byte {
	return s.hash.Sum(nil)
}

// First256 returns the first 256 bits of the hash
func (s *Sha512) First256() []byte {
	return s.Finish()[:32]
}

// First256BigInt converts the first 256 bits of the hash to a big integer
func (s *Sha512) First256BigInt() *big.Int {
	// return Bytes new(big.Int).SetBytes(s.First256())
	return BytesToNumberBE(s.First256())
}
