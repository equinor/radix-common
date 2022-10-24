package pointers

// Ptr Helper function to get the pointer of an instance
func Ptr[T any](i T) *T {
	return &i
}
