# Examples
This document provides practical examples of how to use the library's features. Each section demonstrates a specific use case with clear, concise code snippets.

## Table of Contents
1. [Boolean](#1-boolean-utilities)
2. [Context (ctxutils)](#2-context-ctxutils)
3. [Error (errutils)](#3-errors)
4. [Maps](#4-map)
5. [Pointers](#5-pointer)
6. [Random (rand)](#6-random)
7. [Slice](#7-slice)
8. [Strings](#8-strings)
9. [Structs](#9-structs)
10. [Templates](#10-templates)
11. [URLs](#11-urls)
12. [Math](#12-math)
13. [Fake](#13-fake)
14. [Time](#14-time)
15. [Loggin](#15-logging)
16. [File System Utilities](#16-fsutils)
15. [Loggin](#15-logging)
16. [File System Utilities](#16-fsutils)
17. [Caching](#15-caching)

## 1. Boolean

### Check if the provided string is a true
```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.IsTrue("T"))
	fmt.Println(boolutils.IsTrue("1"))
	fmt.Println(boolutils.IsTrue("TRUE"))
}
```
#### Output:
```
true
true
true
```

### Toggle a boolean value
```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.Toggle(true))
	fmt.Println(boolutils.Toggle(false))
}
```
#### Output:
```
false
true
```

### Check if all values in a slice are true
```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.AllTrue([]bool{true, true, true}))
	fmt.Println(boolutils.AllTrue([]bool{true, false, true}))
	fmt.Println(boolutils.AllTrue([]bool{}))
}
```
#### Output:
```
true
false
false
```

### Check if any value in a slice is true
```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.AnyTrue([]bool{false, true, false}))
	fmt.Println(boolutils.AnyTrue([]bool{false, false, false}))
	fmt.Println(boolutils.AnyTrue([]bool{}))
}
```
#### Output:
```
true
false
false
```

### Check if none of the values in a slice are true
```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.NoneTrue([]bool{false, false, false}))
	fmt.Println(boolutils.NoneTrue([]bool{false, true, false}))
	fmt.Println(boolutils.NoneTrue([]bool{}))
}
```
#### Output:
```
true
false
true
```

### Count the number of true values in a slice
```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.CountTrue([]bool{true, false, true}))
	fmt.Println(boolutils.CountTrue([]bool{false, false, false}))
	fmt.Println(boolutils.CountTrue([]bool{}))
}
```
#### Output:
```
2
0
0
```

### Count the number of false values in a slice
```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.CountFalse([]bool{true, false, true}))
	fmt.Println(boolutils.CountFalse([]bool{false, false, false}))
	fmt.Println(boolutils.CountFalse([]bool{}))
}
```
#### Output:
```
1
3
0
```

### Check if all values in a slice are equal
```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.Equal(true, true, true))
	fmt.Println(boolutils.Equal(false, false, false))
	fmt.Println(boolutils.Equal(true, false, true))
	fmt.Println(boolutils.Equal())
}
```
#### Output:
```
true
true
false
false
```

### Perform a logical AND operation on a slice
```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.And([]bool{true, true, true}))
	fmt.Println(boolutils.And([]bool{true, false, true}))
	fmt.Println(boolutils.And([]bool{}))
}
```
#### Output:
```
true
false
false
```

### Perform a logical OR operation on a slice
```go
package main

import (
	"fmt"

	boolutils "github.com/kashifkhan0771/utils/boolean"
)

func main() {
	fmt.Println(boolutils.Or([]bool{false, true, false}))
	fmt.Println(boolutils.Or([]bool{false, false, false}))
	fmt.Println(boolutils.Or([]bool{}))
}
```
#### Output:
```
true
false
false
```

---
## 2. Context (ctxutils)

### Set and Get a String Value in Context
```go
package main

import (
	"context"
	"fmt"

	"github.com/kashifkhan0771/utils/ctxutils"
)

func main() {
	// Create a context
	ctx := context.Background()

	// Set a string value in context
	ctx = ctxutils.SetStringValue(ctx, ctxutils.ContextKeyString{"userName"}, "JohnDoe")

	// Get the string value from context
	userName, ok := ctxutils.GetStringValue(ctx, ctxutils.ContextKeyString{"userName"})
	if ok {
		fmt.Println(userName)
	}
}
```
#### Output:
```
JohnDoe
```

### Set and Get a Int Value in Context
```go
package main

import (
	"context"
	"fmt"

	"github.com/kashifkhan0771/utils/ctxutils"
)

func main() {
	// Create a context
	ctx := context.Background()

	// Set an integer value in context
	ctx = ctxutils.SetIntValue(ctx, ctxutils.ContextKeyInt{Key: 42}, 100)

	// Get the integer value from context
	value, ok := ctxutils.GetIntValue(ctx, ctxutils.ContextKeyInt{Key: 42})
	if ok {
		fmt.Println(value)
	}
}

```
#### Output:
```
100
```
---

## 3. Errors

### Add Errors to Aggregator
```go
package main

import (
	"fmt"
	"errors"

    "github.com/kashifkhan0771/utils/errutils"
)

func main() {
	// Create a new error aggregator
	agg := errutils.NewErrorAggregator()

	// Add errors to the aggregator
	agg.Add(errors.New("First error"))
	agg.Add(errors.New("Second error"))
	agg.Add(errors.New("Third error"))

	// Retrieve the aggregated error
	if err := agg.Error(); err != nil {
		fmt.Println("Aggregated Error:", err)
	}
}
```
#### Output:
```
Aggregated Error: First error; Second error; Third error
```
### Check if there are any errors
```go
package main

import (
	"fmt"
	"errors"

    "github.com/kashifkhan0771/utils/errutils"
)

func main() {
	// Create a new error aggregator
	agg := errutils.NewErrorAggregator()

	// Add an error
	agg.Add(errors.New("First error"))

	// Check if there are any errors
	if agg.HasErrors() {
		fmt.Println("There are errors")
	} else {
		fmt.Println("No errors")
	}
}
```
#### Output:
```
There are errors
```
---


## 4. Map

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

## 5. Pointer

### DefaultIfNil Example - Return Default Value if Pointer is Nil
```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/pointers"
)

func main() {
	// Example of DefaultIfNil for a string pointer
	var str *string
	defaultStr := "Default String"
	result := pointers.DefaultIfNil(str, defaultStr)
	fmt.Println(result)
}
```
#### Output:
```
Default String
```

### NullableBool Example - Get Value from Bool Pointer
```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/pointers"
)

func main() {
	// Example of NullableBool for a bool pointer
	var flag *bool
	result := pointers.NullableBool(flag)
	fmt.Println(result)
}
```
#### Output:
```
false
```

### NullableTime Example - Get Value from Time Pointer
```go
package main

import (
	"fmt"
	"time"

	"github.com/kashifkhan0771/utils/pointers"
)

func main() {
	// Example of NullableTime for a time.Time pointer
	var t *time.Time
	result := pointers.NullableTime(t)
	fmt.Println(result)
}
```
#### Output:
```
0001-01-01 00:00:00 +0000 UTC
```

### NullableInt Example - Get Value from Int Pointer
```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/pointers"
)

func main() {
	// Example of NullableInt for an int pointer
	var num *int
	result := pointers.NullableInt(num)
	fmt.Println(result)
}
```
#### Output:
```
0
```

### NullableString Example - Get Value from String Pointer
```go
package main

import (
	"fmt"
    
	"github.com/kashifkhan0771/utils/pointers"
)

func main() {
	// Example of NullableString for a string pointer
	var str *string
	result := pointers.NullableString(str)
	fmt.Println(result)  // Output: ""
}
```
#### Output:
```
""
```

---
## 6. Random


### Generate a Random Number

```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/rand"
)

func main() {
	num, err := rand.Number()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Random Number:", num)
}
```
#### Output:
```
Random Number: 8507643814357583841
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

## 7. Slice

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

## 8. Strings

### SubstringSearch
```go
package main

import (
	"fmt"
	"github.com/kashifkhan0771/utils/strings"
)

func main() {
	options := strings.SubstringSearchOptions{
		CaseInsensitive: true,
		ReturnIndexes:   false,
	}
	result := strings.SubstringSearch("Go is Great, Go is Fun!", "go", options)
	fmt.Println(result) // Output: [Go Go]
}
```

### Title
```go
func main() {
	title := strings.Title("hello world")
	fmt.Println(title) // Output: Hello World
}
```

### ToTitle
```go
func main() {
	title := strings.ToTitle("hello world of go", []string{"of", "go"})
	fmt.Println(title) // Output: Hello World of go
}
```

### Tokenize
```go
func main() {
	tokens := strings.Tokenize("hello,world;this is Go", ",;")
	fmt.Println(tokens) // Output: [hello world this is Go]
}
```

### Rot13Encode
```go
func main() {
	encoded := strings.Rot13Encode("hello")
	fmt.Println(encoded) // Output: uryyb
}
```

### Rot13Decode
```go
func main() {
	decoded := strings.Rot13Decode("uryyb")
	fmt.Println(decoded) // Output: hello
}
```

### CaesarEncrypt
```go
func main() {
	encrypted := strings.CaesarEncrypt("hello", 3)
	fmt.Println(encrypted) // Output: khoor
}
```

### CaesarDecrypt
```go
func main() {
	decrypted := strings.CaesarDecrypt("khoor", 3)
	fmt.Println(decrypted) // Output: hello
}
```

### RunLengthEncode
```go
func main() {
	encoded := strings.RunLengthEncode("aaabbbccc")
	fmt.Println(encoded) // Output: 3a3b3c
}
```

### RunLengthDecode
```go
func main() {
	decoded, _ := strings.RunLengthDecode("3a3b3c")
	fmt.Println(decoded) // Output: aaabbbccc
}
```

### IsValidEmail
```go
func main() {
	valid := strings.IsValidEmail("test@example.com")
	fmt.Println(valid) // Output: true
}
```

### SanitizeEmail
```go
func main() {
	email := strings.SanitizeEmail("   test@example.com   ")
	fmt.Println(email) // Output: test@example.com
}
```

### Reverse
```go
func main() {
	reversed := strings.Reverse("hello")
	fmt.Println(reversed) // Output: olleh
}
```

### CommonPrefix
```go
func main() {
	prefix := strings.CommonPrefix("nation", "national", "nasty")
	fmt.Println(prefix) // Output: na
}
```

### CommonSuffix
```go
func main() {
	suffix := strings.CommonSuffix("testing", "running", "jumping")
	fmt.Println(suffix) // Output: ing
}
```
---

## 9. Structs

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


## 10. Templates

### Render HTML and Text Templates:

```go
package main

import (
	"fmt"
	"time"

	"github.com/kashifkhan0771/utils/templates"
)

func main() {
	// Define a sample HTML template
	htmlTmpl := `
	<!DOCTYPE html>
	<html>
		<head><title>{{ title .Title }}</title></head>
		<body>
			<h1>{{ toUpper .Header }}</h1>
			<p>{{ .Content }}</p>
			<p>Generated at: {{ formatDate .GeneratedAt "2006-01-02 15:04:05" }}</p>
		</body>
	</html>
	`

	// Define data for the HTML template
	htmlData := map[string]interface{}{
		"Title":      "template rendering demo",
		"Header":     "welcome to the demo",
		"Content":    "This is a demonstration of the RenderHTMLTemplate function.",
		"GeneratedAt": time.Now(),
	}

	// Render the HTML template
	renderedHTML, err := templates.RenderHTMLTemplate(htmlTmpl, htmlData)
	if err != nil {
		fmt.Println("Error rendering HTML template:", err)
		return
	}

	fmt.Println("Rendered HTML Template:")
	fmt.Println(renderedHTML)

	// Define a sample text template
	textTmpl := `
	Welcome, {{ toUpper .Name }}!
	Today is {{ formatDate .Date "Monday, January 2, 2006" }}.
	{{ if contains .Message "special" }}
	Note: You have a special message!
	{{ end }}
	`

	// Define data for the text template
	textData := map[string]interface{}{
		"Name":    "Alice",
		"Date":    time.Now(),
		"Message": "This is a special announcement.",
	}

	// Render the text template
	renderedText, err := templates.RenderText(textTmpl, textData)
	if err != nil {
		fmt.Println("Error rendering text template:", err)
		return
	}

	fmt.Println("Rendered Text Template:")
	fmt.Println(renderedText)
}
```

---

### Output:

#### Rendered HTML Template:
```html
<!DOCTYPE html>
<html>
	<head><title>Template Rendering Demo</title></head>
	<body>
		<h1>WELCOME TO THE DEMO</h1>
		<p>This is a demonstration of the RenderHTMLTemplate function.</p>
		<p>Generated at: 2024-11-19 14:45:00</p>
	</body>
</html>
```

#### Rendered Text Template:
```text
Welcome, ALICE!
Today is Tuesday, November 19, 2024.
Note: You have a special message!
```

### Available Functions:
Here's a list of all available custom functions from the `customFuncsMap`:

### String Functions:
1. **`toUpper`**: Converts a string to uppercase.
2. **`toLower`**: Converts a string to lowercase.
3. **`title`**: Converts a string to title case (e.g., "hello world" → "Hello World").
4. **`contains`**: Checks if a string contains a specified substring.
5. **`replace`**: Replaces all occurrences of a substring with another string.
6. **`trim`**: Removes leading and trailing whitespace from a string.
7. **`split`**: Splits a string into a slice based on a specified delimiter.
8. **`reverse`**: Reverses a string (supports Unicode characters).
9. **`toString`**: Converts a value of any type to its string representation.

### Date and Time Functions:
10. **`formatDate`**: Formats a `time.Time` object using a custom layout.
11. **`now`**: Returns the current date and time (`time.Time`).

### Numeric and Arithmetic Functions:
12. **`add`**: Adds two integers.
13. **`sub`**: Subtracts the second integer from the first.
14. **`mul`**: Multiplies two integers.
15. **`div`**: Divides the first integer by the second (integer division).
16. **`mod`**: Returns the remainder of dividing the first integer by the second.

### Conditional and Logical Functions:
17. **`isNil`**: Checks if a value is `nil`.
18. **`not`**: Negates a boolean value (e.g., `true` → `false`).

### Debugging Functions:
19. **`dump`**: Returns a detailed string representation of a value (useful for debugging).
20. **`typeOf`**: Returns the type of a value as a string.

### Safe HTML Rendering:
21. **`safeHTML`**: Marks a string as safe HTML, preventing escaping in templates.

---

## 11. URLs

### Build a URL
```go
url, err := BuildURL("https", "example.com", "search", map[string]string{"q": "golang"})
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("URL:", url)
}
// Output: https://example.com/search?q=golang
```

### Add Query Params
```go
url, err := AddQueryParams("http://example.com", map[string]string{"key": "value", "page": "2"})
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Updated URL:", url)
}
// Output: http://example.com?key=value&page=2
```

### Validate a URL
```go
isValid := IsValidURL("https://example.com", []string{"http", "https"})
fmt.Println("Is Valid:", isValid)
// Output: true
```

### Extract Domain
```go
domain, err := ExtractDomain("https://sub.example.com/path?query=value")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Domain:", domain)
}
// Output: example.com
```

### Get Query Parameter
```go
value, err := GetQueryParam("https://example.com?foo=bar&baz=qux", "foo")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Value:", value)
}
// Output: bar
```
---

## 12. Math

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

## 13. Fake

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

###  Generate a random date between 1st January 1970 and the current date
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

###  Generate a random US phone number
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

###  Generates a random US address
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

## 14. Time 

## 1. `StartOfDay`

### Get the start of the day for the given time
```go
package main

import (
	"fmt"
	"time"

	utils "github.com/kashifkhan0771/utils/time"
)

func main() {
	t := time.Now()
	fmt.Println(utils.StartOfDay(t))
}
```
#### Output:
```
2024-12-29 00:00:00 +0500 PKT
```

---

## 2. `EndOfDay`

### Get the end of the day for the given time
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    t := time.Now()
    fmt.Println(utils.EndOfDay(t))
}
```
#### Output:
```
2024-12-29 23:59:59.999999999 +0500 PKT
```

---

## 3. `AddBusinessDays`

### Add business days to a date (skipping weekends)
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    t := time.Date(2024, 12, 27, 0, 0, 0, 0, time.Local) // Friday
    // Add 3 business days
    result := utils.AddBusinessDays(t, 3)
    fmt.Println(result)
}
```
#### Output:
```
2025-01-01 00:00:00 +0500 PKT
```

---

## 4. `IsWeekend`

### Check if a given date is a weekend
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    saturday := time.Date(2024, 12, 28, 0, 0, 0, 0, time.Local)
    monday := time.Date(2024, 12, 30, 0, 0, 0, 0, time.Local)
    
    fmt.Printf("Is Saturday a weekend? %v\n", utils.IsWeekend(saturday))
    fmt.Printf("Is Monday a weekend? %v\n", utils.IsWeekend(monday))
}
```
#### Output:
```
Is Saturday a weekend? true
Is Monday a weekend? false
```

---

## 5. `TimeDifferenceHumanReadable`

### Get human-readable time difference
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    now := time.Now()
    future := now.Add(72 * time.Hour)
    past := now.Add(-48 * time.Hour)
    
    fmt.Println(utils.TimeDifferenceHumanReadable(now, future))
    fmt.Println(utils.TimeDifferenceHumanReadable(now, past))
}
```
#### Output:
```
in 3 day(s)
2 day(s) ago
```

---

## 6. `DurationUntilNext`

### Calculate duration until next specified weekday
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    now := time.Now()
    nextMonday := utils.DurationUntilNext(time.Monday, now)
    fmt.Printf("Duration until next Monday: %v\n", nextMonday)
}
```
#### Output:
```
Duration until next Monday: 24h0m0s
```

---

## 7. `ConvertToTimeZone`

### Convert time to different timezone
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    t := time.Now()
    nyTime, err := utils.ConvertToTimeZone(t, "America/New_York")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(nyTime)
}
```
#### Output:
```
2024-12-29 14:00:00 -0500 EST
```

---

## 8. `HumanReadableDuration`

### Format duration in human-readable format
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    d := 3*time.Hour + 25*time.Minute + 45*time.Second
    fmt.Println(utils.HumanReadableDuration(d))
}
```
#### Output:
```
3h 25m 45s
```

---

## 9. `CalculateAge`

### Calculate age from birthdate
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.Local)
    age := utils.CalculateAge(birthDate)
    fmt.Printf("Age: %d years\n", age)
}
```
#### Output:
```
Age: 34 years
```

---

## 10. `IsLeapYear`

### Check if a year is a leap year
```go
package main

import (
    "fmt"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    fmt.Printf("Is 2024 a leap year? %v\n", utils.IsLeapYear(2024))
    fmt.Printf("Is 2023 a leap year? %v\n", utils.IsLeapYear(2023))
}
```
#### Output:
```
Is 2024 a leap year? true
Is 2023 a leap year? false
```

---

## 11. `NextOccurrence`

### Find next occurrence of a specific time
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    now := time.Now()
    nextNoon := utils.NextOccurrence(12, 0, 0, now)
    fmt.Println("Next noon:", nextNoon)
}
```
#### Output:
```
Next noon: 2024-12-30 12:00:00 +0500 PKT
```

---

## 12. `WeekNumber`

### Get ISO year and week number
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    t := time.Now()
    year, week := utils.WeekNumber(t)
    fmt.Printf("Year: %d, Week: %d\n", year, week)
}
```
#### Output:
```
Year: 2024, Week: 52
```

---

## 13. `DaysBetween`

### Calculate days between two dates
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
    end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.Local)
    days := utils.DaysBetween(start, end)
    fmt.Printf("Days between: %d\n", days)
}
```
#### Output:
```
Days between: 365
```

---

## 14. `IsTimeBetween`

### Check if time is between two other times
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    now := time.Now()
    start := now.Add(-1 * time.Hour)
    end := now.Add(1 * time.Hour)
    fmt.Printf("Is current time between? %v\n", utils.IsTimeBetween(now, start, end))
}
```
#### Output:
```
Is current time between? true
```

---

## 15. `UnixMilliToTime`

### Convert Unix milliseconds to time
```go
package main

import (
    "fmt"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    ms := int64(1703836800000) // 2024-12-29 00:00:00
    t := utils.UnixMilliToTime(ms)
    fmt.Println(t)
}
```
#### Output:
```
2024-12-29 00:00:00 +0000 UTC
```

---

## 16. `SplitDuration`

### Split duration into components
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    d := 50*time.Hour + 30*time.Minute + 15*time.Second
    days, hours, minutes, seconds := utils.SplitDuration(d)
    fmt.Printf("Days: %d, Hours: %d, Minutes: %d, Seconds: %d\n", 
        days, hours, minutes, seconds)
}
```
#### Output:
```
Days: 2, Hours: 2, Minutes: 30, Seconds: 15
```

---

## 17. `GetMonthName`

### Get month name from number
```go
package main

import (
    "fmt"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    monthName, err := utils.GetMonthName(12)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Month 12 is: %s\n", monthName)
}
```
#### Output:
```
Month 12 is: December
```

---

## 18. `GetDayName`

### Get day name from number
```go
package main

import (
    "fmt"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    dayName, err := utils.GetDayName(1)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Day 1 is: %s\n", dayName)
}
```
#### Output:
```
Day 1 is: Monday
```

---

## 19. `FormatForDisplay`

### Format time for display
```go
package main

import (
    "fmt"
    "time"

    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    t := time.Now()
    formatted := utils.FormatForDisplay(t)
    fmt.Println(formatted)
}
```
#### Output:
```
Sunday, 29 Dec 2024
```

---

## 20. `IsToday`

### Check if date is today
```go
package main

import (
    "fmt"
    "time"
	
    utils "github.com/kashifkhan0771/utils/time"
)

func main() {
    now := time.Now()
    tomorrow := now.AddDate(0, 0, 1)
    
    fmt.Printf("Is now today? %v\n", utils.IsToday(now))
    fmt.Printf("Is tomorrow today? %v\n", utils.IsToday(tomorrow))
}
```
#### Output:
```
Is now today? true
Is tomorrow today? false
```

---

## 15. Logging

### Create and use a logger
```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Create a new logger with prefix "MyApp", minimum level INFO, and output to stdout
	logger := logging.NewLogger("MyApp", logging.INFO, os.Stdout)

	// Log messages of different levels
	logger.Debug("This is a debug message.") // Ignored because minLevel is INFO
	logger.Info("Application started.")      // Printed with blue color
	logger.Warn("Low disk space.")           // Printed with yellow color
	logger.Error("Failed to connect to DB.") // Printed with red color
}
```
#### Output:
```
[2025-01-09 12:34:56] [INFO] MyApp: Application started.
[2025-01-09 12:34:56] [WARN] MyApp: Low disk space.
[2025-01-09 12:34:56] [ERROR] MyApp: Failed to connect to DB.
```

### Log without colors (useful for plain text logs)
```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Create a logger and disable colors
	logger := logging.NewLogger("MyApp", logging.DEBUG, os.Stdout)
	logger.disableColors = true

	// Log messages of different levels
	logger.Debug("Debugging without colors.")
	logger.Info("Information without colors.")
	logger.Warn("Warning without colors.")
	logger.Error("Error without colors.")
}
```
#### Output:
```
[2025-01-09 12:34:56] [DEBUG] MyApp: Debugging without colors.
[2025-01-09 12:34:56] [INFO] MyApp: Information without colors.
[2025-01-09 12:34:56] [WARN] MyApp: Warning without colors.
[2025-01-09 12:34:56] [ERROR] MyApp: Error without colors.
```

### Log messages to a file
```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Open a log file for writing
	file, err := os.Create("app.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a logger with file output
	logger := logging.NewLogger("MyApp", logging.DEBUG, file)

	// Log messages
	logger.Debug("Writing debug logs to file.")
	logger.Info("Application log stored in file.")
	logger.Warn("This is a warning.")
	logger.Error("This is an error.")
}
```
#### Output (in `app.log` file):
```
[2025-01-09 12:34:56] [DEBUG] MyApp: Writing debug logs to file.
[2025-01-09 12:34:56] [INFO] MyApp: Application log stored in file.
[2025-01-09 12:34:56] [WARN] MyApp: This is a warning.
[2025-01-09 12:34:56] [ERROR] MyApp: This is an error.
```

### Filter logs by minimum log level
```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Create a logger with minimum level WARN
	logger := logging.NewLogger("MyApp", logging.WARN, os.Stdout)

	// Log messages
	logger.Debug("This is a debug message.") // Ignored
	logger.Info("This is an info message.")  // Ignored
	logger.Warn("This is a warning.")        // Printed
	logger.Error("This is an error.")        // Printed
}
```
#### Output:
```
[2025-01-09 12:34:56] [WARN] MyApp: This is a warning.
[2025-01-09 12:34:56] [ERROR] MyApp: This is an error.
```

### Customize log prefixes
```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Create a logger with a custom prefix
	logger := logging.NewLogger("CustomPrefix", logging.INFO, os.Stdout)

	// Log messages
	logger.Info("This message has a custom prefix.")
}
```
#### Output:
```
[2025-01-09 12:34:56] [INFO] CustomPrefix: This message has a custom prefix.
```

## 16. Fsutils

### Format a file size given in bytes into a human-readable format
```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	sizes := []int64{0, 512, 1024, 1048576, 1073741824, 1099511627776}

	for _, size := range sizes {
		fmt.Println(fsutils.FormatFileSize(size))
	}
}
```
#### Output:
```
0 B
512 B
1.00 KB
1.00 MB
1.00 GB
1.00 TB
```

### Search for files with the specified extension
```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	dir := "/path/to/your/dir"

	txtFiles, err := fsutils.FindFiles(dir, ".txt")
	if err != nil {
		log.Fatalf("Error finding .txt files: %v", err)
	}
	
	fmt.Println("TXT Files:", txtFiles)

	logFiles, err := fsutils.FindFiles(dir, ".log")
	if err != nil {
		log.Fatalf("Error finding .log files: %v", err)
	}

	fmt.Println("LOG Files:", logFiles)

	allFiles, err := fsutils.FindFiles(dir, "")
	if err != nil {
		log.Fatalf("Error finding all files: %v", err)
	}

	fmt.Println("All Files:", allFiles)
}

```
#### Output:
```
TXT Files: [/path/to/your/dir/file1.txt /path/to/your/dir/file2.txt /path/to/your/dir/file4.txt]
LOG Files: [/path/to/your/dir/file3.log]
All Files: [/path/to/your/dir/file1.txt /path/to/your/dir/file2.txt /path/to/your/dir/file3.log /path/to/your/dir/file4.txt]
```

### Calculate the total size (in bytes) of all files in a directory
```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	dir := "/path/to/your/dir"

	size, err := fsutils.GetDirectorySize(dir)
	if err != nil {
		log.Fatalf("Error calculating directory size: %v", err)
	}

	fmt.Printf("The total size of directory \"%s\" is %dB\n", dir, size)
}

```
#### Output:
```
The total size of directory "/path/to/your/dir" is 6406B
```

### Compare two files
```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	file1 := "/path/to/your/file1.txt"
	file2 := "/path/to/your/file2.txt"

	identical, err := fsutils.FilesIdentical(file1, file2)
	if err != nil {
		log.Fatalf("Error comparing files: %v", err)
	}

	if identical {
		fmt.Printf("The files %s and %s are identical\n", file1, file2)
	} else {
		fmt.Printf("The files %s and %s are not identical\n", file1, file2)
	}
}

```
#### Output:
```
The files /path/to/your/file1.txt and /path/to/your/file2.txt are identical
```

### Compare two directories
```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	dir1 := "/path/to/your/dir1"
	dir2 := "/path/to/your/dir2"

	identical, err := fsutils.DirsIdentical(dir1, dir2)
	if err != nil {
		log.Fatalf("Error comparing directories: %v", err)
	}

	if identical {
		fmt.Printf("The directories %s and %s are identical.\n", dir1, dir2)
	} else {
		fmt.Printf("The directories %s and %s are not identical.\n", dir1, dir2)
	}
}

```
#### Output:
```
The directories /path/to/your/dir1 and /path/to/your/dir2 are identical.
```

### Get File Metadata
```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	file := "example.txt"
	metadata, err := fsutils.GetFileMetadata(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf(
		"Name: %s, Size: %d, IsDir: %t, ModTime: %s, Mode: %v, Path: %s, Ext: %s, Owner: %s\n",
		metadata.Name, metadata.Size, 
		metadata.IsDir, metadata.ModTime.String(), 
		metadata.Mode, metadata.Path, 
		metadata.Ext, metadata.Owner,
	)
}

```
#### Output:
```
Name: example.txt, Size: 172, IsDir: false, ModTime: 2025-01-20 15:03:00.189199994 +0100 CET, Mode: -rw-rw-r--, Path: /path/to/your/dir/example.txt, Ext: .txt, Owner: owner
```

### Get Directory Metadata
```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	dir := "example/"
	metadata, err := fsutils.GetFileMetadata(dir)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf(
		"Name: %s, Size: %d, IsDir: %t, ModTime: %s, Mode: %v, Path: %s, Ext: %s, Owner: %s\n",
		metadata.Name, metadata.Size, 
		metadata.IsDir, metadata.ModTime.String(), 
		metadata.Mode, metadata.Path, 
		metadata.Ext, metadata.Owner,
	)
}

```
#### Output:
```
Name: example, Size: 4096, IsDir: true, ModTime: 2025-01-20 15:06:23.057206656 +0100 CET, Mode: drwxrwxr-x, Path: /path/to/your/dir/example, Ext: , Owner: owner
```

### Marshal File's Metadata to JSON
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/fsutils"
)

func main() {
	file := "example.txt"
	metadata, err := fsutils.GetFileMetadata(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	json, err := json.Marshal(&metadata)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(json))
}

```
#### Output:
```json
{
   "name": "example.txt",
   "size": 172,
   "is_dir": false,
   "mod_time": "2025-01-20T15:06:34.812677487+01:00",
   "mode": 436,
   "path": "/path/to/your/dir/example.txt",
   "ext": ".txt",
}
```
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
