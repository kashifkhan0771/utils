package pointers

// NullableInt returns the dereferenced value of the int pointer
// or 0 if the pointer is nil
func NullableInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// NullableInt8 returns the dereferenced value of *int8 if not nil,
// or 0 otherwise.
func NullableInt8(i *int8) int8 {
	if i == nil {
		return 0
	}
	return *i
}

// NullableInt16 returns the dereferenced value of *int16 if not nil,
// or 0 otherwise.
func NullableInt16(i *int16) int16 {
	if i == nil {
		return 0
	}
	return *i
}

// NullableInt32 returns the dereferenced value of *int32 if not nil,
// or 0 otherwise.
func NullableInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

// NullableInt64 returns the dereferenced value of *int64 if not nil,
// or 0 otherwise.
func NullableInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// NullableUint returns the dereferenced value of *uint if not nil,
// or 0 otherwise.
func NullableUint(i *uint) uint {
	if i == nil {
		return 0
	}
	return *i
}

// NullableUint8 returns the dereferenced value of *uint8 if not nil,
// or 0 otherwise.
func NullableUint8(i *uint8) uint8 {
	if i == nil {
		return 0
	}
	return *i
}

// NullableUint16 returns the dereferenced value of *uint16 if not nil,
// or 0 otherwise.
func NullableUint16(i *uint16) uint16 {
	if i == nil {
		return 0
	}
	return *i
}

// NullableUint32 returns the dereferenced value of *uint32 if not nil,
// or 0 otherwise.
func NullableUint32(i *uint32) uint32 {
	if i == nil {
		return 0
	}
	return *i
}

// NullableUint64 returns the dereferenced value of *uint64 if not nil,
// or 0 otherwise.
func NullableUint64(i *uint64) uint64 {
	if i == nil {
		return 0
	}
	return *i
}
