## Random Function Examples

### Generate a Pseudo-Random Number

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/rand"
)

func main() {
	i := rand.Int()
	i64 := rand.Int64()
	fmt.Println("Random Number:", i)
	fmt.Println("Random Number(63-bit):", i64)
}
```

#### Output:

```
Random Number: 1983964840637203872
Random Number(63-bit): 8714503361527813617
```

### Generate a Cryptographically Secure Random Number

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/rand"
)

func main() {
	n, err := rand.SecureNumber()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Cryptographically Secure Random Number:", n)
}
```

#### Output:

```
Cryptographically Secure Random Number: 5251369289452281710
```

### Generate a Random Number in a Range

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/rand"
)

func main() {
	num, err := rand.NumberInRange(10, 50)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Random Number in Range [10, 50]:", num)
}
```

#### Output:

```
Random Number in Range [10, 50]: 37
```

### Generate a Random String

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/rand"
)

func main() {
	str, err := rand.String()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Random String:", str)
}
```

#### Output:

```
Random String: b5fG8TkWz1
```

#### Generate a random string with a custom length (15)

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/rand"
)

func main() {
	str, err := rand.StringWithLength(15)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Random String (Length 15):", str)
}
```

#### Output:

```
Random String (Length 15): J8fwkL2PvXM7NqZ
```

#### Generate a random string using a custom character set

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/rand"
)

func main() {
	charset := "abcdef12345"
	str, err := rand.StringWithCharset(8, charset)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Random String with Custom Charset:", str)
}
```

#### Output:

```
Random String with Custom Charset: 1b2f3c4a
```

---

### Pick a Random Element from a Slice

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/rand"
)

func main() {
	words := []string{"apple", "banana", "cherry", "date", "elderberry"}
	word, err := rand.Pick(words)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Random Word:", word)
}
```

#### Output:

```
Random Word: cherry
```

#### Pick a random integer from a slice

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/rand"
)

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	num, err := rand.Pick(numbers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Random Number:", num)
}
```

#### Output:

```
Random Number: 40
```

---

### Shuffle a Slice

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/rand"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	err := rand.Shuffle(numbers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Shuffled Numbers:", numbers)
}
```

#### Output:

```
Shuffled Numbers: [3 1 5 4 2]
```

#### Shuffle a slice of strings

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/rand"
)

func main() {
	words := []string{"alpha", "beta", "gamma", "delta"}
	err := rand.Shuffle(words)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Shuffled Words:", words)
}
```

#### Output:

```
Shuffled Words: [delta alpha gamma beta]
```

---
