/*
Package strings defines strings helpers.
*/
package strings

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var c = cases.Title(language.English)

// SubstringSearchOptions contains options for substring search.
type SubstringSearchOptions struct {
	CaseInsensitive bool // Perform case-insensitive search
	ReturnIndexes   bool // Return the starting indexes of found substrings
}

// TruncateOptions represents optional parameters for the Truncate function.
type TruncateOptions struct {
	Length   int
	Omission string
}

// SubstringSearch performs substring search in a string and optionally returns indexes.
func SubstringSearch(input, substring string, options SubstringSearchOptions) []string {
	if input == "" || substring == "" {
		return nil
	}

	var (
		searchInput  = input
		searchSubstr = substring

		result []string
	)

	if options.CaseInsensitive {
		searchInput = strings.ToLower(searchInput)
		searchSubstr = strings.ToLower(searchSubstr)
	}
	for start, idx := 0, 0; start < len(searchInput); start += idx + 1 {
		idx = strings.Index(searchInput[start:], searchSubstr)
		if idx == -1 {
			break
		}

		if options.ReturnIndexes {
			result = append(result, strconv.Itoa(start+idx))
		} else {
			result = append(result, input[start+idx:start+idx+len(substring)])
		}
	}

	return result
}

// Title return string in title case with English language-specific title
func Title(input string) string {
	return c.String(input)
}

// ToTitle converts a string to title case, capitalizing the first letter of each word.
// It excludes exceptions specified in the exceptions slice.
func ToTitle(input string, exceptions []string) string {
	if input = strings.TrimSpace(input); len(input) == 0 {
		return ""
	}

	// Lookup-map for word exceptions.
	exceptionsMap := make(map[string]struct{}, len(exceptions))
	for _, e := range exceptions {
		exceptionsMap[e] = struct{}{}
	}

	var output strings.Builder
	output.Grow(len(input)) // pre-allocate builder to avoid growth

	for i, word := range strings.Fields(input) {
		// If this is not the first word, add a space before it
		if i > 0 {
			output.WriteByte(' ')
		}

		// If the word is in the exceptions, write it as is
		if _, skip := exceptionsMap[word]; skip {
			output.WriteString(word)

			continue
		}

		// Capitalize the first letter and lower the rest
		// Decode the first rune to handle multi-byte characters correctly.
		// size is the number of bytes in the first rune.
		firstLetter, size := utf8.DecodeRuneInString(word)
		output.WriteRune(unicode.ToUpper(firstLetter))

		// If the word has more than one rune, write the rest in lowercase
		if len(word) > size {
			output.WriteString(strings.ToLower(word[size:]))
		}
	}

	return output.String()
}

// Tokenize splits a given string into words based on whitespace and custom delimiters.
func Tokenize(input string, customDelimiters string) []string {
	// Create a function to split a string based on custom delimiters.
	customSplit := func(c rune) bool {
		return strings.ContainsRune(customDelimiters, c) || c == ' '
	}

	// Split the string using the custom split function.
	return strings.FieldsFunc(input, customSplit)
}

// Rot13Encode encodes a string using the ROT13 cipher.
func Rot13Encode(input string) string {
	if len(input) == 0 {
		return ""
	}

	encoded := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		char := input[i]
		switch {
		case char >= 'A' && char <= 'Z':
			encoded[i] = 'A' + ((char - 'A' + 13) % 26)
		case char >= 'a' && char <= 'z':
			encoded[i] = 'a' + ((char - 'a' + 13) % 26)
		default:
			encoded[i] = char
		}
	}

	return string(encoded)
}

// Rot13Decode decodes a string encoded with the ROT13 cipher.
func Rot13Decode(input string) string {
	return Rot13Encode(input)
}

// CaesarEncrypt encrypts a string using the Caesar cipher with a given shift.
func CaesarEncrypt(input string, shift int) string {
	if len(input) == 0 {
		return ""
	}

	shifted := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		char := input[i]
		shiftedChar := char

		if char >= 'A' && char <= 'Z' {
			shiftedChar = 'A' + (char-'A'+byte(shift))%26
		} else if char >= 'a' && char <= 'z' {
			shiftedChar = 'a' + (char-'a'+byte(shift))%26
		}

		shifted[i] = shiftedChar
	}

	return string(shifted)
}

// RunLengthEncode takes a string and returns its Run-Length Encoded
// representation or the original string if the encoding did not achieve any compression
func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var encoded strings.Builder
	runes := []rune(input)
	count := 1

	for i := 1; i < len(input); i++ {
		if runes[i] == runes[i-1] {
			count++
		} else {
			encoded.WriteRune(runes[i-1])
			encoded.WriteString(strconv.Itoa(count))
			count = 1
		}
	}

	// Write the last character and its count
	encoded.WriteRune(runes[len(runes)-1])
	encoded.WriteString(strconv.Itoa(count))

	// Return the original string if encoding doesn't save space
	if encoded.Len() >= len(input) {
		return input
	}

	return encoded.String()
}

