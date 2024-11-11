package pointers

// NullableFloat32 returns the dereferenced value of *float32 if not nil,
// or 0.0 otherwise.
func NullableFloat32(f *float32) float32 {
	if f == nil {
		return 0.0
	}
	return *f
}

// NullableFloat64 returns the dereferenced value of *float64 if not nil,
// or 0.0 otherwise.
func NullableFloat64(f *float64) float64 {
	if f == nil {
		return 0.0
	}
	return *f
}
