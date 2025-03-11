### File System Utilities

- **FormatFileSize**: Formats a file size given in bytes into a human-readable string with appropriate units (B, KB, MB, GB, TB).

- **FindFiles**: Searches for files with the specified extension in the given root directory and returns a slice of matching file paths.

- **GetDirectorySize**: Calculates the total size (in bytes) of all files within the specified directory.

- **FilesIdentical**: Compares two files byte by byte to determine if they are identical.

- **DirsIdentical**: Compares two directories to determine if they are identical.

- **GetFileMetadata**: Retrieves metadata for a specified file path. Returns a `FileMetadata` struct that can be marshaled to JSON.

### 17. Caching

The caching package provides utilities for creating caching decorators to enhance the performance of functions by storing computed results. It includes both thread-safe and non-thread-safe implementations.

- **SafeCacheWrapper**: A thread-safe caching decorator that safely memoizes function results in concurrent environments.

  - Uses `sync.Map` to ensure thread-safety
  - Caches all results indefinitely (no eviction)
  - Best suited for pure functions with limited input domains
  - Safe for concurrent access but may impact performance under high contention

- **CacheWrapper**: A non-thread-safe caching decorator that memoizes function results.
  - Caches all results indefinitely (no eviction)
  - Best suited for pure functions with limited input domains
  - Not safe for concurrent access
  - Use SafeCacheWrapper for concurrent scenarios

---

## Examples:

For examples of each function, please checkout [EXAMPLES.md](/fsutils/EXAMPLES.md)

---
