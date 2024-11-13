package pointers

import "testing"

func TestNullableString(t *testing.T) {
	type args struct {
		s *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - s is nil",
			args: args{
				s: nil,
			},
			want: "",
		},
		{
			name: "success - s has a value",
			args: args{
				s: func() *string { v := "Hello, World!"; return &v }(),
			},
			want: "Hello, World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullableString(tt.args.s); got != tt.want {
				t.Errorf("NullableString() = %v, want %v", got, tt.want)
			}
		})
	}
}
