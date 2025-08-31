package queue

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

// Test helpers
func assertQueueEmpty(t *testing.T, q *Queue[int]) {
	t.Helper()
	if !q.IsEmpty() {
		t.Errorf("Queue should be empty, but size is %d", q.Size())
	}
	if q.Size() != 0 {
		t.Errorf("Queue size should be 0, got %d", q.Size())
	}
}

func assertQueueSize(t *testing.T, q *Queue[int], expected int) {
	t.Helper()
	if size := q.Size(); size != expected {
		t.Errorf("Expected queue size %d, got %d", expected, size)
	}
}

func enqueueItems(q *Queue[int], items []int) {
	for _, item := range items {
		q.Enqueue(item)
	}
}

func dequeueAndVerify(t *testing.T, q *Queue[int], expected []int) {
	t.Helper()
	for i, exp := range expected {
		item, err := q.Dequeue()
		if err != nil {
			t.Fatalf("Unexpected error at index %d: %v", i, err)
		}
		if item != exp {
			t.Errorf("At index %d, expected %d, got %d", i, exp, item)
		}
	}
}

func TestNewQueue(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		expected int
	}{
		{"positive capacity", 10, 10},
		{"zero capacity", 0, 16},
		{"negative capacity", -5, 16},
		{"large capacity", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue[int](tt.capacity)
			if q == nil {
				t.Fatal("NewQueue returned nil")
			}
			if cap := q.Capacity(); cap != tt.expected {
				t.Errorf("Expected capacity %d, got %d", tt.expected, cap)
			}
			if size := q.Size(); size != 0 {
				t.Errorf("Expected size 0 for new queue, got %d", size)
			}
			if !q.IsEmpty() {
				t.Error("New queue should be empty")
			}
		})
	}
}

func TestEnqueueDequeue(t *testing.T) {
	tests := []struct {
		name   string
		items  []int
		expect []int
	}{
		{"single item", []int{1}, []int{1}},
		{"multiple items", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"empty queue", []int{}, []int{}},
		{"duplicate items", []int{1, 1, 1}, []int{1, 1, 1}},
		{"large sequence", make([]int, 100), make([]int, 100)},
	}

	// Initialize large sequence test case
	for i := range tests[4].items {
		tests[4].items[i] = i
		tests[4].expect[i] = i
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue[int](4) // Small initial capacity to test growth

			// Enqueue all items
			enqueueItems(q, tt.items)

			assertQueueSize(t, q, len(tt.expect))

			// Dequeue and verify order
			dequeueAndVerify(t, q, tt.expect)

			assertQueueEmpty(t, q)
		})
	}
}

func TestDequeueEmpty(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
	}{
		{"default capacity", 0},
		{"small capacity", 4},
		{"large capacity", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue[int](tt.capacity)

			// Test dequeue on empty queue
			item, err := q.Dequeue()
			if err != ErrEmptyQueue {
				t.Errorf("Expected ErrEmptyQueue, got %v", err)
			}
			if item != 0 {
				t.Errorf("Expected zero value, got %d", item)
			}

			// Verify error message
			if err != nil && err.Error() != "queue is empty" {
				t.Errorf("Expected error message 'queue is empty', got '%s'", err.Error())
			}

			// Test after enqueue/dequeue cycle
			q.Enqueue(42)
			_, _ = q.Dequeue()

			_, err = q.Dequeue()
			if err != ErrEmptyQueue {
				t.Errorf("Expected ErrEmptyQueue after emptying queue, got %v", err)
			}
		})
	}
}

func TestQueueGrowth(t *testing.T) {
	tests := []struct {
		name         string
		initialCap   int
		itemsToAdd   int
		expectedSize int
	}{
		{"grow from default", 4, 10, 10},
		{"grow from small", 2, 20, 20},
		{"no growth needed", 20, 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue[int](tt.initialCap)
			initialCapacity := q.Capacity()

			// Add items to trigger growth
			for i := 0; i < tt.itemsToAdd; i++ {
				q.Enqueue(i)
			}

			if size := q.Size(); size != tt.expectedSize {
				t.Errorf("Expected size %d, got %d", tt.expectedSize, size)
			}

			if tt.itemsToAdd > initialCapacity {
				if newCap := q.Capacity(); newCap <= initialCapacity {
					t.Errorf("Queue should have grown. Initial: %d, Current: %d", initialCapacity, newCap)
				}
			}

			// Verify all items are still accessible
			for i := 0; i < tt.itemsToAdd; i++ {
				item, err := q.Dequeue()
				if err != nil {
					t.Fatalf("Error dequeuing item %d: %v", i, err)
				}
				if item != i {
					t.Errorf("Expected item %d, got %d", i, item)
				}
			}
		})
	}
}

