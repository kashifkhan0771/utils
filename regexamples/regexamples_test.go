package regexamples

import (
	"regexp"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		pattern   string
		n         int
		wantErr   bool
		wantLen   int
		mustMatch bool
	}{
		{
			name:      "zero examples",
			pattern:   `\d+`,
			n:         0,
			wantLen:   0,
			mustMatch: false,
		},
		{
			name:      "digits",
			pattern:   `\d+`,
			n:         5,
			wantLen:   5,
			mustMatch: true,
		},
		{
			name:      "word characters with range",
			pattern:   `\w{3,6}`,
			n:         10,
			wantLen:   10,
			mustMatch: true,
		},
		{
			name:      "alternation",
			pattern:   `(foo|bar|baz)`,
			n:         10,
			wantLen:   10,
			mustMatch: true,
		},
		{
			name:      "literal",
			pattern:   `hello`,
			n:         3,
			wantLen:   3,
			mustMatch: true,
		},
		{
			name:      "optional character",
			pattern:   `colou?r`,
			n:         10,
			wantLen:   10,
			mustMatch: true,
		},
		{
			name:      "alphanumeric character class",
			pattern:   `[a-zA-Z0-9]+`,
			n:         5,
			wantLen:   5,
			mustMatch: true,
		},
		{
			name:      "negated character class",
			pattern:   `[^0-9]{4}`,
			n:         5,
			wantLen:   5,
			mustMatch: true,
		},
		{
			name:      "dot any char",
			pattern:   `.{5}`,
			n:         5,
			wantLen:   5,
			mustMatch: true,
		},
		{
			name:      "anchors are ignored",
			pattern:   `^\d{3}$`,
			n:         5,
			wantLen:   5,
			mustMatch: true,
		},
		{
			name:      "exact repetition",
			pattern:   `a{5}`,
			n:         3,
			wantLen:   3,
			mustMatch: true,
		},
		{
			name:      "email-like pattern",
			pattern:   `[a-z]{3,8}@[a-z]{3,6}\.(com|net|org)`,
			n:         5,
			wantLen:   5,
			mustMatch: true,
		},
		{
			name:      "star quantifier",
			pattern:   `ab*c`,
			n:         10,
			wantLen:   10,
			mustMatch: true,
		},
		{
			name:      "plus quantifier",
			pattern:   `ab+c`,
			n:         10,
			wantLen:   10,
			mustMatch: true,
		},
		{
			name:      "word boundary anchors",
			pattern:   `\bword\b`,
			n:         3,
			wantLen:   3,
			mustMatch: true,
		},
		{
			name:    "negative n",
			pattern: `\d+`,
			n:       -1,
			wantErr: true,
		},
		{
			name:    "invalid pattern",
			pattern: `[invalid`,
			n:       5,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			results, err := Generate(tt.pattern, tt.n)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Generate(%q, %d): expected error, got nil", tt.pattern, tt.n)
				}

				return
			}

			if err != nil {
				t.Fatalf("Generate(%q, %d): unexpected error: %v", tt.pattern, tt.n, err)
			}

			if len(results) != tt.wantLen {
				t.Errorf("Generate(%q, %d): got %d results, want %d", tt.pattern, tt.n, len(results), tt.wantLen)
			}

			if tt.mustMatch {
				re := regexp.MustCompile(tt.pattern)

				for i, s := range results {
					if !re.MatchString(s) {
						t.Errorf("result[%d] = %q does not match pattern %q", i, s, tt.pattern)
					}
				}
			}
		})
	}
}

func TestNewGenerator(t *testing.T) {
	t.Parallel()

	t.Run("invalid pattern returns error", func(t *testing.T) {
		t.Parallel()

		_, err := NewGenerator(`(unclosed`)
		if err == nil {
			t.Error("NewGenerator: expected error for invalid pattern, got nil")
		}
	})

	t.Run("reuse generator", func(t *testing.T) {
		t.Parallel()

		g, err := NewGenerator(`\d{4}`)
		if err != nil {
			t.Fatalf("NewGenerator: unexpected error: %v", err)
		}

		re := regexp.MustCompile(`\d{4}`)

		for range 20 {
			results, err := g.Generate(3)
			if err != nil {
				t.Fatalf("Generate: unexpected error: %v", err)
			}

			for _, s := range results {
				if !re.MatchString(s) {
					t.Errorf("%q does not match pattern", s)
				}
			}
		}
	})
}

func TestGeneratorSetSeed(t *testing.T) {
	t.Parallel()

	pattern := `[a-z]{5}-\d{3}`

	g1, err := NewGenerator(pattern)
	if err != nil {
		t.Fatalf("NewGenerator: unexpected error: %v", err)
	}

	g1.SetSeed(42)
	results1, err := g1.Generate(10)
	if err != nil {
		t.Fatalf("g1.Generate: unexpected error: %v", err)
	}

	g2, err := NewGenerator(pattern)
	if err != nil {
		t.Fatalf("NewGenerator: unexpected error: %v", err)
	}

	g2.SetSeed(42)
	results2, err := g2.Generate(10)
	if err != nil {
		t.Fatalf("g2.Generate: unexpected error: %v", err)
	}

	for i := range results1 {
		if results1[i] != results2[i] {
			t.Errorf("SetSeed(42): result[%d] differs: %q vs %q", i, results1[i], results2[i])
		}
	}
}

func TestPrintableOutput(t *testing.T) {
	t.Parallel()

	// Negated digit class — must stay in printable ASCII range.
	results, err := Generate(`[^0-9]{10}`, 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, s := range results {
		for _, r := range s {
			if r < printableMin || r > printableMax {
				t.Errorf("result %q contains non-printable rune %U", s, r)
			}
		}
	}
}
