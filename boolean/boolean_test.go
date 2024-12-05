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

func TestToggle(t *testing.T) {
	type args struct {
		value bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - toggle true to false",
			args: args{value: true},
			want: false,
		},
		{
			name: "success - toggle false to true",
			args: args{value: false},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Toggle(tt.args.value); got != tt.want {
				t.Errorf("Toggle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllTrue(t *testing.T) {
	type args struct {
		values []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - all values true",
			args: args{values: []bool{true, true, true}},
			want: true,
		},
		{
			name: "success - one value false",
			args: args{values: []bool{true, false, true}},
			want: false,
		},
		{
			name: "success - empty slice",
			args: args{values: []bool{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllTrue(tt.args.values); got != tt.want {
				t.Errorf("AllTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyTrue(t *testing.T) {
	type args struct {
		values []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - at least one value true",
			args: args{values: []bool{false, true, false}},
			want: true,
		},
		{
			name: "success - no true values",
			args: args{values: []bool{false, false, false}},
			want: false,
		},
		{
			name: "success - empty slice",
			args: args{values: []bool{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyTrue(tt.args.values); got != tt.want {
				t.Errorf("AnyTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoneTrue(t *testing.T) {
	type args struct {
		values []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - no true values",
			args: args{values: []bool{false, false, false}},
			want: true,
		},
		{
			name: "success - at least one true value",
			args: args{values: []bool{false, true, false}},
			want: false,
		},
		{
			name: "success - empty slice",
			args: args{values: []bool{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoneTrue(tt.args.values); got != tt.want {
				t.Errorf("NoneTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountTrue(t *testing.T) {
	type args struct {
		values []bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - count true values",
			args: args{values: []bool{true, false, true}},
			want: 2,
		},
		{
			name: "success - no true values",
			args: args{values: []bool{false, false, false}},
			want: 0,
		},
		{
			name: "success - empty slice",
			args: args{values: []bool{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountTrue(tt.args.values); got != tt.want {
				t.Errorf("CountTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountFalse(t *testing.T) {
	type args struct {
		values []bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - count false values",
			args: args{values: []bool{true, false, true}},
			want: 1,
		},
		{
			name: "success - no false values",
			args: args{values: []bool{true, true, true}},
			want: 0,
		},
		{
			name: "success - empty slice",
			args: args{values: []bool{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountFalse(tt.args.values); got != tt.want {
				t.Errorf("CountFalse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	type args struct {
		values []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - all values equal (true)",
			args: args{values: []bool{true, true, true}},
			want: true,
		},
		{
			name: "success - all values equal (false)",
			args: args{values: []bool{false, false, false}},
			want: true,
		},
		{
			name: "failure - mixed values",
			args: args{values: []bool{true, false, true}},
			want: false,
		},
		{
			name: "failure - empty slice",
			args: args{values: []bool{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.values...); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	type args struct {
		values []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - all values true",
			args: args{values: []bool{true, true, true}},
			want: true,
		},
		{
			name: "failure - one value false",
			args: args{values: []bool{true, false, true}},
			want: false,
		},
		{
			name: "failure - empty slice",
			args: args{values: []bool{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := And(tt.args.values); got != tt.want {
				t.Errorf("And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOr(t *testing.T) {
	type args struct {
		values []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - at least one value true",
			args: args{values: []bool{false, true, false}},
			want: true,
		},
		{
			name: "failure - all values false",
			args: args{values: []bool{false, false, false}},
			want: false,
		},
		{
			name: "failure - empty slice",
			args: args{values: []bool{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Or(tt.args.values); got != tt.want {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}
