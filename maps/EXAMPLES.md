## Map Function Examples

### StateMap Example - Set and Get States

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/maps"
)

func main() {
	// Create a new StateMap
	state := maps.NewStateMap()

	// Set a state to true
	state.SetState("isActive", true)

	// Get the state
	if state.IsState("isActive") {
		fmt.Println("The state 'isActive' is true.")
	} else {
		fmt.Println("The state 'isActive' is false.")
	}
}
```

#### Output:

```
The state 'isActive' is true.
```

### StateMap Example - Toggle a State

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/maps"
)

func main() {
	// Create a new StateMap
	state := maps.NewStateMap()

	// Set a state to true
	state.SetState("isActive", true)

	// Toggle the state
	state.ToggleState("isActive")

	// Check if the state is now false
	if !state.IsState("isActive") {
		fmt.Println("The state 'isActive' has been toggled to false.")
	}
}
```

#### Output:

```
The state 'isActive' has been toggled to false.
```

### StateMap Example - Check if State Exists

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/maps"
)

func main() {
	// Create a new StateMap
	state := maps.NewStateMap()

	// Set some states
	state.SetState("isActive", true)
	state.SetState("isVerified", false)

	// Check if the state exists
	if state.HasState("isVerified") {
		fmt.Println("State 'isVerified' exists.")
	} else {
		fmt.Println("State 'isVerified' does not exist.")
	}
}
```

#### Output:

```
State 'isVerified' exists.
```

### Metadata Example - Update and Retrieve Values

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/maps"
)

func main() {
	// Create a new Metadata map
	meta := maps.NewMetadata()

	// Update metadata with key-value pairs
	meta.Update("author", "John Doe")
	meta.Update("version", "1.0.0")

	// Retrieve metadata values
	fmt.Println("Author:", meta.Value("author"))
	fmt.Println("Version:", meta.Value("version"))
}
```

#### Output:

```
Author: John Doe
Version: 1.0.0
```

### Metadata Example - Check if Key Exists

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/maps"
)

func main() {
	// Create a new Metadata map
	meta := maps.NewMetadata()

	// Update metadata with key-value pairs
	meta.Update("author", "John Doe")

	// Check if the key exists
	if meta.Has("author") {
		fmt.Println("Key 'author' exists.")
	} else {
		fmt.Println("Key 'author' does not exist.")
	}

	// Check for a non-existent key
	if meta.Has("publisher") {
		fmt.Println("Key 'publisher' exists.")
	} else {
		fmt.Println("Key 'publisher' does not exist.")
	}
}
```

---