func TestQueueShrinking(t *testing.T) {
	q := NewQueue[int](4)

	// Fill queue to force growth
	for i := 0; i < 40; i++ {
		q.Enqueue(i)
	}

	largeCapacity := q.Capacity()
	if largeCapacity <= 40 {
		t.Errorf("Queue should have grown beyond 40, got %d", largeCapacity)
	}

	// Remove most items to trigger shrinking
	for i := 0; i < 35; i++ {
		_, _ = q.Dequeue()
	}

	// After shrinking, capacity should be smaller
	newCapacity := q.Capacity()
	if newCapacity >= largeCapacity {
		t.Errorf("Queue should have shrunk. Before: %d, After: %d", largeCapacity, newCapacity)
	}
	// Verify remaining items are correct
	for i := 35; i < 40; i++ {
		item, err := q.Dequeue()
		if err != nil {
			t.Fatalf("Error dequeuing remaining item: %v", err)
		}
		if item != i {
			t.Errorf("Expected remaining item %d, got %d", i, item)
		}
	}
}

func TestCircularBehavior(t *testing.T) {
	q := NewQueue[int](4)

	// Test wraparound behavior
	tests := []struct {
		name      string
		operation string
		value     int
		expectErr bool
	}{
		{"enqueue 1", "enqueue", 1, false},
		{"enqueue 2", "enqueue", 2, false},
		{"dequeue first", "dequeue", 1, false},
		{"enqueue 3", "enqueue", 3, false},
		{"enqueue 4", "enqueue", 4, false},
		{"dequeue second", "dequeue", 2, false},
		{"enqueue 5", "enqueue", 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.operation == "enqueue" {
				q.Enqueue(tt.value)
			} else {
				item, err := q.Dequeue()
				if (err != nil) != tt.expectErr {
					t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
				}
				if !tt.expectErr && item != tt.value {
					t.Errorf("Expected %d, got %d", tt.value, item)
				}
			}
		})
	}
}

func TestConcurrentAccess(t *testing.T) {
	q := NewQueue[int](10)
	const numGoroutines = 10
	const itemsPerGoroutine = 100

	var wg sync.WaitGroup

	// Concurrent enqueuing
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for j := 0; j < itemsPerGoroutine; j++ {
				q.Enqueue(start*itemsPerGoroutine + j)
			}
		}(i)
	}

	wg.Wait()

	expectedSize := numGoroutines * itemsPerGoroutine
	if size := q.Size(); size != expectedSize {
		t.Errorf("Expected size %d after concurrent enqueuing, got %d", expectedSize, size)
	}

	// Concurrent dequeuing
	results := make([]int, expectedSize)
	var resultMu sync.Mutex
	var index int

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < itemsPerGoroutine; j++ {
				item, err := q.Dequeue()
				if err != nil {
					t.Errorf("Unexpected error during concurrent dequeue: %v", err)
					return
				}
				resultMu.Lock()
				results[index] = item
				index++
				resultMu.Unlock()
			}
		}()
	}

	wg.Wait()

	if !q.IsEmpty() {
		t.Errorf("Queue should be empty after concurrent dequeuing, size: %d", q.Size())
	}
}

func TestDifferentTypes(t *testing.T) {
	t.Run("string queue", func(t *testing.T) {
		q := NewQueue[string](4)
		items := []string{"hello", "world", "go", "generics"}

		for _, item := range items {
			q.Enqueue(item)
		}

		for _, expected := range items {
			item, err := q.Dequeue()
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if item != expected {
				t.Errorf("Expected %s, got %s", expected, item)
			}
		}
	})

	t.Run("struct queue", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		q := NewQueue[Person](2)
		people := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}

		for _, person := range people {
			q.Enqueue(person)
		}

		for _, expected := range people {
			person, err := q.Dequeue()
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if person != expected {
				t.Errorf("Expected %+v, got %+v", expected, person)
			}
		}
	})
}

func TestEdgeCases(t *testing.T) {
	t.Run("single capacity queue", func(t *testing.T) {
		q := NewQueue[int](1)

		q.Enqueue(42)
		if size := q.Size(); size != 1 {
			t.Errorf("Expected size 1, got %d", size)
		}

		// This should trigger growth since queue is full
		q.Enqueue(43)
		if size := q.Size(); size != 2 {
			t.Errorf("Expected size 2 after growth, got %d", size)
		}

		item, _ := q.Dequeue()
		if item != 42 {
			t.Errorf("Expected 42, got %d", item)
		}

		item, _ = q.Dequeue()
		if item != 43 {
			t.Errorf("Expected 43, got %d", item)
		}
	})

	t.Run("zero value handling", func(t *testing.T) {
		q := NewQueue[int](4)

		// Enqueue zero values
		q.Enqueue(0)
		q.Enqueue(1)
		q.Enqueue(0)

		values := []int{0, 1, 0}
		for _, expected := range values {
			item, err := q.Dequeue()
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if item != expected {
				t.Errorf("Expected %d, got %d", expected, item)
			}
		}
	})
}

