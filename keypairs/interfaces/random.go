package interfaces

// Randomizer is an interface that defines the methods for a randomizer.
type Randomizer interface {
	GenerateBytes(n int) ([]byte, error)
}
