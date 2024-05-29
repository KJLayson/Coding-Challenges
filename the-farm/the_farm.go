package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, num_cows int) (float64, error) {
	amt, error1 := fc.FodderAmount(num_cows)
	factor, error2 := fc.FatteningFactor()
	fmt.Println(amt, factor, num_cows, error1, error2)
	if error1 == nil && error2 == nil {
		total := amt * factor / float64(num_cows)
		return total, error1
	} else if error1 != nil {
		return 0, error1
	} else {
		return 0, error2
	}
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, num_cows int) (float64, error) {
	if num_cows > 0 {
		return DivideFood(fc, num_cows)
	} else {
		return float64(0), errors.New("invalid number of cows")
	}
}

type InvalidCowsError struct {
	num_cows int
	msg      error
}

func (err *InvalidCowsError) validation() {
	if err.num_cows > 0 {
		err.msg = nil
	} else if err.num_cows == 0 {
		str := fmt.Sprintf("%d cows are invalid: no cows don't need food", err.num_cows)
		err.msg = errors.New(str)
	} else {
		str := fmt.Sprintf("%d cows are invalid: there are no negative cows", err.num_cows)
		err.msg = errors.New(str)
	}
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(num_cows int) error {
	new_err := InvalidCowsError{num_cows: num_cows}
	new_err.validation()
	return new_err.msg
}
