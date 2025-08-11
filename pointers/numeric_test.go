package pointers

import "testing"

// ptr is a helper function that returns a pointer to the value of type T.
func ptr[T any](v T) *T { return &v }

type testCase[T any] struct {
	name  string
	input *T
	want  T
}

func TestNullableInt(t *testing.T) {
	tests := []testCase[int]{
		{
			name:  "success - i is nil",
			input: nil,
			want:  0,
		},
		{
			name:  "success - i has a value",
			input: ptr(42),
			want:  42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableInt(tt.input); got != tt.want {
				t.Errorf("NullableInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt8(t *testing.T) {
	tests := []testCase[int8]{
		{
			name:  "success - i is nil",
			input: nil,
			want:  0,
		},
		{
			name:  "success - i has a value",
			input: ptr(int8(42)),
			want:  42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableInt8(tt.input); got != tt.want {
				t.Errorf("NullableInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt16(t *testing.T) {
	tests := []testCase[int16]{
		{
			name:  "success - i is nil",
			input: nil,
			want:  0,
		},
		{
			name:  "success - i has a value",
			input: ptr(int16(42)),
			want:  42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableInt16(tt.input); got != tt.want {
				t.Errorf("NullableInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt32(t *testing.T) {
	tests := []testCase[int32]{
		{
			name:  "success - i is nil",
			input: nil,
			want:  0,
		},
		{
			name:  "success - i has a value",
			input: ptr(int32(42)),
			want:  42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableInt32(tt.input); got != tt.want {
				t.Errorf("NullableInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt64(t *testing.T) {
	tests := []testCase[int64]{
		{
			name:  "success - i is nil",
			input: nil,
			want:  0,
		},
		{
			name:  "success - i has a value",
			input: ptr(int64(42)),
			want:  42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableInt64(tt.input); got != tt.want {
				t.Errorf("NullableInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableUint(t *testing.T) {
	tests := []testCase[uint]{
		{
			name:  "success - i is nil",
			input: nil,
			want:  0,
		},
		{
			name:  "success - i has a value",
			input: ptr(uint(42)),
			want:  42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableUint(tt.input); got != tt.want {
				t.Errorf("NullableUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableUint8(t *testing.T) {
	tests := []testCase[uint8]{
		{
			name:  "success - i is nil",
			input: nil,
			want:  0,
		},
		{
			name:  "success - i has a value",
			input: ptr(uint8(42)),
			want:  42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableUint8(tt.input); got != tt.want {
				t.Errorf("NullableUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableUint16(t *testing.T) {
	tests := []testCase[uint16]{
		{
			name:  "success - i is nil",
			input: nil,
			want:  0,
		},
		{
			name:  "success - i has a value",
			input: ptr(uint16(42)),
			want:  42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableUint16(tt.input); got != tt.want {
				t.Errorf("NullableUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableUint32(t *testing.T) {
	tests := []testCase[uint32]{
		{
			name:  "success - i is nil",
			input: nil,
			want:  0,
		},
		{
			name:  "success - i has a value",
			input: ptr(uint32(42)),
			want:  42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableUint32(tt.input); got != tt.want {
				t.Errorf("NullableUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableUint64(t *testing.T) {
	tests := []testCase[uint64]{
		{
			name:  "success - i is nil",
			input: nil,
			want:  0,
		},
		{
			name:  "success - i has a value",
			input: ptr(uint64(42)),
			want:  42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableUint64(tt.input); got != tt.want {
				t.Errorf("NullableUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFloat32(t *testing.T) {
	tests := []testCase[float32]{
		{
			name:  "success - f is nil",
			input: nil,
			want:  0.0,
		},
		{
			name:  "success - f has a value",
			input: ptr(float32(3.14)),
			want:  3.14,
		},
		{
			name:  "success - f is zero",
			input: ptr(float32(0.0)),
			want:  0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableFloat32(tt.input); got != tt.want {
				t.Errorf("NullableFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFloat64(t *testing.T) {
	tests := []testCase[float64]{
		{
			name:  "success - f is nil",
			input: nil,
			want:  0.0,
		},
		{
			name:  "success - f has a value",
			input: ptr(6.28),
			want:  6.28,
		},
		{
			name:  "success - f is zero",
			input: ptr(0.0),
			want:  0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableFloat64(tt.input); got != tt.want {
				t.Errorf("NullableFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableComplex64(t *testing.T) {
	tests := []testCase[complex64]{
		{
			name:  "success - c is nil",
			input: nil,
			want:  0 + 0i,
		},
		{
			name:  "success - c is not nil and is 0+0i",
			input: ptr(complex64(0 + 0i)),
			want:  0 + 0i,
		},
		{
			name:  "success - c is not nil and is 1+2i",
			input: ptr(complex64(1 + 2i)),
			want:  1 + 2i,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableComplex64(tt.input); got != tt.want {
				t.Errorf("NullableComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableComplex128(t *testing.T) {
	tests := []testCase[complex128]{
		{
			name:  "success - c is nil",
			input: nil,
			want:  0 + 0i,
		},
		{
			name:  "success - c is not nil and is 0+0i",
			input: ptr(0 + 0i),
			want:  0 + 0i,
		},
		{
			name:  "success - c is not nil and is 3+4i",
			input: ptr(3 + 4i),
			want:  3 + 4i,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NullableComplex128(tt.input); got != tt.want {
				t.Errorf("NullableComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}
