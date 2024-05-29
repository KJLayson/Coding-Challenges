package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	var err error

	switch {
	case number < 1 || number > 64:
		return 0, errors.New("invalid square number")
	default:
		big_num := uint64(math.Pow(2, float64(number-1)))
		return big_num, err
	}

}

func Total() uint64 {
	var sum uint64 = 0

	for i := 1; i <= 64; i++ {
		square, err := Square(i)
		if err != nil {
			return 0
		} else {
			sum += square
		}
	}

	return sum
}
