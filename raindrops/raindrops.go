package raindrops

import "fmt"

func Convert(number int) string {
	remainder3 := number % 3
	remainder5 := number % 5
	remainder7 := number % 7

	out_str := ""

	if remainder3 == 0 {
		out_str += "Pling"
	}

	if remainder5 == 0 {
		out_str += "Plang"
	}

	if remainder7 == 0 {
		out_str += "Plong"
	}

	if out_str == "" {
		out_str = fmt.Sprintf("%d", number)
	}

	return out_str
}
