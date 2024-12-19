package interfaces

// KeypairCryptoAlg is an interface that defines the methods for a keypair crypto algorithm.
type KeypairCryptoAlg interface {
	DeriveKeypair(decodedSeed []byte, validator bool) (string, string, error)
	Sign(msg, privKey string) (string, error)
	Validate(msg, pubkey, sig string) bool
}

// NodeDerivationCryptoAlg is an interface that defines the methods for a node derivation crypto algorithm.
type NodeDerivationCryptoAlg interface {
	DerivePublicKeyFromPublicGenerator(pubKey []byte) ([]byte, error)
}
