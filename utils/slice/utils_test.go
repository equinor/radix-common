package slice

import (
	"fmt"
	"strings"
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

func Test_Any(t *testing.T) {
	type testType struct{ val int }
	testData := []testType{{val: 1}, {val: 2}, {val: 3}}
	assert.True(t, Any(testData, func(o testType) bool { return o.val == 2 }))
	assert.False(t, Any(testData, func(o testType) bool { return o.val == 4 }))
	assert.False(t, Any([]testType{}, func(o testType) bool { return o.val == 4 }))
}

func Test_All(t *testing.T) {
	type testType struct{ val int }
	testData := []testType{{val: 1}, {val: 2}, {val: 3}}
	assert.True(t, All(testData, func(o testType) bool { return o.val > 0 }))
	assert.False(t, All(testData, func(o testType) bool { return o.val > 1 }))
	assert.True(t, All([]testType{}, func(o testType) bool { return o.val > 0 }))
}

func Test_FindAll(t *testing.T) {
	type testType struct {
		id   int
		name string
	}
	testData := []testType{{id: 1, name: "foo"}, {id: 2, name: "bar"}, {id: 3, name: "baz"}, {id: 4, name: "foo2"}, {id: 5, name: "bar2"}}
	expected := []testType{{id: 2, name: "bar"}, {id: 3, name: "baz"}, {id: 5, name: "bar2"}}
	actual := FindAll(testData, func(v testType) bool { return strings.HasPrefix(v.name, "ba") })
	assert.ElementsMatch(t, expected, actual)
}
