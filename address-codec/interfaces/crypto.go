package interfaces

// CryptoImplementation defines an interface for implementing cryptographic operations
// required by address codec.
type CryptoImplementation interface {
	DeriveKeypair(decodedSeed []byte, validator bool) (string, string, error)
	Sign(msg, privKey string) (string, error)
	Validate(msg, pubkey, sig string) bool
}
