package caching

import (
	"math/big"
	"sync"
	"testing"
)

type testCase[T any] struct {
	name  string
	input int
	want  T
}

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

	tests := []testCase[*big.Int]{
		{
			name:  "success - calculate factorial of 5",
			input: 5,
			want:  big.NewInt(120),
		},
		{
			name:  "success - calculate factorial of 0",
			input: 0,
			want:  big.NewInt(1),
		},
		{
			name:  "success - repeated call with factorial of 5",
			input: 5,
			want:  big.NewInt(120),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cachedFactorial(tt.input); got.Cmp(tt.want) != 0 {
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

	tests := []testCase[int]{
		{
			name:  "success - double 4",
			input: 4,
			want:  8,
		},
		{
			name:  "success - double 0",
			input: 0,
			want:  0,
		},
		{
			name:  "success - repeated call with double 4",
			input: 4,
			want:  8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := cachedDouble(tt.input); got != tt.want {
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
	const numRoutines = 10

	results := make([]int, numRoutines)
	wg.Add(numRoutines)
	for i := range numRoutines {
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

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func BenchmarkFib(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_ = fib(30)
	}
}

func BenchmarkCachedFib(b *testing.B) {
	cachedFib := CacheWrapper(fib)
	_ = cachedFib(30) // warm-up the cache before running the benchmark
	b.ReportAllocs()
	for b.Loop() {
		_ = cachedFib(30)
	}
}

func BenchmarkSafeCachedFib(b *testing.B) {
	cachedFib := SafeCacheWrapper(fib)
	_ = cachedFib(30) // warm-up the cache before running the benchmark

	b.ReportAllocs()
	for b.Loop() {
		_ = cachedFib(30)
	}
}

func BenchmarkConcurrentSafeCachedFib(b *testing.B) {
	cachedFib := SafeCacheWrapper(fib)
	_ = cachedFib(30) // warm-up the cache before running the benchmark

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = cachedFib(30)
		}
	})
}
