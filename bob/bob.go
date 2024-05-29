package bob

import (
	"fmt"
	"regexp"
	"strings"
)

func Hey(remark string) string {
	remark = strings.ReplaceAll(remark, " ", "")
	reg_question := "Sure."
	reg_yell := "Whoa, chill out!"
	yell_question := "Calm down, I know what I'm doing!"
	silence := "Fine. Be that way!"
	other := "Whatever."

	fmt.Println(remark)
	re_qMark := regexp.MustCompile(`\?$`)
	re_lowcase := regexp.MustCompile(`[a-z]`)
	re_space := regexp.MustCompile(`\s`)
	re_upcase := regexp.MustCompile(`[A-Z]`)
	re_anyChar := regexp.MustCompile(`\S`)
	//re_digit := regexp.MustCompile(`\d`)

	flag_qMark := re_qMark.MatchString(remark)
	flag_allCaps := !re_lowcase.MatchString(remark)
	flag_space := re_space.MatchString(remark)
	flag_upcase := re_upcase.MatchString(remark)
	flag_anyChar := re_anyChar.MatchString(remark)
	//flag_digit := re_digit.MatchString(remark)

	switch {
	case flag_qMark && flag_allCaps && flag_upcase:
		return yell_question
	case flag_qMark:
		return reg_question
	case flag_allCaps && flag_upcase && !flag_space:
		return reg_yell
	case !flag_anyChar:
		return silence
	default:
		return other
	}
}
