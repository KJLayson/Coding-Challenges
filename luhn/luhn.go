package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	rune_id := []rune(id)
	length := len(rune_id)
	var int_id []int

	if length <= 1 {
		//fmt.Printf("error: bad length with id: %s\n", id)
		return false

	}

	for i := 0; i < length; i++ {
		curr_val := rune_id[length-(1+i)]
		int_val, err := strconv.Atoi(string(curr_val))

		if err != nil {
			//fmt.Printf("error: non-digit with id: %s\n", id)
			return false

		} else if (i+1)%2 == 0 {
			int_val *= 2

			if int_val > 9 {
				int_val -= 9
			}
		}

		int_id = append(int_id, int_val)
	}
	sum := 0
	for _, val := range int_id {
		sum += val
	}
	//fmt.Printf("pass with id: %s. sum: %d\n", id, sum)
	return sum%10 == 0
}
