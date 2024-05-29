package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ToLower(input)
	letter_map := map[string]int{}
	for _, char := range input {
		if unicode.IsLetter(char) {
			letter_map[string(char)] += 1
		}
	}

	if len(letter_map) >= 26 {
		return true
	} else {
		return false
	}
}
