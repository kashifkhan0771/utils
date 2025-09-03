# Queue

This package provides a generic, thread-safe FIFO queue implementation for Go, supporting any type with automatic dynamic resizing.

## Features

- **Generic**: Works with any type using Go generics (`Queue[T any]`)
- **Thread-Safe**: All operations are protected by mutex for concurrent access
- **Dynamic Resizing**: Automatic growth (2x) and shrinking (0.5x) based on usage
- **Circular Buffer**: Efficient O(1) enqueue/dequeue operations with minimal memory overhead
- **Memory Efficient**: Zero-value clearing prevents memory leaks
- **FIFO Semantics**: First-In-First-Out queue behavior
- **Error Handling**: Proper error types for empty queue operations

## API

- **NewQueue**: Creates a new queue with specified initial capacity
- **Enqueue**: Adds an item to the end of the queue
- **Dequeue**: Removes and returns the front item (returns error if empty)
- **Peek**: Returns the front item without removing it (returns error if empty)
- **Size**: Returns the current number of elements in the queue
- **Capacity**: Returns the queue's current capacity
- **IsEmpty**: Returns true if the queue contains no elements

## Performance

- **Enqueue**: ~28ns/op with minimal allocations
- **Dequeue**: ~11ns/op with minimal allocations  
- **Peek**: ~5ns/op with zero allocations
- **Memory**: Minimal heap allocations during normal operations

## Examples

For comprehensive examples of each function, please check out [EXAMPLES.md](EXAMPLES.md)

---
