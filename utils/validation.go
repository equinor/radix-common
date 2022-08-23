package utils

import "reflect"

// IsNil check if the object is nil or an interface pointer contain nil
func IsNil(obj interface{}) bool {
	return obj == nil || (reflect.ValueOf(obj).Kind() == reflect.Ptr && reflect.ValueOf(obj).IsNil())
}
