### RateLimiter

The `ratelimiter` package provides utilities to control the rate of operations by limiting how frequently actions can be performed. It implements both **Token Bucket** and **Fixed Window** rate limiters that are safe for concurrent use.

- **TokenBucket**: A thread-safe token bucket rate limiter.

  - Controls operation rate by allowing a fixed number of tokens (operations) per time unit
  - Supports burst capacity to handle short bursts of traffic
  - Provides non-blocking (`Allow()`) and blocking (`Wait()`) methods
  - `Wait()` supports context cancellation for graceful timeout or abort
  - Allows dynamic adjustment of capacity and refill rate at runtime
  - No external dependencies, lightweight and efficient

- **FixedWindow**: A thread-safe fixed window rate limiter.

  - Allows up to a specified number of operations per fixed time window
  - Window resets after each interval period
  - Provides non-blocking (`Allow()`) method for immediate rate limiting decisions
  - Allows dynamic adjustment of limit and interval at runtime
  - Simple and predictable rate limiting behavior
  - Perfect for scenarios requiring strict rate limits per time period

## Examples:
For examples of each function, please checkout [EXAMPLES.md](/ratelimiter/EXAMPLES.md)

---
