package ratelimiter

import (
	"context"
	"testing"
	"time"
)

// TestAllowN_Coverage covers basic AllowN usage and edge cases.
func TestAllowN_Coverage(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		refill   float64
		consume  []int
		want     []bool
	}{
		{
			name:     "consume within capacity",
			capacity: 3,
			refill:   1,
			consume:  []int{1, 2, 1},
			want:     []bool{true, true, false},
		},
		{
			name:     "consume zero tokens always true",
			capacity: 1,
			refill:   1,
			consume:  []int{0, 0, 0},
			want:     []bool{true, true, true},
		},
		{
			name:     "consume more than capacity immediately false",
			capacity: 2,
			refill:   1,
			consume:  []int{3},
			want:     []bool{false},
		},
		{
			name:     "consume negative tokens treated as true",
			capacity: 1,
			refill:   1,
			consume:  []int{-1},
			want:     []bool{true},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rl := NewTokenBucket(tc.capacity, tc.refill)
			for i, n := range tc.consume {
				got := rl.AllowN(n)
				if got != tc.want[i] {
					t.Errorf("AllowN(%d) = %v; want %v", n, got, tc.want[i])
				}
			}
		})
	}
}

// TestRefill tests the refill behavior after sleeping.
func TestRefill(t *testing.T) {
	rl := NewTokenBucket(1, 2) // capacity 1, refill 2 tokens/sec

	if !rl.Allow() {
		t.Fatal("expected Allow() to succeed initially")
	}

	time.Sleep(700 * time.Millisecond) // ~1tokens refill

	if !rl.Allow() {
		t.Fatal("expected Allow() to succeed after refill")
	}
}

