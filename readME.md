# Utils
Common Utilities library for Go

## Overview
Utils is a lightweight, flexible, and reusable library providing utility functions and helpers for common operations in Go applications. With packages designed for managing boolean conversions, handling context values, map operations, slice utilities, string manipulations, and struct comparison, utils enhances Go projects with optimized, clean, and practical solutions.

### Prerequisites
**Go**: The project is written in Golang, so you'll need Go installed (preferably Go 1.16 or later). You can download and install Go from [here](https://go.dev/doc/install).

**Git**: For cloning the repository.

### Installation
To use utils in your project, add it as a module dependency:

#### Clone the Repository
```
go get github.com/kashifkhan0771/utils
```
Alternatively, include it directly in your go.mod file (use the latest release):
```
require github.com/kashifkhan0771/utils v0.3.0
```

## Key Features

**Boolean Utilities (boolean)**: Simple functions for evaluating and converting string values to booleans.

**Context Utilities (context)**: Convenient functions for setting and retrieving typed values from context.

**Map Helpers (maps)**: State management with StateMap, metadata storage with Metadata, and efficient map operations.

**Slice Utilities (slice)**: Duplicate removal for string and integer slices.

**String Manipulation (strings)**: Substring search, case transformations, ROT13/Caesar encoding, email validation, and more.

**Struct Comparison (structs)**: Deep comparison between structs with custom field tags.

## Usage Guide
After adding utils to your project, you can import and utilize the packages as needed. Below is a breakdown of each package and some example usage.

### Boolean (boolean)
Provides functions to handle boolean conversion from strings.

**IsTrue(v string) bool**: Converts strings like "1", "t", "T", "TRUE", "true", or "True" to true, treating all other inputs as false.

Example:
```
package main

import "github.com/kashifkhan0771/utils/boolean"

func main() {
    isTrue := boolean.IsTrue("true")   // returns true
    isFalse := boolean.IsTrue("false") // returns false
    fmt.Println(isTrue, isFalse) // Output: true false
}
```

### Context Utilities (ctxutils)
Typed setters and getters for safely storing and retrieving values from context.

**SetStringValue(ctx context.Context, key ContextKeyString, value string) context.Context**: Stores a string in context.

**GetStringValue(ctx context.Context, key ContextKeyString) (string, bool)**: Retrieves a string from context.

Example:
```
package main

import (
	"context"
	"fmt"

	"github.com/kashifkhan0771/utils/ctxutils"
)

func main() {
	// Define the key using ContextKeyString
	usernameKey := ctxutils.ContextKeyString{Key: "username"}

	// Set a value in the context with the defined key
	ctx := context.Background()
	ctx = ctxutils.SetStringValue(ctx, usernameKey, "shahzad")

	// Retrieve the value from the context using the same key
	if username, ok := ctxutils.GetStringValue(ctx, usernameKey); ok {
		fmt.Println("Username:", username) // Output: Username: shahzad
	}
}
```

### Slice Utilities (slice)
Helpers for common slice operations.

**RemoveDuplicateStr(strSlice []string) []string**: Removes duplicates from a string slice.

**RemoveDuplicateInt(intSlice []int) []int**: Removes duplicates from an integer slice.

Example:
```
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/slice"
)

func main() {
	items := []string{"apple", "banana", "apple", "orange", "banana", "apple", "orange"}
	uniqueItems := slice.RemoveDuplicateStr(items)

	fmt.Println(uniqueItems) // Output: [apple banana orange]
}
```

### Maps (maps)
Efficient state management and metadata handling.

**NewStateMap() StateMap**: Creates a new StateMap.

**ToggleState(stateType string)**: Toggles boolean state in StateMap.

**NewMetadata() Metadata**: Creates a Metadata instance for managing key-value pairs.

### Strings (strings)
Advanced string operations and transformations.

**SubstringSearch(input, substring string, options SubstringSearchOptions) []string**: Searches for substrings with optional case insensitivity and index return.

**Title(input string) string**: Converts a string to title case.

**IsValidEmail(email string) bool**: Checks email format validity.

Example:
```
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/strings"
)

func main() {
	title := strings.Title("hello world") // Converts to title case
	valid := strings.IsValidEmail("example@email.com")
	inValid := strings.IsValidEmail("example.email.com")
	anotherValid := strings.IsValidEmail("example@email.tech")

	fmt.Println(title)        // Output: Hello World
	fmt.Println(valid)        // Output: true
	fmt.Println(inValid)      // Output: false
	fmt.Println(anotherValid) // Output: true
}
```

### Structs
Efficient, tag-based struct comparison.

**CompareStructs(old, new interface{}) ([]Result, error)**: Compares two structs based on custom field tags, returning changes.

Example:
```
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/structs"
)

type Person struct {
	Name string `updateable:"true"`
	Age  int    `updateable:"true"`
}

func main() {
	person1 := Person{Name: "Alice", Age: 25}
	person2 := Person{Name: "Alice", Age: 30}

	differences, _ := structs.CompareStructs(person1, person2)
	fmt.Println(differences) // Output: [{Age 25 30}]
}
```


### Templates
Dynamic rendering of HTML and text templates.

**HTML Template Rendering**: Enhances Go's html/template library by incorporating a comprehensive set of utility functions making it easier to handle inline conditionals, loops, and a variety of string, numeric, logic, date, and debugging operations within templates.

Example:

```
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/templates"
)

func main() {
	htmlTemplate := `<h1>{{toUpper .Title}}</h1><p>{{.Content}}</p>`
	data := map[string]interface{}{
		"Title":   "hello, world",
		"Content": "This is a sample content for HTML template rendering.",
	}

	result, err := templates.RenderHTMLTemplate(htmlTemplate, data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
	// Output: <h1>HELLO, WORLD</h1><p>This is a sample content for HTML template rendering.</p>
}
```

**Text Template Rendering**: Enhances Go's text/template library by incorporating a comprehensive set of utility functions making it easier to handle a variety of string, numeric, logic, date, and debugging operations within templates.

Example:

```
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/templates"
)

func main() {
	textTemplate := `Name: {{.Name | toUpper}}, Age: {{add .Age 5}}`
	data := map[string]interface{}{
		"Name": "john doe",
		"Age":  20,
	}

	result, err := templates.RenderText(textTemplate, data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
	// Output: Name: JOHN DOE, Age: 25
}
```

**Custom Functions**

The templates package offers a set of custom functions that can be directly used within templates to perform various operations.

_String Functions:_

- toUpper, toLower, title, contains, replace, trim, split, reverse, toString

_Date and Time Functions:_

- formatDate, now

_Arithmetic Functions:_

- add, sub, mul, div, mod

_Conditional and Logical Functions:_

- isNil, not

_Debugging Functions:_

- dump, typeOf

Example:

```
{{ "example text" | toUpper }}           // Outputs: EXAMPLE TEXT
{{ formatDate now "2006-01-02" }}        // Outputs the current date in YYYY-MM-DD format
{{ add 10 5 }}                           // Outputs: 15
{{ typeOf .SomeVariable }}               // Outputs the type of .SomeVariable
```


# Contributions
Contributions to this project are welcome! If you would like to contribute, please feel free to open a PR.

Please read the [Contribution Guide](/CONTRIBUTING.md) before opening any new pull request

Together, we can make Utils even better for the Go community!
