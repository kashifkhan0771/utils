package boolean

import "testing"

type testCase[V any, W comparable] struct {
	name   string
	values []V
	value  V
	want   W
}

func TestIsTrue(t *testing.T) {
	t.Parallel()

	tests := []testCase[string, bool]{
		{
			name:  "success - check 1 as true",
			value: "1",
			want:  true,
		},
		{
			name:  "success - check t as true",
			value: "t",
			want:  true,
		},
		{
			name:  "success - check TRUE as true",
			value: "TRUE",
			want:  true,
		},
		{
			name:  "success - check true as true",
			value: "true",
			want:  true,
		},
		{
			name:  "success - check T as true",
			value: "T",
			want:  true,
		},
		{
			name:  "success - check 0 as false",
			value: "0",
			want:  false,
		},
		{
			name:  "success - check any random word as false",
			value: "g",
			want:  false,
		},
		{
			name:  "success - check f as false",
			value: "f",
			want:  false,
		},
		{
			name:  "success - check F as false",
			value: "F",
			want:  false,
		},
		{
			name:  "success - check false as false",
			value: "false",
			want:  false,
		},
		{
			name:  "success - check FALSE as false",
			value: "FALSE",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsTrue(tt.value); got != tt.want {
				t.Errorf("IsTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToggle(t *testing.T) {
	t.Parallel()

	tests := []testCase[bool, bool]{
		{
			name:  "success - toggle true to false",
			value: true,
			want:  false,
		},
		{
			name:  "success - toggle false to true",
			value: false,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Toggle(tt.value); got != tt.want {
				t.Errorf("Toggle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllTrue(t *testing.T) {
	t.Parallel()

	tests := []testCase[bool, bool]{
		{
			name:   "success - all values true",
			values: []bool{true, true, true},
			want:   true,
		},
		{
			name:   "success - one value false",
			values: []bool{true, false, true},
			want:   false,
		},
		{
			name:   "success - empty slice",
			values: []bool{},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := AllTrue(tt.values); got != tt.want {
				t.Errorf("AllTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyTrue(t *testing.T) {
	t.Parallel()

	tests := []testCase[bool, bool]{
		{
			name:   "success - at least one value true",
			values: []bool{false, true, false},
			want:   true,
		},
		{
			name:   "success - no true values",
			values: []bool{false, false, false},
			want:   false,
		},
		{
			name:   "success - empty slice",
			values: []bool{},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := AnyTrue(tt.values); got != tt.want {
				t.Errorf("AnyTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoneTrue(t *testing.T) {
	t.Parallel()

	tests := []testCase[bool, bool]{
		{
			name:   "success - no true values",
			values: []bool{false, false, false},
			want:   true,
		},
		{
			name:   "success - at least one true value",
			values: []bool{false, true, false},
			want:   false,
		},
		{
			name:   "success - empty slice",
			values: []bool{},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NoneTrue(tt.values); got != tt.want {
				t.Errorf("NoneTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountTrue(t *testing.T) {
	t.Parallel()

	tests := []testCase[bool, int]{
		{
			name:   "success - count true values",
			values: []bool{true, false, true},
			want:   2,
		},
		{
			name:   "success - no true values",
			values: []bool{false, false, false},
			want:   0,
		},
		{
			name:   "success - empty slice",
			values: []bool{},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := CountTrue(tt.values); got != tt.want {
				t.Errorf("CountTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountFalse(t *testing.T) {
	t.Parallel()

	tests := []testCase[bool, int]{
		{
			name:   "success - count false values",
			values: []bool{true, false, true},
			want:   1,
		},
		{
			name:   "success - no false values",
			values: []bool{true, true, true},
			want:   0,
		},
		{
			name:   "success - empty slice",
			values: []bool{},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := CountFalse(tt.values); got != tt.want {
				t.Errorf("CountFalse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	t.Parallel()

	tests := []testCase[bool, bool]{
		{
			name:   "success - all values equal (true)",
			values: []bool{true, true, true},
			want:   true,
		},
		{
			name:   "success - all values equal (false)",
			values: []bool{false, false, false},
			want:   true,
		},
		{
			name:   "failure - mixed values",
			values: []bool{true, false, true},
			want:   false,
		},
		{
			name:   "failure - empty slice",
			values: []bool{},
			want:   false,
		},
		{
			name:   "success - single value",
			values: []bool{true},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Equal(tt.values...); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	t.Parallel()

	tests := []testCase[bool, bool]{
		{
			name:   "success - all values true",
			values: []bool{true, true, true},
			want:   true,
		},
		{
			name:   "failure - one value false",
			values: []bool{true, false, true},
			want:   false,
		},
		{
			name:   "failure - empty slice",
			values: []bool{},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := And(tt.values); got != tt.want {
				t.Errorf("And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOr(t *testing.T) {
	t.Parallel()

	tests := []testCase[bool, bool]{
		{
			name:   "success - at least one value true",
			values: []bool{false, true, false},
			want:   true,
		},
		{
			name:   "failure - all values false",
			values: []bool{false, false, false},
			want:   false,
		},
		{
			name:   "failure - empty slice",
			values: []bool{},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Or(tt.values); got != tt.want {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}
