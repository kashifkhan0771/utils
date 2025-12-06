package logging_test

import (
	"bytes"
	"io"
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

func TestLoggerRedaction(t *testing.T) {
	tests := []struct {
		name           string
		rules          map[string]string
		message        string
		wantContains   string
		wantNotContain string
	}{
		{
			name: "success - redact password",
			rules: map[string]string{
				"password:": "***REDACTED***",
			},
			message:        "User logged in with password:mysecretpass123",
			wantContains:   "password:***REDACTED***",
			wantNotContain: "mysecretpass123",
		},
		{
			name: "success - redact credit card",
			rules: map[string]string{
				"credit_card=": "***REDACTED***",
			},
			message:        "Payment processed with credit_card=1234-5678-9876-5432",
			wantContains:   "credit_card=***REDACTED***",
			wantNotContain: "1234-5678-9876-5432",
		},
		{
			name: "success - redact multiple fields",
			rules: map[string]string{
				"password:":    "***REDACTED***",
				"email=":       "***REDACTED***",
				"credit_card=": "***REDACTED***",
			},
			message:        "User email=user@example.com logged in with password:secret123",
			wantContains:   "email=***REDACTED***",
			wantNotContain: "user@example.com",
		},
		{
			name:           "success - no redaction without rules",
			rules:          map[string]string{},
			message:        "User logged in with password:secret",
			wantContains:   "password:secret",
			wantNotContain: "REDACTED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			logger := logging.NewLogger("Test", logging.INFO, buffer)
			logger.SetRedactionRules(tt.rules)

			logger.Info(tt.message)

			output := buffer.String()
			if !strings.Contains(output, tt.wantContains) {
				t.Errorf("Expected output to contain '%v', got: %v", tt.wantContains, output)
			}
			if tt.wantNotContain != "" && strings.Contains(output, tt.wantNotContain) {
				t.Errorf("Expected output NOT to contain '%v', got: %v", tt.wantNotContain, output)
			}
		})
	}
}

func TestLoggerRedactionRegex(t *testing.T) {
	tests := []struct {
		name           string
		patterns       map[string]string
		message        string
		wantContains   string
		wantNotContain string
		wantErr        bool
	}{
		{
			name: "success - redact email with regex",
			patterns: map[string]string{
				`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`: "[EMAIL REDACTED]",
			},
			message:        "Contact us at support@example.com for help",
			wantContains:   "[EMAIL REDACTED]",
			wantNotContain: "support@example.com",
		},
		{
			name: "success - redact credit card numbers",
			patterns: map[string]string{
				`\b\d{4}-\d{4}-\d{4}-\d{4}\b`: "[CARD REDACTED]",
			},
			message:        "Card number: 1234-5678-9012-3456",
			wantContains:   "[CARD REDACTED]",
			wantNotContain: "1234-5678-9012-3456",
		},
		{
			name: "success - redact phone numbers",
			patterns: map[string]string{
				`\b\d{3}-\d{3}-\d{4}\b`: "[PHONE REDACTED]",
			},
			message:        "Call me at 555-123-4567",
			wantContains:   "[PHONE REDACTED]",
			wantNotContain: "555-123-4567",
		},
		{
			name: "error - invalid regex pattern",
			patterns: map[string]string{
				`[invalid(regex`: "REDACTED",
			},
			message: "test message",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			logger := logging.NewLogger("Test", logging.INFO, buffer)

			err := logger.SetRedactionRegex(tt.patterns)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetRedactionRegex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			logger.Info(tt.message)

			output := buffer.String()
			if !strings.Contains(output, tt.wantContains) {
				t.Errorf("Expected output to contain '%v', got: %v", tt.wantContains, output)
			}
			if tt.wantNotContain != "" && strings.Contains(output, tt.wantNotContain) {
				t.Errorf("Expected output NOT to contain '%v', got: %v", tt.wantNotContain, output)
			}
		})
	}
}

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func BenchmarkLogger(b *testing.B) {
	logger := logging.NewLogger("Test", logging.INFO, io.Discard)
	b.ReportAllocs()
	for b.Loop() {
		logger.Info("This is an info message")
	}
}

func BenchmarkLoggerWithRedaction(b *testing.B) {
	logger := logging.NewLogger("Test", logging.INFO, io.Discard)
	logger.SetRedactionRules(map[string]string{
		"password:": "***REDACTED***",
		"email=":    "***REDACTED***",
	})
	b.ReportAllocs()
	for b.Loop() {
		logger.Info("User email=user@example.com logged in with password:secret123")
	}
}
