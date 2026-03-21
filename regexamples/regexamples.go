// Package regexamples generates random strings that match a given regular expression.
package regexamples

import (
	"fmt"
	"math"
	randv2 "math/rand/v2"
	"regexp/syntax"
	"strings"
	"unicode"
)

const (
	// DefaultMaxRepeat is the default cap for unbounded quantifiers (*, +, and {n,} without upper bound).
	DefaultMaxRepeat = 10

	printableMin rune = 32  // space
	printableMax rune = 126 // tilde (~)
)

// Generator holds a compiled pattern and a random source, and can produce
// multiple matching strings efficiently.
type Generator struct {
	re  *syntax.Regexp
	rng *randv2.Rand
}

// NewGenerator compiles pattern and returns a Generator ready to produce
// matching strings. Returns an error if the pattern is invalid.
func NewGenerator(pattern string) (*Generator, error) {
	re, err := syntax.Parse(pattern, syntax.Perl)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %w", err)
	}

	return &Generator{
		re:  re.Simplify(),
		rng: randv2.New(randv2.NewPCG(randv2.Uint64(), randv2.Uint64())), //nolint:gosec
	}, nil
}

// SetSeed resets the generator's random source to a deterministic seed,
// which makes subsequent Generate calls produce the same sequence of strings.
func (g *Generator) SetSeed(seed uint64) {
	// #nosec G404 -- not used for security
	g.rng = randv2.New(randv2.NewPCG(seed, seed))
}

// Generate returns n strings that each match the compiled pattern.
// Returns an error if n is negative or if the pattern cannot be satisfied.
func (g *Generator) Generate(n int) ([]string, error) {
	if n < 0 {
		return nil, fmt.Errorf("n cannot be negative: %d", n)
	}

	if n == 0 {
		return []string{}, nil
	}

	results := make([]string, n)

	for i := range n {
		var sb strings.Builder

		if err := g.generate(g.re, &sb); err != nil {
			return nil, fmt.Errorf("failed to generate example %d: %w", i+1, err)
		}

		results[i] = sb.String()
	}

	return results, nil
}

// Generate is a convenience function that compiles pattern and returns n
// matching strings. For repeated generation from the same pattern prefer
// NewGenerator to avoid recompiling on every call.
func Generate(pattern string, n int) ([]string, error) {
	g, err := NewGenerator(pattern)
	if err != nil {
		return nil, err
	}

	return g.Generate(n)
}

// generate recursively walks the regex AST and writes a matching string into sb.
func (g *Generator) generate(re *syntax.Regexp, sb *strings.Builder) error {
	switch re.Op {
	case syntax.OpLiteral:
		for _, r := range re.Rune {
			if re.Flags&syntax.FoldCase != 0 {
				if g.rng.IntN(2) == 0 {
					sb.WriteRune(unicode.ToUpper(r))
				} else {
					sb.WriteRune(unicode.ToLower(r))
				}
			} else {
				sb.WriteRune(r)
			}
		}

	case syntax.OpCharClass:
		r, err := g.pickFromCharClass(re.Rune)
		if err != nil {
			return err
		}

		sb.WriteRune(r)

	case syntax.OpAnyCharNotNL:
		sb.WriteRune(g.randPrintableExcluding('\n'))

	case syntax.OpAnyChar:
		sb.WriteRune(g.randPrintable())

	case syntax.OpQuest:
		if g.rng.IntN(2) == 0 {
			if err := g.generate(re.Sub[0], sb); err != nil {
				return err
			}
		}

	case syntax.OpStar:
		count := g.rng.IntN(DefaultMaxRepeat + 1)
		for range count {
			if err := g.generate(re.Sub[0], sb); err != nil {
				return err
			}
		}

	case syntax.OpPlus:
		count := g.rng.IntN(DefaultMaxRepeat) + 1
		for range count {
			if err := g.generate(re.Sub[0], sb); err != nil {
				return err
			}
		}

	case syntax.OpRepeat:
		min, max := re.Min, re.Max
		if max == -1 {
			max = min + DefaultMaxRepeat
		}

		count := min
		if max > min {
			count = min + g.rng.IntN(max-min+1)
		}

		for range count {
			if err := g.generate(re.Sub[0], sb); err != nil {
				return err
			}
		}

	case syntax.OpConcat:
		for _, sub := range re.Sub {
			if err := g.generate(sub, sb); err != nil {
				return err
			}
		}

	case syntax.OpAlternate:
		if len(re.Sub) == 0 {
			return nil // empty alternation matches empty string
		}
		idx := g.rng.IntN(len(re.Sub))
		if err := g.generate(re.Sub[idx], sb); err != nil {
			return err
		}

	case syntax.OpCapture:
		if err := g.generate(re.Sub[0], sb); err != nil {
			return err
		}

	case syntax.OpBeginText, syntax.OpEndText,
		syntax.OpBeginLine, syntax.OpEndLine,
		syntax.OpWordBoundary, syntax.OpNoWordBoundary:
		// Anchors and boundaries assert position — they produce no characters.

	case syntax.OpEmptyMatch:
		// Nothing to emit.

	case syntax.OpNoMatch:
		return fmt.Errorf("pattern contains a never-matching expression")

	default:
		return fmt.Errorf("unsupported regex operation: %v", re.Op)
	}

	return nil
}