// RunLengthDecode takes a Run-Length Encoded string and returns
// the decoded string or an error + the orignal encoded string
func RunLengthDecode(encoded string) (string, error) {
	if len(encoded) == 0 {
		return "", nil
	}

	var decoded strings.Builder
	runes := []rune(encoded)
	length := len(runes)

	for i, j := 0, 0; i < length; i = j {
		char := runes[i]
		j = i + 1

		// Check if a number follows the character
		for j < length && runes[j] >= '0' && runes[j] <= '9' {
			j++
		}

		if j > i+1 {
			// Parse the count if a number follows the character
			count, err := strconv.Atoi(string(runes[i+1 : j]))

			// If an error accures return the error and the original string
			if err != nil {
				return encoded, err
			}

			decoded.WriteString(strings.Repeat(string(char), count))
		} else {
			// If no number follows, treat the character as unencoded
			decoded.WriteRune(char)
		}
	}

	return decoded.String(), nil
}

// CaesarDecrypt decrypts a string encrypted with the Caesar cipher and a given shift.
func CaesarDecrypt(input string, shift int) string {
	return CaesarEncrypt(input, -shift)
}

// IsValidEmail checks if a given string is a valid email address.
func IsValidEmail(email string) bool {
	// Regular expression for basic email validation
	const pattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	isValid, err := regexp.MatchString(pattern, email)

	return err == nil && isValid
}

// SanitizeEmail removes leading and trailing whitespace from an email address.
func SanitizeEmail(email string) string {
	return strings.TrimSpace(email)
}

// Reverse returns a reversed version of the input string.
// It correctly handles Unicode characters.
// For example:
//
//	Reverse("hello") returns "olleh"
//	Reverse("世界") returns "界世"
func Reverse(input string) string {
	var (
		runes   = []rune(input)
		result  = make([]rune, len(runes))
		lastIdx = len(runes) - 1
	)

	for idx := range runes {
		result[lastIdx-idx] = runes[idx]
	}

	return string(result)
}

// CommonPrefix returns the longest common prefix of the given strings.
// If no strings are provided, it returns an empty string.
// If only one string is provided, it returns that string.
// For example, CommonPrefix("nation", "national", "nasty") returns "na".
func CommonPrefix(input ...string) string {
	if len(input) == 0 {
		return ""
	}

	if len(input) == 1 {
		return input[0]
	}

	prefix := []rune(input[0])

	if len(prefix) == 0 {
		return ""
	}

	for i := 1; i < len(input); i++ {
		if input[i] == "" {
			return ""
		}

		item := []rune(input[i])
		shortestTextLength := len(prefix)
		if len(prefix) > len(item) {
			shortestTextLength = len(item)
		}
		for j := 0; j < shortestTextLength; j++ {
			if prefix[j] != item[j] {
				prefix = prefix[:j]

				break
			}
		}
	}

	return string(prefix)
}

// CommonSuffix returns the longest common suffix of the given strings.
// If no strings are provided, it returns an empty string.
// If only one string is provided, it returns that string.
// For example, CommonSuffix("testing", "running", "jumping") returns "ing".
func CommonSuffix(input ...string) string {
	if len(input) == 0 {
		return ""
	}

	if len(input) == 1 {
		return input[0]
	}

	suffix := []rune(input[0])

	if len(suffix) == 0 {
		return ""
	}

	for i := 1; i < len(input); i++ {
		if input[i] == "" {
			return ""
		}

		item := []rune(input[i])

		suffixLength := len(suffix)
		itemLength := len(item)

		// Adjust suffix length if current item is shorter
		if itemLength < suffixLength {
			suffix = suffix[suffixLength-itemLength:]
			suffixLength = len(suffix)

		}

		for j, k := suffixLength-1, itemLength-1; j >= 0 && k >= 0; j, k = j-1, k-1 {
			if suffix[j] != item[k] {
				suffix = suffix[j+1:]

				break
			}
		}
	}

	return string(suffix)
}

// Truncate shortens a given input string based on provided options.
// Parameters:
// - input: the original string to truncate.
// - opts: optional settings to specify truncation length and omission suffix.
// If opts is nil or certain fields are unspecified, defaults are applied:
// Length defaults to 12 and Omission defaults to "...".
func Truncate(input string, opts *TruncateOptions) string {
	length := 12
	omission := "..."

	if opts != nil {
		if opts.Length > 0 {
			length = opts.Length
		}
		if opts.Omission != "" {
			omission = opts.Omission
		}
	}

	if len(input) <= length {
		return input
	}

	// Consider omission length in the final string length
	effectiveLength := length - len(omission)
	if effectiveLength <= 0 {
		effectiveLength = 1 // Ensure at least one character from input if possible
	}

	return input[:effectiveLength] + omission
}
