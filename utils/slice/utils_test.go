package slice_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/equinor/radix-common/utils/slice"
	"github.com/stretchr/testify/assert"
)

func Test_PointersOf(t *testing.T) {
	type obj struct{ prop string }

	t.Run("nil returns empty slice", func(t *testing.T) {
		result := slice.PointersOf[[]obj](nil)
		assert.Empty(t, result)
	})
	t.Run("empty slice", func(t *testing.T) {
		result := slice.PointersOf([]obj{})
		assert.Empty(t, result)
	})
	t.Run("single element", func(t *testing.T) {
		src := []obj{{"a"}}
		result := slice.PointersOf(src)
		assert.Len(t, result, 1)
		assert.Equal(t, src[0], *result[0])
	})
	t.Run("multiple elements length and values", func(t *testing.T) {
		src := []obj{{"a"}, {"b"}, {"c"}}
		result := slice.PointersOf(src)
		assert.Len(t, result, len(src))
		for i, p := range result {
			assert.Equal(t, src[i], *p)
		}
	})
	t.Run("pointers are distinct", func(t *testing.T) {
		src := []obj{{"x"}, {"x"}}
		result := slice.PointersOf(src)
		assert.NotSame(t, result[0], result[1])
	})
	t.Run("modifying source does not affect returned pointers", func(t *testing.T) {
		src := []obj{{"original"}}
		result := slice.PointersOf(src)
		src[0].prop = "mutated"
		assert.Equal(t, "original", result[0].prop)
	})
	t.Run("modifying returned pointer does not affect other pointers", func(t *testing.T) {
		src := []obj{{"a"}, {"b"}}
		result := slice.PointersOf(src)
		result[0].prop = "changed"
		assert.Equal(t, "b", result[1].prop)
	})
	t.Run("integer slice", func(t *testing.T) {
		src := []int{1, 2, 3}
		result := slice.PointersOf(src)
		assert.Len(t, result, 3)
		for i, p := range result {
			assert.Equal(t, src[i], *p)
		}
	})
	t.Run("string slice", func(t *testing.T) {
		src := []string{"foo", "bar"}
		result := slice.PointersOf(src)
		assert.Len(t, result, 2)
		assert.Equal(t, "foo", *result[0])
		assert.Equal(t, "bar", *result[1])
	})
	t.Run("order is preserved", func(t *testing.T) {
		src := []int{10, 20, 30, 40}
		result := slice.PointersOf(src)
		for i, p := range result {
			assert.Equal(t, src[i], *p)
		}
	})
}

func Test_Map(t *testing.T) {
	expected := []string{"10", "20", "30"}
	actual := slice.Map([]int{1, 2, 3}, func(v int) string { return fmt.Sprintf("%d", v*10) })
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
		actual := slice.Reduce(testData, accumulation{}, func(agg accumulation, d int) accumulation {
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
		actual := slice.Reduce(testData, map[string]int{}, func(agg map[string]int, v string) map[string]int {
			agg[v] += 1
			return agg
		})
		assert.Equal(t, map[string]int{"a": 2, "b": 1, "c": 3}, actual)
	})
}

func Test_Any(t *testing.T) {
	type testType struct{ val int }
	testData := []testType{{val: 1}, {val: 2}, {val: 3}}
	assert.True(t, slice.Any(testData, func(o testType) bool { return o.val == 2 }))
	assert.False(t, slice.Any(testData, func(o testType) bool { return o.val == 4 }))
	assert.False(t, slice.Any([]testType{}, func(o testType) bool { return o.val == 4 }))
}

func Test_All(t *testing.T) {
	type testType struct{ val int }
	testData := []testType{{val: 1}, {val: 2}, {val: 3}}
	assert.True(t, slice.All(testData, func(o testType) bool { return o.val > 0 }))
	assert.False(t, slice.All(testData, func(o testType) bool { return o.val > 1 }))
	assert.True(t, slice.All([]testType{}, func(o testType) bool { return o.val > 0 }))
}

func Test_FindAll(t *testing.T) {
	type testType struct {
		id   int
		name string
	}
	testData := []testType{{id: 1, name: "foo"}, {id: 2, name: "bar"}, {id: 3, name: "baz"}, {id: 4, name: "foo2"}, {id: 5, name: "bar2"}}
	expected := []testType{{id: 2, name: "bar"}, {id: 3, name: "baz"}, {id: 5, name: "bar2"}}
	actual := slice.FindAll(testData, func(v testType) bool { return strings.HasPrefix(v.name, "ba") })
	assert.ElementsMatch(t, expected, actual)
}

