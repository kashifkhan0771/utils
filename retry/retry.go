package retry

import (
	"context"
	"fmt"
	"time"
)

type Options struct {
	MaxAttempts  int
	TotalTimeout time.Duration
	// Backoff returns how long to wait before the next attempt.
	// attempt is zero-indexed (0 = pause after the first failure).
	Backoff func(attempt int) time.Duration

	// ShouldRetry reports whether the given error is retryable.
	// Return false to abort immediately.
	ShouldRetry func(err error) bool
}

type RetryFunc[T any] func(ctx context.Context) (T, error)

// Do calls fn repeatedly until fn succeeds, ShouldRetry returns false,
// MaxAttempts is reached, or TotalTimeout elapses.
func Do[T any](ctx context.Context, opts Options, fn RetryFunc[T]) (T, error) {
	ctx, cls := context.WithTimeout(ctx, opts.TotalTimeout)
	defer cls()

	var zero T
	for attempt := range opts.MaxAttempts {
		ret, err := fn(ctx)
		if err == nil {
			return ret, nil
		}

		if !opts.ShouldRetry(err) {
			return zero, err
		}

		select {
		case <-time.After(opts.Backoff(attempt)):
		case <-ctx.Done():
			return zero, ctx.Err()
		}

	}

	return zero, fmt.Errorf("max attempts reached")
}

// DoVoid is a convenience wrapper around [Do] for operations that return no value.
func DoVoid(ctx context.Context, opts Options, fn func(ctx context.Context) error) error {
	_, err := Do(ctx, opts, func(ctx context.Context) (struct{}, error) {
		return struct{}{}, fn(ctx)
	})

	return err
}

// FixedBackoff returns a backoff that always waits exactly d.
func FixedBackoff(d time.Duration) func(attempt int) time.Duration {
	return func(attempt int) time.Duration {
		return d
	}
}

// LinearBackoff returns a backoff that waits d * attempt.
func LinearBackoff(d time.Duration) func(attempt int) time.Duration {
	return func(attempt int) time.Duration {
		return d * time.Duration(attempt)
	}
}

// ExponentialBackoff returns a backoff that waits d * 2^attempt.
func ExponentialBackoff(d time.Duration) func(attempt int) time.Duration {
	return func(attempt int) time.Duration {
		return d * time.Duration(1<<attempt)
	}
}
