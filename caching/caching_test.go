package caching

import (
	"math/big"
	"sync"
	"testing"
)

// TestCacheWrapper tests the non-thread-safe caching wrapper.
func TestCacheWrapper(t *testing.T) {
	// Example function: Calculate factorial of a number.
	factorial := func(n int) *big.Int {
		result := big.NewInt(1)
		for i := 2; i <= n; i++ {
			result.Mul(result, big.NewInt(int64(i)))
		}
		return result
	}

	cachedFactorial := CacheWrapper(factorial)

	tests := []struct {
		name string
		arg  int
		want *big.Int
	}{
		{
			name: "success - calculate factorial of 5",
			arg:  5,
			want: big.NewInt(120),
		},
		{
			name: "success - calculate factorial of 0",
			arg:  0,
			want: big.NewInt(1),
		},
		{
			name: "success - repeated call with factorial of 5",
			arg:  5,
			want: big.NewInt(120),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cachedFactorial(tt.arg); got.Cmp(tt.want) != 0 {
				t.Errorf("CacheWrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSafeCacheWrapper tests the thread-safe caching wrapper.
func TestSafeCacheWrapper(t *testing.T) {
	// Example function: Double a number (for simplicity in concurrent tests).
	double := func(n int) int {
		return n * 2
	}

	cachedDouble := SafeCacheWrapper(double)

	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			name: "success - double 4",
			arg:  4,
			want: 8,
		},
		{
			name: "success - double 0",
			arg:  0,
			want: 0,
		},
		{
			name: "success - repeated call with double 4",
			arg:  4,
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cachedDouble(tt.arg); got != tt.want {
				t.Errorf("SafeCacheWrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSafeCacheWrapperConcurrency tests the thread-safe caching in a concurrent environment.
func TestSafeCacheWrapperConcurrency(t *testing.T) {
	// Example function: Square a number.
	square := func(n int) int {
		return n * n
	}

	cachedSquare := SafeCacheWrapper(square)
	var wg sync.WaitGroup

	// Test concurrency with multiple goroutines.
	results := make([]int, 10)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			defer wg.Done()
			results[idx] = cachedSquare(4) // All goroutines calculate square of 4.
		}(i)
	}
	wg.Wait()

	// Verify all results are correct and identical.
	for _, result := range results {
		if result != 16 {
			t.Errorf("SafeCacheWrapperConcurrency() = %v, want %v", result, 16)
		}
	}
}
