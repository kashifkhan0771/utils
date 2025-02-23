package rand

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strings"
)

const (
	// DefaultCharset defines the default characters
	DefaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// DefaultLength defines the default length for random string
	DefaultLength = 10
)

func Number() (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %w", err)
	}

	return n.Int64(), nil
}

// NumberInRange generates a random number between min and max
func NumberInRange(min, max int64) (int64, error) {
	if min > max {
		return 0, fmt.Errorf("min (%d) cannot be greater than max (%d)", min, max)
	}

	// Early return if min equals max
	if min == max {
		return min, nil
	}

	rangeSize := max - min + 1
	// Calculate the largest multiple of rangeSize that fits in MaxInt64
	limit := math.MaxInt64 - (math.MaxInt64 % rangeSize)

	for {
		n, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			return 0, fmt.Errorf("failed to generate random number in range: %w", err)
		}

		if n.Int64() < limit {
			return min + (n.Int64() % rangeSize), nil
		}
		// If we're above the limit, try again to ensure uniform distribution
	}
}

// String generates a random string using the default constants
func String() (string, error) {
	return StringWithLength(DefaultLength)
}

// StringWithLength generates a random string of the specified length using the default charset
func StringWithLength(length int) (string, error) {
	return StringWithCharset(length, DefaultCharset)
}

// Pick returns a random element from the provided slice
func Pick[T any](slice []T) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, fmt.Errorf("cannot pick from empty slice")
	}

	idx, err := NumberInRange(0, int64(len(slice)-1))
	if err != nil {
		return zero, fmt.Errorf("failed to generate random index: %w", err)
	}

	return slice[int(idx)], nil
}

// Shuffle reorders the elements in the provided slice
func Shuffle[T any](slice []T) error {
	if len(slice) == 0 {
		return nil // Nothing to shuffle in an empty slice
	}

	for i := len(slice) - 1; i > 0; i-- {
		j, err := NumberInRange(0, int64(i))
		if err != nil {
			return fmt.Errorf("failed to generate random index: %w", err)
		}
		slice[i], slice[int(j)] = slice[int(j)], slice[i]
	}

	return nil
}

// StringWithCharset generates a random string with the specified length and character set
func StringWithCharset(length int, charset string) (string, error) {
	if length < 0 {
		return "", fmt.Errorf("length cannot be negative: %d", length)
	}

	trimmedCharset := strings.TrimSpace(charset)
	if len(trimmedCharset) == 0 {
		return "", fmt.Errorf("charset cannot be empty or contain only whitespace")
	}

	result := make([]byte, length)
	charsetLength := big.NewInt(int64(len(trimmedCharset)))

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", fmt.Errorf("failed to generate random string: %w", err)
		}
		result[i] = trimmedCharset[n.Int64()]
	}

	return string(result), nil
}
