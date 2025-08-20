package conversion

import "math"

// RoundToTwoDecimals is deprecated; rounds a float64 value to two decimal places.
// Deprecated: Use `math.RoundDecimalPlaces(val float64, places int) float64` in the math-utils package instead.
func RoundToTwoDecimals(val float64) float64 {
	return roundToDecimalPlaces(val, 2)
}

func roundToDecimalPlaces(val float64, places int) float64 {
	p := math.Pow(10, float64(places))

	return math.Round(val*p) / p
}
