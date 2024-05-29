package strand

func ToRNA(dna string) string {
	seq_map := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	str_out := ""
	for _, val := range dna {
		str_out += string(seq_map[val])
	}

	return str_out
}
