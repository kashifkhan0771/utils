// Package conversion provides utilities for converting various units like data sizes, time, and temperature.
//
// This package includes functions for converting:
// - Data sizes (bytes, kilobytes, megabytes, gigabytes)
// - Time (seconds, minutes, hours, days)
// - Temperatures (Celsius, Fahrenheit, Kelvin)

package conversion

// Data size constants
const (
	KiloByte = 1_024
	MegaByte = 1_048_576
	GigaByte = 1_073_741_824
)

// Time constants
const (
	SecondsPerMinute = 60
	MinutesPerHour   = 60
	HoursPerDay      = 24
)

// Temperature constants
const (
	CelsiusFahrenheitOffset = 32
	CelsiusToKelvinOffset   = 273.15

	FahrenheitToKelvinOffset = 459.67

	CelsiusToFahrenheitFactor = 9.0 / 5.0
	FahrenheitToCelsiusFactor = 5.0 / 9.0

	roundingDecimalPlaces = 2
)

// Data size conversion functions

// BytesToKB converts bytes to kilobytes.
// It takes an integer value of bytes and returns the equivalent in kilobytes as a float64.
func BytesToKB(b int64) float64 {
	return float64(b) / KiloByte
}

// KBToBytes converts kilobytes to bytes.
// It takes a float64 value of kilobytes and returns the equivalent in bytes as an int64.
func KBToBytes(kb float64) int64 {
	return int64(kb * KiloByte)
}

// BytesToMB converts bytes to megabytes.
// It takes an integer value of bytes and returns the equivalent in megabytes as a float64.
func BytesToMB(b int64) float64 {
	return float64(b) / MegaByte
}

// MBToBytes converts megabytes to bytes.
// It takes a float64 value of megabytes and returns the equivalent in bytes as an int64.
func MBToBytes(mb float64) int64 {
	return int64(mb * MegaByte)
}

// BytesToGB converts bytes to gigabytes
// It takes an integer value of bytes and returns the equivalent in gigabytes as a float64.
func BytesToGB(b int64) float64 {
	return float64(b) / GigaByte
}

// GBToBytes converts gigabytes to bytes.
// It takes a float64 value of gigabytes and returns the equivalent in bytes as an int64.
func GBToBytes(gb float64) int64 {
	return int64(gb * GigaByte)
}

// Time conversion functions

// SecondsToMinutes converts seconds to minutes.
// It takes an integer value of seconds and returns the equivalent in minutes as a float64.
func SecondsToMinutes(s int64) float64 {
	return float64(s) / SecondsPerMinute
}

// MinutesToSeconds converts minutes to seconds.
// It takes an integer value of minutes and returns the equivalent in seconds as an int64.
func MinutesToSeconds(m int64) int64 {
	return m * SecondsPerMinute
}

// MinutesToHours converts minutes to hours.
// It takes an integer value of minutes and returns the equivalent in hours as a float64.
func MinutesToHours(m int64) float64 {
	return float64(m) / MinutesPerHour
}

// HoursToMinutes converts hours to minutes.
// It takes an integer value of hours and returns the equivalent in minutes as an int64.
func HoursToMinutes(h int64) int64 {
	return h * MinutesPerHour
}

// HoursToDays converts hours to days.
// It takes an integer value of hours and returns the equivalent in days as a float64.
func HoursToDays(h int64) float64 {
	return float64(h) / HoursPerDay
}

// DaysToHours converts days to hours.
// It takes an integer value of days and returns the equivalent in hours as an int64.
func DaysToHours(d int64) int64 {
	return d * HoursPerDay
}

// Temperature conversion functions

// CelsiusToFahrenheit converts Celsius to Fahrenheit.
// It takes a float64 value of Celsius and returns the equivalent in Fahrenheit as a float64.
func CelsiusToFahrenheit(c float64) float64 {
	output := (c * CelsiusToFahrenheitFactor) + CelsiusFahrenheitOffset

	return roundToDecimalPlaces(output, roundingDecimalPlaces)
}

// FahrenheitToCelsius converts Fahrenheit to Celsius.
// It takes a float64 value of Fahrenheit and returns the equivalent in Celsius as a float64.
func FahrenheitToCelsius(f float64) float64 {
	output := (f - CelsiusFahrenheitOffset) * FahrenheitToCelsiusFactor

	return roundToDecimalPlaces(output, roundingDecimalPlaces)
}

// CelsiusToKelvin converts Celsius to Kelvin.
// It takes a float64 value of Celsius and returns the equivalent in Kelvin as a float64.
func CelsiusToKelvin(c float64) float64 {
	output := c + CelsiusToKelvinOffset

	return roundToDecimalPlaces(output, roundingDecimalPlaces)
}

// KelvinToCelsius converts Kelvin to Celsius.
// It takes a float64 value of Kelvin and returns the equivalent in Celsius as a float64.
func KelvinToCelsius(k float64) float64 {
	output := k - CelsiusToKelvinOffset

	return roundToDecimalPlaces(output, roundingDecimalPlaces)
}

// FahrenheitToKelvin converts Fahrenheit to Kelvin.
// It takes a float64 value of Fahrenheit and returns the equivalent in Kelvin as a float64.
func FahrenheitToKelvin(f float64) float64 {
	output := (f + FahrenheitToKelvinOffset) * FahrenheitToCelsiusFactor

	return roundToDecimalPlaces(output, roundingDecimalPlaces)
}

// KelvinToFahrenheit converts Kelvin to Fahrenheit.
// It takes a float64 value of Kelvin and returns the equivalent in Fahrenheit as a float64.
func KelvinToFahrenheit(k float64) float64 {
	output := (k-CelsiusToKelvinOffset)*CelsiusToFahrenheitFactor + CelsiusFahrenheitOffset

	return roundToDecimalPlaces(output, roundingDecimalPlaces)
}
