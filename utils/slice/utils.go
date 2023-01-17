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
