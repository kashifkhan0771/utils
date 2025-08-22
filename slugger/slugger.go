package slugger

import (
	"cmp"
	"slices"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/forPelevin/gomoji"
	"golang.org/x/text/unicode/norm"
)

const defaultSeparator = "-"

// Slugger provides functionality to generate URL-friendly slugs from strings.
// It supports configurable substitutions, optional emoji handling, and custom separators.
type Slugger struct {
	Separator     string            // Default character(s) used to separate words in the slug if not explicitly provided
	Substitutions map[string]string // A map of string replacements to apply before generating the slug
	WithEmoji     bool              // If true, emojis will be included in a slug-friendly format

	mu       sync.RWMutex      // Guards substitutions and replacer; safe for concurrent Slug and updates
	replacer *strings.Replacer // Replacer used to handle substitutions in the input string
}

// New creates and returns a new Slugger instance with optional substitutions and emoji handling.
func New(substitutions map[string]string, withEmoji bool) *Slugger {
	return &Slugger{
		Substitutions: copyMap(substitutions),
		WithEmoji:     withEmoji,
		Separator:     defaultSeparator,
	}
}

// Slug converts the given string `s` into a URL-friendly slug.
// - Applies emoji replacement if enabled
// - Lowercases the string
// - Applies configured substitutions
// - Normalizes to safe ASCII
// - Joins words with the provided separator (or the default)
func (sl *Slugger) Slug(s, separator string) string {
	if s == "" {
		return ""
	}

	if sl.WithEmoji {
		s = gomoji.ReplaceEmojisWithSlug(s)
	}

	s = strings.ToLower(s)

	sl.mu.RLock()
	r := sl.replacer
	sl.mu.RUnlock()
	if r == nil {
		sl.initReplacer()
		sl.mu.RLock()
		r = sl.replacer
		sl.mu.RUnlock()
	}
	if r != nil {
		s = r.Replace(s)
	}

	words := normalizeWordsToSafeASCII(s)
	if len(words) == 0 {
		return ""
	}

	return strings.Join(words, cmp.Or(separator, sl.Separator, defaultSeparator))
}

// AddSubstitution adds a new substitution to the Slugger and resets the replacer cache.
func (sl *Slugger) AddSubstitution(key, value string) {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	if sl.Substitutions == nil {
		sl.Substitutions = make(map[string]string)
	}

	sl.Substitutions[key] = value
	sl.replacer = nil
}

// RemoveSubstitution deletes a substitution by key and resets the replacer cache.
func (sl *Slugger) RemoveSubstitution(key string) {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	if len(sl.Substitutions) == 0 {
		return
	}

	if _, exists := sl.Substitutions[key]; exists {
		delete(sl.Substitutions, key)
		sl.replacer = nil
	}
}

// ReplaceSubstitution updates the value of an existing substitution and resets the replacer cache.
func (sl *Slugger) ReplaceSubstitution(key, newValue string) {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	if len(sl.Substitutions) == 0 {
		return
	}

	if _, exists := sl.Substitutions[key]; exists {
		sl.Substitutions[key] = newValue
		sl.replacer = nil
	}
}

// SetSubstitutions replaces all current substitutions with the provided map and resets the replacer cache.
func (sl *Slugger) SetSubstitutions(substitutions map[string]string) {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	sl.Substitutions = copyMap(substitutions)
	sl.replacer = nil
}

func copyMap(m map[string]string) map[string]string {
	cp := make(map[string]string, len(m))
	for k, v := range m {
		cp[k] = v
	}

	return cp
}

// ligatureReplacer is used to replace common ligatures with their ASCII equivalents.
// Only lowercase is needed because Slug() lowercases before calling this.
// Add more if you run into them (e.g. œ → ae, ß → ss).
var ligatureReplacer = strings.NewReplacer(
	"æ", "ae",
	"ø", "oe",
	"ß", "ss",
)

// normalizeWordsToSafeASCII converts a string into a slice of ASCII-only words:
// - Replaces common ligatures (e.g., æ → ae, ø → oe)
// - Applies NFKD normalization (decomposes diacritics)
// - Keeps only ASCII letters, digits, spaces, and a few safe punctuation chars (-_.)
// - Removes all other characters
func normalizeWordsToSafeASCII(s string) []string {
	const keep = "-_."

	// Handle common ligatures explicitly (belt & suspenders with NFKD)
	s = ligatureReplacer.Replace(s)

	// Decompose once so diacritics become combining marks we can drop
	s = norm.NFKD.String(s)

	s = strings.Map(func(r rune) rune {
		if r <= unicode.MaxASCII && (unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) || strings.ContainsRune(keep, r)) {
			return r
		}

		return -1
	}, s)

	return strings.Fields(s)
}

// initReplacer builds and caches a strings.Replacer from the current substitutions.
// Substitution keys are sorted by length (descending) to ensure longer matches are replaced first.
func (sl *Slugger) initReplacer() {
	sl.mu.Lock()

	if sl.replacer != nil {
		sl.mu.Unlock()

		return
	}
	defer sl.mu.Unlock()

	if len(sl.Substitutions) == 0 {
		return
	}

	// struct to hold key-value pairs for sorting
	type subsKV struct {
		k, v string
		n    int
	}

	subsPairs := make([]subsKV, 0, len(sl.Substitutions))
	for k, v := range sl.Substitutions {
		if k == "" {
			continue
		}
		subsPairs = append(subsPairs,
			subsKV{
				k: k,
				v: strings.ToLower(v),
				n: utf8.RuneCountInString(k),
			})
	}

	if len(subsPairs) == 0 {
		return
	}

	slices.SortFunc(subsPairs, func(a, b subsKV) int {
		if a.n != b.n {
			return cmp.Compare(b.n, a.n) // sort by key length DESC
		}

		return cmp.Compare(a.k, b.k)
	})

	subs := make([]string, 0, len(subsPairs)*2)
	for _, sub := range subsPairs {
		subs = append(subs, strings.ToLower(sub.k), " "+sub.v)
	}

	if sLen := len(subs); sLen < 2 || sLen%2 != 0 {
		return
	}

	sl.replacer = strings.NewReplacer(subs...)
}
