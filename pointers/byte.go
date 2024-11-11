package pointers

// NullableByteSlice returns the dereferenced value of *[]byte if not nil,
// or an empty byte slice otherwise.
func NullableByteSlice(b *[]byte) []byte {
	if b == nil {
		return []byte{}
	}
	return *b
}