// pickFromCharClass selects a random rune from a character class.
//
// The Rune slice in an OpCharClass node stores consecutive [lo, hi] inclusive
// rune-range pairs (as computed by regexp/syntax, including complement ranges
// for negated classes like [^0-9]).
//
// We first try to pick from the intersection of those ranges with the printable
// ASCII band [32, 126]. This keeps output human-readable and avoids emitting
// control characters or high Unicode codepoints from broad negated classes.
// If the intersection is empty we fall back to the full ranges.
func (g *Generator) pickFromCharClass(ranges []rune) (rune, error) {
	// Build the intersection with printable ASCII.
	printable := intersectWithPrintable(ranges)
	if len(printable) > 0 {
		return g.pickFromRanges(printable), nil
	}

	// Fallback: pick directly from the raw ranges (may include non-ASCII).
	total := countRunes(ranges)
	if total == 0 {
		return 0, fmt.Errorf("character class matches no characters")
	}

	return g.pickFromRanges(ranges), nil
}

// intersectWithPrintable clips each [lo, hi] pair in ranges to [printableMin, printableMax]
// and returns the resulting (possibly empty) list of clipped pairs.
func intersectWithPrintable(ranges []rune) []rune {
	var result []rune

	for i := 0; i < len(ranges); i += 2 {
		lo, hi := ranges[i], ranges[i+1]

		if lo > printableMax || hi < printableMin {
			continue // entirely outside printable band
		}

		if lo < printableMin {
			lo = printableMin
		}

		if hi > printableMax {
			hi = printableMax
		}

		result = append(result, lo, hi)
	}

	return result
}

// pickFromRanges picks a uniformly random rune from a list of [lo, hi] pairs.
func (g *Generator) pickFromRanges(ranges []rune) rune {
	total := countRunes(ranges)
	n := g.rng.IntN(total)

	for i := 0; i < len(ranges); i += 2 {
		size := int(ranges[i+1]-ranges[i]) + 1
		if n < size {
			if n < 0 || n > math.MaxInt32 {
				return 0
			}

			return ranges[i] + rune(n)
		}

		n -= size
	}

	return ranges[0]
}

// countRunes returns the total number of runes covered by the [lo, hi] pairs.
func countRunes(ranges []rune) int {
	total := 0
	for i := 0; i < len(ranges); i += 2 {
		total += int(ranges[i+1]-ranges[i]) + 1
	}

	return total
}

func (g *Generator) randPrintable() rune {
	val := g.rng.IntN(int(printableMax - printableMin + 1))
	if val < 0 || val > math.MaxInt32 {
		return 0
	}

	return printableMin + rune(val)
}

func (g *Generator) randPrintableExcluding(exclude rune) rune {
	for {
		r := g.randPrintable()
		if r != exclude {
			return r
		}
	}
}
