package math

import (
	"errors"
	"math"
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
	return math.Pow(float64(base), float64(exp))
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

// Sqrt computes the square root of a number using the standard library's math.Sqrt.
// For negative inputs, it returns the original value and an error.
func Sqrt[T number](x T) (float64, error) {
	if x < 0 {
		return float64(x), errors.New("square root of a negative number is undefined")
	}

	return math.Sqrt(float64(x)), nil
}

// IsPrime checks if a number is prime. Complexity = O(sqrt(n)).
func IsPrime(x int) bool {
	// 1 and 0 are not primes and handle all negative numbers as non-prime.
	if x < 2 || (x != 2 && IsEven(x)) {
		return false
	}

	// 2 is the only even number that is a prime number
	if x == 2 {
		return true
	}

	// Since even numbers are eliminated, iterate over odd divisors only.
	// Use i*i <= x to avoid float conversions and rounding concerns.
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			return false
		}
	}

	return true
}

// Sieve of Eratosthenes algorithm Complexity = O(n log log n).
func PrimeList(n int) []int {
	if n < 2 {
		return []int{}
	}
	list := []int{}
	// Slice of bool to flag each number, intially we assume every number is a prime number except 1 and 0
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	// Basic implementation of Sieve
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i + i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			list = append(list, i)
		}
	}

	return list
}

// Complexity = sqrt(n)
func GetDivisors(n int) []int {
	if n <= 0 {
		return []int{}
	}
	// 2 is a good starting point for the slice capacity
	list := make([]int, 0, 2)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			list = append(list, i)
			// Check for perfect squares
			if i != n/i {
				list = append(list, n/i)
			}
		}
	}
	// The list is not sorted.
	return list
}

// RoundDecimalPlaces rounds a float64 to the specified number of decimal places.
// Negative values for places are clamped to 0 (i.e., rounds to a whole number).
func RoundDecimalPlaces(val float64, places int) float64 {
	p := math.Pow10(max(places, 0))

	return math.Round(val*p) / p
}
