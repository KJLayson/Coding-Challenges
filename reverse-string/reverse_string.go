package reverse

func Reverse(input string) string {
	str_out := ""
	if input != "" {
		runes := []rune(input)
		for i := len(runes) - 1; i >= 0; i-- {
			str_out += string(runes[i])
		}
	}
	return str_out
}
