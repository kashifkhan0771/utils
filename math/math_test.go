package math

import (
	"math"
	"slices"
	"testing"
)

func TestAbsForInts(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - positive input",
			args: args{x: 5},
			want: 5,
		},
		{
			name: "success - negative input",
			args: args{x: -5},
			want: 5,
		},
		{
			name: "success - zero input",
			args: args{x: 0},
			want: 0,
		},
		{
			name: "success - negative zero input",
			args: args{x: -0},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsForFloats(t *testing.T) {
	type args struct {
		x float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "success - positive input",
			args: args{x: 5.0},
			want: 5.0,
		},
		{
			name: "success - negative input",
			args: args{x: -5.3},
			want: 5.3,
		},
		{
			name: "success - zero input",
			args: args{x: 0.0},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsForFloats_NaN(t *testing.T) {
	got := Abs(float32(math.NaN()))
	if !math.IsNaN(float64(got)) {
		t.Errorf("Abs(NaN) = %v, want NaN", got)
	}
}

func TestSignForInt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - positive number",
			args: args{x: 10},
			want: 1,
		},
		{
			name: "success - negative number",
			args: args{x: -10},
			want: -1,
		},
		{
			name: "success - zero",
			args: args{x: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sign(tt.args.x); got != tt.want {
				t.Errorf("Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignForFloat(t *testing.T) {
	type args struct {
		x float32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - positive number",
			args: args{x: 10.3},
			want: 1,
		},
		{
			name: "success - negative number",
			args: args{x: -10.5},
			want: -1,
		},
		{
			name: "success - zero",
			args: args{x: 0.0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sign(tt.args.x); got != tt.want {
				t.Errorf("Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinForInt(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - x < y",
			args: args{x: 5, y: 10},
			want: 5,
		},
		{
			name: "success - x > y",
			args: args{x: 15, y: 10},
			want: 10,
		},
		{
			name: "success - x == y",
			args: args{x: 7, y: 7},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinForFloat(t *testing.T) {
	type args struct {
		x float32
		y float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "success - x < y",
			args: args{x: 5.3, y: 10.0},
			want: 5.3,
		},
		{
			name: "success - x > y",
			args: args{x: 15.3, y: 10.2},
			want: 10.2,
		},
		{
			name: "success - x == y",
			args: args{x: 7.3, y: 7.3},
			want: 7.3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxForInt(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - x > y",
			args: args{x: 15, y: 10},
			want: 15,
		},
		{
			name: "success - x < y",
			args: args{x: 5, y: 10},
			want: 10,
		},
		{
			name: "success - x == y",
			args: args{x: 7, y: 7},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxForFloat(t *testing.T) {
	type args struct {
		x float32
		y float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "success - x > y",
			args: args{x: 15.2, y: 10.2},
			want: 15.2,
		},
		{
			name: "success - x < y",
			args: args{x: 5.3, y: 10.2},
			want: 10.2,
		},
		{
			name: "success - x == y",
			args: args{x: 7.2, y: 7.2},
			want: 7.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClampForInt(t *testing.T) {
	type args struct {
		min   int
		max   int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - value within range",
			args: args{min: 1, max: 10, value: 5},
			want: 5,
		},
		{
			name: "success - value below range",
			args: args{min: 1, max: 10, value: 0},
			want: 1,
		},
		{
			name: "success - value above range",
			args: args{min: 1, max: 10, value: 15},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.args.min, tt.args.max, tt.args.value); got != tt.want {
				t.Errorf("Clamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClampForFloat(t *testing.T) {
	type args struct {
		min   float32
		max   float32
		value float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "success - value within range",
			args: args{min: 1.3, max: 10.2, value: 5.1},
			want: 5.1,
		},
		{
			name: "success - value below range",
			args: args{min: 1.4, max: 10.4, value: 0.0},
			want: 1.4,
		},
		{
			name: "success - value above range",
			args: args{min: 1.2, max: 10.6, value: 15.2},
			want: 10.6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.args.min, tt.args.max, tt.args.value); got != tt.want {
				t.Errorf("Clamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntPow(t *testing.T) {
	type args struct {
		base int
		exp  int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "success - base 2, exp 3",
			args: args{base: 2, exp: 3},
			want: 8,
		},
		{
			name: "success - base 5, exp 0",
			args: args{base: 5, exp: 0},
			want: 1,
		},
		{
			name: "success - base 3, exp 2",
			args: args{base: 3, exp: 2},
			want: 9,
		},
		{
			name: "success - base 2, exp -3",
			args: args{base: 2, exp: -3},
			want: 0.125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntPow(tt.args.base, tt.args.exp); got != tt.want {
				t.Errorf("IntPow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - even number",
			args: args{x: 8},
			want: true,
		},
		{
			name: "success - odd number",
			args: args{x: 7},
			want: false,
		},
		{
			name: "success - zero",
			args: args{x: 0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.x); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOdd(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - odd number",
			args: args{x: 7},
			want: true,
		},
		{
			name: "success - even number",
			args: args{x: 8},
			want: false,
		},
		{
			name: "success - zero",
			args: args{x: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOdd(tt.args.x); got != tt.want {
				t.Errorf("IsOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want args // Expected swapped values
	}{
		{
			name: "success - swap positive numbers",
			args: args{x: 10, y: 20},
			want: args{x: 20, y: 10},
		},
		{
			name: "success - swap negative numbers",
			args: args{x: -10, y: -20},
			want: args{x: -20, y: -10},
		},
		{
			name: "success - swap zero and number",
			args: args{x: 0, y: 5},
			want: args{x: 5, y: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := tt.args.x
			y := tt.args.y
			Swap(&x, &y)
			if x != tt.want.x || y != tt.want.y {
				t.Errorf("Swap() = x: %v, y: %v, want x: %v, y: %v", x, y, tt.want.x, tt.want.y)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - factorial of 0",
			args: args{x: 0},
			want: 1,
		},
		{
			name: "success - factorial of 1",
			args: args{x: 1},
			want: 1,
		},
		{
			name: "success - factorial of 5",
			args: args{x: 5},
			want: 120,
		},
		{
			name: "success - factorial of 7",
			args: args{x: 7},
			want: 5040,
		},
		{
			name: "failure - factorial of -3",
			args: args{x: -3},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Factorial(tt.args.x); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - GCD of two positive numbers",
			args: args{x: 12, y: 18},
			want: 6,
		},
		{
			name: "success - GCD where one number is zero",
			args: args{x: 0, y: 18},
			want: 18,
		},
		{
			name: "success - GCD of two prime numbers",
			args: args{x: 17, y: 19},
			want: 1,
		},
		{
			name: "success - GCD of negative numbers",
			args: args{x: -12, y: -18},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLCM(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - LCM of two positive numbers",
			args: args{x: 4, y: 6},
			want: 12,
		},
		{
			name: "success - LCM where one number is zero",
			args: args{x: 0, y: 5},
			want: 0,
		},
		{
			name: "success - LCM of two prime numbers",
			args: args{x: 7, y: 13},
			want: 91,
		},
		{
			name: "success - LCM of one number being multiple of the other",
			args: args{x: 5, y: 10},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCM(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("LCM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqrtForInts(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "success - square root of 4",
			args: args{x: 4},
			want: 2.0,
		},
		{
			name: "success - square root of 0",
			args: args{x: 0},
			want: 0.0,
		},
		{
			name: "success - square root of 2",
			args: args{x: 2},
			want: 1.414213562373095048801688724209698078569671875376,
		},
		{
			name: "failure - square root of -1",
			args: args{x: -1},
			want: -1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Sqrt(tt.args.x)
			if Abs(got-tt.want) > 1e-6 {
				t.Errorf("Sqrt(%v) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}

func TestSqrtForFloats(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "success - square root of pi",
			args: args{x: math.Pi},
			want: math.SqrtPi,
		},
		{
			name: "success - square root of square root of 2",
			args: args{x: 1.414213562373095048801688724209698078569671875376},
			want: 1.189207115002721066717499970560475915292972092463,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Sqrt(tt.args.x)
			if Abs(got-tt.want) > 1e-6 {
				t.Errorf("Sqrt(%v) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}
func TestIsPrime(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - negative 10",
			args: args{x: -10},
			want: false,
		},
		{
			name: "success - 0",
			args: args{x: 0},
			want: false,
		},
		{
			name: "success - 1",
			args: args{x: 1},
			want: false,
		},
		{
			name: "success - 2",
			args: args{x: 2},
			want: true,
		},
		{
			name: "success - 3",
			args: args{x: 3},
			want: true,
		},
		{
			name: "success - 4",
			args: args{x: 4},
			want: false,
		},
		{
			name: "success - 36473",
			args: args{x: 36473},
			want: true,
		},
		{
			name: "success - 2147483647 (2^31-1)",
			args: args{x: 2147483647},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.x); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrimeList(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success - 2",
			args: args{x: 2},
			want: []int{2},
		},
		{
			name: "success - 0",
			args: args{x: 0},
			want: []int{},
		},
		{
			name: "success - 1",
			args: args{x: 1},
			want: []int{},
		},
		{
			name: "success - 3",
			args: args{x: 3},
			want: []int{2, 3},
		},
		{
			name: "success - 4",
			args: args{x: 4},
			want: []int{2, 3},
		},
		{
			name: "success - 10",
			args: args{x: 10},
			want: []int{2, 3, 5, 7},
		},
		{
			name: "success - 100 ",
			args: args{x: 100},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimeList(tt.args.x); !slices.Equal(got, tt.want) {
				t.Errorf("PrimeList() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetDivisors(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success - 0",
			args: args{x: 0},
			want: []int{},
		},
		{
			name: "success - 1",
			args: args{x: 1},
			want: []int{1},
		},
		{
			name: "success - 2",
			args: args{x: 2},
			want: []int{1, 2},
		},
		{
			name: "success - 3",
			args: args{x: 3},
			want: []int{1, 3},
		},
		{
			name: "success - 4",
			args: args{x: 4},
			want: []int{1, 2, 4},
		},
		{
			name: "success - 10",
			args: args{x: 10},
			want: []int{1, 2, 5, 10},
		},
		{
			name: "success - 49",
			args: args{x: 49},
			want: []int{1, 7, 49},
		},
		{
			name: "success - 48",
			args: args{x: 48},
			want: []int{1, 2, 3, 4, 6, 8, 12, 16, 24, 48},
		},
		{
			name: "success - 100 ",
			args: args{x: 100},
			want: []int{1, 2, 4, 5, 10, 20, 25, 50, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetDivisors(tt.args.x)
			// The slice need to be sorted in order to be compared
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("GetDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundDecimalPlaces(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		places   int
		expected float64
	}{
		{
			name:     "success - round to 2 decimal places",
			value:    3.14159,
			places:   2,
			expected: 3.14,
		},
		{
			name:     "success - round to 3 decimal places",
			value:    2.718281828459045,
			places:   3,
			expected: 2.718,
		},
		{
			name:     "success - round to 0 decimal places",
			value:    1.9999,
			places:   0,
			expected: 2.0,
		},
		{
			name:     "success - round negative number to 1 decimal place",
			value:    -1.2345,
			places:   1,
			expected: -1.2,
		},
		{
			name:     "success - round to 4 decimal places",
			value:    0.123456789,
			places:   4,
			expected: 0.1235,
		},
		{
			name:     "success - round to 5 decimal places",
			value:    0.123456789,
			places:   5,
			expected: 0.12346,
		},
		{
			name:     "success - round to 0 decimal places with negative number",
			value:    3.14159,
			places:   -1,
			expected: 3.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := RoundDecimalPlaces(tt.value, tt.places)
			if result != tt.expected {
				t.Errorf("RoundDecimalPlaces(%v, %d) = %v; want %v", tt.value, tt.places, result, tt.expected)
			}
		})
	}
}

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func BenchmarkIntPow(b *testing.B) {
	b.ReportAllocs()

	const base, exp = 3, 2
	for b.Loop() {
		_ = IntPow(base, exp)
	}
}

func BenchmarkFactorial(b *testing.B) {
	b.ReportAllocs()

	const x = 7
	for b.Loop() {
		_, _ = Factorial(x)
	}
}

func BenchmarkGCD(b *testing.B) {
	b.ReportAllocs()

	const x, y = 17, 19
	for b.Loop() {
		_ = GCD(x, y)
	}
}

func BenchmarkLCM(b *testing.B) {
	b.ReportAllocs()

	const x, y = 4, 6
	for b.Loop() {
		_ = LCM(x, y)
	}
}

func BenchmarkSqrt(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		_, _ = Sqrt(math.Pi)
	}
}

func BenchmarkIsPrime(b *testing.B) {
	b.ReportAllocs()

	const prime = 2147483647
	for b.Loop() {
		_ = IsPrime(prime)
	}
}

func BenchmarkPrimeList(b *testing.B) {
	b.ReportAllocs()

	const n = 1e8
	for b.Loop() {
		_ = PrimeList(n)
	}
}

func BenchmarkGetDivisors(b *testing.B) {
	b.ReportAllocs()

	const n = 2e9
	for b.Loop() {
		_ = GetDivisors(n)
	}
}
