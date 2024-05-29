package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	card_map := map[string]int{
		"ace":   11,
		"king":  10,
		"queen": 10,
		"jack":  10,
		"ten":   10,
		"nine":  9,
		"eight": 8,
		"seven": 7,
		"six":   6,
		"five":  5,
		"four":  4,
		"three": 3,
		"two":   2,
	}

	return card_map[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	pair_value := ParseCard(card1) + ParseCard(card2)
	dealer_value := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case pair_value == 21 && dealer_value < 10:
		return "W"
	case pair_value == 21 && dealer_value >= 10:
		return "S"
	case pair_value >= 17 && pair_value <= 20:
		return "S"
	case pair_value >= 12 && pair_value <= 16 && dealer_value < 7:
		return "S"
	default:
		return "H"
	}

}
