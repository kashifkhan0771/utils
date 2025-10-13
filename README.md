<div align="center">
  <img src="/assets/logo.png" width="75%"/>
</div>

# Utils

A common utilities library for Go.

## Overview

**Utils** is a lightweight, flexible, and reusable library that provides utility functions and helpers for common operations in Go applications. It includes packages for boolean conversions, context value handling, map operations, slice utilities, string manipulations, struct comparisons, and more—enhancing Go projects with optimized, clean, and practical solutions.

## Prerequisites

- **Go**: The project is written in Golang. Ensure you have Go installed (preferably Go 1.24 or later). You can download it [here](https://go.dev/doc/install).
- **Git**: Required for cloning the repository.

## Installation

To use **Utils** in your project, add it as a module dependency:

### Install the Package

```sh
go get github.com/kashifkhan0771/utils
```

Alternatively, include it directly in your `go.mod` file (use the latest release):

```sh
require github.com/kashifkhan0771/utils v0.3.0
```

### Clone the Repository (For Development)

If you want to contribute or modify the library, clone the repository:

```sh
git clone https://github.com/kashifkhan0771/utils.git
cd utils
```

## Key Features

### Utility Packages

| Package Name  | Description                                        | Documentation                 | Examples                          |
| ------------- | -------------------------------------------------- | ----------------------------- | --------------------------------- |
| **boolean**   | Utilities for boolean value checking and toggling  | [README](boolean/README.md)   | [EXAMPLES](boolean/EXAMPLES.md)   |
| **caching**   | Cache management utilities                         | [README](caching/README.md)   | [EXAMPLES](caching/EXAMPLES.md)   |
| **cryptoutils** | A set of cryptographic utility functions for various cryptographic operations            | [README](cryptoutils/README.md)       | [EXAMPLES](cryptoutils/EXAMPLES.md)       |
| **ctxutils**  | Context utilities                                  | [README](ctxutils/README.md)  | [EXAMPLES](ctxutils/EXAMPLES.md)  |
| **errutils**  | Error aggregation and management utilities         | [README](errutils/README.md)  | [EXAMPLES](errutils/EXAMPLES.md)  |
| **fake**      | Fake data generation (UUIDs, addresses, dates)     | [README](fake/README.md)      | [EXAMPLES](fake/EXAMPLES.md)      |
| **fsutils**   | File system utilities (size, metadata, comparison) | [README](fsutils/README.md)   | [EXAMPLES](fsutils/EXAMPLES.md)   |
| **logging**   | Flexible logging system for Golang                 | [README](logging/README.md)   | [EXAMPLES](logging/EXAMPLES.md)   |
| **maps**      | Utilities for state and metadata maps              | [README](maps/README.md)      | [EXAMPLES](maps/EXAMPLES.md)      |
| **math**      | Mathematical utilities and helpers                 | [README](math/README.md)      | [EXAMPLES](math/EXAMPLES.md)      |
| **pointers**  | Helper functions for working with pointer values   | [README](pointers/README.md)  | [EXAMPLES](pointers/EXAMPLES.md)  |
| **rand**      | Random number and string generation utilities      | [README](rand/README.md)      | [EXAMPLES](rand/EXAMPLES.md)      |
| **slice**     | Slice manipulation and de-duplication utilities    | [README](slice/README.md)     | [EXAMPLES](slice/EXAMPLES.md)     |
| **slugger**   | A simple and efficient way to generate URL-friendly slugs from strings             | [README](slugger/README.md)       | [EXAMPLES](slugger/EXAMPLES.md)       |
| **sort**      | Sorting algorithms                                 | [README](sort/README.md)      | [EXAMPLES](sort/EXAMPLES.md)      |
| **stack**     | Stack data structure                               | [README](stack/README.md)     | [EXAMPLES](stack/EXAMPLES.md)     |
| **strings**   | String manipulation and encoding utilities         | [README](strings/README.md)   | [EXAMPLES](strings/EXAMPLES.md)   |
| **structs**   | Struct comparison utilities                        | [README](structs/README.md)   | [EXAMPLES](structs/EXAMPLES.md)   |
| **templates** | Template rendering utilities                       | [README](templates/README.md) | [EXAMPLES](templates/EXAMPLES.md) |
| **timeutils** | Time and date manipulation utilities               | [README](time/README.md)      | [EXAMPLES](time/EXAMPLES.md)      |
| **url**       | URL parsing and manipulation utilities             | [README](url/README.md)       | [EXAMPLES](url/EXAMPLES.md)       |
| **conversion** | Conversion of data types, time, and temperatues   | [README](conversion/README.md) | [EXAMPLES](conversion/EXAMPLES.md)
| **ratelimiter** | Token-bucket rate limiter (allow/wait, adjustable capacity & refill rate) | [README](ratelimiter/README.md) | [EXAMPLES](ratelimiter/EXAMPLES.md) |
| **queue** | Queue data structure| [README](queue/README.md) | [EXAMPLES](queue/EXAMPLES.md) |
| **image** | Utility functions for working with images | [README](image/README.md) | [EXAMPLES](image/EXAMPLES.md) |

## Comparison

| Feature / Utility Area     | `kashifkhan0771/utils`      | `go-commons-lang`                 | `gookit/goutil`                   |
|---------------------------|-----------------------------|----------------------------------|----------------------------------|
| Boolean Utilities         | ✅ Yes                      | ✅ Yes                           | ✅ Yes                           |
| Caching                   | ✅ Yes                      | ❌ No                            | ✅ Yes                           |
| Cryptographic Utilities   | ✅ Yes                      | ✅ Yes                           | ✅ Yes                           |
| Context Utilities         | ✅ Yes                      | ❌ No                            | ❌ No                            |
| Error Aggregation         | ✅ Yes                      | ❌ No                            | ✅ Yes                           |
| Fake Data (UUID, etc.)    | ✅ Yes                      | ✅ Yes                           | ✅ Yes                           |
| Filesystem Utilities      | ✅ Yes                      | ✅ Yes                           | ✅ Yes                           |
| Logging                   | ✅ Yes                      | ❌ No                            | ✅ Yes                           |
| Maps / Metadata Helpers   | ✅ Yes                      | ✅ Yes                           | ✅ Yes                           |
| Math Utilities            | ✅ Yes                      | ✅ Yes                           | ✅ Yes                           |
| Pointer Utilities         | ✅ Yes                      | ❌ No                            | ❌ No                            |
| Conversion Utilities      | ✅ Yes                      | ❌ No                            | ❌ No                            |
| Random Utilities          | ✅ Yes                      | ✅ Yes                           | ✅ Yes                           |
| Slice Utilities           | ✅ Yes                      | ✅ Yes                           | ✅ Yes                           |
| Slugify                   | ✅ Yes                      | ❌ No                            | ✅ Yes                           |
| Sorting                   | ✅ Yes                      | ✅ Yes                           | ✅ Yes                           |
| String Utilities          | ✅ Yes                      | ✅ Yes                           | ✅ Yes                           |
| Struct Comparison         | ✅ Yes                      | ❌ No                            | ✅ Yes                           |
| Template Helpers          | ✅ Yes                      | ❌ No                            | ✅ Yes                           |
| Time Utilities            | ✅ Yes                      | ✅ Yes                           | ✅ Yes                           |
| URL Utilities             | ✅ Yes                      | ❌ No                            | ✅ Yes                           |
| Dependency-Free           | ✅ Yes                      | ❌ No                            | ❌ No                            |
| Small API Surface         | ✅ Yes                      | ❌ No                            | ✅ Yes                           |
| Rate Limiter Utilities    | ✅ Yes                      | ❌ No                            | ❌ No                            |
| Image                     | ✅ Yes                      | ❌ No                            | ❌ No                            |


## Contributions

Contributions are welcome! If you'd like to contribute, feel free to open a pull request.

Before submitting a PR, please review the [Contribution Guide](/CONTRIBUTING.md).

Together, we can make **Utils** even better for the Go community!

## Credits

The image used in this project was sourced from **https://github.com/MariaLetta/free-gophers-pack**.

📷 Image by **[MariaLetta](https://github.com/MariaLetta)**, used under the **[Creative Commons (CC0-1.0)](https://github.com/MariaLetta/free-gophers-pack?tab=CC0-1.0-1-ov-file) license.**

## Contributors

Powered by coffee, code, and these legends ☕💻:

<a href="https://github.com/kashifkhan0771/utils/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=kashifkhan0771/utils" />
</a>

## Star History

<a href="https://www.star-history.com/#kashifkhan0771/utils&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=kashifkhan0771/utils&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=kashifkhan0771/utils&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=kashifkhan0771/utils&type=Date" />
 </picture>
</a>
