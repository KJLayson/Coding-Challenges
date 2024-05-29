package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	var err error

	if n < 1 {
		err = errors.New("number must be greater than 1")

		return 0, err
	}

	tally := 0

	for ; n != 1; tally++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = (3 * n) + 1
		}
	}

	return tally, err

}
