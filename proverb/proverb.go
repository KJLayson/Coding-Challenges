package proverb

import "fmt"

func regular_word(str1, str2 string) string {
	return fmt.Sprintf("For want of a %s the %s was lost.", str1, str2)
}

func final_word(str string) string {
	return fmt.Sprintf("And all for the want of a %s.", str)
}

func Proverb(rhyme []string) []string {
	var s []string
	length := len(rhyme)

	for i := 0; i < length; i++ {
		if i < length-1 {
			s = append(s, regular_word(rhyme[i], rhyme[i+1]))
		} else {
			s = append(s, final_word(rhyme[0]))
		}
	}

	return s
}
