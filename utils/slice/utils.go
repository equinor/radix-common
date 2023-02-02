package slice

import (
	"reflect"
)

// PointersOf Returnes a pointer of
func PointersOf(v interface{}) interface{} {
	in := reflect.ValueOf(v)
	out := reflect.MakeSlice(reflect.SliceOf(reflect.PtrTo(in.Type().Elem())), in.Len(), in.Len())
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

// Applies an reducer function over a slice.
// The specified seed value is used as the initial accumulator value.
func Reduce[TSource, TAccumulate any](source []TSource, seed TAccumulate, accumulator func(TAccumulate, TSource) TAccumulate) TAccumulate {
	for _, v := range source {
		seed = accumulator(seed, v)
	}
	return seed
}

// Determines whether any element of a slice satisfies a condition.
func Any[T any](source []T, predicate func(T) bool) bool {
	for _, v := range source {
		if predicate(v) {
			return true
		}
	}
	return false
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
