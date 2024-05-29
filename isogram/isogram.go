package isogram

import "strings"

func IsIsogram(word string) bool {
	letters := make(map[rune]bool)
	word = strings.ToLower(word)

	for _, letter := range word {
		_, ok := letters[letter]
		if ok && letter != ' ' && letter != '-' {
			return false
		} else {
			letters[letter] = true
		}
	}

	return true
}
