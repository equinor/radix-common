package utils

// BoolPtr returns a pointer to the passed bool. Deprecated - use pointers.Ptr
func BoolPtr(value bool) *bool {
	return &value
}
