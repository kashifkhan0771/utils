package pointers

// NullableString returns the value of the string pointer
// or an empty string if the pointer is nil
func NullableString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// DefaultIfNil returns the value of the pointer if it is not nil,
// or the default value if the pointer is nil
func DefaultIfNil[T any](ptr *T, defaultVal T) T {
	if ptr == nil {
		return defaultVal
	}
	return *ptr
}
