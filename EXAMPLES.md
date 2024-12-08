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

## 1. `Abs`

### Calculate the absolute value of a number
```go
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

## 2. `Sign`

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

## 3. `Min`

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

## 4. `Max`

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

## 5. `Clamp`

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

## 6. `IntPow`

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

## 7. `IsEven`

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

## 8. `IsOdd`

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

## 9. `Swap`

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

## 10. `Factorial`

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

## 11. `GCD`

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

## 12. `LCM`

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
