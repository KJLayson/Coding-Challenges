package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	_, ok := fnb.(FancyNumber)
	if ok {
		i, _ := strconv.Atoi(fnb.Value())
		return i
	} else {
		return 0
	}

}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	number := 0
	_, ok := fnb.(FancyNumber)
	if ok {
		number, _ = strconv.Atoi(fnb.Value())
	}

	return fmt.Sprintf("This is a fancy box containing the number %d.0", number)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch i.(type) {
	case int:
		num, _ := i.(int)
		return DescribeNumber(float64(num))
	case float64:
		num, _ := i.(float64)
		return DescribeNumber(num)
	case NumberBox:
		box, _ := i.(NumberBox)
		return DescribeNumberBox(box)
	case FancyNumberBox:
		box, _ := i.(FancyNumber)
		return DescribeFancyNumberBox(box)
	default:
		return "Return to sender"
	}
}
