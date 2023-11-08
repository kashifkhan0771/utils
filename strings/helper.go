package strings

import "unicode"

func firstLetterToUpper(input string) string {
	if len(input) == 0 {
		return input
	}

	// Convert the first character to uppercase
	return string(unicode.ToUpper(rune(input[0]))) + input[1:]
}
