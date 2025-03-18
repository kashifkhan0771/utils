package slugger

import (
	"strings"
	"testing"
)

func TestSlugger_Slug(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		separator     string
		substitutions map[string]string
		withEmoji     bool
		unique        bool
		expected      string
	}{
		{
			name:      "Basic slug generation",
			input:     "Hello World",
			separator: "-",
			expected:  "hello-world",
		},
		{
			name:      "Basic slug generation using the default separator",
			input:     "Hello World",
			separator: "",
			expected:  "hello-world",
		},
		{
			name:      "Custom separator",
			input:     "Hello World",
			separator: "_",
			expected:  "hello_world",
		},
		{
			name:          "With substitutions",
			input:         "10% or 5€",
			separator:     "-",
			substitutions: map[string]string{"%": "percent", "€": "euro"},
			expected:      "10-percent-or-5-euro",
		},
		{
			name:      "With emoji replacement",
			input:     "Hello 🌍",
			separator: "-",
			withEmoji: true,
			expected:  "hello-globe-showing-europe-africa",
		},
		{
			name:      "With unique slug",
			input:     "Hello World",
			separator: "-",
			unique:    true,
			expected:  "hello-world-", // UUID will be appended
		},
		{
			name:      "With unique slug and custom separator",
			input:     "Hello World",
			separator: "/",
			unique:    true,
			expected:  "hello/world/", // UUID will be appended
		},
		{
			name:      "Normalize to safe ASCII",
			input:     "Wôrķšpáçè ~~sèťtïñğš~~",
			separator: "-",
			expected:  "workspace-settings",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slugger := New(tt.substitutions, tt.withEmoji, tt.unique)
			result := slugger.Slug(tt.input, tt.separator)

			if tt.unique {
				// Check if the result starts with the expected prefix and ends with a UUID
				if !strings.HasPrefix(result, tt.expected) || len(result) <= len(tt.expected) {
					t.Errorf("expected slug to start with %q and include UUID, got %q", tt.expected, result)
				}
			} else {
				if result != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, result)
				}
			}
		})
	}
}

func BenchmarkSlugger_Slug(b *testing.B) {
	slugger := &Slugger{
		Separator: "-",
		WithEmoji: false,
		Substitutions: map[string]string{
			"&": "and",
		},
		Unique: true,
	}

	for n := 0; n < b.N; n++ {
		slugger.Slug("Wôrķšpáçè ~~sèťtïñğš~~", "")
	}
}

func BenchmarkSlugger_Slug_WithEmoji(b *testing.B) {
	slugger := &Slugger{
		Separator: "-",
		WithEmoji: true,
		Substitutions: map[string]string{
			"&": "and",
		},
		Unique: false,
	}

	for n := 0; n < b.N; n++ {
		slugger.Slug("a 😺, 🐈‍⬛, and a 🦁 go to 🏞️", "")
	}
}

func BenchmarkSlugger_Slug_CustomSeparator(b *testing.B) {
	slugger := &Slugger{
		Separator: "_",
		WithEmoji: false,
		Substitutions: map[string]string{
			"&": "and",
		},
		Unique: false,
	}

	for n := 0; n < b.N; n++ {
		slugger.Slug("Wôrķšpáçè ~~sèťtïñğš~~", "|")
	}
}