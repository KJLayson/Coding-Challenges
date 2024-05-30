package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	rune_digits := []rune(digits)
	num_digits := len(rune_digits)

	switch {
	case span > num_digits:
		return 0, errors.New("span cannot be bigger than number of digits")
	case span <= 0 || digits == "" || num_digits == 0:
		return 0, errors.New("span and number of digits must be greater than 0")
	}

	for _, val := range rune_digits {
		if !unicode.IsNumber(val) {
			return 0, errors.New("input must only contain digits")
		}
	}

	series := [][]string{}
	for i := range digits {
		if i <= num_digits-span {
			new_series := []string{}
			for j := i; j < i+span; j++ {
				new_series = append(new_series, string(digits[j]))
			}

			series = append(series, new_series)
		}
	}

	var largest_product int64 = -1

	for _, numbers := range series {
		var product int64 = 0
		for i := range numbers {
			p, _ := strconv.Atoi(numbers[i])
			if i == 0 {
				product = int64(p)
			} else {
				product *= int64(p)
			}
		}

		if product > largest_product {
			largest_product = product
		}
	}

	return largest_product, nil
}
