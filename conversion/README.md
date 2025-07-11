# Conversion

## Overview
This package provides utilities for converting various units like data sizes, time, and temperature.

## Functions

### Data Size Conversion
- `BytesToKB(b int64) float64`: Converts bytes to kilobytes.
- `KBToBytes(kb float64) int64`: Converts kilobytes to bytes.
- `BytesToMB(b int64) float64`: Converts bytes to megabytes.
- `MBToBytes(mb float64) int64`: Converts megabytes to bytes.
- `BytesToGB(b int64) float64`: Converts bytes to gigabytes.
- `GBToBytes(gb float64) int64`: Converts gigabytes to bytes.

### Time Conversion
- `SecondsToMinutes(s int64) float64`: Converts seconds to minutes.
- `MinutesToSeconds(m int64) int64`: Converts minutes to seconds.
- `MinutesToHours(m int64) float64`: Converts minutes to hours.
- `HoursToMinutes(h int64) int64`: Converts hours to minutes.
- `HoursToDays(h int64) float64`: Converts hours to days.
- `DaysToHours(d int64) int64`: Converts days to hours.

### Temperature Conversion
- `CelsiusToFahrenheit(c float64) float64`: Converts Celsius to Fahrenheit.
- `FahrenheitToCelsius(f float64) float64`: Converts Fahrenheit to Celsius.
- `CelsiusToKelvin(c float64) float64`: Converts Celsius to Kelvin.
- `KelvinToCelsius(k float64) float64`: Converts Kelvin to Celsius.
- `FahrenheitToKelvin(f float64) float64`: Converts Fahrenheit to Kelvin.
- `KelvinToFahrenheit(k float64) float64`: Converts Kelvin to Fahrenheit.

## Examples:
---
For examples of each function, please checkout [EXAMPLES.md](/conversion/EXAMPLES.md)
