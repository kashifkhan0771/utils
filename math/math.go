package math

import (
	"errors"
	"log"
)

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

// IntPow calculates base raised to the power of exp.
// Supports both positive and negative exponents. Returns float64 for fractional results.
func IntPow(base, exp int) float64 {
	if base == 0 && exp < 0 {
		log.Fatal("IntPow: Error, base 0 raised to a negative number simplifies to 1/0(Impossible).")
	}
	if exp == 0 {
		return 1 // Any number to the power of 0 is 1
	}

	result := 1
	isNegative := exp < 0

	// Use absolute value of exp for calculations
	if isNegative {
		exp = -exp
	}

	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}

	// If the exponent was negative, return the reciprocal
	if isNegative {
		return 1 / float64(result)
	}

	return float64(result)
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
func Factorial(x int) (int, error) {
	if x < 0 {
		return x, errors.New("factorial of a negative number is undefined")
	}

	if x == 0 || x == 1 {
		return 1, nil
	}

	result := 1
	for i := 2; i <= x; i++ {
		result *= i
	}

	return result, nil
}

// GCD computes the greatest common divisor (GCD) of two integers x and y
// using the Euclidean algorithm.
func GCD(x, y int) int {
	x = Abs(x)
	y = Abs(y)

	for y != 0 {
		x, y = y, x%y
	}

	return x
}

// LCM computes the least common multiple (LCM) of two integers x and y
func LCM(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}

	return (x / GCD(x, y)) * y
}

// Sqrt computes the square root of a number using Newton's method.
func Sqrt[T number](x T) (float64, error) {
	if x < 0 {
		return float64(x), errors.New("square root of a negative number is undefined")
	} else if x == 0 {
		return 0.0, nil
	}

	epsilon := 1e-10 // Precision threshold
	z := float64(x)  // Initial guess

	for {
		nextZ := z - (z*z-float64(x))/(2*z)
		if float64(Abs(nextZ-z)) < epsilon {
			return z, nil
		}
		z = nextZ
	}
}

// check if a number is a prime number complexity = O(sqrt(n)
func IsPrime(x int) (bool, error) {
	// don't handle negative numbers
	if x < 0 {
		return false, errors.New("can't check on negative numbers")
	}
	// both 0 and 1 are not prime numbers
	if x == 0 || x == 1 {
		return false, nil
	}
	// 2 is the only even number that is a prime number
	if x == 2 {
		return true, nil
	}
	fsqrx, err := Sqrt(x)
	if err != nil {
		return false, err
	}
	sqrx := int(fsqrx)

	// if it's an even number that is not 2 then it's not a prime number
	if x%2 == 0 {
		return false, nil
	}
	//since we eliminated all even numbers we can just iterate over only odd numbers
	// I'm iterating up to sqrx+1 just to account for any rounding errors
	for i := 3; i <= sqrx+1; i += 2 {
		if x%i == 0 {
			return false, nil
		}
	}
	return true, nil
}
