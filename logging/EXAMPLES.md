## Logging Function Examples

### Create and use a logger

```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Create a new logger with prefix "MyApp", minimum level INFO, and output to stdout
	logger := logging.NewLogger("MyApp", logging.INFO, os.Stdout)

	// Log messages of different levels
	logger.Debug("This is a debug message.") // Ignored because minLevel is INFO
	logger.Info("Application started.")      // Printed with blue color
	logger.Warn("Low disk space.")           // Printed with yellow color
	logger.Error("Failed to connect to DB.") // Printed with red color
}
```

#### Output:

```
[2025-01-09 12:34:56] [INFO] MyApp: Application started.
[2025-01-09 12:34:56] [WARN] MyApp: Low disk space.
[2025-01-09 12:34:56] [ERROR] MyApp: Failed to connect to DB.
```

### Log without colors (useful for plain text logs)

```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Create a logger and disable colors
	logger := logging.NewLogger("MyApp", logging.DEBUG, os.Stdout)
	logger.disableColors = true

	// Log messages of different levels
	logger.Debug("Debugging without colors.")
	logger.Info("Information without colors.")
	logger.Warn("Warning without colors.")
	logger.Error("Error without colors.")
}
```

#### Output:

```
[2025-01-09 12:34:56] [DEBUG] MyApp: Debugging without colors.
[2025-01-09 12:34:56] [INFO] MyApp: Information without colors.
[2025-01-09 12:34:56] [WARN] MyApp: Warning without colors.
[2025-01-09 12:34:56] [ERROR] MyApp: Error without colors.
```

### Log messages to a file

```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Open a log file for writing
	file, err := os.Create("app.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a logger with file output
	logger := logging.NewLogger("MyApp", logging.DEBUG, file)

	// Log messages
	logger.Debug("Writing debug logs to file.")
	logger.Info("Application log stored in file.")
	logger.Warn("This is a warning.")
	logger.Error("This is an error.")
}
```

#### Output (in `app.log` file):

```
[2025-01-09 12:34:56] [DEBUG] MyApp: Writing debug logs to file.
[2025-01-09 12:34:56] [INFO] MyApp: Application log stored in file.
[2025-01-09 12:34:56] [WARN] MyApp: This is a warning.
[2025-01-09 12:34:56] [ERROR] MyApp: This is an error.
```

### Filter logs by minimum log level

```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Create a logger with minimum level WARN
	logger := logging.NewLogger("MyApp", logging.WARN, os.Stdout)

	// Log messages
	logger.Debug("This is a debug message.") // Ignored
	logger.Info("This is an info message.")  // Ignored
	logger.Warn("This is a warning.")        // Printed
	logger.Error("This is an error.")        // Printed
}
```

#### Output:

```
[2025-01-09 12:34:56] [WARN] MyApp: This is a warning.
[2025-01-09 12:34:56] [ERROR] MyApp: This is an error.
```

### Customize log prefixes

```go
package main

import (
	logging "github.com/kashifkhan0771/utils/logging"
	"os"
)

func main() {
	// Create a logger with a custom prefix
	logger := logging.NewLogger("CustomPrefix", logging.INFO, os.Stdout)

	// Log messages
	logger.Info("This message has a custom prefix.")
}
```

#### Output:

```
[2025-01-09 12:34:56] [INFO] CustomPrefix: This message has a custom prefix.
```

### Redact sensitive information in logs

```go
package main

import (
 logging "github.com/kashifkhan0771/utils/logging"
 "os"
)

func main() {
 // Create a logger
 logger := logging.NewLogger("MyApp", logging.DEBUG, os.Stdout)

 // Set redaction rules for sensitive data
 logger.SetRedactionRules(map[string]string{
  "password:":    "***REDACTED***",
  "credit_card=": "***REDACTED***",
 })

 // Log messages with sensitive data
 logger.Info("User logged in with password:mysecretpass123")
 logger.Error("Payment failed for credit_card=1234-5678-9876-5432")
}
```
#### Output:

```
[2025-01-09 12:34:56] [INFO] MyApp: User logged in with password:***REDACTED***
[2025-01-09 12:34:56] [ERROR] MyApp: Payment failed for credit_card=***REDACTED***
```