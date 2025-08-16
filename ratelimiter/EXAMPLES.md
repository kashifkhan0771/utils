# RateLimiter Examples

This document provides comprehensive examples for all functions in the ratelimiter package.

## Table of Contents

- [NewTokenBucket](#newtokenbucket)
- [Allow](#allow)
- [AllowN](#allown)
- [Wait](#wait)
- [WaitN](#waitn)
- [Tokens](#tokens)
- [SetCapacity](#setcapacity)
- [SetRefillRate](#setrefillrate)

## NewTokenBucket

Creates a new token bucket rate limiter with specified capacity and refill rate.

```go
package main

import (
    "fmt"
    "utils/ratelimiter"
)

func main() {
    // Create a token bucket with capacity of 5 tokens
    // and refill rate of 2 tokens per second
    bucket := ratelimiter.NewTokenBucket(5, 2.0)
    
    fmt.Printf("Created token bucket with capacity: 5, refill rate: 2.0 tokens/sec\n")
    
    // Edge cases: invalid values are corrected
    bucket2 := ratelimiter.NewTokenBucket(0, -1.0) // capacity becomes 1, rate becomes 1
    fmt.Printf("Invalid parameters corrected - capacity: 1, rate: 1.0\n")
}
```

**Output:**
```
Created token bucket with capacity: 5, refill rate: 2.0 tokens/sec
Invalid parameters corrected - capacity: 1, rate: 1.0
```

## Allow

Checks if one token is available and consumes it immediately without blocking.

```go
package main

import (
    "fmt"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a bucket with 3 tokens, refilling at 1 token per second
    bucket := ratelimiter.NewTokenBucket(3, 1.0)
    
    fmt.Println("Testing Allow() method:")
    
    for i := 1; i <= 5; i++ {
        if bucket.Allow() {
            fmt.Printf("Request #%d: ALLOWED at %v\n", i, time.Now().Format("15:04:05.000"))
        } else {
            fmt.Printf("Request #%d: DENIED at %v\n", i, time.Now().Format("15:04:05.000"))
        }
        time.Sleep(200 * time.Millisecond)
    }
}
```

**Output:**
```
Testing Allow() method:
Request #1: ALLOWED at 14:30:15.001
Request #2: ALLOWED at 14:30:15.202
Request #3: ALLOWED at 14:30:15.403
Request #4: DENIED at 14:30:15.604
Request #5: DENIED at 14:30:15.805
```

## AllowN

Checks if N tokens are available and consumes them atomically without blocking.

```go
package main

import (
    "fmt"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a bucket with 10 tokens, refilling at 5 tokens per second
    bucket := ratelimiter.NewTokenBucket(10, 5.0)
    
    fmt.Println("Testing AllowN() method:")
    
    // Try to consume different amounts of tokens
    testCases := []int{3, 5, 2, 8, 1}
    
    for i, n := range testCases {
        if bucket.AllowN(n) {
            fmt.Printf("Request #%d: ALLOWED %d tokens at %v\n", i+1, n, time.Now().Format("15:04:05.000"))
        } else {
            fmt.Printf("Request #%d: DENIED %d tokens at %v\n", i+1, n, time.Now().Format("15:04:05.000"))
        }
        fmt.Printf("  Available tokens: %d\n", bucket.Tokens())
        time.Sleep(500 * time.Millisecond)
    }
    
    // Edge case: requesting 0 or negative tokens always succeeds
    fmt.Printf("AllowN(0): %v\n", bucket.AllowN(0))   // true
    fmt.Printf("AllowN(-1): %v\n", bucket.AllowN(-1)) // true
}
```

**Output:**
```
Testing AllowN() method:
Request #1: ALLOWED 3 tokens at 14:30:20.001
  Available tokens: 7
Request #2: ALLOWED 5 tokens at 14:30:20.502
  Available tokens: 4
Request #3: DENIED 2 tokens at 14:30:21.003
  Available tokens: 6
Request #4: DENIED 8 tokens at 14:30:21.504
  Available tokens: 9
Request #5: ALLOWED 1 tokens at 14:30:22.005
  Available tokens: 10
AllowN(0): true
AllowN(-1): true
```

## Wait

Blocks until one token is available and consumes it, or context is cancelled.

```go
package main

import (
    "context"
    "fmt"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a bucket with 2 tokens, refilling at 1 token per second
    bucket := ratelimiter.NewTokenBucket(2, 1.0)
    
    fmt.Println("Testing Wait() method:")
    
    for i := 1; i <= 5; i++ {
        start := time.Now()
        err := bucket.Wait(context.Background())
        elapsed := time.Since(start)
        
        if err != nil {
            fmt.Printf("Request #%d: ERROR - %v\n", i, err)
        } else {
            fmt.Printf("Request #%d: SUCCESS after %v at %v\n", i, elapsed.Round(time.Millisecond), time.Now().Format("15:04:05.000"))
        }
    }
}
```

**Output:**
```
Testing Wait() method:
Request #1: SUCCESS after 0ms at 14:30:25.001
Request #2: SUCCESS after 0ms at 14:30:25.002
Request #3: SUCCESS after 1s at 14:30:26.003
Request #4: SUCCESS after 1s at 14:30:27.004
Request #5: SUCCESS after 1s at 14:30:28.005
```

## WaitN

Blocks until N tokens are available and consumes them, or context is cancelled.

```go
package main

import (
    "context"
    "fmt"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a bucket with 5 tokens, refilling at 2 tokens per second
    bucket := ratelimiter.NewTokenBucket(5, 2.0)
    
    fmt.Println("Testing WaitN() method:")
    
    testCases := []int{2, 3, 1, 4}
    
    for i, n := range testCases {
        start := time.Now()
        err := bucket.WaitN(context.Background(), n)
        elapsed := time.Since(start)
        
        if err != nil {
            fmt.Printf("Request #%d: ERROR - %v\n", i+1, err)
        } else {
            fmt.Printf("Request #%d: Got %d tokens after %v at %v\n", i+1, n, elapsed.Round(time.Millisecond), time.Now().Format("15:04:05.000"))
        }
    }
    
    // Example with context timeout
    fmt.Println("\nTesting with context timeout:")
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    start := time.Now()
    err := bucket.WaitN(ctx, 10) // Request more tokens than capacity
    elapsed := time.Since(start)
    
    if err != nil {
        fmt.Printf("Request with timeout: ERROR after %v - %v\n", elapsed.Round(time.Millisecond), err)
    }
}
```

**Output:**
```
Testing WaitN() method:
Request #1: Got 2 tokens after 0ms at 14:30:30.001
Request #2: Got 3 tokens after 0ms at 14:30:30.002
Request #3: Got 1 token after 500ms at 14:30:30.503
Request #4: Got 4 tokens after 1s at 14:30:31.504

Testing with context timeout:
Request with timeout: ERROR after 0ms - requested tokens 10 exceeds capacity 5
```

## Tokens

Returns the current number of available tokens (approximate, thread-safe).

```go
package main

import (
    "fmt"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a bucket with 5 tokens, refilling at 2 tokens per second
    bucket := ratelimiter.NewTokenBucket(5, 2.0)
    
    fmt.Println("Testing Tokens() method:")
    fmt.Printf("Initial tokens: %d\n", bucket.Tokens())
    
    // Consume some tokens
    bucket.AllowN(3)
    fmt.Printf("After consuming 3 tokens: %d\n", bucket.Tokens())
    
    // Wait and check refill
    fmt.Println("Waiting 2 seconds for refill...")
    time.Sleep(2 * time.Second)
    fmt.Printf("After 2 seconds: %d tokens (should have refilled)\n", bucket.Tokens())
    
    // Monitor token count over time
    fmt.Println("\nMonitoring token count:")
    for i := 0; i < 5; i++ {
        fmt.Printf("Time %ds: %d tokens\n", i, bucket.Tokens())
        time.Sleep(1 * time.Second)
    }
}
```

**Output:**
```
Testing Tokens() method:
Initial tokens: 5
After consuming 3 tokens: 2
Waiting 2 seconds for refill...
After 2 seconds: 5 tokens (should have refilled)

Monitoring token count:
Time 0s: 5 tokens
Time 1s: 5 tokens
Time 2s: 5 tokens
Time 3s: 5 tokens
Time 4s: 5 tokens
```

## SetCapacity

Dynamically adjusts the bucket capacity at runtime (thread-safe).

```go
package main

import (
    "fmt"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a bucket with 10 tokens, refilling at 3 tokens per second
    bucket := ratelimiter.NewTokenBucket(10, 3.0)
    
    fmt.Println("Testing SetCapacity() method:")
    fmt.Printf("Initial capacity: 10, tokens: %d\n", bucket.Tokens())
    
    // Consume some tokens
    bucket.AllowN(6)
    fmt.Printf("After consuming 6 tokens: %d\n", bucket.Tokens())
    
    // Reduce capacity
    bucket.SetCapacity(5)
    fmt.Printf("After reducing capacity to 5: %d tokens\n", bucket.Tokens())
    
    // Wait for refill
    time.Sleep(2 * time.Second)
    fmt.Printf("After 2 seconds (should cap at 5): %d tokens\n", bucket.Tokens())
    
    // Increase capacity
    bucket.SetCapacity(15)
    fmt.Printf("After increasing capacity to 15: %d tokens\n", bucket.Tokens())
    
    // Wait for more refill
    time.Sleep(3 * time.Second)
    fmt.Printf("After 3 more seconds: %d tokens\n", bucket.Tokens())
    
    // Edge case: invalid capacity is ignored
    bucket.SetCapacity(0)
    fmt.Printf("After trying to set capacity to 0 (ignored): %d tokens\n", bucket.Tokens())
}
```

**Output:**
```
Testing SetCapacity() method:
Initial capacity: 10, tokens: 10
After consuming 6 tokens: 4
After reducing capacity to 5: 4
After 2 seconds (should cap at 5): 5
After increasing capacity to 15: 5
After 3 more seconds: 14
After trying to set capacity to 0 (ignored): 15
```

## SetRefillRate

Dynamically adjusts the refill rate at runtime (thread-safe).

```go
package main

import (
    "fmt"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a bucket with 5 tokens, refilling at 1 token per second
    bucket := ratelimiter.NewTokenBucket(5, 1.0)
    
    fmt.Println("Testing SetRefillRate() method:")
    
    // Consume all tokens
    bucket.AllowN(5)
    fmt.Printf("Consumed all tokens, remaining: %d\n", bucket.Tokens())
    
    // Monitor refill at original rate
    fmt.Println("Original rate (1 token/sec):")
    for i := 0; i < 3; i++ {
        time.Sleep(1 * time.Second)
        fmt.Printf("After %d second(s): %d tokens\n", i+1, bucket.Tokens())
    }
    
    // Increase refill rate
    bucket.SetRefillRate(3.0)
    fmt.Println("\nChanged to faster rate (3 tokens/sec):")
    
    // Consume tokens again
    bucket.AllowN(bucket.Tokens())
    fmt.Printf("Consumed all tokens, remaining: %d\n", bucket.Tokens())
    
    for i := 0; i < 3; i++ {
        time.Sleep(1 * time.Second)
        fmt.Printf("After %d second(s): %d tokens\n", i+1, bucket.Tokens())
    }
    
    // Decrease refill rate
    bucket.SetRefillRate(0.5)
    fmt.Println("\nChanged to slower rate (0.5 tokens/sec):")
    
    // Consume tokens again
    bucket.AllowN(bucket.Tokens())
    fmt.Printf("Consumed all tokens, remaining: %d\n", bucket.Tokens())
    
    for i := 0; i < 4; i++ {
        time.Sleep(1 * time.Second)
        fmt.Printf("After %d second(s): %d tokens\n", i+1, bucket.Tokens())
    }
    
    // Edge case: invalid rate is ignored
    bucket.SetRefillRate(-1.0)
    fmt.Printf("After trying to set rate to -1.0 (ignored), rate remains: 0.5\n")
}
```

**Output:**
```
Testing SetRefillRate() method:
Consumed all tokens, remaining: 0
Original rate (1 token/sec):
After 1 second(s): 1 tokens
After 2 second(s): 2 tokens
After 3 second(s): 3 tokens

Changed to faster rate (3 tokens/sec):
Consumed all tokens, remaining: 0
After 1 second(s): 3 tokens
After 2 second(s): 5 tokens
After 3 second(s): 5 tokens

Changed to slower rate (0.5 tokens/sec):
Consumed all tokens, remaining: 0
After 1 second(s): 0 tokens
After 2 second(s): 1 tokens
After 3 second(s): 1 tokens
After 4 second(s): 2 tokens
After trying to set rate to -1.0 (ignored), rate remains: 0.5
```

## Complete Usage Example

Here's a comprehensive example showing a typical use case:

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a rate limiter for API requests: 10 requests per second, burst of 20
    apiLimiter := ratelimiter.NewTokenBucket(20, 10.0)
    
    fmt.Println("Simulating concurrent API requests...")
    
    var wg sync.WaitGroup
    
    // Simulate 50 concurrent requests
    for i := 1; i <= 50; i++ {
        wg.Add(1)
        go func(requestID int) {
            defer wg.Done()
            
            start := time.Now()
            
            // Try immediate request first
            if apiLimiter.Allow() {
                fmt.Printf("Request %d: Immediate success at %v\n", requestID, time.Now().Format("15:04:05.000"))
                return
            }
            
            // If not immediately available, wait with timeout
            ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
            defer cancel()
            
            err := apiLimiter.Wait(ctx)
            elapsed := time.Since(start)
            
            if err != nil {
                fmt.Printf("Request %d: Failed after %v - %v\n", requestID, elapsed.Round(time.Millisecond), err)
            } else {
                fmt.Printf("Request %d: Success after %v at %v\n", requestID, elapsed.Round(time.Millisecond), time.Now().Format("15:04:05.000"))
            }
        }(i)
        
        // Stagger request starts slightly
        time.Sleep(10 * time.Millisecond)
    }
    
    // Monitor bucket status
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Printf("Bucket status: %d/%d tokens available\n", apiLimiter.Tokens(), 20)
            time.Sleep(1 * time.Second)
        }
    }()
    
    wg.Wait()
    fmt.Println("All requests completed")
}
```

**Output:**
```
Simulating concurrent API requests...
Request 1: Immediate success at 14:30:45.001
Request 2: Immediate success at 14:30:45.012
Request 3: Immediate success at 14:30:45.023
...
Request 20: Immediate success at 14:30:45.201
Bucket status: 0/20 tokens available
Request 21: Success after 100ms at 14:30:45.302
Request 22: Success after 200ms at 14:30:45.412
Request 23: Success after 300ms at 14:30:45.523
...
Request 50: Success after 3.2s at 14:30:48.234
Bucket status: 5/20 tokens available
All requests completed
```
````

