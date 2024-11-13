package pointers

// DefaultIfNil returns the value of the pointer if it is not nil,
// or the default value if the pointer is nil
func DefaultIfNil[T any](ptr *T, defaultVal T) T {
	if ptr == nil {
		return defaultVal
	}
	return *ptr
}
