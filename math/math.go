package math

// number is a type constraint that matches all numeric types (integers and floats).
type number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

// signedNumber is a type constraint that matches all signed numeric types (integers and floats).
type signedNumber interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Abs returns the absolute value of a number.
// For negative inputs, it returns -x; otherwise, it returns x.
func Abs[T number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Sign determines the sign of a signed number.
// It returns 1 if x is positive, -1 if x is negative, and 0 if x is zero.
func Sign[T signedNumber](x T) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

// Min returns the smaller of two numbers x and y.
func Min[T number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Max returns the larger of two numbers x and y.
func Max[T number](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// Clamp restricts a value within a specified range [min, max].
// If value is less than min, it returns min; if greater than max, it returns max.
func Clamp[T number](min, max, value T) T {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

// IntPow calculates the integer power of a base raised to an exponent using exponentiation by squaring.
func IntPow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

// IsEven checks if an integer x is even.
func IsEven(x int) bool {
	return x%2 == 0
}

// IsOdd checks if an integer x is odd.
func IsOdd(x int) bool {
	return x%2 != 0
}

// Swap exchanges the values of two variables x and y.
func Swap[T any](x, y *T) {
	tmp := *x
	*x = *y
	*y = tmp
}

// Factorial calculates the factorial of a non-negative integer x.
func Factorial(x int) int {
	if x == 0 || x == 1 {
		return 1
	}

	result := 1
	for i := 2; i <= x; i++ {
		result *= i
	}
	return result
}

// GCD computes the greatest common divisor (GCD) of two integers x and y
// using the Euclidean algorithm.
func GCD(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	// The greatest common divisor must be positive
	if x >= 0 {
		return x
	}

	return -x
}

// LCM computes the least common multiple (LCM) of two integers x and y
func LCM(x, y int) int {
	return (x / GCD(x, y)) * y
}
