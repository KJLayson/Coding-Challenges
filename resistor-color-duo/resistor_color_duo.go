package resistorcolorduo

import "strconv"

func Value(colors []string) int {
	color_map := map[string]string{
		"black":  "0",
		"brown":  "1",
		"red":    "2",
		"orange": "3",
		"yellow": "4",
		"green":  "5",
		"blue":   "6",
		"violet": "7",
		"grey":   "8",
		"white":  "9",
	}
	color1 := color_map[colors[0]]
	color2 := color_map[colors[1]]

	str_val := color1 + color2
	int_val, err := strconv.Atoi(str_val)

	if err == nil {
		return int_val
	} else {
		return 0
	}
}
