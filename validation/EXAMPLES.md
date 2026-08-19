## Validation Function Examples

### StringNotEmpty

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/validation"
)

func main() {
	if err := validation.StringNotEmpty("name", "jane"); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("valid")
	}

	if err := validation.StringNotEmpty("name", ""); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("valid")
	}
}
```

#### Output:

```
valid
invalid: name: must not be empty
```

---

### StringIn

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/validation"
)

func main() {
	if err := validation.StringIn("role", "admin", "admin", "user"); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("valid")
	}

	if err := validation.StringIn("role", "superuser", "admin", "user"); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("valid")
	}
}
```

#### Output:

```
valid
invalid: role: must be one of [admin, user]
```

---

### IntBetween

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/validation"
)

func main() {
	if err := validation.IntBetween("age", 25, 0, 150); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("valid")
	}

	if err := validation.IntBetween("age", -5, 0, 150); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("valid")
	}
}
```

#### Output:

```
valid
invalid: age: integer must be between 0 and 150
```

---

### IsEmail

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/validation"
)

func main() {
	if err := validation.IsEmail("email", "jane@example.com"); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("valid")
	}

	if err := validation.IsEmail("email", "not-an-email"); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("valid")
	}
}
```

#### Output:

```
valid
invalid: email: mail: missing '@' or angle-addr
```

---

### IsURL

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/validation"
)

func main() {
	if err := validation.IsURL("url", "https://example.com"); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("valid")
	}

	// Restrict to specific schemes
	if err := validation.IsURL("url", "ftp://example.com", "http", "https"); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("valid")
	}
}
```

#### Output:

```
valid
invalid: url: must be a valid URL
```

---

### IsSlug

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/validation"
)

func main() {
	if err := validation.IsSlug("slug", "my-new-post"); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("valid")
	}

	if err := validation.IsSlug("slug", "My Post"); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("valid")
	}
}
```

#### Output:

```
valid
invalid: slug: must be a valid slug
```

---

### Errors

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/validation"
)

func main() {
	var errs validation.Errors

	errs.Add(validation.StringNotEmpty("name", ""))
	errs.Add(validation.IsEmail("email", "not-an-email"))
	errs.Add(validation.IntBetween("age", -1, 0, 150))

	if err := errs.AsError(); err != nil {
		fmt.Println(err)
	}
}
```

#### Output:

```
name: must not be empty; email: mail: missing '@' or angle-addr; age: integer must be between 0 and 150
```

---

### Validator Interface and ValidateStruct

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/validation"
)

type CreateUserInput struct {
	Name  string
	Email string
	Age   int
}

// Validate implements the validation.Validator interface.
func (i CreateUserInput) Validate() error {
	var errs validation.Errors

	errs.Add(validation.StringNotEmpty("name", i.Name))
	errs.Add(validation.IsEmail("email", i.Email))
	errs.Add(validation.IntBetween("age", i.Age, 0, 150))

	return errs.AsError()
}

func main() {
	input := CreateUserInput{
		Name:  "",
		Email: "not-an-email",
		Age:   200,
	}

	if err := validation.ValidateStruct(input); err != nil {
		fmt.Println("invalid input:", err)
	}
}
```

#### Output:

```
invalid input: name: must not be empty; email: mail: missing '@' or angle-addr; age: integer must be between 0 and 150
```

---