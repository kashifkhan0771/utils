package sort

// number is a type constraint that matches all numeric types (integers and floats).
type number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}
