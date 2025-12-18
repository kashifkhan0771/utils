## Examples:


### Open URL in Default Browser

```go
err := browser.OpenURL("https://example.com")
if err != nil {
    fmt.Println("Error opening browser:", err)
} else {
    fmt.Println("Browser opened successfully")
}
```

### Open URL with Path

```go
err := browser.OpenURL("https://example.com/path/to/page")
if err != nil {
    fmt.Println("Error:", err)
}
```

### Open URL with Query Parameters

```go
err := browser.OpenURL("https://example.com/search?q=golang&page=1")
if err != nil {
    fmt.Println("Error:", err)
}
```

### Open Local File

```go
// Open a local HTML file
err := browser.OpenURL("file:///path/to/file.html")
if err != nil {
    fmt.Println("Error:", err)
}
```

---