// Test that queue operations don't panic under any conditions
func TestNoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Queue operation caused panic: %v", r)
		}
	}()

	q := NewQueue[int](1)

	// Try various operations that might cause panics
	_, _ = q.Dequeue() // Should return error, not panic
	_, _ = q.Peek()    // Should return error, not panic
	q.Size()           // Should work
	q.Capacity()       // Should work
	q.IsEmpty()        // Should work

	// Operations after enqueue
	q.Enqueue(1)
	_, _ = q.Peek()
	q.Size()
	q.Capacity()
	q.IsEmpty()
	_, _ = q.Dequeue()

	// Multiple dequeues
	_, _ = q.Dequeue()
	_, _ = q.Dequeue()
}

// Fuzz test for queue operations (Go 1.18+)
func FuzzQueueOperations(f *testing.F) {
	// Seed with some initial test cases
	f.Add(int8(1), int8(0)) // enqueue=1, dequeue=0
	f.Add(int8(0), int8(1)) // enqueue=0, dequeue=1
	f.Add(int8(5), int8(3)) // enqueue=5, dequeue=3

	f.Fuzz(func(t *testing.T, enqueueOps, dequeueOps int8) {
		// Limit operations to reasonable range
		if enqueueOps < 0 || enqueueOps > 100 {
			return
		}
		if dequeueOps < 0 || dequeueOps > 100 {
			return
		}

		q := NewQueue[int](4)

		// Perform enqueue operations
		for i := int8(0); i < enqueueOps; i++ {
			q.Enqueue(int(i))
		}

		// Perform dequeue operations
		successfulDequeues := int8(0)
		for i := int8(0); i < dequeueOps; i++ {
			_, err := q.Dequeue()
			if err == nil {
				successfulDequeues++
			} else if err != ErrEmptyQueue {
				t.Fatalf("Unexpected error: %v", err)
			}
		}

		// Verify queue state is consistent
		expectedSize := enqueueOps - successfulDequeues
		if expectedSize < 0 {
			expectedSize = 0
		}

		if q.Size() != int(expectedSize) {
			t.Errorf("Size mismatch: expected %d, got %d", expectedSize, q.Size())
		}

		if (q.Size() == 0) != q.IsEmpty() {
			t.Errorf("IsEmpty() inconsistent with Size(): size=%d, isEmpty=%v", q.Size(), q.IsEmpty())
		}
	})
}

// Memory leak tests
func TestMemoryLeakAfterDequeueAll(t *testing.T) {
	// Test that dequeuing all elements properly clears references to prevent memory leaks
	q := NewQueue[*int](4)
	const numElements = 1000

	// Track initial memory
	var m1 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	// Create and enqueue many pointer elements
	for i := 0; i < numElements; i++ {
		x := new(int)
		*x = i
		q.Enqueue(x)
	}

	// Dequeue all elements
	for i := 0; i < numElements; i++ {
		_, err := q.Dequeue()
		if err != nil {
			t.Fatalf("Unexpected error at %d: %v", i, err)
		}
	}

	// Force garbage collection
	runtime.GC()
	runtime.GC() // Second GC to ensure everything is cleaned up

	// Check final memory
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	// The queue should be empty
	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}

	// Memory usage should not have grown significantly
	// Allow for some variance in memory measurement
	memoryGrowth := m2.HeapAlloc - m1.HeapAlloc
	t.Logf("Memory growth: %d bytes", memoryGrowth)

	// This is a soft check - we expect minimal growth
	if memoryGrowth > 100*1024 { // 100KB threshold
		t.Logf("Warning: Potential memory leak detected. Memory grew by %d bytes", memoryGrowth)
	}
}

// Peek method tests
func TestPeekEmptyQueue(t *testing.T) {
	q := NewQueue[int](4)

	// Test peek on completely empty queue
	value, err := q.Peek()
	if err != ErrEmptyQueue {
		t.Errorf("Expected ErrEmptyQueue, got %v", err)
	}
	if value != 0 {
		t.Errorf("Expected zero value (0), got %d", value)
	}

	// Verify queue state unchanged after peek error
	if !q.IsEmpty() {
		t.Error("Queue should still be empty after failed peek")
	}
	if q.Size() != 0 {
		t.Errorf("Queue size should be 0 after failed peek, got %d", q.Size())
	}
}

