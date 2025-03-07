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

### 1. Boolean
- **IsTrue**: Checks if the provided string represents a true value (e.g., "T", "1", "TRUE").
- **Toggle**: Toggles the given boolean value, returning its negation. (e.g., `true` becomes `false` and vice versa).
- **AllTrue**: Checks if all the values in a slice of booleans are `true`. Returns `false` if the slice is empty.
- **AnyTrue**: Checks if at least one value in a slice of booleans is `true`. Returns `false` if the slice is empty.
- **NoneTrue**: Checks if none of the values in a slice of booleans are `true`. Returns `true` if the slice is empty.
- **CountTrue**: Counts the number of `true` values in a slice of booleans. Returns `0` for an empty slice.
- **CountFalse**: Counts the number of `false` values in a slice of booleans. Returns `0` for an empty slice.
- **Equal**: Checks if all the values in a variadic boolean argument are equal. Returns `true` if the slice contains only one or no elements.
- **And**: Performs a logical AND operation on all the values in a slice of booleans. Returns `true` only if all values are `true`. Returns `false` for an empty slice.
- **Or**: Performs a logical OR operation on all the values in a slice of booleans. Returns `true` if at least one value is `true`. Returns `false` for an empty slice.


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
- **RunLengthEncode** Encodes a string using Run-length-encoding.
- **RunLengthDecode** Decodes a string that has been encoded using Run-length-encoding. Ruturns the original string and an error if the encoding failed

### 9. Structs
- **CompareStructs**: Compares two structs and returns the differences between them.

### 10. Templates
- **RenderHTMLTemplate**: Renders an HTML template with dynamic values.
- **RenderText**: Renders a text template with dynamic values.

### 11. URLs
- **Parse a URL**: Parses a URL into its components like scheme, host, path, query, and fragment.

- **Build a URL from Components**: Constructs a complete URL by combining base URL, path, and query parameters.

- **Resolve Relative URLs**: Resolves a relative URL against a base URL to form an absolute URL.

- **URL Encoding**: Encodes URL components to ensure they are properly formatted for inclusion in URLs.

### 12. Math
- **Abs**: Returns the absolute value of the given number. Works for both integers and floating-point numbers. If the input is negative, it returns its positive equivalent; otherwise, it returns the number as is.

- **Sign**: Determines the sign of a signed number. Returns `1` if the number is positive, `-1` if negative, and `0` if the number is zero.

- **Min**: Returns the smaller of two numbers. Works for all number types including integers and floating-point numbers.

- **Max**: Returns the larger of two numbers. Works for all number types including integers and floating-point numbers.

- **Clamp**: Restricts a given value to be within a specified range. If the value is below the minimum, it returns the minimum; if above the maximum, it returns the maximum.

- **IntPow**: Calculates base raised to the power of exp. Supports both positive and negative exponents. Returns float64 for fractional results.

- **IsEven**: Checks if the given integer is even. Returns `true` for even numbers and `false` otherwise.

- **IsOdd**: Checks if the given integer is odd. Returns `true` for odd numbers and `false` otherwise.

- **Swap**: Swaps the values of two variables in place. It uses pointers to modify the original variables.

- **Factorial**: Computes the factorial of a non-negative integer. Factorial of `n` is defined as the product of all integers from `1` to `n`. For `0` and `1`, the result is `1`. Factorial returns an error on invalid input.

- **GCD**: Finds the greatest common divisor (GCD) of two integers using the Euclidean algorithm. If one of the inputs is `0`, the other input is returned.

- **LCM**: Finds the least common multiple (LCM) of two integers.

### 13. Fake
- **RandomUUID**: Generates a random UUID of version 4 and variant 2.

- **RandomDate**: Generates a random date between 1st January 1970 and the current date.

- **RandomPhoneNumber**: Generates a random US phone number in the format (XXX) XXX-XXXX.

- **RandomAddress**: Generates a random US address including street, city, state, and ZIP code.

### 14. Time Utilities

- **StartOfDay**: Returns a time.Time set to the beginning (00:00:00) of the given day.

- **EndOfDay**: Returns a time.Time set to the last moment (23:59:59.999999999) of the given day.

