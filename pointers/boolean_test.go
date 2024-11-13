package pointers

import "testing"

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
