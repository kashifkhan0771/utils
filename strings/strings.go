/*
Package strings defines strings helpers.
*/
package strings

import (
	"regexp"
	"strconv"
	"strings"

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
	var result []string
	var lowerInput, lowerSubstring string

	if options.CaseInsensitive {
		lowerInput = strings.ToLower(input)
		lowerSubstring = strings.ToLower(substring)
	} else {
		lowerInput = input
		lowerSubstring = substring
	}

	startIndex := 0
	for {
		index := strings.Index(lowerInput[startIndex:], lowerSubstring)
		if index == -1 {
			break
		}

		if options.ReturnIndexes {
			result = append(result, input[index+startIndex:])
		} else {
			result = append(result, input[startIndex+index:startIndex+index+len(substring)])
		}

		startIndex += index + 1
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
	// Split the input string into words
	words := strings.Fields(input)

	// Create a maps of exceptions for faster lookup
	exceptionMap := make(map[string]bool)
	for _, exception := range exceptions {
		exceptionMap[exception] = true
	}

	// Iterate through words and capitalize the first letter if not in exceptions
	for i, word := range words {
		if !exceptionMap[word] {
			// Convert the first character to uppercase.
			words[i] = firstLetterToUpper(strings.ToLower(word))
		}
	}

	// Join the words back together into a single string
	return strings.Join(words, " ")
}

// Tokenize splits a given string into words based on whitespace and custom delimiters.
func Tokenize(input string, customDelimiters string) []string {
	// Create a function to split a string based on custom delimiters.
	customSplit := func(c rune) bool {
		return strings.ContainsRune(customDelimiters, c) || c == ' '
	}

	// Split the string using the custom split function.
	tokens := strings.FieldsFunc(input, customSplit)

	return tokens
}

// Rot13Encode encodes a string using the ROT13 cipher.
func Rot13Encode(input string) string {
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
	i := 0

	for i < length {
		char := runes[i]
		j := i + 1

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

		// Move to the next character group
		i = j
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
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	isValid, err := regexp.MatchString(pattern, email)
	if err != nil {
		return false
	}

	return isValid
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
	runes := []rune(input)
	inputLength := len(runes)
	result := make([]rune, inputLength)
	lastCharacterIndex := inputLength - 1

	for index, character := range runes {
		result[lastCharacterIndex-index] = character
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
