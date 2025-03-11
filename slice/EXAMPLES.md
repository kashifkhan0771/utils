## Slice Function Examples

### Remove Duplicates from String Slices

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/slice"
)

func main() {
	strings := []string{"apple", "banana", "apple", "cherry", "banana", "date"}
	uniqueStrings := slice.RemoveDuplicateStr(strings)
	fmt.Println("Unique Strings:", uniqueStrings)
}
```

#### Output:

```
Unique Strings: [apple banana cherry date]
```

### Remove Duplicates from Integer Slices

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/slice"
)

func main() {
	numbers := []int{1, 2, 3, 2, 1, 4, 5, 3, 4}
	uniqueNumbers := slice.RemoveDuplicateInt(numbers)
	fmt.Println("Unique Numbers:", uniqueNumbers)
}
```

#### Output:

```
Unique Numbers: [1 2 3 4 5]
```

---
