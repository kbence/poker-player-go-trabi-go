package main

func findStraight(cards []Card) int {
	for i := 1; i < 10; i++ {
		if findByPos(cards, i) && findByPos(cards, i+1) && findByPos(cards, i+2) && findByPos(cards, i+3) && findByPos(cards, i+4) {
			return i;
		}
	}

	return -1;
}

func findByPos(cards []Card, position int) bool {
	for _, card := range cards {
		if(card.Position()==position || (card.Position()==14 && position==1)) return true
	}

	return false
}


func GetPokerLevel(cards []Card) (int, int) {
	pokerTypes := make(map[string]int)
	pokerCards := make(map[string]int)
	pokerColors := make(map[string]int)
	cardPosition := 0
	level := 1
	for _, c := range cards {
		if(pokerCards[c.Rank] == 2 && pokerTypes["drill"] == 0) {
			pokerTypes["drill"]++
		}
		if(pokerCards[c.Rank] == 1) {
			pokerTypes["pairs"]++
		}
		pokerColors[c.Suit]++
		if(pokerColors[c.Suit] == 5) {
			pokerTypes["flush"]++
		}
		pokerCards[c.Rank]++
		if(pokerCards["10"] > 0 && pokerCards["J"] > 0 && pokerCards["Q"] > 0 && pokerCards["K"] > 0 && pokerCards["A"] > 0 && pokerColors[c.Suit] == 5) {
			return 10, 1
		}
		if(pokerCards[c.Rank] == 4) {
			return 8, 1
		}
		if(cardPosition < c.Position()) {
			cardPosition = c.Position()
		}
	}
	if(pokerTypes["pairs"] == 1) {
		level = 2
	}
	if(pokerTypes["pairs"] == 2 || pokerTypes["pairs"] == 3) {
		level = 3
	}
	if(pokerTypes["drill"] == 1) {
		level = 4
	}
	if(pokerTypes["flush"] != 0) {
		level = 6
	}
	if(level == 1) {
		return 1, cardPosition
	}
	straight := findStraight(cards)
	if straight > 0 {
		return 5, straight
	}

	return level, 1
}
