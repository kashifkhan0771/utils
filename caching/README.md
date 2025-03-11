### Caching

The caching package provides utilities for creating caching decorators to enhance the performance of functions by storing computed results. It includes both thread-safe and non-thread-safe implementations.

- **SafeCacheWrapper**: A thread-safe caching decorator that safely memoizes function results in concurrent environments.

  - Uses `sync.Map` to ensure thread-safety
  - Caches all results indefinitely (no eviction)
  - Best suited for pure functions with limited input domains
  - Safe for concurrent access but may impact performance under high contention

- **CacheWrapper**: A non-thread-safe caching decorator that memoizes function results.
  - Caches all results indefinitely (no eviction)
  - Best suited for pure functions with limited input domains
  - Not safe for concurrent access
  - Use SafeCacheWrapper for concurrent scenarios

## Examples:

For examples of each function, please checkout [EXAMPLES.md](/caching/EXAMPLES.md)

---
