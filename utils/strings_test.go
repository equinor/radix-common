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
