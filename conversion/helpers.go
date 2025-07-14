package conversion

import "math"

func RoundToTwoDecimals(val float64) float64 {
	return math.Round(val*100) / 100
}
