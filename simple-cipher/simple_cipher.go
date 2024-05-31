package cipher

import (
	"math"
	"regexp"
	"strings"
)

type shift struct {
	distance int
}
type vigenere struct {
	distances []int
}

func NewCaesar() Cipher {
	return &shift{3}
}

func NewShift(distance int) Cipher {
	if math.Abs(float64(distance)) > 25 || distance == 0 {
		return nil
	} else {

		return &shift{distance}
	}
}

func (c shift) Encode(input string) string {
	input = cleanInput(input)
	code := ""

	for _, char := range input {
		alphaNum := int(char) + c.distance
		code += calcLetter(alphaNum)
	}

	return code
}

func (c shift) Decode(input string) string {
	input = cleanInput(input)
	code := ""

	for _, char := range input {
		alphaNum := int(char) - c.distance
		code += calcLetter(alphaNum)
	}

	return code
}

func NewVigenere(key string) Cipher {
	if key == "" || regexp.MustCompile(`[^a-z]`).MatchString(key) {
		return nil
	}

	f_nonA := false
	distances := []int{}

	for _, char := range key {
		if char != 'a' {
			f_nonA = true
		}

		distances = append(distances, int(char-'a'))
	}

	if !f_nonA {
		return nil
	} else {
		return &vigenere{distances: distances}
	}
}

func (v vigenere) Encode(input string) string {
	input, code, i, num_keys := initVig(input, v)

	for _, char := range input {
		alphaNum := int(char) + v.distances[i]
		code += calcLetter(alphaNum)
		i = vigInc(i, num_keys)
	}

	return code
}

func (v vigenere) Decode(input string) string {
	input, code, i, num_keys := initVig(input, v)

	for _, char := range input {
		alphaNum := int(char) - v.distances[i]
		code += calcLetter(alphaNum)
		i = vigInc(i, num_keys)
	}

	return code
}

func cleanInput(input string) string {
	input = strings.ToLower(input)
	re := regexp.MustCompile(`[^a-z]`)
	input = re.ReplaceAllString(input, "")

	return input
}

func calcLetter(alphaNum int) string {
	switch {
	case alphaNum < 'a':
		alphaNum += 26
	case alphaNum > 'z':
		alphaNum -= 26
	}

	new_char := string(rune(alphaNum))

	return new_char
}

func initVig(input string, v vigenere) (string, string, int, int) {
	input = cleanInput(input)
	code := ""
	i := 0
	num_keys := len(v.distances) - 1

	return input, code, i, num_keys
}

func vigInc(i int, num_keys int) int {
	i++
	if i > num_keys {
		i = 0
	}

	return i
}
