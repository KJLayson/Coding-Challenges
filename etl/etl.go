package etl

import "strings"

// input:  map[int][]string{1: {"A", "E", "I", "O", "U"}},
// expect: map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1},

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}

	for key, val := range in {
		for _, char := range val {
			char = strings.ToLower(char)
			out[char] = key
		}
	}

	return out
}
