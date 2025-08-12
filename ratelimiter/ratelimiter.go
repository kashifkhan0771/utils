package ratelimiter

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	mu         sync.Mutex
	capacity   int
	tokens     float64
	refillRate float64
	last       time.Time
}

func NewTokenBucket(capacity int, refillRate float64) *TokenBucket {
	if capacity < 1 {
		capacity = 1
	}
	if refillRate <= 0 {
		refillRate = 1
	}

	return &TokenBucket{
		capacity:   capacity,
		tokens:     float64(capacity),
		refillRate: refillRate,
		last:       time.Now(),
	}
}

// refill adds tokens according to elapsed time.
// Caller must hold t.mu
func (t *TokenBucket) refill(now time.Time) {
	if now.Before(t.last) {
		t.last = now

		return
	}
	elapsed := now.Sub(t.last).Seconds()
	if elapsed <= 0 {
		return
	}
	add := elapsed * t.refillRate
	// Add tokens directly to the float field, preserving fractional tokens
	t.tokens += add
	// Cap at capacity
	if t.tokens > float64(t.capacity) {
		t.tokens = float64(t.capacity)
	}
	t.last = now
}

func (t *TokenBucket) Allow() bool {
	return t.AllowN(1)
}

// AllowN checks if n tokens are available and consumes them atomically.
func (t *TokenBucket) AllowN(n int) bool {
	if n <= 0 {
		return true
	}
	now := time.Now()
	t.mu.Lock()
	defer t.mu.Unlock()
	t.refill(now)
	if t.tokens >= float64(n) {
		t.tokens -= float64(n)

		return true
	}

	return false
}

// nextAvailableDuration returns how long until n tokens are available from now.
// Caller must hold the mutex.
func (t *TokenBucket) nextAvailableDuration(n int, now time.Time) time.Duration {
	if n <= 0 {
		return 0
	}
	if t.tokens >= float64(n) {
		return 0
	}

	needed := float64(n) - t.tokens
	secs := needed / t.refillRate

	return time.Duration(secs * float64(time.Second))
}

// Wait blocks until one token is available or ctx is done.
func (t *TokenBucket) Wait(ctx context.Context) error {
	return t.WaitN(ctx, 1)
}

// WaitN blocks until n tokens are available and consumes them, or ctx is done.
func (t *TokenBucket) WaitN(ctx context.Context, n int) error {
	if n <= 0 {
		return nil
	}
	t.mu.Lock()
	cap := t.capacity
	t.mu.Unlock()
	if n > cap {
		return fmt.Errorf("requested tokens %d exceeds capacity %v", n, t.capacity)
	}
	for {
		now := time.Now()
		t.mu.Lock()
		t.refill(now)
		if t.tokens >= float64(n) {
			t.tokens -= float64(n)
			t.mu.Unlock()

			return nil
		}

		d := t.nextAvailableDuration(n, now)
		t.mu.Unlock()
		if d <= 0 {
			// When nextAvailableDuration returns 0, it means tokens should be available,
			// but there might be timing issues or race conditions. Sleep briefly to
			// prevent busy-waiting while maintaining responsiveness to cancellation.
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(5 * time.Millisecond):
				// Yield CPU with a slightly longer sleep to prevent excessive busy-waiting
				// while still being responsive to context cancellation
			}

			continue
		}
		timer := time.NewTimer(d)
		select {
		case <-ctx.Done():
			timer.Stop()

			return ctx.Err()
		case <-timer.C:
			// Timer expired, loop to check tokens again

		}
	}
}

// Tokens returns the current available tokens (approximate). This is safe and does not consume tokens.
func (t *TokenBucket) Tokens() float64 {
	now := time.Now()
	t.mu.Lock()
	defer t.mu.Unlock()
	t.refill(now)

	return t.tokens
}

// SetCapacity lets you adjust the capacity at runtime. If new capacity is smaller,
// tokens are trimmed to the new capacity. Thread-safe.
func (t *TokenBucket) SetCapacity(cap int) {
	if cap < 1 {
		return
	}
	now := time.Now()
	t.mu.Lock()
	defer t.mu.Unlock()
	t.refill(now)
	t.capacity = cap
	if t.tokens > float64(t.capacity) {
		t.tokens = float64(t.capacity)
	}
}

// SetRefillRate adjusts refill rate (tokens per second). Thread-safe.
func (t *TokenBucket) SetRefillRate(rate float64) {
	if rate <= 0 {
		return
	}
	now := time.Now()
	t.mu.Lock()
	defer t.mu.Unlock()
	t.refill(now)
	t.refillRate = rate
}
