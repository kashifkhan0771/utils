// Package caching provides utilities for creating caching decorators to
// enhance the performance of functions by storing computed results.
// It includes both thread-safe and non-thread-safe implementations.

package caching

import "sync"

// CacheWrapper is a non-thread-safe caching decorator.
func CacheWrapper[T comparable, R any](fn func(T) R) func(T) R {
	cache := make(map[T]R)

	return func(input T) R {
		// Check if the result is already cached
		if result, exists := cache[input]; exists {
			return result
		}
		// Call the function and store the result in the cache
		result := fn(input)
		cache[input] = result

		return result
	}
}

// SafeCacheWrapper is a thread-safe caching decorator.
func SafeCacheWrapper[T comparable, R any](fn func(T) R) func(T) R {
	var cache sync.Map

	return func(input T) R {
		// Check if the result is already cached
		if result, exists := cache.Load(input); exists {
			return result.(R) // Type-safe due to generics
		}
		// Call the function and store the result in the cache
		result := fn(input)
		cache.Store(input, result)

		return result
	}
}
