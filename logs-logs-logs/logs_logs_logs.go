package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {

	codes := map[string]string{
		"U+2757":  "recommendation",
		"U+1F50D": "search",
		"U+2600":  "weather",
	}

	for _, char := range log {
		code := fmt.Sprintf("%U", char)
		val, exists := codes[code]
		if exists {
			return val
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	old := fmt.Sprintf("%c", oldRune)
	new := fmt.Sprintf("%c", newRune)

	new_string := strings.ReplaceAll(log, old, new)

	return new_string
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	numberOfRunes := utf8.RuneCountInString(log)

	if numberOfRunes <= limit {
		return true
	} else {
		return false
	}
}
