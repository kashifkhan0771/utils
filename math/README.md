### Math

- **Abs**: Returns the absolute value of the given number. Works for both integers and floating-point numbers. If the input is negative, it returns its positive equivalent; otherwise, it returns the number as is.

- **Sign**: Determines the sign of a signed number. Returns `1` if the number is positive, `-1` if negative, and `0` if the number is zero.

- **Min**: Returns the smaller of two numbers. Works for all number types including integers and floating-point numbers.

- **Max**: Returns the larger of two numbers. Works for all number types including integers and floating-point numbers.

- **Clamp**: Restricts a given value to be within a specified range. If the value is below the minimum, it returns the minimum; if above the maximum, it returns the maximum.

- **IntPow**: Calculates base raised to the power of exp. Supports both positive and negative exponents. Returns float64 for fractional results.

- **IsEven**: Checks if the given integer is even. Returns `true` for even numbers and `false` otherwise.

- **IsOdd**: Checks if the given integer is odd. Returns `true` for odd numbers and `false` otherwise.

- **Swap**: Swaps the values of two variables in place. It uses pointers to modify the original variables.

- **Factorial**: Computes the factorial of a non-negative integer. Factorial of `n` is defined as the product of all integers from `1` to `n`. For `0` and `1`, the result is `1`. Factorial returns an error on invalid input.

- **GCD**: Finds the greatest common divisor (GCD) of two integers using the Euclidean algorithm. If one of the inputs is `0`, the other input is returned.

- **LCM**: Finds the least common multiple (LCM) of two integers.

- **Sqrt**: Finds square root of the given number. Works for both integers and floating-point numbers. If the input is negative, it returns the initial given number.
+ **Sqrt**: Finds the square root of the given number. Works for both integers and floating-point numbers. If the input is negative, it returns an error along with the original negative number.

## Examples:

For examples of each function, please checkout [EXAMPLES.md](/math/EXAMPLES.md)

---
