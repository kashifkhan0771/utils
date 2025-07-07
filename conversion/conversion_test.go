package conversion_test

import (
	"testing"

	"github.com/kashifkhan0771/utils/conversion"
)

func TestBytesToK(t *testing.T) {
	tests := []struct {
		name string
		arg  int64
		want float64
	}{
		{
			name: "calculate bytes to kilobytes",
			arg:  1024,
			want: 1.0,
		},
		{
			name: "zero bytes should return zero kilobytes",
			arg:  0,
			want: 0.0,
		},
		{
			name: "partial kilobyte (512 bytes)",
			arg:  512,
			want: 0.5,
		},
		{
			name: "large number of bytes",
			arg:  10 * 1024 * 1024,
			want: 10240.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.BytesToKB(tt.arg)
			if got != tt.want {
				t.Errorf("BytesToK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKBToBytes(t *testing.T) {
	tests := []struct {
		name string
		arg  float64
		want int64
	}{
		{
			name: "1 KB to bytes",
			arg:  1.0,
			want: 1024,
		},
		{
			name: "0 KB to bytes",
			arg:  0.0,
			want: 0,
		},
		{
			name: "0.5 KB to bytes",
			arg:  0.5,
			want: 512,
		},
		{
			name: "10.25 KB to bytes",
			arg:  10.25,
			want: 10496, // 10.25 * 1024 = 10496
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.KBToBytes(tt.arg)
			if got != tt.want {
				t.Errorf("KBToBytes(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestBytesToMB(t *testing.T) {
	tests := []struct {
		name string
		arg  int64
		want float64
	}{
		{
			name: "1048576 bytes to megabytes",
			arg:  1048576,
			want: 1.0,
		},
		{
			name: "0 bytes to megabytes",
			arg:  0,
			want: 0.00,
		},
		{
			name: "2621440 bytes to megabytes",
			arg:  2621440,
			want: 2.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.BytesToMB(tt.arg)
			if got != tt.want {
				t.Errorf("BytesToMB(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestMBToBytes(t *testing.T) {
	tests := []struct {
		name string
		arg  float64
		want int64
	}{
		{
			name: "1 MB to bytes",
			arg:  1.0,
			want: 1048576,
		},
		{
			name: "0 MB to bytes",
			arg:  0.0,
			want: 0,
		},
		{
			name: "2.5 MB to bytes",
			arg:  2.5,
			want: 2621440,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.MBToBytes(tt.arg)
			if got != tt.want {
				t.Errorf("MBToBytes(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestBytesToGB(t *testing.T) {

	tests := []struct {
		name string
		arg  int64
		want float64
	}{
		{
			name: "1073741824 bytes to gigabytes",
			arg:  1073741824,
			want: 1.0,
		},
		{
			name: "0 bytes to gigabytes",
			arg:  0,
			want: 0.0,
		},
		{
			name: "107374182400 bytes to gigabytes",
			arg:  107374182400,
			want: 100.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.BytesToGB(tt.arg)
			if got != tt.want {
				t.Errorf("BytesToGB(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestGBToBytes(t *testing.T) {
	tests := []struct {
		name string
		arg  float64
		want int64
	}{
		{
			name: "1 GB to bytes",
			arg:  1.0,
			want: 1073741824,
		},
		{
			name: "0 GB to bytes",
			arg:  0.0,
			want: 0,
		},
		{
			name: "2.5 GB to bytes",
			arg:  2.5,
			want: 2684354560,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.GBToBytes(tt.arg)
			if got != tt.want {
				t.Errorf("GBToBytes(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

// Time conversion functions
func TestSecondsToMinutes(t *testing.T) {
	tests := []struct {
		name string
		arg  int64
		want float64
	}{
		{
			name: "60 seconds to minutes",
			arg:  60,
			want: 1.0,
		},
		{
			name: "0 seconds to minutes",
			arg:  0,
			want: 0.0,
		},
		{
			name: "30 seconds to minutes",
			arg:  30,
			want: 0.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.SecondsToMinutes(tt.arg)
			if got != tt.want {
				t.Errorf("SecondsToMinutes(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestMinutesToSeconds(t *testing.T) {
	tests := []struct {
		name string
		arg  int64
		want int64
	}{
		{
			name: "1 minute to seconds",
			arg:  1,
			want: 60,
		},
		{
			name: "0 minutes to seconds",
			arg:  0,
			want: 0,
		},
		{
			name: "0.5 minutes to seconds",
			arg:  0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.MinutesToSeconds(tt.arg)
			if got != tt.want {
				t.Errorf("MinutesToSeconds(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestMinutesToHours(t *testing.T) {
	tests := []struct {
		name string
		arg  int64
		want float64
	}{
		{
			name: "60 minutes to hours",
			arg:  60,
			want: 1.0,
		},
		{
			name: "0 minutes to hours",
			arg:  0,
			want: 0.0,
		},
		{
			name: "30 minutes to hours",
			arg:  30,
			want: 0.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.MinutesToHours(tt.arg)
			if got != tt.want {
				t.Errorf("MinutesToHours(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestHoursToMinutes(t *testing.T) {
	tests := []struct {
		name string
		arg  int64
		want int64
	}{
		{
			name: "1 hour to minutes",
			arg:  1,
			want: 60,
		},
		{
			name: "0 hours to minutes",
			arg:  0,
			want: 0,
		},
		{
			name: "2 hours to minutes",
			arg:  2,
			want: 120,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.HoursToMinutes(tt.arg)
			if got != tt.want {
				t.Errorf("HoursToMinutes(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestHoursToDays(t *testing.T) {
	tests := []struct {
		name string
		arg  int64
		want float64
	}{
		{
			name: "24 hours to days",
			arg:  24,
			want: 1.0,
		},
		{
			name: "0 hours to days",
			arg:  0,
			want: 0.0,
		},
		{
			name: "48 hours to days",
			arg:  48,
			want: 2.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.HoursToDays(tt.arg)
			if got != tt.want {
				t.Errorf("HoursToDays(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestDaysToHours(t *testing.T) {
	tests := []struct {
		name string
		arg  int64
		want int64
	}{
		{
			name: "1 day to hours",
			arg:  1,
			want: 24,
		},
		{
			name: "0 days to hours",
			arg:  0,
			want: 0,
		},
		{
			name: "2 days to hours",
			arg:  2,
			want: 48,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.DaysToHours(tt.arg)
			if got != tt.want {
				t.Errorf("DaysToHours(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

// Temperature conversion functions
func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name string
		arg  float64
		want float64
	}{
		{
			name: "0 Celsius to Fahrenheit",
			arg:  0,
			want: 32,
		},
		{
			name: "100 Celsius to Fahrenheit",
			arg:  100,
			want: 212,
		},
		{
			name: "-40 Celsius to Fahrenheit",
			arg:  -40,
			want: -40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.CelsiusToFahrenheit(tt.arg)
			if got != tt.want {
				t.Errorf("CelsiusToFahrenheit(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		name string
		arg  float64
		want float64
	}{
		{
			name: "32 Fahrenheit to Celsius",
			arg:  32,
			want: 0,
		},
		{
			name: "212 Fahrenheit to Celsius",
			arg:  212,
			want: 100,
		},
		{
			name: "-40 Fahrenheit to Celsius",
			arg:  -40,
			want: -40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.FahrenheitToCelsius(tt.arg)
			if got != tt.want {
				t.Errorf("FahrenheitToCelsius(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		name string
		arg  float64
		want float64
	}{
		{
			name: "0 Celsius to Kelvin",
			arg:  0,
			want: 273.15,
		},
		{
			name: "100 Celsius to Kelvin",
			arg:  100,
			want: 373.15,
		},
		{
			name: "-273.15 Celsius to Kelvin",
			arg:  -273.15,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.CelsiusToKelvin(tt.arg)
			if got != tt.want {
				t.Errorf("CelsiusToKelvin(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestKelvinToCelsius(t *testing.T) {
	tests := []struct {
		name string
		arg  float64
		want float64
	}{
		{
			name: "273.15 Kelvin to Celsius",
			arg:  273.15,
			want: 0,
		},
		{
			name: "373.15 Kelvin to Celsius",
			arg:  373.15,
			want: 100,
		},
		{
			name: "0 Kelvin to Celsius",
			arg:  0,
			want: -273.15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.KelvinToCelsius(tt.arg)
			if got != tt.want {
				t.Errorf("KelvinToCelsius(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	tests := []struct {
		name string
		arg  float64
		want float64
	}{
		{
			name: "0 Fahrenheit to Kelvin",
			arg:  0,
			want: 255.37,
		},
		{
			name: "100 Fahrenheit to Kelvin",
			arg:  100,
			want: 310.93,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.FahrenheitToKelvin(tt.arg)
			if got != tt.want {
				t.Errorf("FahrenheitToCelsius(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	tests := []struct {
		name string
		arg  float64
		want float64
	}{
		{
			name: "255.37 Kelvin to Fahrenheit",
			arg:  255.37,
			want: 0,
		},
		{
			name: "310.93 Kelvin to Fahrenheit",
			arg:  310.93,
			want: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conversion.KelvinToFahrenheit(tt.arg)
			if got != tt.want {
				t.Errorf("KelvinToFahrenheit(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}
