## Math Function Examples

## `Abs`

### Calculate the absolute value of a number

```go
import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	fmt.Println(utils.Abs(-5))
	fmt.Println(utils.Abs(10))
	fmt.Println(utils.Abs(0))
}
```

#### Output:

```
5
10
0
```

---

## `Sign`

### Determine the sign of a number

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	fmt.Println(utils.Sign(15))  // Positive number
	fmt.Println(utils.Sign(-10)) // Negative number
	fmt.Println(utils.Sign(0))   // Zero
}
```

#### Output:

```
1
-1
0
```

---

## `Min`

### Find the smaller of two numbers

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	fmt.Println(utils.Min(10, 20))
	fmt.Println(utils.Min(25, 15))
	fmt.Println(utils.Min(7, 7))
}
```

#### Output:

```
10
15
7
```

---

## `Max`

### Find the larger of two numbers

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	fmt.Println(utils.Max(10, 20))
	fmt.Println(utils.Max(25, 15))
	fmt.Println(utils.Max(7, 7))
}
```

#### Output:

```
20
25
7
```

---

## `Clamp`

### Clamp a value to stay within a range

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	fmt.Println(utils.Clamp(1, 10, 5))  // Value within range
	fmt.Println(utils.Clamp(1, 10, 0))  // Value below range
	fmt.Println(utils.Clamp(1, 10, 15)) // Value above range
}
```

#### Output:

```
5
1
10
```

---

## `IntPow`

### Compute integer powers

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	fmt.Println(utils.IntPow(2, 3))  // 2^3
	fmt.Println(utils.IntPow(5, 0))  // 5^0
	fmt.Println(utils.IntPow(3, 2))  // 3^2
	fmt.Println(utils.IntPow(2, -3))  // 3^(-3)
}
```

#### Output:

```
8
1
9
0.125
```

---

## `IsEven`

### Check if a number is even

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	fmt.Println(utils.IsEven(8))  // Even number
	fmt.Println(utils.IsEven(7))  // Odd number
	fmt.Println(utils.IsEven(0))  // Zero
}
```

#### Output:

```
true
false
true
```

---

## `IsOdd`

### Check if a number is odd

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	fmt.Println(utils.IsOdd(7))  // Odd number
	fmt.Println(utils.IsOdd(8))  // Even number
	fmt.Println(utils.IsOdd(0))  // Zero
}
```

#### Output:

```
true
false
false
```

---

## `Swap`

### Swap two variables

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	x, y := 10, 20
	utils.Swap(&x, &y)
	fmt.Println(x, y)
}
```

#### Output:

```
20 10
```

---

## `Factorial`

### Calculate the factorial of a number

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
    result, err := utils.Factorial(5)
    if err != nil {
        fmt.Printf("%v\n", err)
    }

    fmt.Println(result)
}
```

#### Output:

```
120
```

---

## `GCD`

### Find the greatest common divisor of two numbers

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	fmt.Println(utils.GCD(12, 18))
	fmt.Println(utils.GCD(17, 19)) // Prime numbers
	fmt.Println(utils.GCD(0, 5))   // Zero input
}
```

#### Output:

```
6
1
5
```

---

## `LCM`

### Find the least common multiple of two numbers

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	fmt.Println(utils.LCM(4, 6))
	fmt.Println(utils.LCM(7, 13))  // Prime numbers
	fmt.Println(utils.LCM(0, 5))   // Zero input
}
```

#### Output:

```
12
91
0
```

---

## `Sqrt`

### Find the square root of a given number.

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	fmt.Println(utils.Sqrt(4))
	fmt.Println(utils.Sqrt(2))
}
```

#### Output:

```
2
1.4142135623730951
```
## `IsPrime`

### Find if a number is prime or not.

```go
package main

import (
	"fmt"

	utils "github.com/kashifkhan0771/utils/math"
)

func main() {
	fmt.Println(utils.IsPrime(5))
	fmt.Println(utils.IsPrime(4))
}
```

#### Output:

```
true
false
```
---
