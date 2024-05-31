package atbash

import (
	"regexp"
	"strings"
	"unicode"
)

func Atbash(s string) string {
	// Rune value for a = 97
	// using 96 so that when we subtract the base (starting value for lowercase letters), it starts at 1 instead of 0.
	const base = 96

	s = strings.ToLower(s)

	re := regexp.MustCompile(`[\W\s]`)
	s = re.ReplaceAllString(s, "")

	str_output := ""

	for i, char := range s {
		if i%5 == 0 && i != 0 {
			str_output += " "
		}

		if unicode.IsLetter(char) {
			code := int(char) - base
			shiftKey := 27 - code
			value := shiftKey + base

			if value > 'z' {
				value -= 26
			}

			str_output += string(rune(value))

		} else {
			str_output += string(char)
		}

	}

	return str_output
}
