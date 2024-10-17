package errors

// ErrNotValidJson is an error that occurs when the json is not valid.
type ErrNotValidJson struct{}

// Error returns the error message for the ErrNotValidJson error.
func (e *ErrNotValidJson) Error() string {
	return "not a valid json"
}
