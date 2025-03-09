package errors

import "fmt"

// InvalidInputError is returned when user-input data is invalid.
type InvalidInputError struct {
	Field  string
	Reason string
}

// Error implements error interface.
func (e InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input in field %s: %s", e.Field, e.Reason)
}
