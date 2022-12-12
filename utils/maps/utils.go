package maps

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
