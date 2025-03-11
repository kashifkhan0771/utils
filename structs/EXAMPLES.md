## Structs Function Examples

### Compare two structs

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/structs"
)

// Define a struct with fields tagged as `updateable`.
type User struct {
	ID        int    `updateable:"false"`       // Not updateable
	Name      string `updateable:"true"`        // Updateable, uses default field name
	Email     string `updateable:"email_field"` // Updateable, uses a custom tag name
	Age       int    `updateable:"true"`        // Updateable, uses default field name
	IsAdmin   bool   // Not tagged, so not updateable
}

func main() {
	oldUser := User{
		ID:      1,
		Name:    "Alice",
		Email:   "alice@example.com",
		Age:     25,
		IsAdmin: false,
	}

	newUser := User{
		ID:      1,
		Name:    "Alice Johnson",
		Email:   "alice.johnson@example.com",
		Age:     26,
		IsAdmin: true,
	}

	// Compare the two struct instances.
	results, err := structs.CompareStructs(oldUser, newUser)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the results.
	for _, result := range results {
		fmt.Printf("Field: %s, Old Value: %v, New Value: %v\n", result.FieldName, result.OldValue, result.NewValue)
	}
}
```

### Output:

```
Field: Name, Old Value: Alice, New Value: Alice Johnson
Field: email_field, Old Value: alice@example.com, New Value: alice.johnson@example.com
Field: Age, Old Value: 25, New Value: 26
```

---
