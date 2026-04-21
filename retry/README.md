### Retry

The `retry` package provides a generic, context-aware way to retry fallible operations with configurable backoff strategies and timeout handling.

#### **Options**

- **`MaxAttempts int`**: Total number of times the operation will be tried before giving up.
- **`TotalTimeout time.Duration`**: Maximum time allowed across all attempts, including backoff delays.
- **`Backoff func(attempt int) time.Duration`**: Returns how long to wait before the next attempt. `attempt` is zero-indexed.
- **`ShouldRetry func(err error) bool`**: Reports whether the given error is retryable. Return `false` to abort immediately.

#### **Functions**

- **`Do[T any](ctx context.Context, opts Options, fn RetryFunc[T]) (T, error)`**:  
  Calls `fn` repeatedly until it succeeds, `ShouldRetry` returns `false`, `MaxAttempts` is reached, or `TotalTimeout` elapses.

- **`DoVoid(ctx context.Context, opts Options, fn func(ctx context.Context) error) error`**:  
  Convenience wrapper around `Do` for operations that return no value.

#### **Backoff Strategies**

- **`FixedBackoff(d time.Duration) func(attempt int) time.Duration`**:  
  Waits exactly `d` between every attempt.

- **`LinearBackoff(d time.Duration) func(attempt int) time.Duration`**:  
  Waits `d * attempt`. Grows linearly: `0, d, 2d, 3d, …`

- **`ExponentialBackoff(d time.Duration) func(attempt int) time.Duration`**:  
  Waits `d * 2^attempt`. Doubles on each failure: `d, 2d, 4d, 8d, …`

#### **Notes**
- `TotalTimeout` is enforced via a derived context passed to every attempt. If the deadline is exceeded mid-backoff, `Do` returns `context.DeadlineExceeded` immediately.
- A non-retryable error returned from `ShouldRetry` is returned as-is, without wrapping.
- `LinearBackoff` produces a zero-length first pause (`attempt 0`). Prefer `FixedBackoff` if an immediate first retry is undesirable.

## Examples:
For examples of each function, please check out [EXAMPLES.md](/retry/EXAMPLES.md)

---
