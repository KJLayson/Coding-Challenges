package rotationalcipher

import (
	"unicode"
)

// Rune values for first and last letter
// a - 97, z - 122
// A - 65, Z - 90
func RotationalCipher(plain string, shiftKey int) string {
	const a = 97
	const z = 122
	const A = 65
	const Z = 90

	str_output := ""

	for _, char := range plain {

		if unicode.IsLetter(char) {

			if unicode.IsUpper(char) {
				value := shiftKey + int(char)

				if value > Z {
					diff := value - Z - 1
					value = A + diff
				}

				str_output += string(rune(value))

			} else {

				value := shiftKey + int(char)

				if value > z {
					diff := value - z - 1
					value = a + diff
				}

				str_output += string(rune(value))

			}

		} else {
			str_output += string(char)
		}

	}

	return str_output
}
