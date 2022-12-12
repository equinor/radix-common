package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PointersOf(t *testing.T) {
	type testObj struct{ prop string }
	obj1 := testObj{prop: "obj1prop"}
	obj2 := testObj{prop: "obj2prop"}
	expected := []*testObj{&obj1, &obj2}
	actual := PointersOf([]testObj{obj1, obj2})
	assert.Equal(t, expected, actual)
}

func Test_Map(t *testing.T) {
	expected := []string{"10", "20", "30"}
	actual := Map([]int{1, 2, 3}, func(v int) string { return fmt.Sprintf("%d", v*10) })
	assert.Equal(t, expected, actual)
}
