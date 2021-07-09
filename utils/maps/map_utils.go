package maps

//GetKeysFromByteMap Returns keys
func GetKeysFromByteMap(mapData map[string][]byte) []string {
	var keys []string
	for k := range mapData {
		k := k
		keys = append(keys, k)
	}

	return keys
}

//GetKeysFromStringMap Returns keys
func GetKeysFromStringMap(mapData map[string]string) []string {
	var keys []string
	for k := range mapData {
		k := k
		keys = append(keys, k)
	}
	return keys
}

//MergeStringMaps Merge two maps, preferring the right over the left
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
