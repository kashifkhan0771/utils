## Errors Function Examples

### Add Errors to Aggregator

```go
package main

import (
	"fmt"
	"errors"

    "github.com/kashifkhan0771/utils/errutils"
)

func main() {
	// Create a new error aggregator
	agg := errutils.NewErrorAggregator()

	// Add errors to the aggregator
	agg.Add(errors.New("First error"))
	agg.Add(errors.New("Second error"))
	agg.Add(errors.New("Third error"))

	// Retrieve the aggregated error
	if err := agg.Error(); err != nil {
		fmt.Println("Aggregated Error:", err)
	}
}
```

#### Output:

```
Aggregated Error: First error; Second error; Third error
```

### Check if there are any errors

```go
package main

import (
	"fmt"
	"errors"

    "github.com/kashifkhan0771/utils/errutils"
)

func main() {
	// Create a new error aggregator
	agg := errutils.NewErrorAggregator()

	// Add an error
	agg.Add(errors.New("First error"))

	// Check if there are any errors
	if agg.HasErrors() {
		fmt.Println("There are errors")
	} else {
		fmt.Println("No errors")
	}
}
```

#### Output:

```
There are errors
```

---
