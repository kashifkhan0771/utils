<div align="center">
  <img src="/assets/logo.png" width="75%"/>
</div>

# Utils

A common utilities library for Go.

## Overview

**Utils** is a lightweight, flexible, and reusable library that provides utility functions and helpers for common operations in Go applications. It includes packages for boolean conversions, context value handling, map operations, slice utilities, string manipulations, struct comparisons, and more—enhancing Go projects with optimized, clean, and practical solutions.

## Prerequisites

- **Go**: The project is written in Golang. Ensure you have Go installed (preferably Go 1.23.3 or later). You can download it [here](https://go.dev/doc/install).
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
| **sort**      | Sorting algorithms                                 | [README](sort/README.md)      | [EXAMPLES](sort/EXAMPLES.md)      |
| **strings**   | String manipulation and encoding utilities         | [README](strings/README.md)   | [EXAMPLES](strings/EXAMPLES.md)   |
| **structs**   | Struct comparison utilities                        | [README](structs/README.md)   | [EXAMPLES](structs/EXAMPLES.md)   |
| **templates** | Template rendering utilities                       | [README](templates/README.md) | [EXAMPLES](templates/EXAMPLES.md) |
| **timeutils** | Time and date manipulation utilities               | [README](time/README.md)      | [EXAMPLES](time/EXAMPLES.md)      |
| **url**       | URL parsing and manipulation utilities             | [README](url/README.md)       | [EXAMPLES](url/EXAMPLES.md)       |
| **slugger**       | A simple and efficient way to generate URL-friendly slugs from strings            | [README](slugger/README.md)       | [EXAMPLES](slugger/EXAMPLES.md)       |

## Contributions

Contributions are welcome! If you'd like to contribute, feel free to open a pull request.

Before submitting a PR, please review the [Contribution Guide](/CONTRIBUTING.md).

Together, we can make **Utils** even better for the Go community!

## Credits

The image used in this project was sourced from **https://github.com/MariaLetta/free-gophers-pack**.

📷 Image by **[MariaLetta](https://github.com/MariaLetta)**, used under the **[Creative Commons (CC0-1.0)](https://github.com/MariaLetta/free-gophers-pack?tab=CC0-1.0-1-ov-file) license.**
