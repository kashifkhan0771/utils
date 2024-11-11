package pointers

import "time"

// NullableTime returns the dereferenced value of *time.Time if not nil,
// or a zero time.Time otherwise.
func NullableTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}
