package pointers

import "time"

// DefaultIfNil returns the value of the pointer if it is not nil,
// or the default value if the pointer is nil
func DefaultIfNil[T any](ptr *T, defaultVal T) T {
	if ptr == nil {
		return defaultVal
	}

	return *ptr
}

// NullableBool returns the value of the bool pointer
// or false if the pointer is nil
func NullableBool(b *bool) bool {
	if b == nil {
		return false
	}

	return *b
}

// NullableTime returns the dereferenced value of *time.Time if not nil,
// or a zero time.Time otherwise.
func NullableTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}

	return *t
}
