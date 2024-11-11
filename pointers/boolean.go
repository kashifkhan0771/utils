package pointers

// NullableBool returns the value of the bool pointer
// or false if the pointer is nil
func NullableBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
