### Logging

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

### Redaction

The `logging` package can redact sensitive values before writing logs. Two methods are provided:

- `SetRedactionRules(rules map[string]string)`
  - Treats each map key as a literal pattern prefix (e.g., `password:` or `api_key=`).
  - Matches the key plus the following non-space characters and replaces them with the key concatenated with the provided replacement.
  - Example: `{"password:": "***REDACTED***"}` turns `password:secret` into `password:***REDACTED***`.

- `SetRedactionRegex(patterns map[string]string) error`
  - Accepts full regular expressions as keys and replacement strings as values.
  - Compiles each regex and returns an error if any pattern is invalid.
  - The replacement string replaces the matched substring.

#### **Notes**

- If the `minLevel` is set to `DEBUG`, all log messages will be displayed.
- Logs are automatically flushed to the configured output as soon as they're written.
- To log without colors (e.g., for testing), set the `disableColors` field to `true` in the `Logger` instance.

## Examples:

For examples of each function, please checkout [EXAMPLES.md](/logging/EXAMPLES.md)

---
