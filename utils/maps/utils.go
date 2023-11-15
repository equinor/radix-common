package maps

import (
	"fmt"
	"strings"

	"github.com/equinor/radix-common/utils/slice"
)

// GetKeysFromByteMap Returns keys. Deprecated - use GetKeysFromMap
func GetKeysFromByteMap(mapData map[string][]byte) []string {
	return GetKeysFromMap(mapData)
}

// GetKeysFromStringMap Returns keys. Deprecated - use GetKeysFromMap
func GetKeysFromStringMap(mapData map[string]string) []string {
	return GetKeysFromMap(mapData)
}

// GetKeysFromMap Returns keys
func GetKeysFromMap[T comparable, V any](mapData map[T]V) []T {
	var keys []T
	for k := range mapData {
		k := k
		keys = append(keys, k)
	}
	return keys
}

// MergeStringMaps Merge two maps, preferring the right over the left.
//
// Deprecated: use MergeMaps
func MergeStringMaps(left, right map[string]string) map[string]string {
	return MergeMaps(left, right)
}

// MergeMaps combines given maps into one map.
// Maps are merged from the beginning (index position 1) of the sources slice,
// and in case of a key conflict, values from the
// source map with the highest index position wins.
func MergeMaps[T comparable, V any](sources ...map[T]V) map[T]V {
	result := make(map[T]V)

	for _, source := range sources {
		for key, rVal := range source {
			rVal := rVal
			result[key] = rVal
		}
	}

	return result
}

// FromString Converts a string with comma separated key-vaults to a map[string]string
func FromString(keyValuePairs string) map[string]string {
	return slice.Reduce(strings.Split(keyValuePairs, ","), make(map[string]string), func(acc map[string]string, pair string) map[string]string {
		keyValue := strings.Split(pair, "=")
		if len(keyValue) != 2 {
			return acc
		}
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])
		if len(key) > 0 && len(value) > 0 {
			acc[key] = value
		}
		return acc
	})
}

// ToString Converts map[string]string to a string with comma separated key-vaults, spaces are trimmed
func ToString(keyValues map[string]string) string {
	return strings.Join(slice.Reduce(GetKeysFromMap(keyValues), make([]string, 0), func(acc []string, key string) []string {
		trimmedKey := strings.TrimSpace(key)
		if len(trimmedKey) == 0 {
			return acc
		}
		return append(acc, fmt.Sprintf("%s=%s", trimmedKey, strings.TrimSpace(keyValues[key])))
	}), ",")
}