// TestWaitN tests blocking wait for tokens.
func TestWaitN(t *testing.T) {
	rl := NewTokenBucket(1, 1) // 1 token/sec

	if !rl.Allow() {
		t.Fatal("expected Allow() to succeed initially")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	start := time.Now()
	if err := rl.WaitN(ctx, 1); err != nil {
		t.Fatalf("expected WaitN to succeed, got %v", err)
	}
	elapsed := time.Since(start)

	if elapsed < 800*time.Millisecond {
		t.Fatalf("expected WaitN to block ~1s, blocked only %v", elapsed)
	}
}

// TestWaitContextCancel tests Wait returns error on context cancellation.
func TestWaitContextCancel(t *testing.T) {
	rl := NewTokenBucket(1, 0.5) // slow refill

	if !rl.Allow() {
		t.Fatal("expected Allow() to succeed initially")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	err := rl.Wait(ctx)
	if err == nil {
		t.Fatal("expected Wait to return context error, got nil")
	}
}

// TestTokens tests Tokens() returns approximate tokens available.
func TestTokens(t *testing.T) {
	rl := NewTokenBucket(5, 10) // capacity 5, fast refill

	for i := 0; i < 3; i++ {
		if !rl.Allow() {
			t.Fatalf("expected Allow() to succeed at iteration %d", i)
		}
	}

	tokens := rl.Tokens()
	if tokens < 1.0 || tokens > 5.0 {
		t.Fatalf("unexpected token count: %v", tokens)
	}
}

// TestSetters tests dynamic changes to capacity and refill rate.
func TestSetters(t *testing.T) {
	rl := NewTokenBucket(2, 1)

	if !rl.Allow() {
		t.Fatal("expected first Allow() to succeed")
	}
	if !rl.Allow() {
		t.Fatal("expected second Allow() to succeed")
	}

	rl.SetCapacity(4)
	rl.SetRefillRate(10)

	time.Sleep(120 * time.Millisecond)

	if rl.Tokens() <= 0 {
		t.Fatal("expected tokens after faster refill and capacity increase")
	}
}

// TestNextAvailableDuration tests nextAvailableDuration helper indirectly via WaitN timing.
func TestNextAvailableDuration(t *testing.T) {
	rl := NewTokenBucket(1, 1)

	if !rl.Allow() {
		t.Fatal("expected initial token")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	start := time.Now()
	err := rl.WaitN(ctx, 1)
	if err != nil {
		t.Fatalf("unexpected error in WaitN: %v", err)
	}
	dur := time.Since(start)

	if dur < 900*time.Millisecond {
		t.Fatalf("expected WaitN to wait approx 1s, waited %v", dur)
	}
}

// TestSetCapacityBelowOne tests that capacity less than 1 is ignored.
func TestSetCapacityBelowOne(t *testing.T) {
	rl := NewTokenBucket(1, 1)
	rl.SetCapacity(0) // should not change capacity

	if rl.capacity != 1 {
		t.Fatalf("capacity changed unexpectedly to %v", rl.capacity)
	}
}

// TestSetRefillRateZeroOrNegative tests that zero or negative refill rate is ignored.
func TestSetRefillRateZeroOrNegative(t *testing.T) {
	rl := NewTokenBucket(1, 1)
	rl.SetRefillRate(0) // ignored
	if rl.refillRate != 1 {
		t.Fatalf("refillRate changed unexpectedly to %v", rl.refillRate)
	}
	rl.SetRefillRate(-5) // ignored
	if rl.refillRate != 1 {
		t.Fatalf("refillRate changed unexpectedly to %v", rl.refillRate)
	}
}

// TestRefillTimeTravel ensures refill time cannot go backwards.
func TestRefillTimeTravel(t *testing.T) {
	rl := NewTokenBucket(2, 1)
	rl.mu.Lock()
	// artificially move last refill into the future
	rl.last = rl.last.Add(10 * time.Second)
	rl.mu.Unlock()

	// Consume one token
	ok := rl.Allow()
	if !ok {
		t.Fatal("expected Allow() to succeed even with time travel")
	}
}

// TestNewTokenBucketEdgeCases tests edge cases in constructor.
func TestNewTokenBucketEdgeCases(t *testing.T) {
	// Test capacity < 1 defaults to 1
	rl := NewTokenBucket(0, 1)
	if rl.capacity != 1 {
		t.Errorf("expected capacity 1 for input 0, got %v", rl.capacity)
	}

	rl = NewTokenBucket(-5, 1)
	if rl.capacity != 1 {
		t.Errorf("expected capacity 1 for input -5, got %v", rl.capacity)
	}

	// Test refillRate <= 0 defaults to 1
	rl = NewTokenBucket(5, 0)
	if rl.refillRate != 1 {
		t.Errorf("expected refillRate 1 for input 0, got %v", rl.refillRate)
	}

	rl = NewTokenBucket(5, -2.5)
	if rl.refillRate != 1 {
		t.Errorf("expected refillRate 1 for input -2.5, got %v", rl.refillRate)
	}
}

// TestWaitNEdgeCases tests edge cases in WaitN.
func TestWaitNEdgeCases(t *testing.T) {
	rl := NewTokenBucket(2, 1)

	// Test n <= 0 returns immediately
	ctx := context.Background()
	err := rl.WaitN(ctx, 0)
	if err != nil {
		t.Errorf("expected WaitN(0) to succeed immediately, got %v", err)
	}

	err = rl.WaitN(ctx, -5)
	if err != nil {
		t.Errorf("expected WaitN(-5) to succeed immediately, got %v", err)
	}

	// Test n > capacity returns error immediately
	err = rl.WaitN(ctx, 3)
	if err == nil {
		t.Error("expected WaitN(3) to return error for capacity 2")
	}
}

// TestWaitEdgeCases tests edge cases in Wait.
func TestWaitEdgeCases(t *testing.T) {
	rl := NewTokenBucket(1, 1)

	// Consume the initial token
	rl.Allow()

	// Test Wait with cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately

	err := rl.Wait(ctx)
	if err == nil {
		t.Error("expected Wait to return error for cancelled context")
	}
}

// TestSetCapacityWithCurrentTokens tests capacity changes when tokens > new capacity.
func TestSetCapacityWithCurrentTokens(t *testing.T) {
	rl := NewTokenBucket(5, 10) // start with capacity 5

	// Wait for refill to ensure we have tokens
	time.Sleep(100 * time.Millisecond)

	// Reduce capacity below current tokens
	rl.SetCapacity(2)

	// Should cap tokens at new capacity
	if rl.Tokens() > 2.0 {
		t.Errorf("expected tokens <= 2 after capacity reduction, got %v", rl.Tokens())
	}
}

// TestConcurrentAccess tests thread safety.
func TestConcurrentAccess(t *testing.T) {
	rl := NewTokenBucket(10, 5)

	done := make(chan bool, 2)

	// Goroutine 1: consuming tokens
	go func() {
		for i := 0; i < 20; i++ {
			rl.Allow()
			time.Sleep(10 * time.Millisecond)
		}
		done <- true
	}()

	// Goroutine 2: checking tokens and changing settings
	go func() {
		for i := 0; i < 10; i++ {
			rl.Tokens()
			rl.SetCapacity(15)
			rl.SetRefillRate(10)
			time.Sleep(20 * time.Millisecond)
		}
		done <- true
	}()

	// Wait for both goroutines
	<-done
	<-done

	// Should not panic or race
}

// TestAllowAfterRefill ensures tokens are refilled correctly after time passes.
func TestAllowAfterRefill(t *testing.T) {
	rl := NewTokenBucket(2, 1)
	rl.AllowN(2)
	time.Sleep(1100 * time.Millisecond)
	if !rl.Allow() {
		t.Error("expected Allow() to succeed after refill")
	}
}

// TestSetCapacityIncreaseAndDecrease tests increasing and decreasing capacity.
func TestSetCapacityIncreaseAndDecrease(t *testing.T) {
	rl := NewTokenBucket(2, 1)
	rl.SetCapacity(5)
	if rl.capacity != 5 {
		t.Errorf("expected capacity to be 5, got %v", rl.capacity)
	}
	rl.SetCapacity(1)
	if rl.capacity != 1 {
		t.Errorf("expected capacity to be 1, got %v", rl.capacity)
	}
}

// TestSetRefillRateIncreaseAndDecrease tests increasing and decreasing refill rate.
func TestSetRefillRateIncreaseAndDecrease(t *testing.T) {
	rl := NewTokenBucket(2, 1)
	rl.SetRefillRate(5)
	if rl.refillRate != 5 {
		t.Errorf("expected refillRate to be 5, got %v", rl.refillRate)
	}
	rl.SetRefillRate(0.1)
	if rl.refillRate != 0.1 {
		t.Errorf("expected refillRate to be 0.1, got %v", rl.refillRate)
	}
}

// TestWaitNWithExactTokens ensures WaitN succeeds immediately if enough tokens are present.
func TestWaitNWithExactTokens(t *testing.T) {
	rl := NewTokenBucket(3, 1)
	ctx := context.Background()
	err := rl.WaitN(ctx, 3)
	if err != nil {
		t.Errorf("expected WaitN to succeed with exact tokens, got %v", err)
	}
}

// TestWaitNImpossibleRequest ensures WaitN returns error if n > capacity.
func TestWaitNImpossibleRequest(t *testing.T) {
	rl := NewTokenBucket(2, 1)
	ctx := context.Background()
	err := rl.WaitN(ctx, 5)
	if err == nil {
		t.Error("expected WaitN to fail for n > capacity")
	}
}

// TestTokensNeverNegative ensures tokens never go negative.
func TestTokensNeverNegative(t *testing.T) {
	rl := NewTokenBucket(1, 1)
	rl.Allow()
	rl.Allow()
	if rl.Tokens() < 0 {
		t.Errorf("tokens should never be negative, got %v", rl.Tokens())
	}
}

// TestFixedWindow_Allow tests the Allow method of the FixedWindow rate limiter.
// It verifies that requests are allowed up to the limit, blocked after the limit,
// and allowed again after the interval resets. It also checks default values for invalid parameters.
func TestFixedWindow_Allow(t *testing.T) {
	tests := []struct {
		name     string
		limit    int
		interval time.Duration
		actions  []struct {
			sleep time.Duration
			want  bool
		}
	}{
		{
			name:     "allow up to limit, block after, reset after interval",
			limit:    2,
			interval: 50 * time.Millisecond,
			actions: []struct {
				sleep time.Duration
				want  bool
			}{
				{0, true},
				{0, true},
				{0, false},
				{60 * time.Millisecond, true},
			},
		},
		{
			name:     "default values for invalid parameters",
			limit:    0,
			interval: 0,
			actions: []struct {
				sleep time.Duration
				want  bool
			}{
				{0, true},
				{0, false},
				{1100 * time.Millisecond, true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter := NewFixedWindow(tt.limit, tt.interval)
			for i, act := range tt.actions {
				if act.sleep > 0 {
					time.Sleep(act.sleep)
				}
				got := limiter.Allow()
				if got != act.want {
					t.Errorf("action %d: Allow() = %v, want %v", i, got, act.want)
				}
			}
		})
	}
}

// TestFixedWindow_Setters tests the SetInterval and SetLimit methods of the FixedWindow rate limiter.
// It verifies that changing the interval or limit updates the limiter's behavior as expected,
// including resetting the window and increasing or decreasing the available requests.
func TestFixedWindow_Setters(t *testing.T) {
	type step struct {
		action string
		value  interface{}
		sleep  time.Duration
		want   bool
	}
	tests := []struct {
		name  string
		steps []step
	}{
		{
			name: "SetInterval resets window",
			steps: []step{
				{"allow", nil, 0, true},
				{"allow", nil, 0, false},
				{"sleep", nil, 25 * time.Millisecond, false}, // Wait for window to expire
				{"setInterval", 5 * time.Millisecond, 0, false},
				{"allow", nil, 0, true}, // Window should be reset since it expired
			},
		},
		{
			name: "SetLimit increases available requests",
			steps: []step{
				{"allow", nil, 0, true},   // count=1, limit=1
				{"allow", nil, 0, false},  // count=1, limit=1 (blocked)
				{"setLimit", 3, 0, false}, // count=1, limit=3
				{"allow", nil, 0, true},   // count=2, limit=3 (allowed)
				{"allow", nil, 0, true},   // count=3, limit=3 (allowed)
				{"allow", nil, 0, false},  // count=3, limit=3 (blocked)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter := NewFixedWindow(1, 20*time.Millisecond)
			for i, s := range tt.steps {
				switch s.action {
				case "allow":
					got := limiter.Allow()
					if got != s.want {
						t.Errorf("step %d: Allow() = %v, want %v", i, got, s.want)
					}
				case "setInterval":
					limiter.SetInterval(s.value.(time.Duration))
				case "setLimit":
					limiter.SetLimit(s.value.(int))
				case "sleep":
					time.Sleep(s.sleep)
				}
			}
		})
	}
}

// TestNewFixedWindow tests the NewFixedWindow constructor for the FixedWindow rate limiter.
// It checks that valid and invalid parameters are handled correctly, including default values,
// and that the initial state of the limiter is as expected.
func TestNewFixedWindow(t *testing.T) {
	tests := []struct {
		name             string
		limit            int
		interval         time.Duration
		expectedLimit    int
		expectedInterval time.Duration
	}{
		{
			name:             "valid parameters",
			limit:            10,
			interval:         5 * time.Second,
			expectedLimit:    10,
			expectedInterval: 5 * time.Second,
		},
		{
			name:             "limit less than 1 defaults to 1",
			limit:            0,
			interval:         5 * time.Second,
			expectedLimit:    1,
			expectedInterval: 5 * time.Second,
		},
		{
			name:             "negative limit defaults to 1",
			limit:            -5,
			interval:         5 * time.Second,
			expectedLimit:    1,
			expectedInterval: 5 * time.Second,
		},
		{
			name:             "zero interval defaults to 1 second",
			limit:            10,
			interval:         0,
			expectedLimit:    10,
			expectedInterval: 1 * time.Second,
		},
		{
			name:             "negative interval defaults to 1 second",
			limit:            10,
			interval:         -5 * time.Second,
			expectedLimit:    10,
			expectedInterval: 1 * time.Second,
		},
		{
			name:             "both invalid parameters use defaults",
			limit:            0,
			interval:         0,
			expectedLimit:    1,
			expectedInterval: 1 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fw := NewFixedWindow(tt.limit, tt.interval)

			if fw.limit != tt.expectedLimit {
				t.Errorf("NewFixedWindow() limit = %v, want %v", fw.limit, tt.expectedLimit)
			}
			if fw.interval != tt.expectedInterval {
				t.Errorf("NewFixedWindow() interval = %v, want %v", fw.interval, tt.expectedInterval)
			}
			if fw.count != 0 {
				t.Errorf("NewFixedWindow() count = %v, want 0", fw.count)
			}
			// windowEndTime should be approximately now + interval
			now := time.Now()
			expectedEnd := now.Add(tt.expectedInterval)
			if fw.windowEndTime.Before(now) || fw.windowEndTime.After(expectedEnd.Add(10*time.Millisecond)) {
				t.Errorf("NewFixedWindow() windowEndTime is not properly set")
			}
		})
	}
}

// TestFixedWindow_Allow_WindowReset tests that the FixedWindow rate limiter resets its window after the interval.
// It verifies that requests are allowed up to the limit, blocked after the limit, and allowed again after the window resets.
func TestFixedWindow_Allow_WindowReset(t *testing.T) {
	limit := 3
	interval := 50 * time.Millisecond
	fw := NewFixedWindow(limit, interval)

	// Use up the limit
	for i := 0; i < limit; i++ {
		if !fw.Allow() {
			t.Errorf("Allow() should return true for request %d", i+1)
		}
	}

	// Next request should be denied
	if fw.Allow() {
		t.Error("Allow() should return false when limit exceeded")
	}

	// Wait for window to reset
	time.Sleep(interval + 10*time.Millisecond)

	// Should be able to make requests again
	for i := 0; i < limit; i++ {
		if !fw.Allow() {
			t.Errorf("Allow() should return true after window reset for request %d", i+1)
		}
	}

	// And should be denied again
	if fw.Allow() {
		t.Error("Allow() should return false when limit exceeded after reset")
	}
}

// TestFixedWindow_Allow_ConcurrentAccess tests the FixedWindow rate limiter under concurrent access.
// It verifies that the limiter allows up to the limit and blocks further requests, even when accessed by multiple goroutines.
func TestFixedWindow_Allow_ConcurrentAccess(t *testing.T) {
	limit := 100
	interval := 100 * time.Millisecond
	fw := NewFixedWindow(limit, interval)

	// Use channels to coordinate goroutines
	start := make(chan struct{})
	results := make(chan bool, limit*2)

	// Start multiple goroutines that will try to acquire tokens
	for i := 0; i < limit*2; i++ {
		go func() {
			<-start // Wait for signal to start
			results <- fw.Allow()
		}()
	}

	// Signal all goroutines to start
	close(start)

	// Collect results
	allowed := 0
	denied := 0
	for i := 0; i < limit*2; i++ {
		if <-results {
			allowed++
		} else {
			denied++
		}
	}

	// Should have exactly 'limit' allowed and 'limit' denied
	if allowed != limit {
		t.Errorf("Expected %d allowed requests, got %d", limit, allowed)
	}
	if denied != limit {
		t.Errorf("Expected %d denied requests, got %d", limit, denied)
	}
}

// TestFixedWindow_SetInterval tests the SetInterval method of the FixedWindow rate limiter.
// It verifies that changing the interval updates the limiter's behavior and resets the window if expired.
// It also checks that invalid intervals default to 1 second.
func TestFixedWindow_SetInterval(t *testing.T) {
	fw := NewFixedWindow(2, 100*time.Millisecond)

	// Use up the limit
	if !fw.Allow() {
		t.Fatal("First request should be allowed")
	}
	if !fw.Allow() {
		t.Fatal("Second request should be allowed")
	}
	if fw.Allow() {
		t.Fatal("Should be rate limited")
	}

	// Wait for window to expire first
	time.Sleep(110 * time.Millisecond)

	// Change interval to a shorter duration - this will reset the window since it expired
	newInterval := 20 * time.Millisecond
	fw.SetInterval(newInterval)

	// Should be allowed now since window was reset
	if !fw.Allow() {
		t.Error("Should be allowed after SetInterval when window expired")
	}

	// Test with invalid interval
	fw.SetInterval(-5 * time.Second)
	if fw.interval != 1*time.Second {
		t.Errorf("Invalid interval should default to 1 second, got %v", fw.interval)
	}
}

// TestFixedWindow_SetInterval_WindowReset tests that SetInterval resets the window if it has expired.
// It verifies that after changing the interval, the limiter allows requests up to the new limit.
func TestFixedWindow_SetInterval_WindowReset(t *testing.T) {
	fw := NewFixedWindow(2, 50*time.Millisecond) // Shorter initial interval

	// Use some of the limit
	fw.Allow()

	// Wait for window to expire
	time.Sleep(60 * time.Millisecond)

	// Change interval - this should reset the window since it expired
	fw.SetInterval(100 * time.Millisecond)

	// Count should be reset, so we should get the full limit again
	if !fw.Allow() {
		t.Error("First request should be allowed after window reset")
	}
	if !fw.Allow() {
		t.Error("Second request should be allowed after window reset")
	}
	if fw.Allow() {
		t.Error("Should be rate limited after using new limit")
	}
}

// TestFixedWindow_SetLimit tests the SetLimit method of the FixedWindow rate limiter.
// It verifies that increasing the limit allows more requests and that invalid limits default to 1.
func TestFixedWindow_SetLimit(t *testing.T) {
	fw := NewFixedWindow(2, 100*time.Millisecond)

	// Use up the initial limit
	if !fw.Allow() {
		t.Fatal("First request should be allowed")
	}
	if !fw.Allow() {
		t.Fatal("Second request should be allowed")
	}
	if fw.Allow() {
		t.Fatal("Should be rate limited")
	}

	// Increase the limit
	fw.SetLimit(4)

	// Should be able to make more requests now (2 more to reach new limit of 4)
	if !fw.Allow() {
		t.Error("Third request should be allowed after increasing limit")
	}
	if !fw.Allow() {
		t.Error("Fourth request should be allowed after increasing limit")
	}
	if fw.Allow() {
		t.Error("Should be rate limited after reaching new limit")
	}

	// Test with invalid limit
	fw.SetLimit(-5)
	if fw.limit != 1 {
		t.Errorf("Invalid limit should default to 1, got %d", fw.limit)
	}
}

// TestFixedWindow_SetLimit_WindowReset tests that SetLimit resets the window if it has expired.
// It verifies that increasing the limit allows more requests within the same window.
func TestFixedWindow_SetLimit_WindowReset(t *testing.T) {
	fw := NewFixedWindow(3, 500*time.Millisecond)

	// Use some of the limit
	fw.Allow()
	fw.Allow()

	// Wait a bit but not enough for window to reset naturally
	time.Sleep(50 * time.Millisecond)

	// Change limit - this should reset the window if window has expired
	fw.SetLimit(5)

	// In this case, window hasn't expired, so count should remain
	// We should be able to make 1 more request (used 2, limit now 5)
	if !fw.Allow() {
		t.Error("Should be allowed within new higher limit")
	}
}

// TestFixedWindow_SetLimit_LowerLimit tests lowering the limit below the current count.
// It verifies that the limiter blocks requests when the count exceeds the new lower limit,
// and allows requests again after the window resets.
func TestFixedWindow_SetLimit_LowerLimit(t *testing.T) {
	fw := NewFixedWindow(5, 100*time.Millisecond)

	// Use some of the limit
	fw.Allow()
	fw.Allow()
	fw.Allow()

	// Lower the limit to below current count
	fw.SetLimit(2)

	// Should be rate limited since current count (3) > new limit (2)
	if fw.Allow() {
		t.Error("Should be rate limited when current count exceeds new limit")
	}

	// Wait for window to reset
	time.Sleep(110 * time.Millisecond)

	// Should work with new limit now
	if !fw.Allow() {
		t.Error("First request should be allowed after window reset with new limit")
	}
	if !fw.Allow() {
		t.Error("Second request should be allowed after window reset with new limit")
	}
	if fw.Allow() {
		t.Error("Should be rate limited at new limit")
	}
}

// TestFixedWindow_EdgeCases tests various edge cases for the FixedWindow rate limiter.
// It checks zero and negative limits and intervals, as well as rapid successive calls.
func TestFixedWindow_EdgeCases(t *testing.T) {
	t.Run("zero and negative limits", func(t *testing.T) {
		fw1 := NewFixedWindow(0, time.Second)
		fw2 := NewFixedWindow(-1, time.Second)

		// Both should have limit of 1
		if !fw1.Allow() {
			t.Error("Zero limit should default to 1")
		}
		if fw1.Allow() {
			t.Error("Should be rate limited after first request")
		}

		if !fw2.Allow() {
			t.Error("Negative limit should default to 1")
		}
		if fw2.Allow() {
			t.Error("Should be rate limited after first request")
		}
	})

	t.Run("zero and negative intervals", func(t *testing.T) {
		fw1 := NewFixedWindow(1, 0)
		fw2 := NewFixedWindow(1, -time.Second)

		// Both should have interval of 1 second
		if fw1.interval != time.Second {
			t.Errorf("Zero interval should default to 1 second, got %v", fw1.interval)
		}
		if fw2.interval != time.Second {
			t.Errorf("Negative interval should default to 1 second, got %v", fw2.interval)
		}
	})

	t.Run("rapid successive calls", func(t *testing.T) {
		fw := NewFixedWindow(1000, 100*time.Millisecond)

		// Make rapid successive calls
		allowed := 0
		for i := 0; i < 1500; i++ {
			if fw.Allow() {
				allowed++
			}
		}

		// Should have allowed exactly 1000
		if allowed != 1000 {
			t.Errorf("Expected 1000 allowed requests, got %d", allowed)
		}
	})
}
