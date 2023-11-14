package maps

import (
	"fmt"
	"testing"

	"github.com/equinor/radix-common/utils"
	"github.com/stretchr/testify/assert"
)

func Test_GetKeysFromMap(t *testing.T) {
	a := make(map[string][]byte)
	a["a"] = []byte("x")
	a["b"] = []byte("y")
	a["c"] = []byte("z")

	assert.True(t, utils.ArrayEqualElements([]string{"a", "b", "c"}, GetKeysFromByteMap(a)))
}

func Test_GetKeysFromMapForStrings(t *testing.T) {
	a := make(map[string]string)
	a["a"] = "x"
	a["b"] = "y"
	a["c"] = "z"

	assert.True(t, utils.ArrayEqualElements([]string{"a", "b", "c"}, GetKeysFromMap(a)))
}

func Test_GetKeysFromMapForBool(t *testing.T) {
	a := make(map[string]bool)
	a["a"] = true
	a["b"] = false
	a["c"] = false

	assert.True(t, utils.ArrayEqualElements([]string{"a", "b", "c"}, GetKeysFromMap(a)))
}

func Test_GetKeysFromMapForStruct(t *testing.T) {
	a := make(map[int]struct{})
	a[7] = struct{}{}
	a[3] = struct{}{}
	a[5] = struct{}{}

	assert.ElementsMatch(t, []int{7, 3, 5}, GetKeysFromMap(a))
}

func TestMergeStringMaps(t *testing.T) {
	empty := make(map[string]string)
	expect := map[string]string{
		"a": "a", "x": "x", "b": "y",
	}

	map1 := map[string]string{"a": "a", "b": "c"}
	map2 := map[string]string{"x": "x", "b": "y"}

	result := MergeStringMaps(map1, map2)
	assert.Equal(t, expect, result)
	result = MergeStringMaps(map2, map1)
	assert.NotEqual(t, expect, result)

	result = MergeStringMaps(nil, map1)
	assert.Equal(t, map1, result)
	result = MergeStringMaps(map2, nil)
	assert.Equal(t, map2, result)

	result = MergeStringMaps(nil, nil)
	assert.Equal(t, empty, result)
}

func TestMergeMaps(t *testing.T) {

	scenarios := []struct {
		sources  []map[string]string
		expected map[string]string
	}{
		{
			sources:  []map[string]string{{"a": "a", "b": "c"}, {"x": "x", "X": "X", "b": "y"}, {"z": "z", "X": "Y"}},
			expected: map[string]string{"a": "a", "x": "x", "b": "y", "X": "Y", "z": "z"},
		},
		{
			sources:  []map[string]string{{"a": "a", "b": "c"}},
			expected: map[string]string{"a": "a", "b": "c"},
		},
		{
			sources:  []map[string]string{{"a": "a", "b": "c"}, nil},
			expected: map[string]string{"a": "a", "b": "c"},
		},
		{
			sources:  []map[string]string{nil},
			expected: map[string]string{},
		},
		{
			sources:  nil,
			expected: map[string]string{},
		},
	}

	for _, scenario := range scenarios {
		actual := MergeMaps(scenario.sources...)
		assert.Equal(t, scenario.expected, actual)
	}
}

func Test_ConvertToMap(t *testing.T) {
	scenarios := []struct {
		source   string
		expected map[string]string
	}{
		{source: "a=b,c=d", expected: map[string]string{"a": "b", "c": "d"}},
		{source: "a=b ,c=d, e=f ,  p=k", expected: map[string]string{"a": "b", "c": "d", "e": "f", "p": "k"}},
		{source: "a=b,c=,=f,=,a", expected: map[string]string{"a": "b"}},
		{source: "a=b,,==,=f=,-", expected: map[string]string{"a": "b"}},
		{source: "", expected: map[string]string{}},
	}
	for i, scenario := range scenarios {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			actual := FromString(scenario.source)
			assert.Equal(t, scenario.expected, actual)
		})
	}
}

func Test_ConvertToString(t *testing.T) {
	scenarios := []struct {
		source   map[string]string
		expected string
	}{
		{source: map[string]string{"a": "b", "c": "d"}, expected: "a=b,c=d"},
		{source: map[string]string{"a": "b  "}, expected: "a=b"},
		{source: map[string]string{"a ": "b  "}, expected: "a=b"},
		{source: map[string]string{" a": " b  "}, expected: "a=b"},
		{source: map[string]string{"": "b  ", "c": "d"}, expected: "c=d"},
		{source: map[string]string{"": "b  ", "  ": "d"}, expected: ""},
		{source: map[string]string{}, expected: ""},
	}
	for i, scenario := range scenarios {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			actual := ToString(scenario.source)
			assert.Equal(t, scenario.expected, actual)
		})
	}
}
