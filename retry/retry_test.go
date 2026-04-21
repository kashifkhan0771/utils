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
	_, err := Do(context.Background(), defaultOpts(), func(ctx context.Context) (string, error) {
		return "partial", errors.New("fail")
	})
	if err == nil {
		t.Fatal("expected error")
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
	for _, attempt := range []int{0, 1, 2, 5} {
		if got := b(attempt); got != 100*time.Millisecond {
			t.Fatalf("attempt %d: expected 100ms, got %v", attempt, got)
		}
	}
}

func TestLinearBackoff(t *testing.T) {
	b := LinearBackoff(100 * time.Millisecond)
	cases := map[int]time.Duration{
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
	cases := map[int]time.Duration{
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
		Do(context.Background(), opts, func(ctx context.Context) (string, error) {
			return "ok", nil
		})
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
		Do(context.Background(), opts, func(ctx context.Context) (string, error) {
			return "", errors.New("fail")
		})
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
		Do(context.Background(), opts, func(ctx context.Context) (string, error) {
			attempt++
			if attempt < 5 {
				return "", errors.New("transient")
			}
			return "ok", nil
		})
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
		Do(context.Background(), opts, func(ctx context.Context) (string, error) {
			return "", permanent
		})
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
		DoVoid(context.Background(), opts, func(ctx context.Context) error {
			return nil
		})
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
	for _, maxAttempts := range []int{1, 5, 10, 50} {
		b.Run(fmt.Sprintf("attempts=%d", maxAttempts), func(b *testing.B) {
			opts := Options{
				MaxAttempts:  maxAttempts,
				TotalTimeout: 30 * time.Second,
				Backoff:      FixedBackoff(0),
				ShouldRetry:  func(err error) bool { return true },
			}

			b.ResetTimer()
			for b.Loop() {
				Do(context.Background(), opts, func(ctx context.Context) (string, error) {
					return "", errors.New("fail")
				})
			}
		})
	}
}

func BenchmarkDo_BackoffStrategies(b *testing.B) {
	strategies := map[string]func(int) time.Duration{
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
				Do(context.Background(), opts, func(ctx context.Context) (string, error) {
					return "", errors.New("fail")
				})
			}
		})
	}
}
