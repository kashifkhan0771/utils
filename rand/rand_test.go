package rand

import (
	"strings"
	"testing"
)

func TestNumber(t *testing.T) {
	// Generate random numbers to ensure they are different
	n1, err := Number()
	if err != nil {
		t.Errorf("Number() error = %v", err)
		return
	}

	n2, err := Number()
	if err != nil {
		t.Errorf("Number() error = %v", err)
		return
	}

	// Numbers should be different (theres an extremely small chance they could be the same)
	if n1 == n2 {
		t.Errorf("Generated numbers are equal: %v", n1)
	}
}

func TestNumberInRange(t *testing.T) {
	tests := []struct {
		name    string
		min     int64
		max     int64
		wantErr bool
	}{
		{
			name:    "success - valid range",
			min:     1,
			max:     100,
			wantErr: false,
		},
		{
			name:    "success - same min and max",
			min:     5,
			max:     5,
			wantErr: false,
		},
		{
			name:    "fail - invalid range",
			min:     100,
			max:     1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NumberInRange(tt.min, tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("NumberInRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				if got < tt.min || got > tt.max {
					t.Errorf("NumberInRange() = %v, want between %v and %v", got, tt.min, tt.max)
				}
			}
		})
	}
}

func TestString(t *testing.T) {
	s1, err := String()
	if err != nil {
		t.Errorf("String() error = %v", err)
		return
	}

	s2, err := String()
	if err != nil {
		t.Errorf("String() error = %v", err)
		return
	}

	if len(s1) != DefaultLength {
		t.Errorf("String() length = %v, want %v", len(s1), DefaultLength)
	}

	if s1 == s2 {
		t.Errorf("Generated strings are equal: %v", s1)
	}
}

func TestStringWithLength(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantLen int
		wantErr bool
	}{
		{
			name:    "success - positive length",
			length:  15,
			wantLen: 15,
			wantErr: false,
		},
		{
			name:    "success - zero length",
			length:  0,
			wantLen: 0,
			wantErr: false,
		},
		{
			name:    "fail - negative length",
			length:  -1,
			wantLen: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringWithLength(tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringWithLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && len(got) != tt.wantLen {
				t.Errorf("StringWithLength() length = %v, want %v", len(got), tt.wantLen)
			}
		})
	}
}

func TestPick(t *testing.T) {
	tests := []struct {
		name    string
		slice   []string
		wantErr bool
	}{
		{
			name:    "success - non-empty slice",
			slice:   []string{"a", "b", "c"},
			wantErr: false,
		},
		{
			name:    "fail - empty slice",
			slice:   []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Pick(tt.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pick() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && !contains(tt.slice, got) {
				t.Errorf("Pick() returned value %v not found in slice", got)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	shuffled := make([]int, len(original))
	copy(shuffled, original)

	err := Shuffle(shuffled)
	if err != nil {
		t.Errorf("Shuffle() error = %v", err)
		return
	}

	// Check that all elements are still present
	if len(shuffled) != len(original) {
		t.Errorf("Shuffle() changed slice length")
		return
	}

	// Check that the order is different (there's a very small chance this could fail)
	different := false
	for i := range original {
		if original[i] != shuffled[i] {
			different = true
			break
		}
	}

	if !different {
		t.Errorf("Shuffle() did not change the order of elements")
	}
}

func TestStringWithCharset(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		charset string
		wantLen int
		wantErr bool
	}{
		{
			name:    "success - numeric only charset",
			length:  10,
			charset: "0123456789",
			wantLen: 10,
			wantErr: false,
		},
		{
			name:    "fail - empty charset",
			length:  10,
			charset: "",
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "fail - negative length",
			length:  -1,
			charset: "abc",
			wantLen: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringWithCharset(tt.length, tt.charset)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringWithCharset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				if len(got) != tt.wantLen {
					t.Errorf("StringWithCharset() length = %v, want %v", len(got), tt.wantLen)
				}

				// Verify all characters are from the charset
				for _, c := range got {
					if !strings.ContainsRune(tt.charset, c) {
						t.Errorf("StringWithCharset() contains character %c not in charset", c)
					}
				}
			}
		})
	}
}

// Helper function for TestPick
func contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
