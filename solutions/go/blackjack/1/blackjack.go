package blackjack

var blackjack = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	if value, ok := blackjack[card]; ok {
		return value
	}

	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	cardSum := ParseCard(card1) + ParseCard(card2)
	dealerCardSum := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case cardSum == 21:
		switch dealerCard {
		case "ace", "jack", "queen", "king":
			return "S"
		default:
			return "W"
		}
	case cardSum >= 17 && cardSum <= 20:
		return "S"
	case cardSum >= 12 && cardSum <= 16:
		switch {
		case dealerCardSum >= 7:
			return "H"
		default:
			return "S"
		}
	default:
		return "H"
	}
}
