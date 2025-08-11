package slugger

import (
	"maps"
	"slices"
	"strings"
	"unicode"

	"github.com/forPelevin/gomoji"
	"golang.org/x/text/unicode/norm"
)

type Slugger struct {
	Separator     string            // Default character(s) used to separate words in the slug if not explicitly provided
	Substitutions map[string]string // A map of string replacements to apply before generating the slug
	WithEmoji     bool              // If true, emojis will be included in a slug-friendly format
}

func New(substitutions map[string]string, withEmoji bool) *Slugger {
	return &Slugger{
		Separator:     "-",
		Substitutions: substitutions,
		WithEmoji:     withEmoji,
	}
}

// Slug generates a slugified version of the input string `s` using the provided `separator`.
func (slugger *Slugger) Slug(s, separator string) string {
	if separator == "" {
		separator = slugger.Separator
	}

	if slugger.WithEmoji {
		s = gomoji.ReplaceEmojisWithSlug(s)
	}

	s = strings.ToLower(s)

	sortedKeys := slices.Sorted(maps.Keys(slugger.Substitutions))
	for _, oldValue := range sortedKeys {
		newValue := slugger.Substitutions[oldValue]
		s = strings.ReplaceAll(s, oldValue, " "+newValue)
	}

	safe := normalizeToSafeASCII(s)
	words := strings.Split(safe, " ")
	var slugBuilder strings.Builder

	for i := range words {
		slugBuilder.WriteString(strings.ToLower(words[i]))

		if i != len(words)-1 {
			slugBuilder.WriteString(separator)
		}
	}

	return slugBuilder.String()
}

func normalizeToSafeASCII(s string) string {
	normalized := norm.NFKD.String(s)
	// remaining safe ASCII characters
	const keep string = "-_."

	var sb strings.Builder
	for _, r := range normalized {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) || strings.ContainsRune(keep, r) {
			sb.WriteRune(r)
		}
	}

	// Remove extra spaces
	return strings.Join(strings.Fields(sb.String()), " ")
}
