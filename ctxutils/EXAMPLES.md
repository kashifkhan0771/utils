## Context (ctxutils) Function Examples

### Set and Get a String Value in Context

```go
package main

import (
	"context"
	"fmt"

	"github.com/kashifkhan0771/utils/ctxutils"
)

func main() {
	// Create a context
	ctx := context.Background()

	// Set a string value in context
	ctx = ctxutils.SetStringValue(ctx, ctxutils.ContextKeyString{"userName"}, "JohnDoe")

	// Get the string value from context
	userName, ok := ctxutils.GetStringValue(ctx, ctxutils.ContextKeyString{"userName"})
	if ok {
		fmt.Println(userName)
	}
}
```

#### Output:

```
JohnDoe
```

### Set and Get a Int Value in Context

```go
package main

import (
	"context"
	"fmt"

	"github.com/kashifkhan0771/utils/ctxutils"
)

func main() {
	// Create a context
	ctx := context.Background()

	// Set an integer value in context
	ctx = ctxutils.SetIntValue(ctx, ctxutils.ContextKeyInt{Key: 42}, 100)

	// Get the integer value from context
	value, ok := ctxutils.GetIntValue(ctx, ctxutils.ContextKeyInt{Key: 42})
	if ok {
		fmt.Println(value)
	}
}

```

#### Output:

```
100
```

---
