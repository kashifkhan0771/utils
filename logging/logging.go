// Package logging provides a simple logging library with support for
// multiple log levels, custom prefixes, colored output, and customizable
// output streams.
package logging

import (
	"fmt"
	"io"
	"os"
	"time"
)

// ANSI color codes for different log levels
const (
	ColorReset  = "\033[0m"  // Reset to default color
	ColorBlue   = "\033[34m" // Blue color for INFO level logs
	ColorGreen  = "\033[32m" // Green color for DEBUG level logs
	ColorYellow = "\033[33m" // Yellow color for WARN level logs
	ColorRed    = "\033[31m" // Red color for ERROR level logs
)

// LogLevel represents the severity level of a log message.
type LogLevel int

// Defined log levels in increasing order of severity.
const (
	DEBUG LogLevel = iota // DEBUG is used for detailed information during development.
	INFO                  // INFO is used for general informational messages.
	WARN                  // WARN is used for warnings that are not critical.
	ERROR                 // ERROR is used for critical error messages.
)

// Logger is a configurable logging instance that supports multiple log levels,
// optional colored output, and custom prefixes for log messages.
type Logger struct {
	minLevel      LogLevel  // Minimum log level for messages to be logged
	prefix        string    // Prefix to prepend to all log messages
	output        io.Writer // Output destination for log messages (e.g., os.Stdout)
	disableColors bool      // Flag to disable color codes (useful for testing or non-ANSI terminals)
}

// NewLogger creates and returns a new Logger instance with the specified prefix,
// minimum log level, and output destination. If output is nil, os.Stdout is used
// as the default destination.
//
// Parameters:
//   - prefix: A string prefix that will appear in all log messages.
//   - minLevel: The minimum log level for a message to be logged.
//   - output: The io.Writer to which log messages will be written.
//
// Returns:
//
//	A pointer to a configured Logger instance.
func NewLogger(prefix string, minLevel LogLevel, output io.Writer) *Logger {
	if output == nil {
		output = os.Stdout // Default to standard output
	}

	return &Logger{
		minLevel: minLevel,
		prefix:   prefix,
		output:   output,
	}
}

// log handles the core logic of logging messages. It applies the appropriate
// color coding (if enabled), formats the log message with a timestamp and prefix,
// and writes it to the configured output destination.
//
// Parameters:
//   - level: The LogLevel of the message being logged.
//   - message: The actual log message to be recorded.
func (l *Logger) log(level LogLevel, message string) {
	if level < l.minLevel {
		return
	}

	// Determine the color and level string for the log level
	color := ""
	levelStr := ""
	switch level {
	case INFO:
		color = ColorBlue
		levelStr = "INFO"
	case DEBUG:
		color = ColorGreen
		levelStr = "DEBUG"
	case WARN:
		color = ColorYellow
		levelStr = "WARN"
	case ERROR:
		color = ColorRed
		levelStr = "ERROR"
	}

	// Disable color if the flag is set
	if l.disableColors {
		color = ""
	}

	// Format the timestamp for the log message
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Construct the formatted log message
	logMessage := fmt.Sprintf("%s[%s] [%s] %s: %s%s\n", color, timestamp, levelStr, l.prefix, message, ColorReset)

	// Write the log message to the configured output
	_, err := fmt.Fprint(l.output, logMessage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write log: %v\n", err)
	}
}

// Info logs a message at the INFO level.
//
// Parameters:
//   - message: The informational message to log.
func (l *Logger) Info(message string) {
	l.log(INFO, message)
}

// Debug logs a message at the DEBUG level.
//
// Parameters:
//   - message: The debug message to log.
func (l *Logger) Debug(message string) {
	l.log(DEBUG, message)
}

// Warn logs a message at the WARN level.
//
// Parameters:
//   - message: The warning message to log.
func (l *Logger) Warn(message string) {
	l.log(WARN, message)
}

// Error logs a message at the ERROR level.
//
// Parameters:
//   - message: The error message to log.
func (l *Logger) Error(message string) {
	l.log(ERROR, message)
}
