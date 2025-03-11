## Templates Function Examples

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
  <head>
    <title>Template Rendering Demo</title>
  </head>
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
