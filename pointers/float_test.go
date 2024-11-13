package pointers

import (
	"testing"
)

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
				f: func() *float64 { v := float64(6.28); return &v }(),
			},
			want: 6.28,
		},
		{
			name: "success - f is zero",
			args: args{
				f: func() *float64 { v := float64(0.0); return &v }(),
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
