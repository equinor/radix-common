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

func EqualStringsAsPtr(s1, s2 *string) bool {
	return ((s1 == nil) == (s2 == nil)) && (s1 != nil && strings.EqualFold(*s1, *s2))
}

//EqualStringMaps Compare two string maps
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

//EqualStringLists Compare two string lists
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

//ShortenString Get string without n-last chars
func ShortenString(s string, charsToCut int) string {
	return s[:len(s)-charsToCut]
}
