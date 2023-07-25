package mergo

import (
	"testing"

	"dario.cat/mergo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_BoolPtrTransformer(t *testing.T) {

	type testBool struct {
		V *bool
	}
	boolPtr := func(v bool) *bool { return &v }

	scenarios := []struct {
		name     string
		dst      testBool
		src      testBool
		expected testBool
	}{
		{name: "use value from src #1", dst: testBool{V: boolPtr(true)}, src: testBool{V: boolPtr(false)}, expected: testBool{V: boolPtr(false)}},
		{name: "use value from src #2", dst: testBool{V: boolPtr(false)}, src: testBool{V: boolPtr(true)}, expected: testBool{V: boolPtr(true)}},
		{name: "use value from src #3", dst: testBool{}, src: testBool{V: boolPtr(false)}, expected: testBool{V: boolPtr(false)}},
		{name: "use value from src #4", dst: testBool{}, src: testBool{V: boolPtr(true)}, expected: testBool{V: boolPtr(true)}},
		{name: "use value from dst #1", dst: testBool{V: boolPtr(true)}, src: testBool{}, expected: testBool{V: boolPtr(true)}},
		{name: "use value from dst #2", dst: testBool{V: boolPtr(false)}, src: testBool{}, expected: testBool{V: boolPtr(false)}},
	}

	for _, scenario := range scenarios {
		err := mergo.Merge(&scenario.dst, &scenario.src, mergo.WithOverride, mergo.WithTransformers(BoolPtrTransformer{}))
		require.NoError(t, err, scenario.name)
		assert.Equal(t, scenario.expected, scenario.dst, scenario.name)
	}

}
