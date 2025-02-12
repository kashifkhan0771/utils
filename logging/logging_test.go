package logging_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/kashifkhan0771/utils/logging"
)

func TestLogger(t *testing.T) {
	type args struct {
		level   logging.LogLevel
		message string
	}
	tests := []struct {
		name       string
		minLevel   logging.LogLevel
		args       args
		wantOutput string
	}{
		{
			name:       "success - log INFO when minLevel is INFO",
			minLevel:   logging.INFO,
			args:       args{level: logging.INFO, message: "This is an info message"},
			wantOutput: "INFO",
		},
		{
			name:       "success - do not log DEBUG when minLevel is INFO",
			minLevel:   logging.INFO,
			args:       args{level: logging.DEBUG, message: "This is a debug message"},
			wantOutput: "",
		},
		{
			name:       "success - log DEBUG when minLevel is DEBUG",
			minLevel:   logging.DEBUG,
			args:       args{level: logging.DEBUG, message: "This is a debug message"},
			wantOutput: "DEBUG",
		},
		{
			name:       "success - log WARN when minLevel is WARN",
			minLevel:   logging.WARN,
			args:       args{level: logging.WARN, message: "This is a warning message"},
			wantOutput: "WARN",
		},
		{
			name:       "success - do not log INFO when minLevel is WARN",
			minLevel:   logging.WARN,
			args:       args{level: logging.INFO, message: "This is an info message"},
			wantOutput: "",
		},
		{
			name:       "success - log ERROR when minLevel is ERROR",
			minLevel:   logging.ERROR,
			args:       args{level: logging.ERROR, message: "This is an error message"},
			wantOutput: "ERROR",
		},
		{
			name:       "success - do not log WARN when minLevel is ERROR",
			minLevel:   logging.ERROR,
			args:       args{level: logging.WARN, message: "This is a warning message"},
			wantOutput: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare a buffer to capture log output
			buffer := &bytes.Buffer{}
			logger := logging.NewLogger("Test", tt.minLevel, buffer)

			// Log the message
			switch tt.args.level {
			case logging.INFO:
				logger.Info(tt.args.message)
			case logging.DEBUG:
				logger.Debug(tt.args.message)
			case logging.WARN:
				logger.Warn(tt.args.message)
			case logging.ERROR:
				logger.Error(tt.args.message)
			}

			// Verify the output
			output := buffer.String()
			if tt.wantOutput == "" {
				if output != "" {
					t.Errorf("Expected no output, got: %v", output)
				}
			} else {
				if !strings.Contains(output, tt.wantOutput) {
					t.Errorf("Expected output to contain '%v', got: %v", tt.wantOutput, output)
				}
			}
		})
	}
}

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func BenchmarkLogger(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		logger := logging.NewLogger("Test", logging.INFO, nil)
		logger.Info("This is an info message")
	}
}
