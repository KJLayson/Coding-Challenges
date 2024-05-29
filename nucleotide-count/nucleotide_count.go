package dna

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	var err error
	h := Histogram{
		'A': 0,
		'T': 0,
		'C': 0,
		'G': 0,
	}

	for _, val := range d {
		if val != 'A' && val != 'T' && val != 'C' && val != 'G' {
			return nil, errors.New("invalid sequence")
		} else {
			h[val] += 1
		}
	}

	return h, err
}
