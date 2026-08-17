package slice

import (
	"reflect"
	"slices"
)

// PointersOf returns a slice of pointers to each element in the provided slice.
func PointersOf(v any) any {
	in := reflect.ValueOf(v)
	out := reflect.MakeSlice(reflect.SliceOf(reflect.PointerTo(in.Type().Elem())), in.Len(), in.Len())
	for i := 0; i < in.Len(); i++ {
		out.Index(i).Set(in.Index(i).Addr())
	}
	return out.Interface()
}

// Projects each element of a slice into a new form.
func Map[T, V any](source []T, mapper func(T) V) []V {
	result := make([]V, len(source))

	for i, v := range source {
		result[i] = mapper(v)
	}

	return result
}

// Applies an accumulator function over a slice.
// The specified seed value is used as the initial accumulator value.
func Reduce[TSource, TAccumulation any](source []TSource, seed TAccumulation, accumulator func(TAccumulation, TSource) TAccumulation) TAccumulation {
	for _, v := range source {
		seed = accumulator(seed, v)
	}
	return seed
}

// Determines whether any element of a slice satisfies a condition.
func Any[T any](source []T, predicate func(T) bool) bool {
	return slices.ContainsFunc(source, predicate)
}

// Determines whether all elements of a slice satisfy a condition.
func All[T any](source []T, predicate func(T) bool) bool {
	return !Any(source, func(v T) bool { return !predicate(v) })
}

// Retrieves all the elements that match the conditions defined by the specified predicate.
func FindAll[T any](source []T, predicate func(T) bool) []T {
	var foundElements []T
	for _, v := range source {
		if predicate(v) {
			foundElements = append(foundElements, v)
		}
	}
	return foundElements
}

// Returns the index of the first element matching the predicate or -1 on fail
func FindIndex[T any](source []T, predicate func(T) bool) int {
	for i, v := range source {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// Returns the first element matching the predicate.
// 'ok' is true if an element was found and false if not.
func FindFirst[T any](source []T, predicate func(T) bool) (element T, ok bool) {
	for _, v := range source {
		if predicate(v) {
			element = v
			ok = true
			return
		}
	}
	return
}

// ElementsMatch reports whether two slices are equal: the same length and all
// elements exist in both slices, ignoring the order of the elements.
// If there are duplicate elements, the number of appearances of each of them in both slices should match.
func ElementsMatch[S ~[]E, E comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make(map[E]int, len(s1))
	for _, v := range s1 {
		counts[v]++
	}

	for _, v := range s2 {
		if counts[v] == 0 {
			return false
		}
		counts[v]--
	}

	return true
}
