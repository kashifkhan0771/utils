/*
Package boolean defines boolean helpers.
*/
package boolean

import (
	"strconv"
)

// IsTrue checks if the provided string is a true value
// It accepts 1, t, T, TRUE, true, True. All other values are considered false.
func IsTrue(v string) bool {
	b, _ := strconv.ParseBool(v)

	return b
}

/*
Toggle negates the given boolean value.

Returns:
  - true if the input is false.
  - false if the input is true.
*/
func Toggle(value bool) bool {
	return !value
}

/*
AllTrue checks if all values in the slice are true.

Returns:
  - false if the slice is empty or contains at least one false value.
  - true if all values are true.
*/
func AllTrue(values []bool) bool {
	if len(values) == 0 {
		return false // Empty slices are considered not "all true".
	}

	for _, b := range values {
		if !b {
			return false
		}
	}

	return true
}

/*
AnyTrue checks if at least one value in the slice is true.

Returns:
  - false if the slice is empty or contains no true values.
  - true if at least one value is true.
*/
func AnyTrue(values []bool) bool {
	if len(values) == 0 {
		return false // Empty slices are considered not "any true".
	}

	for _, b := range values {
		if b {
			return true
		}
	}

	return false
}

/*
NoneTrue checks if none of the values in the slice are true.

Returns:
  - true if the slice is empty or contains no true values.
  - false if at least one value is true.
*/
func NoneTrue(values []bool) bool {
	if len(values) == 0 {
		return true // Empty slices are considered "none true".
	}

	for _, b := range values {
		if b {
			return false
		}
	}

	return true
}

/*
CountTrue counts the number of true values in the slice.

Returns:
  - The number of elements in the slice that are true.
*/
func CountTrue(values []bool) int {
	count := 0
	for _, b := range values {
		if b == true {
			count++ // Increment for each true value.
		}
	}

	return count
}

/*
CountFalse counts the number of false values in the slice.

Returns:
  - The number of elements in the slice that are false.
*/
func CountFalse(values []bool) int {
	count := 0
	for _, b := range values {
		if b == false {
			count++ // Increment for each false value.
		}
	}

	return count
}

/*
Equal checks if all values in the variadic argument list are equal.

Returns:
  - true if all values are the same (either all true or all false).
  - false if there is any inconsistency or the argument list is empty.
*/
func Equal(values ...bool) bool {
	if len(values) == 0 {
		return false // Empty input is considered not equal.
	}

	first := values[0]
	for _, v := range values[1:] {
		if v != first {
			return false
		}
	}

	return true
}

/*
And performs a logical AND across all values in the slice.

Returns:
  - false if the slice is empty or any value is false.
  - true if all values are true.
*/
func And(values []bool) bool {
	for _, b := range values {
		if !b {
			return false // Short-circuit if any value is false.
		}
	}

	return len(values) > 0 // Ensure the slice is not empty.
}

/*
Or performs a logical OR across all values in the slice.

Returns:
  - true if at least one value is true.
  - false if the slice is empty or all values are false.
*/
func Or(values []bool) bool {
	for _, b := range values {
		if b {
			return true // Short-circuit if any value is true.
		}
	}

	return false // No true values found, or the slice is empty.
}
