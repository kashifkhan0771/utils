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
