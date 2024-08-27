package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointers(t *testing.T) {
	var p = Ptr(1337)
	v := Val(p)

	assert.Equal(t, 1337, v)
	assert.Equal(t, 0, Val[int](nil))
	assert.Equal(t, "", Val[string](nil))
	assert.Equal(t, false, Val[bool](nil))
}
