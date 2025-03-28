## Strings Function Examples

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

### Truncate

```go
func main() {
	short := Truncate("This is a long string that needs to be truncated", &TruncateOptions{Length: 20, Omission: "°°°"})
	fmt.Println(short) // Output: This is a long strin°°°

	defaultShort := Truncate("Short example", nil)
fmt.Println(defaultShort) // Output: Short examp...
}
```

---
