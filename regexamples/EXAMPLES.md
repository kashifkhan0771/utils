## Regex Examples Function Examples

### Generate examples for a digit pattern

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/regexamples"
)

func main() {
	examples, err := regexamples.Generate(`\d{4}`, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, s := range examples {
		fmt.Println(s)
	}
}
```

#### Output:

```
7392
0154
8820
3467
5901
```

---

### Generate examples for a word pattern with a range quantifier

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/regexamples"
)

func main() {
	examples, err := regexamples.Generate(`\w{3,6}`, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, s := range examples {
		fmt.Println(s)
	}
}
```

#### Output:

```
aB3
kQz19w
mR7
x2Lp
Tz4kWn
```

---

### Generate examples for an alternation pattern

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/regexamples"
)

func main() {
	examples, err := regexamples.Generate(`(foo|bar|baz)`, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, s := range examples {
		fmt.Println(s)
	}
}
```

#### Output:

```
bar
foo
baz
foo
bar
```

---

### Generate examples for an email-like pattern

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/regexamples"
)

func main() {
	examples, err := regexamples.Generate(`[a-z]{4,8}@[a-z]{3,6}\.(com|net|org)`, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, s := range examples {
		fmt.Println(s)
	}
}
```

#### Output:

```
jktr@mxpw.com
abcdefgh@zlo.net
pwqr@abcde.org
mnop@xyz.com
lkjh@qrst.net
```

---

### Reuse a Generator for multiple batches

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/regexamples"
)

func main() {
	g, err := regexamples.NewGenerator(`[A-Z]{2}\d{4}`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for range 3 {
		batch, err := g.Generate(3)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(batch)
	}
}
```

#### Output:

```
[KT8301 PX4720 BN0093]
[ZA1154 MQ6688 RL2279]
[CF9900 YD3341 WS7712]
```

---

### Deterministic generation with SetSeed

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/regexamples"
)

func main() {
	g, err := regexamples.NewGenerator(`[a-z]{5}-\d{3}`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	g.SetSeed(42)

	examples, err := g.Generate(3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, s := range examples {
		fmt.Println(s)
	}
}
```

#### Output:

```
mxkqr-581
azpwt-047
lbnvc-923
```

---