- **AddBusinessDays**: Adds specified number of business days (excluding weekends) to a given date.

- **IsWeekend**: Determines if a given date falls on a weekend (Saturday or Sunday).

- **TimeDifferenceHumanReadable**: Converts time difference between two dates into a human-friendly string (e.g., "in 2 days" or "3 hours ago").

- **DurationUntilNext**: Calculates duration until the next occurrence of a specific weekday.

- **ConvertToTimeZone**: Converts a time.Time to a different timezone based on location name.

- **HumanReadableDuration**: Formats a duration into a human-readable string (e.g., "2h 30m 45s").

- **CalculateAge**: Computes age in years given a birth date.

- **IsLeapYear**: Checks if a given year is a leap year.

- **NextOccurrence**: Finds the next occurrence of a specific time on the same or next day.

- **WeekNumber**: Returns the ISO year and week number for a given date.

- **DaysBetween**: Calculates the number of days between two dates.

- **IsTimeBetween**: Checks if a time falls between two other times.

- **UnixMilliToTime**: Converts Unix milliseconds timestamp to time.Time.

- **SplitDuration**: Breaks down a duration into days, hours, minutes, and seconds.

- **GetMonthName**: Returns the English name of a month given its number (1-12).

- **GetDayName**: Returns the English name of a day given its number (0-6).

- **FormatForDisplay**: Formats a date in a readable format (e.g., "Monday, 2 Jan 2006").

- **IsToday**: Checks if a given date is today.

### 15. Logging

The `logging` package provides a simple, flexible, and color-coded logging system for Golang. Below are the main features and methods of the package:

#### **Logger Constructor**
- **`NewLogger(prefix string, minLevel LogLevel, output io.Writer) *Logger`**:  
  Creates a new logger instance.  
  - **`prefix`**: A string prefix added to all log messages.  
  - **`minLevel`**: The minimum log level to output (`DEBUG`, `INFO`, `WARN`, `ERROR`). Messages below this level are ignored.  
  - **`output`**: The destination for log output (e.g., `os.Stdout`, `os.Stderr`, or any `io.Writer`). Defaults to `os.Stdout` if `nil`.  

#### **Log Levels**
- **DEBUG**: Used for detailed debug information.
- **INFO**: General informational messages.
- **WARN**: Warnings about potential issues.
- **ERROR**: Critical errors.

#### **Logger Methods**
- **`Info(message string)`**:  
  Logs an informational message with the log level **INFO**.  

- **`Debug(message string)`**:  
  Logs a debug message with the log level **DEBUG**.  

- **`Warn(message string)`**:  
  Logs a warning message with the log level **WARN**.  

- **`Error(message string)`**:  
  Logs an error message with the log level **ERROR**.  

#### **Key Features**
- **Color-Coded Logs**:  
  Log messages are color-coded based on the log level:  
  - **DEBUG**: Green  
  - **INFO**: Blue  
  - **WARN**: Yellow  
  - **ERROR**: Red  

- **Timestamped Logs**:  
  Each log message includes a timestamp in the format `YYYY-MM-DD HH:MM:SS`.  

- **Log Filtering by Level**:  
  Logs below the specified minimum level (`minLevel`) are ignored.  

- **Custom Output**:  
  Logs can be directed to any `io.Writer`, allowing flexible output destinations (e.g., files, network connections).  

- **Disable Colors**:  
  The `disableColors` field in the `Logger` struct can be set to `true` to disable color codes (useful for testing or plain-text logs).  

#### **Notes**
- If the `minLevel` is set to `DEBUG`, all log messages will be displayed.
- Logs are automatically flushed to the configured output as soon as they're written.  
- To log without colors (e.g., for testing), set the `disableColors` field to `true` in the `Logger` instance.

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
    * Uses `sync.Map` to ensure thread-safety
    * Caches all results indefinitely (no eviction)
    * Best suited for pure functions with limited input domains
    * Safe for concurrent access but may impact performance under high contention

  - **CacheWrapper**: A non-thread-safe caching decorator that memoizes function results. 
    * Caches all results indefinitely (no eviction)
    * Best suited for pure functions with limited input domains
    * Not safe for concurrent access
    * Use SafeCacheWrapper for concurrent scenarios

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
