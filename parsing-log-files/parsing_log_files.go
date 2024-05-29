package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	valid := map[string]bool{
		"[TRC]": true,
		"[DBG]": true,
		"[INF]": true,
		"[WRN]": true,
		"[ERR]": true,
		"[FTL]": true,
	}
	re := regexp.MustCompile(`^\[(...)]`)
	s := re.FindString(text)
	_, ok := valid[s]
	if ok {
		return true
	} else {
		return false
	}
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[*=~-]*?>`)
	s := re.Split(text, -1)
	return s
}

func CountQuotedPasswords(lines []string) int {
	tally := 0
	re := regexp.MustCompile(`".*?password.*?"`)
	for _, str := range lines {
		str = strings.ToLower(str)
		if re.MatchString(str) {
			tally += 1
		}
	}

	return tally
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\S*`)
	s := re.ReplaceAllString(text, "")
	return s
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User[^\S]\s*(?P<user>\S*)`)
	for i, str := range lines {
		s := re.FindStringSubmatch(str)
		if len(s) > 0 {
			lines[i] = fmt.Sprintf("[USR] %s %s", s[1], str)
		}
	}

	return lines
}
