package resistorcolor

var color_map map[string]int

// Colors returns the list of all colors.
func Colors() []string {
	if color_map == nil {
		color_map = map[string]int{
			"black":  0,
			"brown":  1,
			"red":    2,
			"orange": 3,
			"yellow": 4,
			"green":  5,
			"blue":   6,
			"violet": 7,
			"grey":   8,
			"white":  9,
		}
	}

	var color_list []string

	for i := range color_map {
		color_list = append(color_list, i)
	}

	return color_list
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	if color_map == nil {
		color_map = map[string]int{
			"black":  0,
			"brown":  1,
			"red":    2,
			"orange": 3,
			"yellow": 4,
			"green":  5,
			"blue":   6,
			"violet": 7,
			"grey":   8,
			"white":  9,
		}
	}

	return color_map[color]
}
