package pointers

// NullableString returns the value of the string pointer
// or an empty string if the pointer is nil
func NullableString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