func Test_FindIndex(t *testing.T) {
	type testType struct {
		id   int
		name string
	}
	testData := []testType{{id: 1, name: "foo"}, {id: 2, name: "bar"}, {id: 3, name: "baz"}, {id: 4, name: "foo2"}, {id: 5, name: "bar2"}}

	assert.Equal(t, 2, slice.FindIndex(testData, func(v testType) bool { return v.name == "baz" }))
	assert.Equal(t, -1, slice.FindIndex(testData, func(v testType) bool { return v.name == "oof" }))
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
	found, ok := slice.FindFirst(testData, func(v testType) bool { return v.name == "foo" })
	assert.True(t, ok)
	assert.Equal(t, e2, found)
	found, ok = slice.FindFirst(testData, func(v testType) bool { return v.name == "bar" })
	assert.True(t, ok)
	assert.Equal(t, e1, found)
	found, ok = slice.FindFirst(testData, func(v testType) bool { return v.name == "none" })
	assert.False(t, ok)
	assert.Equal(t, testType{}, found)

	// Test with slice of struct pointers
	testDataPtr := []*testType{&e1, &e2, &e3}
	foundPtr, ok := slice.FindFirst(testDataPtr, func(v *testType) bool { return v.name == "foo" })
	assert.True(t, ok)
	assert.Equal(t, &e2, foundPtr)
	foundPtr, ok = slice.FindFirst(testDataPtr, func(v *testType) bool { return v.name == "bar" })
	assert.True(t, ok)
	assert.Equal(t, &e1, foundPtr)
	foundPtr, ok = slice.FindFirst(testDataPtr, func(v *testType) bool { return v.name == "none" })
	assert.False(t, ok)
	assert.Nil(t, foundPtr)
}

func Test_ElementsMatch(t *testing.T) {
	t.Run("both nil", func(t *testing.T) {
		assert.True(t, slice.ElementsMatch[[]string](nil, nil))
	})
	t.Run("nil and empty slice", func(t *testing.T) {
		assert.True(t, slice.ElementsMatch([]string{}, nil))
		assert.True(t, slice.ElementsMatch(nil, []string{}))
	})
	t.Run("both empty", func(t *testing.T) {
		assert.True(t, slice.ElementsMatch([]string{}, []string{}))
	})
	t.Run("same elements same order", func(t *testing.T) {
		assert.True(t, slice.ElementsMatch([]string{"a", "b", "c"}, []string{"a", "b", "c"}))
	})
	t.Run("same elements different order", func(t *testing.T) {
		assert.True(t, slice.ElementsMatch([]string{"c", "a", "b"}, []string{"a", "b", "c"}))
	})
	t.Run("different lengths", func(t *testing.T) {
		assert.False(t, slice.ElementsMatch([]string{"a", "b"}, []string{"a", "b", "c"}))
		assert.False(t, slice.ElementsMatch([]string{"a", "b", "c"}, []string{"a", "b"}))
	})
	t.Run("same length different elements", func(t *testing.T) {
		assert.False(t, slice.ElementsMatch([]string{"a", "b", "c"}, []string{"a", "b", "d"}))
	})
	t.Run("duplicates match", func(t *testing.T) {
		assert.True(t, slice.ElementsMatch([]string{"a", "a", "b"}, []string{"a", "b", "a"}))
	})
	t.Run("duplicate count mismatch", func(t *testing.T) {
		assert.False(t, slice.ElementsMatch([]string{"a", "a", "b"}, []string{"a", "b", "b"}))
	})
	t.Run("all duplicates equal", func(t *testing.T) {
		assert.True(t, slice.ElementsMatch([]string{"x", "x", "x"}, []string{"x", "x", "x"}))
	})
	t.Run("all duplicates different count", func(t *testing.T) {
		assert.False(t, slice.ElementsMatch([]string{"x", "x", "x"}, []string{"x", "x"}))
	})
	t.Run("integer slices equal", func(t *testing.T) {
		assert.True(t, slice.ElementsMatch([]int{3, 1, 2}, []int{1, 2, 3}))
	})
	t.Run("integer slices not equal", func(t *testing.T) {
		assert.False(t, slice.ElementsMatch([]int{1, 2, 3}, []int{1, 2, 4}))
	})
	t.Run("one empty one not", func(t *testing.T) {
		assert.False(t, slice.ElementsMatch([]string{"a"}, []string{}))
		assert.False(t, slice.ElementsMatch([]string{}, []string{"a"}))
	})
	t.Run("superset is not equal", func(t *testing.T) {
		assert.False(t, slice.ElementsMatch([]string{"a", "b"}, []string{"a", "b", "a"}))
	})
}
