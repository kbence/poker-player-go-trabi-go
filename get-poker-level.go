package main

func findStraight(cards []Card, suitCheck func(Card) bool) int {
	for i := 1; i <= 10; i++ {
		if findByPos(cards, i, suitCheck) && findByPos(cards, i+1, suitCheck) && findByPos(cards, i+2, suitCheck) && findByPos(cards, i+3, suitCheck) && findByPos(cards, i+4, suitCheck) {
			return i
		}
	}

	return -1
}

func findByPos(cards []Card, position int, suitCheck func(Card) bool) bool {
	for _, card := range cards {
		if suitCheck(card) && (card.Position() == position || (card.Position() == 14 && position == 1)) {
			return true
		}
	}

	return false
}

func GetPokerLevel(cards []Card) (int, int) {
	suits := []string{"clubs", "spades", "hearts", "diamonds"}
	for _, suit := range suits {
		straight := findStraight(cards, func(card Card) bool { return card.Suit == suit })
		if straight == 10 {
			return 10, straight
		}
		if straight > 0 {
			return 9, straight
		}
	}

	pokerTypes := make(map[string]int)
	pokerCards := make(map[string]int)
	pokerColors := make(map[string]int)
	cardPosition := 0
	level := 1
	for _, c := range cards {
		if pokerCards[c.Rank] == 2 && pokerTypes["drill"] == 0 {
			pokerTypes["drill"]++
		}
		if pokerCards[c.Rank] == 1 {
			pokerTypes["pairs"]++
		}
		pokerColors[c.Suit]++
		if pokerColors[c.Suit] == 5 {
			pokerTypes["flush"]++
		}
		pokerCards[c.Rank]++
		if pokerCards[c.Rank] == 4 {
			return 8, 1
		}
		if cardPosition < c.Position() {
			cardPosition = c.Position()
		}
	}
	if pokerTypes["pairs"] == 1 {
		return 2, cardPosition
	}
	if pokerTypes["pairs"] == 2 || pokerTypes["pairs"] == 3 {
		level = 3
	}
	if pokerTypes["drill"] == 1 {
		level = 4
	}
	if pokerTypes["flush"] != 0 {
		level = 6
	}
	straight := findStraight(cards, func(Card) bool { return true })
	if straight > 0 {
		return 5, straight
	}
	if level == 1 {
		return 1, cardPosition
	}

	return level, 1
}
