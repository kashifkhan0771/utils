// Package conversion provides utilities for converting various units like data sizes, time, and temperature.
//
// This package includes functions for converting:
// - Data sizes (bytes, kilobytes, megabytes, gigabytes)
// - Time (seconds, minutes, hours, days)
// - Temperatures (Celsius, Fahrenheit, Kelvin)

package conversion

// Data size conversion functions

// BytesToKB converts bytes to kilobytes.
// It takes an integer value of bytes and returns the equivalent in kilobytes as a float64.
func BytesToKB(b int64) float64 {
	return float64(b) / 1024
}

// KBToBytes converts kilobytes to bytes.
// It takes a float64 value of kilobytes and returns the equivalent in bytes as an int64.
func KBToBytes(kb float64) int64 {
	return int64(kb * 1024)
}

// BytesToMB converts bytes to megabytes.
// It takes an interger value of bytes and return a float64 value of megabytes
func BytesToMB(b int64) float64 {
	return float64(b) / 1048576
}

// MBToBytes converts megabytes to bytes.
// It takes a float64 value of megabytes and returns the equivalent in bytes as an int64.
func MBToBytes(mb float64) int64 {
	return int64(mb * 1048576)
}

// BytesToGB converts bytes to gigabytes
// It takes an integer value of bytes and return the equivalent as a float64
func BytesToGB(b int64) float64 {
	return float64(b) / 1073741824
}

// GBToBytes converts gigabytes to bytes.
// It takes a float64 value of gigabytes and returns the equivalent in bytes as an int64.
func GBToBytes(gb float64) int64 {
	return int64(gb * 1073741824)
}

// Time conversion functions

// SecondsToMinutes converts seconds to minutes.
// It takes an integer value of seconds and returns the equivalent in minutes as a float64.
func SecondsToMinutes(s int64) float64 {
	return float64(s) / 60
}

// MinutesToSeconds converts minutes to seconds.
// It takes an integer value of minutes and returns the equivalent in seconds as an int64.
func MinutesToSeconds(m int64) int64 {
	return int64(m * 60)
}

// MinutesToHours converts minutes to hours.
// It takes an integer value of minutes and returns the equivalent in hours as a float64.
func MinutesToHours(m int64) float64 {
	return float64(m) / 60
}

// HoursToMinutes converts hours to minutes.
// It takes an integer value of hours and returns the equivalent in minutes as an int64.
func HoursToMinutes(h int64) int64 {
	return int64(h * 60)
}

// HoursToDays converts hours to days.
// It takes an integer value of hours and returns the equivalent in days as a float64.
func HoursToDays(h int64) float64 {
	return float64(h) / 24
}

// DaysToHours converts days to hours.
// It takes an integer value of days and returns the equivalent in days as an int64.
func DaysToHours(d int64) int64 {
	return int64(d * 24)
}

// Temperature conversion functions

// CelsiusToFahrenheit converts Celsius to Fahrenheit.
// It takes a float64 value of Celsius and returns the equivalent in Fahrenheit as a float64.
func CelsiusToFahrenheit(c float64) float64 {
	return RoundToTwoDecimals((c * 9 / 5) + 32)
}

// FahrenheitToCelsius converts Fahrenheit to Celsius.
// It takes a float64 value of Fahrenheit and returns the equivalent in Celsius as a float64.
func FahrenheitToCelsius(f float64) float64 {
	return RoundToTwoDecimals((f - 32) * 5 / 9)
}

// CelsiusToKelvin converts Celsius to Kelvin.
// It takes a float64 value of Celsius and returns the equivalent in Kelvin as a float64.
func CelsiusToKelvin(c float64) float64 {
	return RoundToTwoDecimals(c + 273.15)
}

// KelvinToCelsius converts Kelvin to Celsius.
// It takes a float64 value of Kelvin and returns the equivalent in Celsius as a float64.
func KelvinToCelsius(k float64) float64 {
	return RoundToTwoDecimals(k - 273.15)
}

// FahrenheitToKelvin converts Fahrenheit to Kelvin.
// It takes a float64 value of Fahrenheit and returns the equivalent in Kelvin as a float64.
func FahrenheitToKelvin(f float64) float64 {
	return RoundToTwoDecimals((f + 459.67) * 5 / 9)
}

// KelvinToFahrenheit converts Kelvin to Fahrenheit.
// It takes a float64 value of Kelvin and returns the equivalent in Fahrenheit as a float64.
func KelvinToFahrenheit(k float64) float64 {
	return RoundToTwoDecimals((k-273.15)*9.0/5.0 + 32)
}
