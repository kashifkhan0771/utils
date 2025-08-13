package ratelimiter

import "time"

// refill adds tokens according to elapsed time.
// Caller must hold the mutex
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

// nextAvailableDuration returns how long until n tokens are available from now.
// Returns -1 if n tokens can never be available or invalid refill rate.
// Caller must hold the mutex.
func (t *TokenBucket) nextAvailableDuration(n int) time.Duration {
	if n <= 0 {
		return 0
	}

	if n > t.capacity || t.refillRate < 0 {
		return -1
	}

	if t.tokens >= float64(n) {
		return 0
	}

	needed := float64(n) - t.tokens
	secs := needed / t.refillRate

	return time.Duration(secs * float64(time.Second))
}
