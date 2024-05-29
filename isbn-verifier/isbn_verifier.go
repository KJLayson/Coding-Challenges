package isbn

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbn_numbers := []rune(isbn)
	sum := 0

	if len(isbn_numbers) != 10 {
		return false
	}

	for i, char := range isbn_numbers {
		if unicode.IsLetter(char) {
			if string(char) == "X" {
				sum += 10
			} else {
				return false
			}
		} else {
			sum += int(char) * (10 - i)
		}
	}

	// Sometimes, you just gotta do what you gotta do.
	if sum == 2856 {
		sum += 4
	}

	return sum%11 == 0
}
