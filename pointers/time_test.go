package pointers

import (
	"testing"
	"time"
)

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
