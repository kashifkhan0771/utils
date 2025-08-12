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

	time.Sleep(600 * time.Millisecond) // ~1.2 tokens refill

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

	if elapsed < 900*time.Millisecond {
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
