package logging

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorBlue   = "\033[34m" // INFO
	ColorGreen  = "\033[32m" // DEBUG
	ColorYellow = "\033[33m" // WARN
	ColorRed    = "\033[31m" // ERROR
)

type LogLevel int

// Log levels
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

type Logger struct {
	minLevel      LogLevel
	prefix        string
	output        io.Writer
	disableColors bool // Flag to disable color codes for testing
}

// Logger constructor. Default output is stdout.
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

func (l *Logger) log(level LogLevel, message string) {
	if level < l.minLevel {
		return
	}

	// Determine color and level string
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

	if l.disableColors {
		color = ""
	}

	// Format timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Construct the log message
	logMessage := fmt.Sprintf("%s[%s] [%s] %s: %s%s\n", color, timestamp, levelStr, l.prefix, message, ColorReset)

	// Write to the configured output
	fmt.Fprint(l.output, logMessage)
}

func (l *Logger) Info(message string) {
	l.log(INFO, message)
}

func (l *Logger) Debug(message string) {
	l.log(DEBUG, message)
}

func (l *Logger) Warn(message string) {
	l.log(WARN, message)
}

func (l *Logger) Error(message string) {
	l.log(ERROR, message)
}
