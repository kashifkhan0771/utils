## Retry Examples

### Basic usage
```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kashifkhan0771/utils/retry"
)

func main() {
	opts := retry.Options{
		MaxAttempts:  3,
		TotalTimeout: 10 * time.Second,
		Backoff:      retry.FixedBackoff(500 * time.Millisecond),
		ShouldRetry:  func(err error) bool { return true },
	}

	result, err := retry.Do(context.Background(), opts, func(ctx context.Context) (string, error) {
		return callAPI(ctx)
	})
	fmt.Println(result, err)
}
```
#### Output:
```
"response" <nil>
```

---

### Aborting on non-retryable errors
```go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kashifkhan0771/utils/retry"
)

var ErrNotFound = errors.New("not found")

func main() {
	opts := retry.Options{
		MaxAttempts:  5,
		TotalTimeout: 10 * time.Second,
		Backoff:      retry.FixedBackoff(500 * time.Millisecond),
		ShouldRetry: func(err error) bool {
			return !errors.Is(err, ErrNotFound)
		},
	}

	_, err := retry.Do(context.Background(), opts, func(ctx context.Context) (string, error) {
		return "", ErrNotFound
	})
	fmt.Println(err)
}
```
#### Output:
```
not found
```

---

### Using DoVoid for side-effecting operations
```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kashifkhan0771/utils/retry"
)

func main() {
	opts := retry.Options{
		MaxAttempts:  3,
		TotalTimeout: 10 * time.Second,
		Backoff:      retry.FixedBackoff(200 * time.Millisecond),
		ShouldRetry:  func(err error) bool { return true },
	}

	err := retry.DoVoid(context.Background(), opts, func(ctx context.Context) error {
		return sendNotification(ctx)
	})
	fmt.Println(err)
}
```
#### Output:
```
<nil>
```

---

### Backoff strategies compared
```go
package main

import (
	"fmt"
	"time"

	"github.com/kashifkhan0771/utils/retry"
)

func main() {
	fixed := retry.FixedBackoff(1 * time.Second)
	linear := retry.LinearBackoff(1 * time.Second)
	exponential := retry.ExponentialBackoff(1 * time.Second)

	for attempt := range 4 {
		fmt.Printf("attempt %d  fixed=%-6s  linear=%-6s  exponential=%s\n",
			attempt,
			fixed(attempt),
			linear(attempt),
			exponential(attempt),
		)
	}
}
```
#### Output:
```
attempt 0  fixed=1s      linear=0s      exponential=1s
attempt 1  fixed=1s      linear=1s      exponential=2s
attempt 2  fixed=1s      linear=2s      exponential=4s
attempt 3  fixed=1s      linear=3s      exponential=8s
```

---

### Respecting context cancellation
```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kashifkhan0771/utils/retry"
)

func main() {
	opts := retry.Options{
		MaxAttempts:  10,
		TotalTimeout: 2 * time.Second,
		Backoff:      retry.ExponentialBackoff(1 * time.Second),
		ShouldRetry:  func(err error) bool { return true },
	}

	_, err := retry.Do(context.Background(), opts, func(ctx context.Context) (string, error) {
		return "", fmt.Errorf("service unavailable")
	})
	fmt.Println(err)
}
```
#### Output:
```
context deadline exceeded
```
