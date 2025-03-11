## URLs Function Examples

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
