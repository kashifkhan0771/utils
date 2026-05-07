package retry

import (
	"context"
	"fmt"
	"math"
	"time"
)

type Options struct {
	MaxAttempts  uint
	TotalTimeout time.Duration
	// Backoff returns how long to wait before the next attempt.
	// attempt is zero-indexed (0 = pause after the first failure).
	Backoff func(attempt uint) time.Duration

	// ShouldRetry reports whether the given error is retryable.
	// Return false to abort immediately.
	ShouldRetry func(err error) bool
}

type RetryFunc[T any] func(ctx context.Context) (T, error)

// Do calls fn repeatedly until fn succeeds, ShouldRetry returns false,
// MaxAttempts is reached, or TotalTimeout elapses.
func Do[T any](ctx context.Context, opts Options, fn RetryFunc[T]) (T, error) {
	var zero T
	if ctx == nil {
		return zero, fmt.Errorf("retry: context must not be nil")
	}

	if opts.TotalTimeout > 0 {
		var cls context.CancelFunc
		ctx, cls = context.WithTimeout(ctx, opts.TotalTimeout)
		defer cls()
	}

	shouldRetry := opts.ShouldRetry
	if shouldRetry == nil {
		shouldRetry = func(error) bool { return true }
	}
	backoff := opts.Backoff
	if backoff == nil {
		backoff = func(uint) time.Duration { return 0 }
	}

	var lastErr error
	for attempt := range opts.MaxAttempts {
		ret, err := fn(ctx)
		if err == nil {
			return ret, nil
		}
		lastErr = err

		if !shouldRetry(err) {
			return zero, err
		}

		select {
		case <-time.After(backoff(attempt)):
		case <-ctx.Done():
			return zero, ctx.Err()
		}
	}

	return zero, fmt.Errorf("retry: max attempts reached: %w", lastErr)
}

// DoVoid is a convenience wrapper around [Do] for operations that return no value.
func DoVoid(ctx context.Context, opts Options, fn func(ctx context.Context) error) error {
	_, err := Do(ctx, opts, func(ctx context.Context) (struct{}, error) {
		return struct{}{}, fn(ctx)
	})

	return err
}

// FixedBackoff returns a backoff that always waits exactly d.
func FixedBackoff(d time.Duration) func(attempt uint) time.Duration {
	return func(attempt uint) time.Duration {
		return d
	}
}

// LinearBackoff returns a backoff that waits d * attempt.
func LinearBackoff(d time.Duration) func(attempt uint) time.Duration {
	return func(attempt uint) time.Duration {
		maxDuration := time.Duration(math.MaxInt64)
		if attempt > math.MaxInt64 {
			return maxDuration
		}
		dur := time.Duration(attempt)
		if d > 0 && dur > maxDuration/d {
			return maxDuration
		}

		return d * dur
	}
}

// ExponentialBackoff returns a backoff that waits min(d * 2^attempt, 2^63 - 1).
func ExponentialBackoff(d time.Duration) func(attempt uint) time.Duration {
	return func(attempt uint) time.Duration {
		maxDuration := time.Duration(1<<63 - 1)
		if attempt == 0 || d < 0 {
			return d
		}

		if attempt >= 63 {
			return maxDuration
		}

		multiplier := time.Duration(1) << attempt
		if d > maxDuration/multiplier {
			return maxDuration
		}

		return d * multiplier
	}
}
