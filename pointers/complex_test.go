package pointers

import (
	"testing"
)

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
				c: func() *complex128 { v := complex128(0 + 0i); return &v }(),
			},
			want: 0 + 0i,
		},
		{
			name: "success - c is not nil and is 3+4i",
			args: args{
				c: func() *complex128 { v := complex128(3 + 4i); return &v }(),
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
