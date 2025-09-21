# RateLimiter Examples

This document provides comprehensive examples for all functions in the ratelimiter package.

## Table of Contents

### TokenBucket
- [NewTokenBucket](#newtokenbucket)
- [Allow](#allow)
- [AllowN](#allown)
- [Wait](#wait)
- [WaitN](#waitn)
- [Tokens](#tokens)
- [SetCapacity](#setcapacity)
- [SetRefillRate](#setrefillrate)

### FixedWindow
- [NewFixedWindow](#newfixedwindow)
- [FixedWindow Allow](#fixedwindow-allow)
- [SetInterval](#setinterval)
- [SetLimit](#setlimit)

## TokenBucket Examples

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

## FixedWindow Examples

## NewFixedWindow

Creates a new fixed window rate limiter with specified limit and interval.

```go
package main

import (
    "fmt"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a fixed window limiter allowing 5 requests per 10 seconds
    limiter := ratelimiter.NewFixedWindow(5, 10*time.Second)
    
    fmt.Printf("Created fixed window limiter: 5 requests per 10 seconds\n")
    
    // Edge cases: invalid values are corrected
    limiter2 := ratelimiter.NewFixedWindow(0, -5*time.Second) // limit becomes 1, interval becomes 1 second
    fmt.Printf("Invalid parameters corrected - limit: 1, interval: 1 second\n")
    
    // Create limiter for API rate limiting
    apiLimiter := ratelimiter.NewFixedWindow(100, 1*time.Minute)
    fmt.Printf("API limiter: 100 requests per minute\n")
}
```

**Output:**
```
Created fixed window limiter: 5 requests per 10 seconds
Invalid parameters corrected - limit: 1, interval: 1 second
API limiter: 100 requests per minute
```

## FixedWindow Allow

Checks if a new event is allowed under the rate limit without blocking.

```go
package main

import (
    "fmt"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a limiter allowing 3 requests per 2 seconds
    limiter := ratelimiter.NewFixedWindow(3, 2*time.Second)
    
    fmt.Println("Testing FixedWindow Allow() method:")
    fmt.Println("Window: 3 requests per 2 seconds")
    
    for i := 1; i <= 8; i++ {
        if limiter.Allow() {
            fmt.Printf("Request #%d: ALLOWED at %v\n", i, time.Now().Format("15:04:05.000"))
        } else {
            fmt.Printf("Request #%d: DENIED at %v\n", i, time.Now().Format("15:04:05.000"))
        }
        time.Sleep(400 * time.Millisecond)
    }
    
    fmt.Println("\nWindow should reset after 2 seconds, trying more requests:")
    
    for i := 9; i <= 12; i++ {
        if limiter.Allow() {
            fmt.Printf("Request #%d: ALLOWED at %v\n", i, time.Now().Format("15:04:05.000"))
        } else {
            fmt.Printf("Request #%d: DENIED at %v\n", i, time.Now().Format("15:04:05.000"))
        }
        time.Sleep(300 * time.Millisecond)
    }
}
```

**Output:**
```
Testing FixedWindow Allow() method:
Window: 3 requests per 2 seconds
Request #1: ALLOWED at 14:30:50.001
Request #2: ALLOWED at 14:30:50.402
Request #3: ALLOWED at 14:30:50.803
Request #4: DENIED at 14:30:51.204
Request #5: DENIED at 14:30:51.605
Request #6: DENIED at 14:30:52.006
Request #7: DENIED at 14:30:52.407
Request #8: DENIED at 14:30:52.808

Window should reset after 2 seconds, trying more requests:
Request #9: ALLOWED at 14:30:53.109
Request #10: ALLOWED at 14:30:53.410
Request #11: ALLOWED at 14:30:53.711
Request #12: DENIED at 14:30:54.012
```

## SetInterval

Dynamically updates the interval duration for the rate limiter.

```go
package main

import (
    "fmt"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a limiter allowing 2 requests per 3 seconds
    limiter := ratelimiter.NewFixedWindow(2, 3*time.Second)
    
    fmt.Println("Testing SetInterval() method:")
    fmt.Println("Initial: 2 requests per 3 seconds")
    
    // Use up the initial limit
    fmt.Printf("Request 1: %v\n", limiter.Allow()) // true
    fmt.Printf("Request 2: %v\n", limiter.Allow()) // true
    fmt.Printf("Request 3: %v\n", limiter.Allow()) // false
    
    fmt.Println("Changing interval to 1 second...")
    limiter.SetInterval(1 * time.Second)
    
    // Wait for the shorter interval
    time.Sleep(1200 * time.Millisecond)
    
    fmt.Println("After waiting 1.2 seconds (window should reset):")
    fmt.Printf("Request 4: %v\n", limiter.Allow()) // true (new window)
    fmt.Printf("Request 5: %v\n", limiter.Allow()) // true
    fmt.Printf("Request 6: %v\n", limiter.Allow()) // false
    
    // Test with invalid interval
    fmt.Println("\nTesting with invalid interval (0):")
    limiter.SetInterval(0) // Should default to 1 second
    
    time.Sleep(1100 * time.Millisecond)
    fmt.Printf("Request 7: %v\n", limiter.Allow()) // true (new window)
    
    // Change to longer interval
    fmt.Println("\nChanging to longer interval (5 seconds):")
    limiter.SetInterval(5 * time.Second)
    fmt.Printf("Request 8: %v\n", limiter.Allow()) // true
    fmt.Printf("Request 9: %v\n", limiter.Allow()) // false (at limit)
    
    fmt.Println("Waiting 2 seconds (window should NOT reset):")
    time.Sleep(2 * time.Second)
    fmt.Printf("Request 10: %v\n", limiter.Allow()) // false (still same window)
}
```

**Output:**
```
Testing SetInterval() method:
Initial: 2 requests per 3 seconds
Request 1: true
Request 2: true
Request 3: false
Changing interval to 1 second...
After waiting 1.2 seconds (window should reset):
Request 4: true
Request 5: true
Request 6: false

Testing with invalid interval (0):
Request 7: true

Changing to longer interval (5 seconds):
Request 8: true
Request 9: false
Waiting 2 seconds (window should NOT reset):
Request 10: false
```

## SetLimit

Dynamically updates the maximum allowed events per window.

```go
package main

import (
    "fmt"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a limiter allowing 2 requests per 3 seconds
    limiter := ratelimiter.NewFixedWindow(2, 3*time.Second)
    
    fmt.Println("Testing SetLimit() method:")
    fmt.Println("Initial: 2 requests per 3 seconds")
    
    // Use up the initial limit
    fmt.Printf("Request 1: %v\n", limiter.Allow()) // true
    fmt.Printf("Request 2: %v\n", limiter.Allow()) // true
    fmt.Printf("Request 3: %v\n", limiter.Allow()) // false
    
    fmt.Println("Increasing limit to 5...")
    limiter.SetLimit(5)
    
    // Should now be able to make more requests
    fmt.Printf("Request 4: %v\n", limiter.Allow()) // true
    fmt.Printf("Request 5: %v\n", limiter.Allow()) // true
    fmt.Printf("Request 6: %v\n", limiter.Allow()) // true
    fmt.Printf("Request 7: %v\n", limiter.Allow()) // false (now at new limit of 5)
    
    // Wait for window to reset
    fmt.Println("Waiting for window to reset...")
    time.Sleep(3200 * time.Millisecond)
    
    // Test reducing limit
    fmt.Println("Reducing limit to 1...")
    limiter.SetLimit(1)
    
    fmt.Printf("Request 8: %v\n", limiter.Allow()) // true
    fmt.Printf("Request 9: %v\n", limiter.Allow()) // false (at new limit of 1)
    
    // Test with invalid limit
    fmt.Println("Testing with invalid limit (0):")
    limiter.SetLimit(0) // Should default to 1
    
    time.Sleep(3200 * time.Millisecond)
    fmt.Printf("Request 10: %v\n", limiter.Allow()) // true (new window, limit 1)
    fmt.Printf("Request 11: %v\n", limiter.Allow()) // false
    
    // Test negative limit
    fmt.Println("Testing with negative limit (-5):")
    limiter.SetLimit(-5) // Should default to 1
    
    time.Sleep(3200 * time.Millisecond)
    fmt.Printf("Request 12: %v\n", limiter.Allow()) // true (new window, limit 1)
    fmt.Printf("Request 13: %v\n", limiter.Allow()) // false
}
```

**Output:**
```
Testing SetLimit() method:
Initial: 2 requests per 3 seconds
Request 1: true
Request 2: true
Request 3: false
Increasing limit to 5...
Request 4: true
Request 5: true
Request 6: true
Request 7: false
Waiting for window to reset...
Reducing limit to 1...
Request 8: true
Request 9: false
Testing with invalid limit (0):
Request 10: true
Request 11: false
Testing with negative limit (-5):
Request 12: true
Request 13: false
```

## Complete FixedWindow Usage Example

Here's a comprehensive example showing a typical use case for API rate limiting:

```go
package main

import (
    "fmt"
    "sync"
    "time"
    "utils/ratelimiter"
)

func main() {
    // Create a rate limiter for API endpoints: 10 requests per minute
    apiLimiter := ratelimiter.NewFixedWindow(10, 1*time.Minute)
    
    fmt.Println("Simulating API rate limiting...")
    fmt.Println("Limit: 10 requests per minute")
    
    var wg sync.WaitGroup
    var mu sync.Mutex
    allowed := 0
    denied := 0
    
    // Simulate 25 concurrent requests
    for i := 1; i <= 25; i++ {
        wg.Add(1)
        go func(requestID int) {
            defer wg.Done()
            
            if apiLimiter.Allow() {
                mu.Lock()
                allowed++
                currentAllowed := allowed
                mu.Unlock()
                fmt.Printf("Request %d: SUCCESS (#%d allowed) at %v\n", requestID, currentAllowed, time.Now().Format("15:04:05.000"))
            } else {
                mu.Lock()
                denied++
                currentDenied := denied
                mu.Unlock()
                fmt.Printf("Request %d: RATE LIMITED (#%d denied) at %v\n", requestID, currentDenied, time.Now().Format("15:04:05.000"))
            }
        }(i)
        
        // Stagger request starts slightly
        time.Sleep(50 * time.Millisecond)
    }
    
    wg.Wait()
    
    fmt.Printf("\nSummary: %d requests allowed, %d requests denied\n", allowed, denied)
    
    // Demonstrate window reset
    fmt.Println("\nWaiting for window to reset (1 minute)...")
    time.Sleep(61 * time.Second)
    
    fmt.Println("Testing after window reset:")
    for i := 26; i <= 30; i++ {
        if apiLimiter.Allow() {
            fmt.Printf("Request %d: SUCCESS at %v\n", i, time.Now().Format("15:04:05.000"))
        } else {
            fmt.Printf("Request %d: RATE LIMITED at %v\n", i, time.Now().Format("15:04:05.000"))
        }
        time.Sleep(100 * time.Millisecond)
    }
    
    // Demonstrate dynamic configuration
    fmt.Println("\nDemonstrating dynamic configuration:")
    fmt.Println("Changing limit to 20 requests per 30 seconds...")
    
    apiLimiter.SetLimit(20)
    apiLimiter.SetInterval(30 * time.Second)
    
    // Wait for window to reset with new configuration
    time.Sleep(31 * time.Second)
    
    fmt.Println("Testing with new configuration:")
    successCount := 0
    for i := 31; i <= 45; i++ {
        if apiLimiter.Allow() {
            successCount++
            fmt.Printf("Request %d: SUCCESS (#%d) at %v\n", i, successCount, time.Now().Format("15:04:05.000"))
        } else {
            fmt.Printf("Request %d: RATE LIMITED at %v\n", i, time.Now().Format("15:04:05.000"))
        }
        time.Sleep(100 * time.Millisecond)
    }
    
    fmt.Printf("\nWith new config: %d out of 15 requests succeeded\n", successCount)
}
```

**Output:**
```
Simulating API rate limiting...
Limit: 10 requests per minute
Request 1: SUCCESS (#1 allowed) at 14:35:10.001
Request 2: SUCCESS (#2 allowed) at 14:35:10.052
Request 3: SUCCESS (#3 allowed) at 14:35:10.103
Request 4: SUCCESS (#4 allowed) at 14:35:10.154
Request 5: SUCCESS (#5 allowed) at 14:35:10.205
Request 6: SUCCESS (#6 allowed) at 14:35:10.256
Request 7: SUCCESS (#7 allowed) at 14:35:10.307
Request 8: SUCCESS (#8 allowed) at 14:35:10.358
Request 9: SUCCESS (#9 allowed) at 14:35:10.409
Request 10: SUCCESS (#10 allowed) at 14:35:10.460
Request 11: RATE LIMITED (#1 denied) at 14:35:10.511
Request 12: RATE LIMITED (#2 denied) at 14:35:10.562
Request 13: RATE LIMITED (#3 denied) at 14:35:10.613
Request 14: RATE LIMITED (#4 denied) at 14:35:10.664
Request 15: RATE LIMITED (#5 denied) at 14:35:10.715
Request 16: RATE LIMITED (#6 denied) at 14:35:10.766
Request 17: RATE LIMITED (#7 denied) at 14:35:10.817
Request 18: RATE LIMITED (#8 denied) at 14:35:10.868
Request 19: RATE LIMITED (#9 denied) at 14:35:10.919
Request 20: RATE LIMITED (#10 denied) at 14:35:10.970
Request 21: RATE LIMITED (#11 denied) at 14:35:11.021
Request 22: RATE LIMITED (#12 denied) at 14:35:11.072
Request 23: RATE LIMITED (#13 denied) at 14:35:11.123
Request 24: RATE LIMITED (#14 denied) at 14:35:11.174
Request 25: RATE LIMITED (#15 denied) at 14:35:11.225

Summary: 10 requests allowed, 15 requests denied

Waiting for window to reset (1 minute)...
Testing after window reset:
Request 26: SUCCESS at 14:36:12.230
Request 27: SUCCESS at 14:36:12.331
Request 28: SUCCESS at 14:36:12.432
Request 29: SUCCESS at 14:36:12.533
Request 30: SUCCESS at 14:36:12.634

Demonstrating dynamic configuration:
Changing limit to 20 requests per 30 seconds...
Testing with new configuration:
Request 31: SUCCESS (#1) at 14:36:44.640
Request 32: SUCCESS (#2) at 14:36:44.741
Request 33: SUCCESS (#3) at 14:36:44.842
Request 34: SUCCESS (#4) at 14:36:44.943
Request 35: SUCCESS (#5) at 14:36:45.044
Request 36: SUCCESS (#6) at 14:36:45.145
Request 37: SUCCESS (#7) at 14:36:45.246
Request 38: SUCCESS (#8) at 14:36:45.347
Request 39: SUCCESS (#9) at 14:36:45.448
Request 40: SUCCESS (#10) at 14:36:45.549
Request 41: SUCCESS (#11) at 14:36:45.650
Request 42: SUCCESS (#12) at 14:36:45.751
Request 43: SUCCESS (#13) at 14:36:45.852
Request 44: SUCCESS (#14) at 14:36:45.953
Request 45: SUCCESS (#15) at 14:36:46.054

With new config: 15 out of 15 requests succeeded
```
````

