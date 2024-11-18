package pointers

import (
	"testing"
	"time"
)

func TestDefaultIfNil(t *testing.T) {
	tests := []struct {
		name       string
		ptr        interface{}
		defaultVal interface{}
		want       interface{}
	}{
		{
			name:       "success - ptr is nil with int",
			ptr:        (*int)(nil), // explicitly passing a nil *int pointer
			defaultVal: 42,
			want:       42,
		},
		{
			name:       "success - ptr is not nil with int",
			ptr:        new(int),
			defaultVal: 42,
			want:       0, // default int value
		},
		{
			name:       "success - ptr is not nil with int and custom value",
			ptr:        func() *int { x := 100; return &x }(),
			defaultVal: 42,
			want:       100,
		},
		{
			name:       "success - ptr is nil with string",
			ptr:        (*string)(nil), // explicitly passing a nil *string pointer
			defaultVal: "default value",
			want:       "default value",
		},
		{
			name:       "success - ptr is not nil with string",
			ptr:        new(string),
			defaultVal: "default value",
			want:       "",
		},
		{
			name:       "success - ptr is not nil with string and custom value",
			ptr:        func() *string { s := "hello"; return &s }(),
			defaultVal: "default value",
			want:       "hello",
		},
		{
			name:       "success - ptr is nil with bool",
			ptr:        (*bool)(nil), // explicitly passing a nil *bool pointer
			defaultVal: false,
			want:       false,
		},
		{
			name:       "success - ptr is not nil with bool",
			ptr:        func() *bool { b := true; return &b }(),
			defaultVal: false,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use type assertion to call DefaultIfNil correctly
			switch v := tt.ptr.(type) {
			case *int:
				got := DefaultIfNil(v, tt.defaultVal.(int))
				if got != tt.want.(int) {
					t.Errorf("DefaultIfNil() = %v, want %v", got, tt.want)
				}
			case *string:
				got := DefaultIfNil(v, tt.defaultVal.(string))
				if got != tt.want.(string) {
					t.Errorf("DefaultIfNil() = %v, want %v", got, tt.want)
				}
			case *bool:
				got := DefaultIfNil(v, tt.defaultVal.(bool))
				if got != tt.want.(bool) {
					t.Errorf("DefaultIfNil() = %v, want %v", got, tt.want)
				}
			default:
				t.Errorf("Unsupported type %T", v)
			}
		})
	}
}

func TestNullableBool(t *testing.T) {
	type args struct {
		b *bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - b is nil",
			args: args{
				b: nil,
			},
			want: false,
		},
		{
			name: "success - b is not nil and is false",
			args: args{
				b: new(bool), // new(bool) initializes to false
			},
			want: false,
		},
		{
			name: "success - b is not nil and is true",
			args: args{
				b: func() *bool { v := true; return &v }(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableBool(tt.args.b); got != tt.want {
				t.Errorf("NullableBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullableTime(t *testing.T) {
	now := time.Now()
	type args struct {
		t *time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "success - t is nil",
			args: args{
				t: nil,
			},
			want: time.Time{},
		},
		{
			name: "success - t has a value",
			args: args{
				t: func() *time.Time {
					return &now
				}(),
			},
			want: now,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableTime(tt.args.t); got != tt.want {
				t.Errorf("NullableTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
