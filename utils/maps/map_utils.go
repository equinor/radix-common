package maps

//GetKeysFromByteMap Returns keys. Deprecated - use GetKeysFromMap
func GetKeysFromByteMap(mapData map[string][]byte) []string {
	return GetKeysFromMap(mapData)
}

//GetKeysFromStringMap Returns keys. Deprecated - use GetKeysFromMap
func GetKeysFromStringMap(mapData map[string]string) []string {
	return GetKeysFromMap(mapData)
}

//GetKeysFromMap Returns keys
func GetKeysFromMap[T any](mapData map[string]T) []string {
	var keys []string
	for k := range mapData {
		k := k
		keys = append(keys, k)
	}
	return keys
}

//MergeStringMaps Merge two maps, preferring the right over the left. Deprecated - use MergeMaps
func MergeStringMaps(left, right map[string]string) map[string]string {
	result := make(map[string]string)

	for key, rVal := range right {
		rVal := rVal
		result[key] = rVal
	}

	for key, lVal := range left {
		lVal := lVal
		if _, present := right[key]; !present {
			result[key] = lVal
		}
	}

	return result
}

//MergeMaps Merge two maps, preferring the right over the left
func MergeMaps[T comparable, V any](left, right map[T]V) map[T]V {
	result := make(map[T]V)

	for key, rVal := range right {
		rVal := rVal
		result[key] = rVal
	}

	for key, lVal := range left {
		lVal := lVal
		if _, present := right[key]; !present {
			result[key] = lVal
		}
	}

	return result
}
