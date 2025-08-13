package ratelimiter

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// TokenBucket implements a thread-safe token bucket rate limiter that allows
// bursts up to a certain capacity while maintaining a steady refill rate.
type TokenBucket struct {
	mu         sync.Mutex
	capacity   int
	tokens     float64
	refillRate float64
	last       time.Time
}

// NewTokenBucket creates a new TokenBucket with the specified capacity and refill rate.
// The capacity determines the maximum number of tokens that can be stored,
// and refillRate specifies how many tokens are added per second.
// If capacity is less than 1 or equal to 0, it defaults to 1.
// If refillRate is less than or equal to 0, it defaults to 1.
// Returns a TokenBucket that starts with full capacity.
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

// Allow checks if one token is available and consumes it if so.
// Returns true if the token was successfully consumed, false otherwise.
// This is a convenience method that calls AllowN(1).
func (t *TokenBucket) Allow() bool {

	return t.AllowN(1)
}

// AllowN checks if n tokens are available and consumes them atomically.
// Returns true if n tokens were successfully consumed, false otherwise.
// This method is thread-safe and non-blocking.
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

// Wait blocks until one token is available and consumes it, or until the context is cancelled.
// Returns an error if the context is cancelled before a token becomes available.
// This is a convenience method that calls WaitN(ctx, 1).
func (t *TokenBucket) Wait(ctx context.Context) error {

	return t.WaitN(ctx, 1)
}

// WaitN blocks until n tokens are available and consumes them, or until the context is cancelled.
// Returns an error if the requested number of tokens exceeds the bucket capacity,
// or if the context is cancelled before the tokens become available.
// This method is thread-safe and will wait as long as necessary (respecting context cancellation).
func (t *TokenBucket) WaitN(ctx context.Context, n int) error {

	if n <= 0 {

		return nil
	}

	// Read capacity under lock to avoid race with SetCapacity
	t.mu.Lock()
	cap := t.capacity
	t.mu.Unlock()

	if n > cap {
		return fmt.Errorf("requested tokens %d exceeds capacity %v", n, cap)
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

		t.mu.Unlock()
		// Call nextAvailableDuration while holding the mutex
		d := t.nextAvailableDuration(n)

		// Handle the case where nextAvailableDuration returns -1 (impossible request)
		if d == -1 {
			return fmt.Errorf("requested tokens %d cannot be fulfilled", n)
		}

		if d <= 0 {
			// When nextAvailableDuration returns 0, it means tokens should be available,
			// but there might be timing issues or race conditions. Sleep briefly to
			// prevent busy-waiting while maintaining responsiveness to cancellation.
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(1 * time.Millisecond):
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

// Tokens returns the current number of available tokens as a float64.
// This method is thread-safe and does not consume any tokens.
// The returned value is approximate and may change immediately after the call returns.
func (t *TokenBucket) Tokens() float64 {
	now := time.Now()
	t.mu.Lock()
	defer t.mu.Unlock()
	t.refill(now)

	return t.tokens
}

// SetCapacity adjusts the bucket capacity at runtime.
// If the new capacity is smaller than the current number of tokens,
// the token count is reduced to match the new capacity.
// This method is thread-safe.
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

// SetRefillRate adjusts the token refill rate (tokens per second) at runtime.
// The rate must be positive; invalid rates are ignored.
// This method is thread-safe.
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
