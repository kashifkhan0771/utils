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
				i: func() *int { v := int(42); return &v }(),
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
