package triangle

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	const min_len = 0

	a_plus_b := a + b
	a_plus_c := a + c
	b_plus_c := b + c

	switch {
	case a <= min_len || b <= min_len || c <= min_len:
		k = NaT
	case a_plus_b < c || a_plus_c < b || b_plus_c < a:
		k = NaT
	case a == b && b == c:
		k = Equ
	case a != b && a != c && b != c:
		k = Sca
	default:
		k = Iso
	}

	return k
}
