package boolean

import "testing"

func TestIsTrue(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - check 1 as true",
			args: args{v: "1"},
			want: true,
		},
		{
			name: "success - check t as true",
			args: args{v: "t"},
			want: true,
		},
		{
			name: "success - check TRUE as true",
			args: args{v: "TRUE"},
			want: true,
		},
		{
			name: "success - check true as true",
			args: args{v: "true"},
			want: true,
		},
		{
			name: "success - check T as true",
			args: args{v: "T"},
			want: true,
		},
		{
			name: "success - check 0 as false",
			args: args{v: "0"},
			want: false,
		},
		{
			name: "success - check any random word as false",
			args: args{v: "g"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTrue(tt.args.v); got != tt.want {
				t.Errorf("IsTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}
