package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "error", errors.New("invalid input")
	}

	romanNumerals := ""

	for input > 0 {
		switch {
		case input >= 1000:
			romanNumerals += "M"
			input -= 1000

		case input >= 900:
			romanNumerals += "CM"
			input -= 900

		case input >= 500:
			romanNumerals += "D"
			input -= 500

		case input >= 400:
			romanNumerals += "CD"
			input -= 400

		case input >= 100:
			romanNumerals += "C"
			input -= 100

		case input >= 90:
			romanNumerals += "XC"
			input -= 90

		case input >= 50:
			romanNumerals += "L"
			input -= 50

		case input >= 40:
			romanNumerals += "XL"
			input -= 40

		case input >= 10:
			romanNumerals += "X"
			input -= 10

		case input >= 9:
			romanNumerals += "IX"
			input -= 9

		case input >= 5:
			romanNumerals += "V"
			input -= 5

		case input >= 4:
			romanNumerals += "IV"
			input -= 4

		case input >= 1:
			romanNumerals += "I"
			input -= 1
		}
	}

	return romanNumerals, nil
}
