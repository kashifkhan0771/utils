package pointers

// NullableComplex64 returns the dereferenced value of *complex64 if not nil,
// or 0+0i otherwise.
func NullableComplex64(c *complex64) complex64 {
	if c == nil {
		return 0 + 0i
	}
	return *c
}

// NullableComplex128 returns the dereferenced value of *complex128 if not nil,
// or 0+0i otherwise.
func NullableComplex128(c *complex128) complex128 {
	if c == nil {
		return 0 + 0i
	}
	return *c
}
