package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	card1Parsed := ParseCard(card1)
	card2Parsed := ParseCard(card2)
	handSum := card1Parsed + card2Parsed
	dealerCardParsed := ParseCard(dealerCard)

	switch {
	case card1Parsed == 11 && card2Parsed == 11:
		return "P"
	case (handSum == 21) && !(dealerCardParsed == 11 || dealerCardParsed == 10):
		return "W"
	case (handSum == 21) && (dealerCardParsed == 11 || dealerCardParsed == 10):
		return "S"
	case handSum >= 17 && handSum <= 20:
		return "S"
	case (handSum >= 12 && handSum <= 16) && dealerCardParsed < 7:
		return "S"
	case (handSum >= 12 && handSum <= 16) && dealerCardParsed >= 7:
		return "H"
	case handSum <= 11:
		return "H"
	default:
		return "N"
	}
}
