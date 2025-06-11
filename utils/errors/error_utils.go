package errors

import (
	"errors"
)

// Deprecated: Use errors.Join(errs) from golang standard library
// Concat Creates a single error from a list of errors
func Concat(errs []error) error {
	return errors.Join(errs...)
}

// Deprecated: Use errors.Join() and error.Is() instead
// Contains Check if error is contained in slice
func Contains(errs []error, err error) bool {
	for _, a := range errs {
		if a.Error() == err.Error() {
			return true
		}
	}
	return false
}
