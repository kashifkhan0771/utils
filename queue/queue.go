package queue

import (
	"errors"
	"sync"
)

// ErrEmptyQueue is returned when the queue is empty.
var ErrEmptyQueue = errors.New("queue is empty")

const (
	// DefaultCapacity is the initial capacity for a new queue.
	DefaultCapacity = 16
	MinCapacity     = 16
)

// Queue is a generic, thread-safe FIFO queue implementation.
// It uses a circular buffer internally and automatically grows or shrinks as needed.
//
// Type Parameters:
//
//	T: The type of elements stored in the queue.
type Queue[T any] struct {
	data []T        // underlying slice storing queue elements
	head int        // index of the front element
	tail int        // index for the next enqueue
	size int        // current number of elements in the queue
	mu   sync.Mutex // mutex to ensure thread safety
}

// NewQueue creates a new Queue with the given capacity.
func NewQueue[T any](capacity int) *Queue[T] {
	if capacity <= 0 {
		capacity = DefaultCapacity
	}

	return &Queue[T]{
		data: make([]T, capacity),
	}
}

// Enqueue adds an item to the end of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.data) == q.size {
		q.grow()
	}

	q.data[q.tail] = item
	q.tail = (q.tail + 1) % len(q.data)
	q.size++
}

// Dequeue removes and returns the front item.
// Returns ErrEmptyQueue if the queue is empty.
func (q *Queue[T]) Dequeue() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.size == 0 {
		var zero T
		return zero, ErrEmptyQueue
	}

	element := q.data[q.head]
	var zero T
	q.data[q.head] = zero
	q.head = (q.head + 1) % len(q.data)
	q.size--

	// Shrink if much less than a quarter full and capacity > 32.
	if q.size > 0 && q.size*4 < len(q.data) && len(q.data) > 32 {
		q.shrink()
	}

	return element, nil
}

// Peek returns the front item without removing it.
// Returns ErrEmptyQueue if the queue is empty.
func (q *Queue[T]) Peek() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.size == 0 {
		var zero T
		return zero, ErrEmptyQueue
	}

	return q.data[q.head], nil
}

// Size returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.size
}

// Capacity returns the queue's capacity.
func (q *Queue[T]) Capacity() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.data)
}

// IsEmpty returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.size == 0
}
