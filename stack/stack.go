package stack

import "sync"

// Stack is a generic thread-safe stack data structure.
type Stack[T any] struct {
	data []T
	mu   sync.RWMutex
}

// New creates a new Stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, value)
}

// Pop removes and returns the top element of the stack.
// Returns the zero value and false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.data) == 0 {
		var zero T

		return zero, false
	}

	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return val, true
}

// Peek returns the element at position n from the top without removing it.
// n=0 returns the top element, n=1 returns the second element from top, etc.
// Returns the zero value and false if the index is out of bounds.
func (s *Stack[T]) Peek(n int) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if (len(s.data) - n - 1) < 0 {
		var zero T

		return zero, false
	}

	return s.data[len(s.data)-n-1], true
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data)
}
