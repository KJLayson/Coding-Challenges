package protein

import "errors"

const stop = "STOP"

var ErrStop = errors.New(stop)
var ErrInvalidBase = errors.New("invalid base")

var codes = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": stop,
	"UAG": stop,
	"UGA": stop,
}

func FromRNA(rna string) ([]string, error) {
	list_codons := []string{}
	list_proteins := []string{}
	new_code := ""
	idx := 0

	for _, char := range rna {
		new_code += string(char)
		idx++

		if idx > 2 {
			idx = 0
			list_codons = append(list_codons, new_code)
			new_code = ""
		}
	}

	for _, codon := range list_codons {
		protein, err := FromCodon(codon)
		if err == ErrInvalidBase {
			return nil, err
		} else if err == ErrStop {
			return list_proteins, nil
		} else {
			list_proteins = append(list_proteins, protein)
		}
	}

	return list_proteins, nil
}

func FromCodon(codon string) (string, error) {
	protein, ok := codes[codon]
	if !ok {
		return "", ErrInvalidBase
	} else if protein == stop {
		return "", ErrStop
	} else {
		return protein, nil
	}
}
