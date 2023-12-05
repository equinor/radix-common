package errors

import (
	"fmt"
	"strings"
)

// Deprecated: Use errors.Join(errs) from golang standard library
// Concat Creates a single error from a list of errors
func Concat(errs []error) error {
	var errstrings []string
	for _, err := range errs {
		if err != nil {
			errstrings = append(errstrings, err.Error())
		}
	}

	if len(errstrings) > 0 {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
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
