package retry

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

// helpers

func defaultOpts() Options {
	return Options{
		MaxAttempts:  3,
		TotalTimeout: 5 * time.Second,
		Backoff:      FixedBackoff(10 * time.Millisecond),
		ShouldRetry:  func(err error) bool { return true },
	}
}

func counter(failTimes int, returnVal string) (func(context.Context) (string, error), *int) {
	attempts := 0
	return func(ctx context.Context) (string, error) {
		attempts++
		if attempts <= failTimes {
			return "", errors.New("transient")
		}
		return returnVal, nil
	}, &attempts
}

func TestDo_SuccessFirstAttempt(t *testing.T) {
	result, err := Do(context.Background(), defaultOpts(), func(ctx context.Context) (string, error) {
		return "ok", nil
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != "ok" {
		t.Fatalf("expected 'ok', got %q", result)
	}
}

func TestDo_SuccessAfterRetries(t *testing.T) {
	fn, attempts := counter(2, "ok")
	result, err := Do(context.Background(), defaultOpts(), fn)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != "ok" {
		t.Fatalf("expected 'ok', got %q", result)
	}
	if *attempts != 3 {
		t.Fatalf("expected 3 attempts, got %d", *attempts)
	}
}

func TestDo_MaxAttemptsReached(t *testing.T) {
	fn, attempts := counter(999, "")
	_, err := Do(context.Background(), defaultOpts(), fn)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if *attempts != 3 {
		t.Fatalf("expected 3 attempts, got %d", *attempts)
	}
	if err.Error() != "retry: max attempts reached: transient" {
		t.Fatalf("expected wrapped max-attempts error, got %v", err)
	}
}

func TestDo_ZeroMaxAttempts(t *testing.T) {
	opts := defaultOpts()
	opts.MaxAttempts = 0

	called := 0
	_, err := Do(context.Background(), opts, func(ctx context.Context) (string, error) {
		called++
		return "", errors.New("should not run")
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "retry: MaxAttempts must be greater than 0" {
		t.Fatalf("unexpected error: %v", err)
	}
	if called != 0 {
		t.Fatalf("expected fn not to be called, got %d calls", called)
	}
}

func TestDo_NoBackoffAfterFinalAttempt(t *testing.T) {
	var backoffAttempts []uint
	opts := Options{
		MaxAttempts:  2,
		TotalTimeout: 5 * time.Second,
		Backoff: func(attempt uint) time.Duration {
			backoffAttempts = append(backoffAttempts, attempt)
			return 0
		},
		ShouldRetry: func(err error) bool { return true },
	}

	fnCalls := 0
	_, err := Do(context.Background(), opts, func(ctx context.Context) (string, error) {
		fnCalls++
		return "", errors.New("transient")
	})

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "retry: max attempts reached: transient" {
		t.Fatalf("expected wrapped last error, got %v", err)
	}
	if fnCalls != 2 {
		t.Fatalf("expected fn to run twice, got %d", fnCalls)
	}
	// One inter-attempt backoff only; must not call Backoff after the final failure.
	if len(backoffAttempts) != 1 || backoffAttempts[0] != 0 {
		t.Fatalf("expected exactly one Backoff call with attempt 0, got %v", backoffAttempts)
	}
}

func TestDo_FinalAttemptPreservesLastErrorOverTimeout(t *testing.T) {
	// Previously the final failure slept, so a short remaining deadline could
	// replace the real last error with context.DeadlineExceeded.
	transient := errors.New("transient")
	opts := Options{
		MaxAttempts:  1,
		TotalTimeout: 30 * time.Millisecond,
		Backoff:      FixedBackoff(200 * time.Millisecond),
		ShouldRetry:  func(err error) bool { return true },
	}

	_, err := Do(context.Background(), opts, func(ctx context.Context) (string, error) {
		return "", transient
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected last error, got timeout: %v", err)
	}
	if !errors.Is(err, transient) {
		t.Fatalf("expected transient wrapped error, got %v", err)
	}
}

func TestDo_ShouldRetryFalseStopsEarly(t *testing.T) {
	permanent := errors.New("permanent")
	opts := defaultOpts()
	opts.ShouldRetry = func(err error) bool { return !errors.Is(err, permanent) }

	attempts := 0
	_, err := Do(context.Background(), opts, func(ctx context.Context) (string, error) {
		attempts++
		return "", permanent
	})

	if !errors.Is(err, permanent) {
		t.Fatalf("expected permanent error, got %v", err)
	}
	if attempts != 1 {
		t.Fatalf("expected 1 attempt, got %d", attempts)
	}
}

func TestDo_TotalTimeoutExceeded(t *testing.T) {
	opts := Options{
		MaxAttempts:  10,
		TotalTimeout: 50 * time.Millisecond,
		Backoff:      FixedBackoff(30 * time.Millisecond),
		ShouldRetry:  func(err error) bool { return true },
	}

	start := time.Now()
	_, err := Do(context.Background(), opts, func(ctx context.Context) (string, error) {
		return "", errors.New("transient")
	})

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected DeadlineExceeded, got %v", err)
	}
	if elapsed := time.Since(start); elapsed > 200*time.Millisecond {
		t.Fatalf("loop ran too long: %v", elapsed)
	}
}

func TestDo_ParentContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	opts := Options{
		MaxAttempts:  10,
		TotalTimeout: 5 * time.Second,
		Backoff:      FixedBackoff(30 * time.Millisecond),
		ShouldRetry:  func(err error) bool { return true },
	}

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	_, err := Do(ctx, opts, func(ctx context.Context) (string, error) {
		return "", errors.New("transient")
	})

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected Canceled, got %v", err)
	}
}

func TestDo_ZeroValueReturnedOnError(t *testing.T) {
	result, err := Do(context.Background(), defaultOpts(), func(ctx context.Context) (string, error) {
		return "partial", errors.New("fail")
	})
	if err == nil {
		t.Fatal("expected error")
	}

	if result != "" {
		t.Fatalf("expected zero value, got %v", result)
	}
}

func TestDoVoid_Success(t *testing.T) {
	called := 0
	err := DoVoid(context.Background(), defaultOpts(), func(ctx context.Context) error {
		called++
		return nil
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if called != 1 {
		t.Fatalf("expected 1 call, got %d", called)
	}
}

func TestDoVoid_RetriesAndFails(t *testing.T) {
	called := 0
	err := DoVoid(context.Background(), defaultOpts(), func(ctx context.Context) error {
		called++
		return errors.New("fail")
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if called != 3 {
		t.Fatalf("expected 3 calls, got %d", called)
	}
}

func TestFixedBackoff(t *testing.T) {
	b := FixedBackoff(100 * time.Millisecond)
	for _, attempt := range []uint{0, 1, 2, 5} {
		if got := b(attempt); got != 100*time.Millisecond {
			t.Fatalf("attempt %d: expected 100ms, got %v", attempt, got)
		}
	}
}

func TestLinearBackoff(t *testing.T) {
	b := LinearBackoff(100 * time.Millisecond)
	cases := map[uint]time.Duration{
		0: 0,
		1: 100 * time.Millisecond,
		2: 200 * time.Millisecond,
		5: 500 * time.Millisecond,
	}
	for attempt, expected := range cases {
		if got := b(attempt); got != expected {
			t.Fatalf("attempt %d: expected %v, got %v", attempt, expected, got)
		}
	}
}

func TestExponentialBackoff(t *testing.T) {
	b := ExponentialBackoff(100 * time.Millisecond)
	cases := map[uint]time.Duration{
		0: 100 * time.Millisecond,
		1: 200 * time.Millisecond,
		2: 400 * time.Millisecond,
		3: 800 * time.Millisecond,
	}
	for attempt, expected := range cases {
		if got := b(attempt); got != expected {
			t.Fatalf("attempt %d: expected %v, got %v", attempt, expected, got)
		}
	}
}

func BenchmarkDo_AlwaysSucceeds(b *testing.B) {
	opts := Options{
		MaxAttempts:  3,
		TotalTimeout: 5 * time.Second,
		Backoff:      FixedBackoff(0),
		ShouldRetry:  func(err error) bool { return true },
	}

	b.ResetTimer()
	for b.Loop() {
		_, err := Do(context.Background(), opts, func(ctx context.Context) (string, error) {
			return "ok", nil
		})

		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDo_AlwaysFails(b *testing.B) {
	opts := Options{
		MaxAttempts:  3,
		TotalTimeout: 5 * time.Second,
		Backoff:      FixedBackoff(0),
		ShouldRetry:  func(err error) bool { return true },
	}

	b.ResetTimer()
	for b.Loop() {
		_, err := Do(context.Background(), opts, func(ctx context.Context) (string, error) {
			return "", errors.New("fail")
		})

		if err == nil {
			b.Fatal("expected error, got nil")
		}
	}
}

func BenchmarkDo_SuccessOnLastAttempt(b *testing.B) {
	opts := Options{
		MaxAttempts:  5,
		TotalTimeout: 5 * time.Second,
		Backoff:      FixedBackoff(0),
		ShouldRetry:  func(err error) bool { return true },
	}

	b.ResetTimer()
	for b.Loop() {
		attempt := 0
		_, err := Do(context.Background(), opts, func(ctx context.Context) (string, error) {
			attempt++
			if attempt < 5 {
				return "", errors.New("transient")
			}
			return "ok", nil
		})

		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDo_ShouldRetryFalse(b *testing.B) {
	permanent := errors.New("permanent")
	opts := Options{
		MaxAttempts:  10,
		TotalTimeout: 5 * time.Second,
		Backoff:      FixedBackoff(0),
		ShouldRetry:  func(err error) bool { return !errors.Is(err, permanent) },
	}

	b.ResetTimer()
	for b.Loop() {
		_, err := Do(context.Background(), opts, func(ctx context.Context) (string, error) {
			return "", permanent
		})

		if !errors.Is(err, permanent) {
			b.Fatalf("expected permanent error, got %v", err)
		}
	}
}

func BenchmarkDoVoid_AlwaysSucceeds(b *testing.B) {
	opts := Options{
		MaxAttempts:  3,
		TotalTimeout: 5 * time.Second,
		Backoff:      FixedBackoff(0),
		ShouldRetry:  func(err error) bool { return true },
	}

	b.ResetTimer()
	for b.Loop() {
		err := DoVoid(context.Background(), opts, func(ctx context.Context) error {
			return nil
		})

		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFixedBackoff(b *testing.B) {
	f := FixedBackoff(100 * time.Millisecond)
	b.ResetTimer()
	for b.Loop() {
		f(3)
	}
}

func BenchmarkLinearBackoff(b *testing.B) {
	f := LinearBackoff(100 * time.Millisecond)
	b.ResetTimer()
	for b.Loop() {
		f(3)
	}
}

func BenchmarkExponentialBackoff(b *testing.B) {
	f := ExponentialBackoff(100 * time.Millisecond)
	b.ResetTimer()
	for b.Loop() {
		f(3)
	}
}

func BenchmarkDo_MaxAttempts(b *testing.B) {
	for _, maxAttempts := range []uint{1, 5, 10, 50} {
		b.Run(fmt.Sprintf("attempts=%d", maxAttempts), func(b *testing.B) {
			opts := Options{
				MaxAttempts:  maxAttempts,
				TotalTimeout: 30 * time.Second,
				Backoff:      FixedBackoff(0),
				ShouldRetry:  func(err error) bool { return true },
			}

			b.ResetTimer()
			for b.Loop() {
				_, err := Do(context.Background(), opts, func(ctx context.Context) (string, error) {
					return "", errors.New("fail")
				})

				if err == nil {
					b.Fatal("expected error, got nil")
				}
			}
		})
	}
}

func BenchmarkDo_BackoffStrategies(b *testing.B) {
	strategies := map[string]func(uint) time.Duration{
		"fixed":       FixedBackoff(0),
		"linear":      LinearBackoff(0),
		"exponential": ExponentialBackoff(0),
	}

	for name, backoff := range strategies {
		b.Run(name, func(b *testing.B) {
			opts := Options{
				MaxAttempts:  3,
				TotalTimeout: 5 * time.Second,
				Backoff:      backoff,
				ShouldRetry:  func(err error) bool { return true },
			}

			b.ResetTimer()
			for b.Loop() {
				_, err := Do(context.Background(), opts, func(ctx context.Context) (string, error) {
					return "", errors.New("fail")
				})

				if err == nil {
					b.Fatal("expected error, got nil")
				}
			}
		})
	}
}
