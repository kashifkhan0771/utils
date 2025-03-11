<div align="center">
  <img src="/assets/logo.png" width="75%"/>
</div>

# Utils

Common Utilities library for Go

## Overview

Utils is a lightweight, flexible, and reusable library providing utility functions and helpers for common operations in Go applications. With packages designed for managing boolean conversions, handling context values, map operations, slice utilities, string manipulations, and struct comparison, utils enhances Go projects with optimized, clean, and practical solutions.

### Prerequisites

**Go**: The project is written in Golang, so you'll need Go installed (preferably Go 1.23.3 or later). You can download and install Go from [here](https://go.dev/doc/install).

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

## Utility Packages

Here’s the completed table with descriptions, documentation, and examples for **errutils** and any missing parts:

| Package Name  | Description                                       | Documentation                 | Examples                          |
| ------------- | ------------------------------------------------- | ----------------------------- | --------------------------------- |
| **boolean**   | Utilities for boolean value checking and toggling | [README](boolean/README.md)   | [EXAMPLES](boolean/EXAMPLES.md)   |
| **caching**   | Cache management utilities                        | coming soon                   | coming soon                       |
| **ctxutils**  | Context utilities                                 | [README](ctxutils/README.md)  | [EXAMPLES](ctxutils/EXAMPLES.md)  |
| **errutils**  | Error aggregation and management utilities        | [README](errutils/README.md)  | [EXAMPLES](errutils/EXAMPLES.md)  |
| **maps**      | Utilities for state and metadata maps             | [README](maps/README.md)      | [EXAMPLES](maps/EXAMPLES.md)      |
| **pointers**  | Helper functions for working with pointer values  | [README](pointers/README.md)  | [EXAMPLES](pointers/EXAMPLES.md)  |
| **rand**      | Random number and string generation utilities     | [README](rand/README.md)      | [EXAMPLES](rand/EXAMPLES.md)      |
| **slice**     | Slice manipulation and de-duplication utilities   | [README](slice/README.md)     | [EXAMPLES](slice/EXAMPLES.md)     |
| **strings**   | String manipulation and encoding utilities        | [README](strings/README.md)   | [EXAMPLES](strings/EXAMPLES.md)   |
| **structs**   | Struct comparison utilities                       | [README](structs/README.md)   | [EXAMPLES](structs/EXAMPLES.md)   |
| **templates** | Template rendering utilities                      | [README](templates/README.md) | [EXAMPLES](templates/EXAMPLES.md) |
| **url**       | URL parsing and manipulation utilities            | [README](url/README.md)       | [EXAMPLES](url/EXAMPLES.md)       |
| **math**      | Mathematical utilities and helpers                | [README](math/README.md)      | [EXAMPLES](math/EXAMPLES.md)      |
| **fake**      | Fake data generation (UUIDs, addresses, dates)    | [README](fake/README.md)      | [EXAMPLES](fake/EXAMPLES.md)      |
| **timeutils** | Time and date manipulation utilities              | [README](time/README.md)      | [EXAMPLES](time/EXAMPLES.md)      |
| **logging**   | Flexible logging system for Golang                | [README](logging/README.md)   | [EXAMPLES](logging/EXAMPLES.md)   |

### 16. File System Utilities

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

For examples of each function, please checkout [EXAMPLES.md](/EXAMPLES.md)

---

# Contributions

Contributions to this project are welcome! If you would like to contribute, please feel free to open a PR.

Please read the [Contribution Guide](/CONTRIBUTING.md) before opening any new pull request

Together, we can make Utils even better for the Go community!

---

# Credits

The image used in this project was sourced from **https://github.com/MariaLetta/free-gophers-pack**.  
📷 Image by **[MariaLetta](https://github.com/MariaLetta)**, used under **[Creative Commons(CCO-1.0)](https://github.com/MariaLetta/free-gophers-pack?tab=CC0-1.0-1-ov-file) license.**.
