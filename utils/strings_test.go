package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EqualLists(t *testing.T) {
	testScenarios := []struct {
		name          string
		list1         []string
		list2         []string
		expectedEqual bool
	}{
		{
			"Equal lists with same order",
			[]string{"1", "2a", "3bc"},
			[]string{"1", "2a", "3bc"},
			true,
		},
		{
			"Equal lists with different orders",
			[]string{"1", "3bc", "2a"},
			[]string{"2a", "1", "3bc"},
			true,
		},
		{
			"Equal lists with one item",
			[]string{"1a"},
			[]string{"1a"},
			true,
		},
		{
			"Empty lists",
			[]string{},
			[]string{},
			true,
		},
		{
			"Not equal lists with one item",
			[]string{"1a"},
			[]string{"2b"},
			false,
		},
		{
			"Not equal lists",
			[]string{"1a", "2b"},
			[]string{"2b", "3cd"},
			false,
		},
	}
	t.Run("", func(t *testing.T) {
		t.Parallel()
		for _, scenario := range testScenarios {
			equals := EqualStringLists(scenario.list1, scenario.list2)
			assert.Equal(t, scenario.expectedEqual, equals, scenario.name)
		}
	})
}

func Test_array_equals(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}

	assert.True(t, ArrayEqual(a, b))
	assert.True(t, ArrayEqualElements(a, b))
}

func Test_array_same_element_diff_order(t *testing.T) {
	a := []string{"a", "c", "b"}
	b := []string{"a", "b", "c"}

	assert.False(t, ArrayEqual(a, b))
	assert.True(t, ArrayEqualElements(a, b))
}

func Test_array_same_len_diff_element(t *testing.T) {
	a := []string{"a", "c", "z"}
	b := []string{"a", "b", "z"}

	assert.False(t, ArrayEqual(a, b))
	assert.False(t, ArrayEqualElements(a, b))
}

func Test_array_diff_len_and_elem(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c", "d"}

	assert.False(t, ArrayEqual(a, b))
	assert.False(t, ArrayEqualElements(a, b))
}

func Test_array_same_len_diff_repeated_element(t *testing.T) {
	a := []string{"a", "a", "a"}
	b := []string{"a", "b", "c"}

	assert.False(t, ArrayEqual(a, b))
	assert.False(t, ArrayEqualElements(a, b))
}

func Test_Contains(t *testing.T) {
	a := []string{"a", "b", "c", "å"}

	assert.True(t, ContainsString(a, "a"))
	assert.True(t, ContainsString(a, "b"))
	assert.True(t, ContainsString(a, "c"))
	assert.False(t, ContainsString(a, "d"))
	assert.False(t, ContainsString(a, "a˚"))
}
