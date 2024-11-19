# Utils
Common Utilities library for Go

## Overview
Utils is a lightweight, flexible, and reusable library providing utility functions and helpers for common operations in Go applications. With packages designed for managing boolean conversions, handling context values, map operations, slice utilities, string manipulations, and struct comparison, utils enhances Go projects with optimized, clean, and practical solutions.

### Prerequisites
**Go**: The project is written in Golang, so you'll need Go installed (preferably Go 1.16 or later). You can download and install Go from [here](https://go.dev/doc/install).

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

### 1. Boolean
- **IsTrue**: Checks if the provided string represents a true value (e.g., "T", "1", "TRUE").

### 2. Context (ctxutils)
- **SetStringValue**: Sets a string value in the context with a custom key.
- **GetStringValue**: Retrieves a string value from the context using the specified key.
- **SetIntValue**: Sets an integer value in the context with a custom key.
- **GetIntValue**: Retrieves an integer value from the context using the specified key.

### 3. Errors (errutils)
- **NewErrorAggregator**: Creates a new error aggregator to collect multiple errors.
- **Add**: Adds an error to the aggregator.
- **Error**: Retrieves the aggregated error message.
- **HasErrors**: Checks if there are any errors in the aggregator.

### 4. Maps
- **NewStateMap**: Creates a new map for managing state flags.
- **SetState**: Sets a specific state to true or false.
- **IsState**: Checks the current value of a state.
- **ToggleState**: Toggles the value of a state.
- **HasState**: Checks if a state exists in the map.
- **NewMetadata**: Creates a new metadata map for key-value storage.
- **Update**: Updates or adds a key-value pair to the metadata map.
- **Value**: Retrieves the value of a key from the metadata map.
- **Has**: Checks if a key exists in the metadata map.

### 5. Pointers
- **DefaultIfNil**: Returns a default value if the pointer is nil.
- **NullableBool**: Returns the value of a boolean pointer or false if nil.
- **NullableTime**: Returns the value of a time pointer or a zero time value if nil.
- **NullableInt**: Returns the value of an integer pointer or zero if nil.
- **NullableString**: Returns the value of a string pointer or an empty string if nil.

### 6. Random (rand)
- **Number**: Generates a random number.
- **NumberInRange**: Generates a random number within a specified range.
- **String**: Generates a random alphanumeric string.
- **StringWithLength**: Generates a random string of a custom length.
- **StringWithCharset**: Generates a random string using a custom character set.
- **Pick**: Picks a random element from a given slice.
- **Shuffle**: Shuffles the elements of a slice randomly.

### 7. Slice
- **RemoveDuplicateStr**: Removes duplicate values from a slice of strings.
- **RemoveDuplicateInt**: Removes duplicate values from a slice of integers.

### 8. Strings
- **SubstringSearch**: Searches for substrings in a string with customizable options.
- **Title**: Converts a string to title case (capitalizes the first letter of each word).
- **ToTitle**: Converts a string to title case, ignoring specified words.
- **Tokenize**: Splits a string into tokens using specified delimiters.
- **Rot13Encode**: Encodes a string using the ROT13 cipher.
- **Rot13Decode**: Decodes a string using the ROT13 cipher.
- **CaesarEncrypt**: Encrypts a string using the Caesar cipher with a given shift.
- **CaesarDecrypt**: Decrypts a string encrypted using the Caesar cipher.
- **IsValidEmail**: Checks if a string is a valid email address.
- **SanitizeEmail**: Removes leading and trailing spaces from an email string.
- **Reverse**: Reverses the characters in a string.
- **CommonPrefix**: Finds the longest common prefix of a set of strings.
- **CommonSuffix**: Finds the longest common suffix of a set of strings.

### 9. Structs
- **CompareStructs**: Compares two structs and returns the differences between them.

### 10. Templates
- **RenderHTMLTemplate**: Renders an HTML template with dynamic values.
- **RenderText**: Renders a text template with dynamic values.

Here are the one-line descriptions for the URL-related code examples:

---

### 11. URLs
- **Parse a URL**: Parses a URL into its components like scheme, host, path, query, and fragment.

- **Build a URL from Components**: Constructs a complete URL by combining base URL, path, and query parameters.

- **Resolve Relative URLs**: Resolves a relative URL against a base URL to form an absolute URL.

- **URL Encoding**: Encodes URL components to ensure they are properly formatted for inclusion in URLs.

---
## Examples:

For examples of each function, please checkout [EXAMPLES.md](/EXAMPLES.md)

---


# Contributions
Contributions to this project are welcome! If you would like to contribute, please feel free to open a PR.

Please read the [Contribution Guide](/CONTRIBUTING.md) before opening any new pull request

Together, we can make Utils even better for the Go community!
