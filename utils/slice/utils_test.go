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

func Test_Reduce(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		type accumulation struct {
			lessThan5 int
			equalGt5  int
		}
		testData := []int{1, 1, 2, 5, 6, 7, 7, 9, 4, 10, 12}
		expected := accumulation{lessThan5: 4, equalGt5: 7}
		actual := Reduce(testData, accumulation{}, func(agg accumulation, d int) accumulation {
			switch {
			case d >= 5:
				agg.equalGt5 += 1
			default:
				agg.lessThan5 += 1
			}
			return agg
		})
		assert.Equal(t, expected, actual)
	})

	t.Run("test 2", func(t *testing.T) {
		testData := []string{"a", "a", "b", "c", "c", "c"}
		actual := Reduce(testData, map[string]int{}, func(agg map[string]int, v string) map[string]int {
			agg[v] += 1
			return agg
		})
		assert.Equal(t, map[string]int{"a": 2, "b": 1, "c": 3}, actual)
	})
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

func Test_FindIndex(t *testing.T) {
	type testType struct {
		id   int
		name string
	}
	testData := []testType{{id: 1, name: "foo"}, {id: 2, name: "bar"}, {id: 3, name: "baz"}, {id: 4, name: "foo2"}, {id: 5, name: "bar2"}}

	assert.Equal(t, 2, FindIndex(testData, func(v testType) bool { return v.name == "baz" }))
	assert.Equal(t, -1, FindIndex(testData, func(v testType) bool { return v.name == "oof" }))
}

func Test_FindFirst(t *testing.T) {
	type testType struct {
		id   int
		name string
	}

	e1 := testType{id: 1, name: "bar"}
	e2 := testType{id: 2, name: "foo"}
	e3 := testType{id: 3, name: "bar"}

	// Test with slice of structs
	testData := []testType{e1, e2, e3}
	found, ok := FindFirst(testData, func(v testType) bool { return v.name == "foo" })
	assert.True(t, ok)
	assert.Equal(t, e2, found)
	found, ok = FindFirst(testData, func(v testType) bool { return v.name == "bar" })
	assert.True(t, ok)
	assert.Equal(t, e1, found)
	found, ok = FindFirst(testData, func(v testType) bool { return v.name == "none" })
	assert.False(t, ok)
	assert.Equal(t, testType{}, found)

	// Test with slice of struct pointers
	testDataPtr := []*testType{&e1, &e2, &e3}
	foundPtr, ok := FindFirst(testDataPtr, func(v *testType) bool { return v.name == "foo" })
	assert.True(t, ok)
	assert.Equal(t, &e2, foundPtr)
	foundPtr, ok = FindFirst(testDataPtr, func(v *testType) bool { return v.name == "bar" })
	assert.True(t, ok)
	assert.Equal(t, &e1, foundPtr)
	foundPtr, ok = FindFirst(testDataPtr, func(v *testType) bool { return v.name == "none" })
	assert.False(t, ok)
	assert.Nil(t, foundPtr)
}
