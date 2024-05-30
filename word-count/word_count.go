package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	re_cleanWS := regexp.MustCompile(`\\.?`)
	re_afterClean := regexp.MustCompile(`[a-zA-Z0-9']*[^\W\s]`)
	histogram := Frequency{}

	phrase = strings.ToLower(phrase)

	phrase_clean := re_cleanWS.ReplaceAllString(phrase, "")
	words := re_afterClean.FindAllString(phrase_clean, -1)

	for _, word := range words {
		if word[:1] == "'" {
			word = word[1:]
		}

		histogram[word] += 1
	}

	return histogram
}
