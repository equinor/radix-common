package utils

import "reflect"

// IsNil check if the object is nil or an interface pointer contain nil
func IsNil(obj any) bool {
	return obj == nil || (reflect.ValueOf(obj).Kind() == reflect.Pointer && reflect.ValueOf(obj).IsNil())
}
