package errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Concat(t *testing.T) {
	var err1 error
	err2 := fmt.Errorf("Some non-nil error")
	var errs = []error{err1, err2}
	errCat := Concat(errs)
	assert.NotEqual(t, errCat, nil)

	err1 = nil
	err2 = nil
	errs = []error{err1, err2}
	errCat = Concat(errs)
	assert.Equal(t, errCat, nil)
}
