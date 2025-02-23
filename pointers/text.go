package pointers

// NullableString returns the value of the string pointer
// or an empty string if the pointer is nil
func NullableString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

// NullableByteSlice returns the dereferenced value of *[]byte if not nil,
// or an empty byte slice otherwise.
func NullableByteSlice(b *[]byte) []byte {
	if b == nil {
		return []byte{}
	}

	return *b
}