func TestPeekDoesNotModifyQueue(t *testing.T) {
	q := NewQueue[int](4)
	items := []int{1, 2, 3, 4, 5}

	// Enqueue all items
	for _, item := range items {
		q.Enqueue(item)
	}

	initialSize := q.Size()
	expectedFront := items[0]

	// Perform multiple peeks
	for i := 0; i < 5; i++ {
		value, err := q.Peek()
		if err != nil {
			t.Fatalf("Peek attempt %d failed: %v", i+1, err)
		}
		if value != expectedFront {
			t.Errorf("Peek attempt %d: expected %d, got %d", i+1, expectedFront, value)
		}

		// Verify queue state unchanged
		if q.Size() != initialSize {
			t.Errorf("After peek %d: size changed from %d to %d", i+1, initialSize, q.Size())
		}
		if q.IsEmpty() {
			t.Errorf("After peek %d: queue became empty", i+1)
		}
	}

	// Verify dequeue still returns correct order after multiple peeks
	for j, expected := range items {
		value, err := q.Dequeue()
		if err != nil {
			t.Fatalf("Dequeue at index %d failed: %v", j, err)
		}
		if value != expected {
			t.Errorf("After peeks, dequeue at index %d: expected %d, got %d", j, expected, value)
		}
	}
}

// Thread-safety tests
func TestConcurrentHeavyLoad(t *testing.T) {
	// Test queue behavior under concurrent load
	q := NewQueue[int](10)
	const numGoroutines = 50
	const operationsPerGoroutine = 1000

	var wg sync.WaitGroup
	var enqueueOps, dequeueOps, peekOps int64
	var errors int64

	// Concurrent enqueuing
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			base := workerID * operationsPerGoroutine
			for j := 0; j < operationsPerGoroutine; j++ {
				q.Enqueue(base + j)
				atomic.AddInt64(&enqueueOps, 1)
			}
		}(i)
	}

	// Concurrent dequeuing
	for i := 0; i < numGoroutines/2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				_, err := q.Dequeue()
				if err == ErrEmptyQueue {
					// Expected when queue is empty
				} else if err != nil {
					t.Errorf("Unexpected error during dequeue: %v", err)
					atomic.AddInt64(&errors, 1)
				} else {
					atomic.AddInt64(&dequeueOps, 1)
				}
			}
		}()
	}

	// Concurrent peeking
	for i := 0; i < numGoroutines/4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				_, err := q.Peek()
				if err == ErrEmptyQueue {
					// Expected when queue is empty
				} else if err != nil {
					t.Errorf("Unexpected error during peek: %v", err)
					atomic.AddInt64(&errors, 1)
				} else {
					atomic.AddInt64(&peekOps, 1)
				}
			}
		}()
	}

	wg.Wait()

	// Verify no unexpected errors occurred
	finalErrors := atomic.LoadInt64(&errors)
	if finalErrors > 0 {
		t.Errorf("Encountered %d unexpected errors during concurrent operations", finalErrors)
	}

	// Verify queue state is consistent
	finalSize := q.Size()
	finalEnqueueOps := atomic.LoadInt64(&enqueueOps)
	finalDequeueOps := atomic.LoadInt64(&dequeueOps)
	expectedMinSize := int(finalEnqueueOps - finalDequeueOps)
	if finalSize < 0 || (expectedMinSize > 0 && finalSize < expectedMinSize) {
		t.Errorf("Queue size inconsistent: got %d, enqueued %d, dequeued %d",
			finalSize, finalEnqueueOps, finalDequeueOps)
	}

	finalPeekOps := atomic.LoadInt64(&peekOps)
	t.Logf("Operations completed - Enqueued: %d, Dequeued: %d, Peeked: %d, Final size: %d",
		finalEnqueueOps, finalDequeueOps, finalPeekOps, finalSize)
}

// BenchmarkEnqueue measures the performance of enqueuing items into the queue.
func BenchmarkEnqueue(b *testing.B) {
	q := NewQueue[int](16)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	q := NewQueue[int](b.N)

	// Pre-fill queue
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = q.Dequeue()
	}
}

func BenchmarkEnqueueDequeue(b *testing.B) {
	q := NewQueue[int](16)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
		_, _ = q.Dequeue()
	}
}

func BenchmarkPeek(b *testing.B) {
	q := NewQueue[int](16)
	q.Enqueue(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = q.Peek()
	}
}

func BenchmarkMixedOperations(b *testing.B) {
	q := NewQueue[int](16)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			q.Enqueue(i)
		} else if i%3 == 1 && !q.IsEmpty() {
			_, _ = q.Dequeue()
		} else {
			_, _ = q.Peek()
		}
	}
}
