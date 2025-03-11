## Pointer Function Examples

### DefaultIfNil Example - Return Default Value if Pointer is Nil

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/pointers"
)

func main() {
	// Example of DefaultIfNil for a string pointer
	var str *string
	defaultStr := "Default String"
	result := pointers.DefaultIfNil(str, defaultStr)
	fmt.Println(result)
}
```

#### Output:

```
Default String
```

### NullableBool Example - Get Value from Bool Pointer

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/pointers"
)

func main() {
	// Example of NullableBool for a bool pointer
	var flag *bool
	result := pointers.NullableBool(flag)
	fmt.Println(result)
}
```

#### Output:

```
false
```

### NullableTime Example - Get Value from Time Pointer

```go
package main

import (
	"fmt"
	"time"

	"github.com/kashifkhan0771/utils/pointers"
)

func main() {
	// Example of NullableTime for a time.Time pointer
	var t *time.Time
	result := pointers.NullableTime(t)
	fmt.Println(result)
}
```

#### Output:

```
0001-01-01 00:00:00 +0000 UTC
```

### NullableInt Example - Get Value from Int Pointer

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/pointers"
)

func main() {
	// Example of NullableInt for an int pointer
	var num *int
	result := pointers.NullableInt(num)
	fmt.Println(result)
}
```

#### Output:

```
0
```

### NullableString Example - Get Value from String Pointer

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/pointers"
)

func main() {
	// Example of NullableString for a string pointer
	var str *string
	result := pointers.NullableString(str)
	fmt.Println(result)  // Output: ""
}
```

#### Output:

```
""
```

---
