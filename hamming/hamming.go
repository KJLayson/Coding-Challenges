package hamming

import "errors"

func Distance(a, b string) (int, error) {
	seq_a := []rune(a)
	seq_b := []rune(b)

	if len(seq_a) != len(seq_b) && seq_a != nil {
		return 0, errors.New("sequence mismatch")
	}

	tally := 0

	for i := range seq_a {
		if seq_a[i] != seq_b[i] {
			tally++
		}
	}

	return tally, nil
}
