package anagram

import "strings"

func breakdown(sub_map map[rune]int, sub, candidate string) bool {
	candidate = strings.ToLower(candidate)
	candidate_map := map[rune]int{}

	if sub == candidate {
		return false
	}

	for _, char := range candidate {
		candidate_map[char] += 1
	}

	if len(sub_map) == len(candidate_map) {
		for key, val := range sub_map {
			if val != candidate_map[key] {
				return false
			}
		}

		return true

	} else {
		return false
	}
}

func Detect(subject string, candidates []string) []string {
	subject_lower := strings.ToLower(subject)
	anagrams := []string{}
	subject_map := map[rune]int{}

	for _, char := range subject_lower {
		subject_map[char] += 1
	}

	for _, word := range candidates {
		if breakdown(subject_map, subject_lower, word) {
			anagrams = append(anagrams, word)
		}
	}

	return anagrams
}
