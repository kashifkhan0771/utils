## Fake Function Examples

### Generate a random UUID of version 4 and variant 2

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/fake"
)

func main() {
	uuid, err := fake.RandomUUID()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(uuid)
}

```

#### Output:

```
93a540eb-46e4-4e52-b0d5-cb63a7c361f9
```

### Generate a random date between 1st January 1970 and the current date

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/fake"
)

func main() {
	date, err := fake.RandomDate()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(date)
}

```

#### Output:

```
2006-06-13 21:31:17.312528419 +0200 CEST
```

### Generate a random US phone number

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/fake"
)

func main() {
	num, err := fake.RandomPhoneNumber()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(num)
}

```

#### Output:

```
+1 (965) 419-5534
```

### Generates a random US address

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/fake"
)

func main() {
	address, err := fake.RandomAddress()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(address)
}

```

#### Output:

```
81 Broadway, Rivertown, CT 12345, USA
```

---
