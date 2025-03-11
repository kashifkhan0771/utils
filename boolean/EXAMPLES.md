## Boolean Function Examples

### Check if the provided string is a true

```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.IsTrue("T"))
	fmt.Println(boolutils.IsTrue("1"))
	fmt.Println(boolutils.IsTrue("TRUE"))
}
```

#### Output:

```
true
true
true
```

### Toggle a boolean value

```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.Toggle(true))
	fmt.Println(boolutils.Toggle(false))
}
```

#### Output:

```
false
true
```

### Check if all values in a slice are true

```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.AllTrue([]bool{true, true, true}))
	fmt.Println(boolutils.AllTrue([]bool{true, false, true}))
	fmt.Println(boolutils.AllTrue([]bool{}))
}
```

#### Output:

```
true
false
false
```

### Check if any value in a slice is true

```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.AnyTrue([]bool{false, true, false}))
	fmt.Println(boolutils.AnyTrue([]bool{false, false, false}))
	fmt.Println(boolutils.AnyTrue([]bool{}))
}
```

#### Output:

```
true
false
false
```

### Check if none of the values in a slice are true

```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.NoneTrue([]bool{false, false, false}))
	fmt.Println(boolutils.NoneTrue([]bool{false, true, false}))
	fmt.Println(boolutils.NoneTrue([]bool{}))
}
```

#### Output:

```
true
false
true
```

### Count the number of true values in a slice

```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.CountTrue([]bool{true, false, true}))
	fmt.Println(boolutils.CountTrue([]bool{false, false, false}))
	fmt.Println(boolutils.CountTrue([]bool{}))
}
```

#### Output:

```
2
0
0
```

### Count the number of false values in a slice

```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.CountFalse([]bool{true, false, true}))
	fmt.Println(boolutils.CountFalse([]bool{false, false, false}))
	fmt.Println(boolutils.CountFalse([]bool{}))
}
```

#### Output:

```
1
3
0
```

### Check if all values in a slice are equal

```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.Equal(true, true, true))
	fmt.Println(boolutils.Equal(false, false, false))
	fmt.Println(boolutils.Equal(true, false, true))
	fmt.Println(boolutils.Equal())
}
```

#### Output:

```
true
true
false
false
```

### Perform a logical AND operation on a slice

```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.And([]bool{true, true, true}))
	fmt.Println(boolutils.And([]bool{true, false, true}))
	fmt.Println(boolutils.And([]bool{}))
}
```

#### Output:

```
true
false
false
```

### Perform a logical OR operation on a slice

```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.Or([]bool{false, true, false}))
	fmt.Println(boolutils.Or([]bool{false, false, false}))
	fmt.Println(boolutils.Or([]bool{}))
}
```

#### Output:

```
true
false
false
```

---
