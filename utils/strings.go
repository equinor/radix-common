package utils

import (
	"strings"
)

// TernaryString operator
func TernaryString(condition bool, trueValue, falseValue string) string {
	return map[bool]string{true: trueValue, false: falseValue}[condition]
}

// StringPtr returns a pointer to the passed string.
func StringPtr(s string) *string {
	return &s
}

// StringUnPtr returns a string from a string pointer.
func StringUnPtr(s *string) string {
	return StringUnPtrDefault(s, "")
}

// StringUnPtrDefault returns a string from a string pointer, or default value, if nil.
func StringUnPtrDefault(s *string, defaultValue string) string {
	if s == nil {
		return defaultValue
	}
	return *s
}

func EqualStringsAsPtr(s1, s2 *string) bool {
	return ((s1 == nil) == (s2 == nil)) && (s1 != nil && strings.EqualFold(*s1, *s2))
}

// EqualStringMaps Compare two string maps
func EqualStringMaps(map1, map2 map[string]string) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, val1 := range map1 {
		val2, ok := map2[key]
		if !ok || !strings.EqualFold(val1, val2) {
			return false
		}
	}
	return true
}

// EqualStringLists Compare two string lists
func EqualStringLists(list1, list2 []string) bool {
	if len(list1) != len(list2) {
		return false
	}
	list1Map := map[string]bool{}
	for _, val := range list1 {
		list1Map[val] = false
	}
	for _, val := range list2 {
		if _, ok := list1Map[val]; !ok {
			return false
		}
	}
	return true
}

// ShortenString Get string without n-last chars
func ShortenString(s string, charsToCut int) string {
	return s[:len(s)-charsToCut]
}

// ArrayEqual tells whether a and b contain the same elements at the same index.
// A nil argument is equivalent to an empty slice.
func ArrayEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// ArrayEqualElements tells whether a and b contain the same elements. Elements does not need to be in same index
// A nil argument is equivalent to an empty slice.
func ArrayEqualElements(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	bmap := make(map[int]string)
	for i, w := range b {
		bmap[i] = w
	}

	for _, v := range a {
		containsV := false
		for i, w := range bmap {
			if w == v {
				containsV = true
				delete(bmap, i)
				break
			}
		}
		if !containsV {
			return false
		}
	}
	return true
}

// ToLowerCase Convert all strings in a list to lower case
func ToLowerCase(slice []string) []string {
	lowerSlice := make([]string, len(slice))

	for i, s := range slice {
		lowerSlice[i] = strings.ToLower(s)
	}

	return lowerSlice
}

// ContainsString return if a string is contained in the slice
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}
