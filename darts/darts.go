package darts

import (
	"math"
)

func Score(x, y float64) int {
	const outside = 0
	const outer = 1
	const middle = 5
	const inner = 10
	score := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	switch {
	case score > 10:
		return outside
	case score <= 10 && score > 5:
		return outer
	case score > 1:
		return middle
	default:
		return inner

	}
}
