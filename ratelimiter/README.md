### RateLimiter

The `ratelimiter` package provides utilities to control the rate of operations by limiting how frequently actions can be performed. It implements a **Token Bucket** rate limiter that is safe for concurrent use.

- **TokenBucket**: A thread-safe token bucket rate limiter.

  - Controls operation rate by allowing a fixed number of tokens (operations) per time unit
  - Supports burst capacity to handle short bursts of traffic
  - Provides non-blocking (`Allow()`) and blocking (`Wait()`) methods
  - `Wait()` supports context cancellation for graceful timeout or abort
  - Allows dynamic adjustment of capacity and refill rate at runtime
  - No external dependencies, lightweight and efficient

## Examples:
For examples of each function, please checkout [EXAMPLES.md](/ratelimiter/EXAMPLES.md)

---
