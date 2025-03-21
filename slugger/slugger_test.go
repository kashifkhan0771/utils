package slugger

import (
	"testing"
)

func TestSlugger_Slug(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		separator     string
		substitutions map[string]string
		withEmoji     bool
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
			name:      "Normalize to safe ASCII",
			input:     "Wôrķšpáçè ~~sèťtïñğš~~",
			separator: "-",
			expected:  "workspace-settings",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slugger := New(tt.substitutions, tt.withEmoji)
			result := slugger.Slug(tt.input, tt.separator)

			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
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
	}

	for n := 0; n < b.N; n++ {
		slugger.Slug("Wôrķšpáçè ~~sèťtïñğš~~", "|")
	}
}
