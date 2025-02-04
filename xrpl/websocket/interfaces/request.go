package interfaces

type Request interface {
	Method() string
	Validate() error
	APIVersion() int
}
