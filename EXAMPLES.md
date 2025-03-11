# Examples

This document provides practical examples of how to use the library's features. Each section demonstrates a specific use case with clear, concise code snippets.

## Table of Contents

1. [Boolean](/boolean/EXAMPLES.md)
2. [Context (ctxutils)](/ctxutils/EXAMPLES.md)
3. [Error (errutils)](/errutils/EXAMPLES.md)
4. [Maps](/maps/EXAMPLES.md)
5. [Pointers](/pointers/EXAMPLES.md)
6. [Random (rand)](/rand/EXAMPLES.md)
7. [Slice](/slice/EXAMPLES.md)
8. [Strings](/strings/EXAMPLES.md)
9. [Structs](/structs/EXAMPLES.md)
10. [Templates](/templates/EXAMPLES.md)
11. [URLs](/url/EXAMPLES.md)
12. [Math](/math/EXAMPLES.md)
13. [Fake](/fake/EXAMPLES.md)
14. [Time](/time/EXAMPLES.md)
15. [Logging](/logging/EXAMPLES.md)
16. [File System Utilities](/fsutils/EXAMPLES.md)
17. [Caching](#15-caching)

# 17. Caching

## `CacheWrapper`

### A non-thread-safe caching decorator

```go
package main

import (
	"fmt"
	"math/big"
	"github.com/kashifkhan0771/utils/math"
)

// Example function: Compute factorial
func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func main() {
	cachedFactorial := utils.CacheWrapper(factorial)
	fmt.Println(cachedFactorial(10))
}
```

#### Output:

```
3628800
```

---

## SafeCacheWrapper

### A thread-safe caching decorator

```go
package main

import (
	"fmt"
	"math/big"
	"github.com/kashifkhan0771/utils/math"
)

// Example function: Compute factorial
func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func main() {
	cachedFactorial := utils.SafeCacheWrapper(factorial)
	fmt.Println(cachedFactorial(10))
}
```

#### Output:

```
3628800
```
