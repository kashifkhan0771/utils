package pointers

import "testing"

func TestNullableInt(t *testing.T) {
	type args struct {
		i *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - i is nil",
			args: args{
				i: nil,
			},
			want: 0,
		},
		{
			name: "success - i has a value",
			args: args{
				i: func() *int { v := 42; return &v }(),
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableInt(tt.args.i); got != tt.want {
				t.Errorf("NullableInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt8(t *testing.T) {
	type args struct {
		i *int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "success - i is nil",
			args: args{
				i: nil,
			},
			want: 0,
		},
		{
			name: "success - i has a value",
			args: args{
				i: func() *int8 { v := int8(42); return &v }(),
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableInt8(tt.args.i); got != tt.want {
				t.Errorf("NullableInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt16(t *testing.T) {
	type args struct {
		i *int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "success - i is nil",
			args: args{
				i: nil,
			},
			want: 0,
		},
		{
			name: "success - i has a value",
			args: args{
				i: func() *int16 { v := int16(42); return &v }(),
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableInt16(tt.args.i); got != tt.want {
				t.Errorf("NullableInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt32(t *testing.T) {
	type args struct {
		i *int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "success - i is nil",
			args: args{
				i: nil,
			},
			want: 0,
		},
		{
			name: "success - i has a value",
			args: args{
				i: func() *int32 { v := int32(42); return &v }(),
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableInt32(tt.args.i); got != tt.want {
				t.Errorf("NullableInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableInt64(t *testing.T) {
	type args struct {
		i *int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "success - i is nil",
			args: args{
				i: nil,
			},
			want: 0,
		},
		{
			name: "success - i has a value",
			args: args{
				i: func() *int64 { v := int64(42); return &v }(),
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableInt64(tt.args.i); got != tt.want {
				t.Errorf("NullableInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableUint(t *testing.T) {
	type args struct {
		i *uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "success - i is nil",
			args: args{
				i: nil,
			},
			want: 0,
		},
		{
			name: "success - i has a value",
			args: args{
				i: func() *uint { v := uint(42); return &v }(),
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableUint(tt.args.i); got != tt.want {
				t.Errorf("NullableUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableUint8(t *testing.T) {
	type args struct {
		i *uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "success - i is nil",
			args: args{
				i: nil,
			},
			want: 0,
		},
		{
			name: "success - i has a value",
			args: args{
				i: func() *uint8 { v := uint8(42); return &v }(),
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableUint8(tt.args.i); got != tt.want {
				t.Errorf("NullableUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableUint16(t *testing.T) {
	type args struct {
		i *uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "success - i is nil",
			args: args{
				i: nil,
			},
			want: 0,
		},
		{
			name: "success - i has a value",
			args: args{
				i: func() *uint16 { v := uint16(42); return &v }(),
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableUint16(tt.args.i); got != tt.want {
				t.Errorf("NullableUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableUint32(t *testing.T) {
	type args struct {
		i *uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "success - i is nil",
			args: args{
				i: nil,
			},
			want: 0,
		},
		{
			name: "success - i has a value",
			args: args{
				i: func() *uint32 { v := uint32(42); return &v }(),
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableUint32(tt.args.i); got != tt.want {
				t.Errorf("NullableUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableUint64(t *testing.T) {
	type args struct {
		i *uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "success - i is nil",
			args: args{
				i: nil,
			},
			want: 0,
		},
		{
			name: "success - i has a value",
			args: args{
				i: func() *uint64 { v := uint64(42); return &v }(),
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableUint64(tt.args.i); got != tt.want {
				t.Errorf("NullableUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFloat32(t *testing.T) {
	type args struct {
		f *float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "success - f is nil",
			args: args{
				f: nil,
			},
			want: 0.0,
		},
		{
			name: "success - f has a value",
			args: args{
				f: func() *float32 { v := float32(3.14); return &v }(),
			},
			want: 3.14,
		},
		{
			name: "success - f is zero",
			args: args{
				f: func() *float32 { v := float32(0.0); return &v }(),
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableFloat32(tt.args.f); got != tt.want {
				t.Errorf("NullableFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableFloat64(t *testing.T) {
	type args struct {
		f *float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "success - f is nil",
			args: args{
				f: nil,
			},
			want: 0.0,
		},
		{
			name: "success - f has a value",
			args: args{
				f: func() *float64 { v := 6.28; return &v }(),
			},
			want: 6.28,
		},
		{
			name: "success - f is zero",
			args: args{
				f: func() *float64 { v := 0.0; return &v }(),
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableFloat64(tt.args.f); got != tt.want {
				t.Errorf("NullableFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableComplex64(t *testing.T) {
	type args struct {
		c *complex64
	}
	tests := []struct {
		name string
		args args
		want complex64
	}{
		{
			name: "success - c is nil",
			args: args{
				c: nil,
			},
			want: 0 + 0i,
		},
		{
			name: "success - c is not nil and is 0+0i",
			args: args{
				c: func() *complex64 { v := complex64(0 + 0i); return &v }(),
			},
			want: 0 + 0i,
		},
		{
			name: "success - c is not nil and is 1+2i",
			args: args{
				c: func() *complex64 { v := complex64(1 + 2i); return &v }(),
			},
			want: 1 + 2i,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableComplex64(tt.args.c); got != tt.want {
				t.Errorf("NullableComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableComplex128(t *testing.T) {
	type args struct {
		c *complex128
	}
	tests := []struct {
		name string
		args args
		want complex128
	}{
		{
			name: "success - c is nil",
			args: args{
				c: nil,
			},
			want: 0 + 0i,
		},
		{
			name: "success - c is not nil and is 0+0i",
			args: args{
				c: func() *complex128 { v := 0 + 0i; return &v }(),
			},
			want: 0 + 0i,
		},
		{
			name: "success - c is not nil and is 3+4i",
			args: args{
				c: func() *complex128 { v := 3 + 4i; return &v }(),
			},
			want: 3 + 4i,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableComplex128(tt.args.c); got != tt.want {
				t.Errorf("NullableComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}
